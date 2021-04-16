// Code generated by go-swagger; DO NOT EDIT.

package international_standing_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new international standing orders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for international standing orders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateInternationalStandingOrderConsents(params *CreateInternationalStandingOrderConsentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInternationalStandingOrderConsentsCreated, error)

	CreateInternationalStandingOrders(params *CreateInternationalStandingOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInternationalStandingOrdersCreated, error)

	GetInternationalStandingOrderConsentsConsentID(params *GetInternationalStandingOrderConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInternationalStandingOrderConsentsConsentIDOK, error)

	GetInternationalStandingOrdersInternationalStandingOrderPaymentID(params *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateInternationalStandingOrderConsents creates international standing order consents
*/
func (a *Client) CreateInternationalStandingOrderConsents(params *CreateInternationalStandingOrderConsentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInternationalStandingOrderConsentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInternationalStandingOrderConsentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateInternationalStandingOrderConsents",
		Method:             "POST",
		PathPattern:        "/international-standing-order-consents",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateInternationalStandingOrderConsentsReader{formats: a.formats},
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
	success, ok := result.(*CreateInternationalStandingOrderConsentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateInternationalStandingOrderConsents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateInternationalStandingOrders creates international standing orders
*/
func (a *Client) CreateInternationalStandingOrders(params *CreateInternationalStandingOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInternationalStandingOrdersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInternationalStandingOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateInternationalStandingOrders",
		Method:             "POST",
		PathPattern:        "/international-standing-orders",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateInternationalStandingOrdersReader{formats: a.formats},
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
	success, ok := result.(*CreateInternationalStandingOrdersCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateInternationalStandingOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInternationalStandingOrderConsentsConsentID gets international standing order consents
*/
func (a *Client) GetInternationalStandingOrderConsentsConsentID(params *GetInternationalStandingOrderConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInternationalStandingOrderConsentsConsentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInternationalStandingOrderConsentsConsentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInternationalStandingOrderConsentsConsentId",
		Method:             "GET",
		PathPattern:        "/international-standing-order-consents/{ConsentId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInternationalStandingOrderConsentsConsentIDReader{formats: a.formats},
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
	success, ok := result.(*GetInternationalStandingOrderConsentsConsentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInternationalStandingOrderConsentsConsentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInternationalStandingOrdersInternationalStandingOrderPaymentID gets international standing orders
*/
func (a *Client) GetInternationalStandingOrdersInternationalStandingOrderPaymentID(params *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInternationalStandingOrdersInternationalStandingOrderPaymentId",
		Method:             "GET",
		PathPattern:        "/international-standing-orders/{InternationalStandingOrderPaymentId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDReader{formats: a.formats},
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
	success, ok := result.(*GetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInternationalStandingOrdersInternationalStandingOrderPaymentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}