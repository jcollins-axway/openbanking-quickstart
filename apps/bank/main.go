package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	logrus "github.com/sirupsen/logrus"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Spec string

const (
	OBUK Spec = "obuk"
	OBBR Spec = "obbr"
)

type Config struct {
	Port         int           `env:"PORT" envDefault:"8070"`
	ClientID     string        `env:"CLIENT_ID,required"`
	IssuerURL    *url.URL      `env:"ISSUER_URL,required"`
	Timeout      time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA       string        `env:"ROOT_CA,required"`
	CertFile     string        `env:"CERT_FILE,required"`
	KeyFile      string        `env:"KEY_FILE,required"`
	Spec         Spec          `env:"SPEC,required"`
	SeedFilePath string
}

func (c *Config) ClientConfig() acpclient.Config {
	return acpclient.Config{
		ClientID:  c.ClientID,
		IssuerURL: c.IssuerURL,
		Scopes:    []string{"introspect_openbanking_tokens"},
		Timeout:   c.Timeout,
		CertFile:  c.CertFile,
		KeyFile:   c.KeyFile,
		RootCA:    c.RootCA,
	}
}

func LoadConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, err
	}

	switch config.Spec {
	case OBUK:
		config.SeedFilePath = fmt.Sprintf("data/%s-data.json", OBUK)
	case OBBR:
		config.SeedFilePath = fmt.Sprintf("data/%s-data.json", OBBR)
	}

	return config, err
}

type Server struct {
	Config  Config
	Client  acpclient.Client
	Storage Storage
	// PaymentQueue PaymentQueue

	MakeGetAccountsHandler         GetEndpointLogicFactory
	MakeGetAccountsInternalHandler GetEndpointLogicFactory

	MakeGetBalancesHandler         GetEndpointLogicFactory
	MakeGetBalancesInternalHandler GetEndpointLogicFactory

	MakeGetTransactionsHandler GetEndpointLogicFactory

	MakeCreatePaymentHandler CreateEndpointLogicFactory
	MakeGetPaymentHandler    GetEndpointLogicFactory
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if server.Client, err = acpclient.New(server.Config.ClientConfig()); err != nil {
		return server, errors.Wrapf(err, "failed to init acp client")
	}

	if server.Storage, err = NewUserRepo(server.Config.SeedFilePath); err != nil {
		return server, errors.Wrapf(err, "failed to init repo")
	}

	switch server.Config.Spec {
	case OBUK:
		server.MakeGetAccountsHandler = NewOBUKGetAccountsHandler
		server.MakeGetAccountsInternalHandler = NewOBUKGetAccountsInternalHandler
		server.MakeGetBalancesHandler = NewOBUKGetBalancesHandler
		server.MakeGetBalancesInternalHandler = NewOBUKGetBalancesInternalHandler
		server.MakeGetTransactionsHandler = NewOBUKGetTransactionsHandler
		server.MakeCreatePaymentHandler = NewOBUKCreatePaymentHandler
		server.MakeGetPaymentHandler = NewOBUKGetPaymentHandler
	case OBBR:
		server.MakeGetAccountsHandler = NewOBBRGetAccountsHandler
		server.MakeGetAccountsInternalHandler = NewOBBRGetAccountsInternalHandler
		server.MakeCreatePaymentHandler = NewOBBRCreatePaymentHandler
	default:
		return server, errors.Wrapf(err, "unsupported spec %s", server.Config.Spec)
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()

	r.GET("/internal/accounts", s.Get(s.MakeGetAccountsInternalHandler))

	switch s.Config.Spec {
	case OBUK:
		r.GET("/accounts", s.Get(s.MakeGetAccountsHandler))
		r.GET("/balances", s.Get(s.MakeGetBalancesHandler))
		r.GET("/internal/balances", s.Get(s.MakeGetBalancesInternalHandler))
		r.GET("/transactions", s.Get(s.MakeGetTransactionsHandler))
		r.POST("/domestic-payments", s.Post(s.MakeCreatePaymentHandler))
		r.GET("/domestic-payments/:DomesticPaymentId", s.Get(s.MakeGetPaymentHandler))

	case OBBR:
		r.GET("/accounts/v1/accounts", s.Get(s.MakeGetAccountsHandler))
		r.GET("/payments/v1/pix/payments", s.Post(s.MakeCreatePaymentHandler))

	default:
		return fmt.Errorf("unsupported spec %s", s.Config.Spec)
	}

	return r.Run(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)))
}

func main() {
	var (
		server Server
		err    error
	)

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if server, err = NewServer(); err != nil {
		logrus.WithError(err).Fatalf("failed to init server")
	}

	if err = server.Start(); err != nil {
		logrus.WithError(err).Fatalf("failed to start server")
	}
}
