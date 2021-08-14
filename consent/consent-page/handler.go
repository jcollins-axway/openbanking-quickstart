package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	ID          string
	State       string
	ConsentType string
}

func NewLoginRequest(c *gin.Context) LoginRequest {
	return LoginRequest{
		ID:          c.Query("login_id"),
		State:       c.Query("login_state"),
		ConsentType: c.Query("consent_type"),
	}
}

func (l *LoginRequest) Validate() error {
	if l.ID == "" || l.State == "" || l.ConsentType == "" {
		return errors.New("login_id / login_state / consent_type missing")
	}

	return nil
}

func (s *Server) GetTemplateNameForSpec(basename string) string {
	switch s.Config.Spec {
	case OBUK:
		return string(OBUK) + "-" + basename
	case OBBR:
		return string(OBBR) + "-" + basename
	}

	return basename
}

func (s *Server) WithConsentHandler(c *gin.Context) (ConsentHandler, LoginRequest, error) {
	var (
		loginRequest = NewLoginRequest(c)
		handler      ConsentHandler
		err          error
		ok           bool
	)

	if err = loginRequest.Validate(); err != nil {
		return handler, loginRequest, err
	}

	if handler, ok = s.GetConsentHandler(loginRequest); !ok {
		return handler, loginRequest, fmt.Errorf("invalid consent type %s", loginRequest.ConsentType)
	}

	return handler, loginRequest, nil
}

func (s *Server) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			handler      ConsentHandler
			loginRequest LoginRequest
			err          error
		)

		if handler, loginRequest, err = s.WithConsentHandler(c); err != nil {
			RenderInvalidRequestError(c, s.Trans, err)
			return
		}

		handler.GetConsent(c, loginRequest)
	}
}

func (s *Server) Post() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			handler ConsentHandler

			loginRequest LoginRequest
			err          error
		)

		if handler, loginRequest, err = s.WithConsentHandler(c); err != nil {
			RenderInvalidRequestError(c, s.Trans, err)
			return
		}

		s.PostConsent(c, loginRequest, handler)
	}
}

func (s *Server) PostConsent(c *gin.Context, loginRequest LoginRequest, consentHandler ConsentHandler) {
	var (
		redirect string
		err      error
	)

	action := c.PostForm("action")

	switch action {
	case "confirm":
		redirect, err = consentHandler.ConfirmConsent(c, loginRequest)
	case "deny":
		redirect, err = consentHandler.DenyConsent(c, loginRequest)
	default:
		RenderInvalidRequestError(c, s.Trans, fmt.Errorf("invalid form action: %s", action))
		return
	}

	if err != nil {
		RenderInternalServerError(c, s.Trans, errors.Wrapf(err, "failed to %s consent: %+v", c.PostForm("action"), loginRequest.ConsentType))
		return
	}

	logrus.Debugf("redirect to: %s", redirect)

	c.Redirect(http.StatusFound, redirect)
}

func (s *Server) GetConsentHandler(loginRequest LoginRequest) (ConsentHandler, bool) {
	switch loginRequest.ConsentType {
	case "domestic_payment", "payments":
		return s.PaymentConsentHandler, true
	case "account_access", "consents":
		return s.AccountAccessConsentHandler, true
	default:
		return nil, false
	}
}

type ConsentHandler interface {
	GetConsent(c *gin.Context, loginRequest LoginRequest)
	ConfirmConsent(c *gin.Context, loginRequest LoginRequest) (redirect string, err error)
	DenyConsent(c *gin.Context, loginRequest LoginRequest) (redirect string, err error)
}
