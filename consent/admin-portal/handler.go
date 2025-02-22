package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp-client-go/client/clients"
	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
)

const (
	AccountAccessConsentType                 = "account_access"
	DomesticPaymentConsentType               = "domestic_payment"
	DomesticScheduledPaymentConsentType      = "domestic_scheduled_payment"
	DomesticStandingOrderConsentType         = "domestic_standing_order"
	InternationalPaymentConsentType          = "international_payment"
	InternationalScheduledPaymentConsentType = "international_scheduled_payment"
	InternationalStandingOrderConsentType    = "international_standing_order"
	FilePaymentConsentType                   = "file_payment"
)

var ConsentTypes = []string{
	AccountAccessConsentType,
	DomesticPaymentConsentType,
	DomesticScheduledPaymentConsentType,
	DomesticStandingOrderConsentType,
	InternationalPaymentConsentType,
	InternationalStandingOrderConsentType,
	InternationalScheduledPaymentConsentType,
	FilePaymentConsentType,
}

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

type Client struct {
	ID       string                                 `json:"client_id"`
	Name     string                                 `json:"client_name,omitempty"`
	Consents []*models.OpenbankingConsentWithClient `json:"consents"`
}

type ListClientsResponse struct {
	Clients []Client `json:"clients"`
}

func (s *Server) ListClients() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			cs                  *clients.ListClientsSystemOK
			consents            *openbanking.ListOBConsentsOK
			clientsWithConsents []Client
			err                 error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if cs, err = s.Client.Clients.ListClientsSystem(
			clients.NewListClientsSystemParamsWithContext(c).
				WithTid(s.Client.TenantID).
				WithAid(s.Config.SystemClientsServerID),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to list clients from acp: %+v", err))
			return
		}

		for _, oc := range cs.Payload.Clients {
			if consents, err = s.Client.Openbanking.ListOBConsents(
				openbanking.NewListOBConsentsParamsWithContext(c).
					WithTid(s.Client.TenantID).
					WithAid(s.Config.SystemClientsServerID).
					WithConsentsRequest(&models.ConsentsRequest{
						ClientID: oc.ClientID,
					}),
				nil,
			); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to list consents for client: %s, err: %+v", oc.ClientID, err))
				return
			}

			if !oc.System {
				clientsWithConsents = append(clientsWithConsents, Client{
					ID:       oc.ClientID,
					Name:     oc.ClientName,
					Consents: consents.Payload.Consents,
				})
			}
		}

		resp := ListClientsResponse{Clients: clientsWithConsents}

		c.JSON(http.StatusOK, &resp)
	}
}

func (s *Server) RevokeConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id  = c.Param("id")
			err error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if _, err = s.Client.Openbanking.RevokeOpenbankingConsent(
			openbanking.NewRevokeOpenbankingConsentParamsWithContext(c).
				WithTid(s.Client.TenantID).
				WithAid(s.Config.SystemClientsServerID).
				WithConsentID(id),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to revoke account access consent: %+v", err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (s *Server) RevokeConsentsForClient() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id  = c.Param("id")
			err error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if _, err = s.Client.Openbanking.RevokeOpenbankingConsents(
			openbanking.NewRevokeOpenbankingConsentsParamsWithContext(c).
				WithTid(s.Client.TenantID).
				WithAid(s.Config.SystemClientsServerID).
				WithConsentTypes(ConsentTypes).
				WithClientID(&id),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to revoke consents for client: %s, err: %+v", id, err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (s *Server) IntrospectToken(c *gin.Context) error {
	var err error

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if _, err = s.IntrospectClient.IntrospectToken(c, token); err != nil {
		return fmt.Errorf("failed to introspect client: %w", err)
	}

	return nil
}
