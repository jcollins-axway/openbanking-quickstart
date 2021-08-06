// Code generated by go-swagger; DO NOT EDIT.

package domestic_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new domestic payments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for domestic payments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDomesticPaymentConsents(params *CreateDomesticPaymentConsentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDomesticPaymentConsentsCreated, error)

	CreateDomesticPayments(params *CreateDomesticPaymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDomesticPaymentsCreated, error)

	GetDomesticPaymentConsentsConsentID(params *GetDomesticPaymentConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDomesticPaymentConsentsConsentIDOK, error)

	GetDomesticPaymentConsentsConsentIDFundsConfirmation(params *GetDomesticPaymentConsentsConsentIDFundsConfirmationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDomesticPaymentConsentsConsentIDFundsConfirmationOK, error)

	GetDomesticPaymentsDomesticPaymentID(params *GetDomesticPaymentsDomesticPaymentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDomesticPaymentsDomesticPaymentIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDomesticPaymentConsents creates domestic payment consents
*/
func (a *Client) CreateDomesticPaymentConsents(params *CreateDomesticPaymentConsentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDomesticPaymentConsentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDomesticPaymentConsentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateDomesticPaymentConsents",
		Method:             "POST",
		PathPattern:        "/domestic-payment-consents",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDomesticPaymentConsentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDomesticPaymentConsentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDomesticPaymentConsents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateDomesticPayments creates domestic payments
*/
func (a *Client) CreateDomesticPayments(params *CreateDomesticPaymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDomesticPaymentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDomesticPaymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateDomesticPayments",
		Method:             "POST",
		PathPattern:        "/domestic-payments",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDomesticPaymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDomesticPaymentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDomesticPayments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDomesticPaymentConsentsConsentID gets domestic payment consents
*/
func (a *Client) GetDomesticPaymentConsentsConsentID(params *GetDomesticPaymentConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDomesticPaymentConsentsConsentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomesticPaymentConsentsConsentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDomesticPaymentConsentsConsentId",
		Method:             "GET",
		PathPattern:        "/domestic-payment-consents/{ConsentId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomesticPaymentConsentsConsentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomesticPaymentConsentsConsentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDomesticPaymentConsentsConsentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDomesticPaymentConsentsConsentIDFundsConfirmation gets domestic payment consents funds confirmation
*/
func (a *Client) GetDomesticPaymentConsentsConsentIDFundsConfirmation(params *GetDomesticPaymentConsentsConsentIDFundsConfirmationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDomesticPaymentConsentsConsentIDFundsConfirmationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomesticPaymentConsentsConsentIDFundsConfirmationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDomesticPaymentConsentsConsentIdFundsConfirmation",
		Method:             "GET",
		PathPattern:        "/domestic-payment-consents/{ConsentId}/funds-confirmation",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomesticPaymentConsentsConsentIDFundsConfirmationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomesticPaymentConsentsConsentIDFundsConfirmationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDomesticPaymentConsentsConsentIdFundsConfirmation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDomesticPaymentsDomesticPaymentID gets domestic payments
*/
func (a *Client) GetDomesticPaymentsDomesticPaymentID(params *GetDomesticPaymentsDomesticPaymentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDomesticPaymentsDomesticPaymentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomesticPaymentsDomesticPaymentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDomesticPaymentsDomesticPaymentId",
		Method:             "GET",
		PathPattern:        "/domestic-payments/{DomesticPaymentId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomesticPaymentsDomesticPaymentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomesticPaymentsDomesticPaymentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDomesticPaymentsDomesticPaymentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}