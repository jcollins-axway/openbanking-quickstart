// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obbr/accounts/models"
)

// AccountsGetAccountsAccountIDTransactionsReader is a Reader for the AccountsGetAccountsAccountIDTransactions structure.
type AccountsGetAccountsAccountIDTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountsGetAccountsAccountIDTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountsGetAccountsAccountIDTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountsGetAccountsAccountIDTransactionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountsGetAccountsAccountIDTransactionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccountsGetAccountsAccountIDTransactionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountsGetAccountsAccountIDTransactionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewAccountsGetAccountsAccountIDTransactionsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAccountsGetAccountsAccountIDTransactionsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAccountsGetAccountsAccountIDTransactionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccountsGetAccountsAccountIDTransactionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAccountsGetAccountsAccountIDTransactionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountsGetAccountsAccountIDTransactionsOK creates a AccountsGetAccountsAccountIDTransactionsOK with default headers values
func NewAccountsGetAccountsAccountIDTransactionsOK() *AccountsGetAccountsAccountIDTransactionsOK {
	return &AccountsGetAccountsAccountIDTransactionsOK{}
}

/* AccountsGetAccountsAccountIDTransactionsOK describes a response with status code 200, with default header values.

Dados da lista de transações da conta identificada por accountId obtidos com sucesso.
*/
type AccountsGetAccountsAccountIDTransactionsOK struct {
	XFapiInteractionID string

	Payload *models.ResponseAccountTransactions
}

func (o *AccountsGetAccountsAccountIDTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/transactions][%d] accountsGetAccountsAccountIdTransactionsOK  %+v", 200, o.Payload)
}
func (o *AccountsGetAccountsAccountIDTransactionsOK) GetPayload() *models.ResponseAccountTransactions {
	return o.Payload
}

func (o *AccountsGetAccountsAccountIDTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.ResponseAccountTransactions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetAccountsAccountIDTransactionsBadRequest creates a AccountsGetAccountsAccountIDTransactionsBadRequest with default headers values
func NewAccountsGetAccountsAccountIDTransactionsBadRequest() *AccountsGetAccountsAccountIDTransactionsBadRequest {
	return &AccountsGetAccountsAccountIDTransactionsBadRequest{}
}

/* AccountsGetAccountsAccountIDTransactionsBadRequest describes a response with status code 400, with default header values.

A requisição foi malformada, omitindo atributos obrigatórios, seja no payload ou através de atributos na URL.
*/
type AccountsGetAccountsAccountIDTransactionsBadRequest struct {
	Payload *models.ResponseError
}

func (o *AccountsGetAccountsAccountIDTransactionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/transactions][%d] accountsGetAccountsAccountIdTransactionsBadRequest  %+v", 400, o.Payload)
}
func (o *AccountsGetAccountsAccountIDTransactionsBadRequest) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *AccountsGetAccountsAccountIDTransactionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetAccountsAccountIDTransactionsUnauthorized creates a AccountsGetAccountsAccountIDTransactionsUnauthorized with default headers values
func NewAccountsGetAccountsAccountIDTransactionsUnauthorized() *AccountsGetAccountsAccountIDTransactionsUnauthorized {
	return &AccountsGetAccountsAccountIDTransactionsUnauthorized{}
}

/* AccountsGetAccountsAccountIDTransactionsUnauthorized describes a response with status code 401, with default header values.

Cabeçalho de autenticação ausente/inválido ou token inválido
*/
type AccountsGetAccountsAccountIDTransactionsUnauthorized struct {
	Payload *models.ResponseError
}

func (o *AccountsGetAccountsAccountIDTransactionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/transactions][%d] accountsGetAccountsAccountIdTransactionsUnauthorized  %+v", 401, o.Payload)
}
func (o *AccountsGetAccountsAccountIDTransactionsUnauthorized) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *AccountsGetAccountsAccountIDTransactionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetAccountsAccountIDTransactionsForbidden creates a AccountsGetAccountsAccountIDTransactionsForbidden with default headers values
func NewAccountsGetAccountsAccountIDTransactionsForbidden() *AccountsGetAccountsAccountIDTransactionsForbidden {
	return &AccountsGetAccountsAccountIDTransactionsForbidden{}
}

/* AccountsGetAccountsAccountIDTransactionsForbidden describes a response with status code 403, with default header values.

O token tem escopo incorreto ou uma política de segurança foi violada
*/
type AccountsGetAccountsAccountIDTransactionsForbidden struct {
	Payload *models.ResponseError
}

func (o *AccountsGetAccountsAccountIDTransactionsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/transactions][%d] accountsGetAccountsAccountIdTransactionsForbidden  %+v", 403, o.Payload)
}
func (o *AccountsGetAccountsAccountIDTransactionsForbidden) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *AccountsGetAccountsAccountIDTransactionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetAccountsAccountIDTransactionsNotFound creates a AccountsGetAccountsAccountIDTransactionsNotFound with default headers values
func NewAccountsGetAccountsAccountIDTransactionsNotFound() *AccountsGetAccountsAccountIDTransactionsNotFound {
	return &AccountsGetAccountsAccountIDTransactionsNotFound{}
}

/* AccountsGetAccountsAccountIDTransactionsNotFound describes a response with status code 404, with default header values.

O recurso solicitado não existe ou não foi implementado
*/
type AccountsGetAccountsAccountIDTransactionsNotFound struct {
	Payload *models.ResponseError
}

func (o *AccountsGetAccountsAccountIDTransactionsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/transactions][%d] accountsGetAccountsAccountIdTransactionsNotFound  %+v", 404, o.Payload)
}
func (o *AccountsGetAccountsAccountIDTransactionsNotFound) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *AccountsGetAccountsAccountIDTransactionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetAccountsAccountIDTransactionsMethodNotAllowed creates a AccountsGetAccountsAccountIDTransactionsMethodNotAllowed with default headers values
func NewAccountsGetAccountsAccountIDTransactionsMethodNotAllowed() *AccountsGetAccountsAccountIDTransactionsMethodNotAllowed {
	return &AccountsGetAccountsAccountIDTransactionsMethodNotAllowed{}
}

/* AccountsGetAccountsAccountIDTransactionsMethodNotAllowed describes a response with status code 405, with default header values.

O consumidor tentou acessar o recurso com um método não suportado
*/
type AccountsGetAccountsAccountIDTransactionsMethodNotAllowed struct {
	Payload *models.ResponseError
}

func (o *AccountsGetAccountsAccountIDTransactionsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/transactions][%d] accountsGetAccountsAccountIdTransactionsMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *AccountsGetAccountsAccountIDTransactionsMethodNotAllowed) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *AccountsGetAccountsAccountIDTransactionsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetAccountsAccountIDTransactionsNotAcceptable creates a AccountsGetAccountsAccountIDTransactionsNotAcceptable with default headers values
func NewAccountsGetAccountsAccountIDTransactionsNotAcceptable() *AccountsGetAccountsAccountIDTransactionsNotAcceptable {
	return &AccountsGetAccountsAccountIDTransactionsNotAcceptable{}
}

/* AccountsGetAccountsAccountIDTransactionsNotAcceptable describes a response with status code 406, with default header values.

A solicitação continha um cabeçalho Accept diferente dos tipos de mídia permitidos ou um conjunto de caracteres diferente de UTF-8
*/
type AccountsGetAccountsAccountIDTransactionsNotAcceptable struct {
	Payload *models.ResponseError
}

func (o *AccountsGetAccountsAccountIDTransactionsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/transactions][%d] accountsGetAccountsAccountIdTransactionsNotAcceptable  %+v", 406, o.Payload)
}
func (o *AccountsGetAccountsAccountIDTransactionsNotAcceptable) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *AccountsGetAccountsAccountIDTransactionsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetAccountsAccountIDTransactionsTooManyRequests creates a AccountsGetAccountsAccountIDTransactionsTooManyRequests with default headers values
func NewAccountsGetAccountsAccountIDTransactionsTooManyRequests() *AccountsGetAccountsAccountIDTransactionsTooManyRequests {
	return &AccountsGetAccountsAccountIDTransactionsTooManyRequests{}
}

/* AccountsGetAccountsAccountIDTransactionsTooManyRequests describes a response with status code 429, with default header values.

A operação foi recusada, pois muitas solicitações foram feitas dentro de um determinado período ou o limite global de requisições concorrentes foi atingido
*/
type AccountsGetAccountsAccountIDTransactionsTooManyRequests struct {
	Payload *models.ResponseError
}

func (o *AccountsGetAccountsAccountIDTransactionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/transactions][%d] accountsGetAccountsAccountIdTransactionsTooManyRequests  %+v", 429, o.Payload)
}
func (o *AccountsGetAccountsAccountIDTransactionsTooManyRequests) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *AccountsGetAccountsAccountIDTransactionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetAccountsAccountIDTransactionsInternalServerError creates a AccountsGetAccountsAccountIDTransactionsInternalServerError with default headers values
func NewAccountsGetAccountsAccountIDTransactionsInternalServerError() *AccountsGetAccountsAccountIDTransactionsInternalServerError {
	return &AccountsGetAccountsAccountIDTransactionsInternalServerError{}
}

/* AccountsGetAccountsAccountIDTransactionsInternalServerError describes a response with status code 500, with default header values.

Ocorreu um erro no gateway da API ou no microsserviço
*/
type AccountsGetAccountsAccountIDTransactionsInternalServerError struct {
	Payload *models.ResponseError
}

func (o *AccountsGetAccountsAccountIDTransactionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/transactions][%d] accountsGetAccountsAccountIdTransactionsInternalServerError  %+v", 500, o.Payload)
}
func (o *AccountsGetAccountsAccountIDTransactionsInternalServerError) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *AccountsGetAccountsAccountIDTransactionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsGetAccountsAccountIDTransactionsDefault creates a AccountsGetAccountsAccountIDTransactionsDefault with default headers values
func NewAccountsGetAccountsAccountIDTransactionsDefault(code int) *AccountsGetAccountsAccountIDTransactionsDefault {
	return &AccountsGetAccountsAccountIDTransactionsDefault{
		_statusCode: code,
	}
}

/* AccountsGetAccountsAccountIDTransactionsDefault describes a response with status code -1, with default header values.

Dados da lista de transações da conta identificada por accountId obtidos com sucesso.
*/
type AccountsGetAccountsAccountIDTransactionsDefault struct {
	_statusCode        int
	XFapiInteractionID string

	Payload *models.ResponseAccountTransactions
}

// Code gets the status code for the accounts get accounts account Id transactions default response
func (o *AccountsGetAccountsAccountIDTransactionsDefault) Code() int {
	return o._statusCode
}

func (o *AccountsGetAccountsAccountIDTransactionsDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/transactions][%d] accountsGetAccountsAccountIdTransactions default  %+v", o._statusCode, o.Payload)
}
func (o *AccountsGetAccountsAccountIDTransactionsDefault) GetPayload() *models.ResponseAccountTransactions {
	return o.Payload
}

func (o *AccountsGetAccountsAccountIDTransactionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.ResponseAccountTransactions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
