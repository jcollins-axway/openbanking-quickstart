// Code generated by go-swagger; DO NOT EDIT.

package file_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
)

// CreateFilePaymentsReader is a Reader for the CreateFilePayments structure.
type CreateFilePaymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFilePaymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateFilePaymentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFilePaymentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateFilePaymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateFilePaymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFilePaymentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateFilePaymentsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateFilePaymentsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateFilePaymentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateFilePaymentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateFilePaymentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFilePaymentsCreated creates a CreateFilePaymentsCreated with default headers values
func NewCreateFilePaymentsCreated() *CreateFilePaymentsCreated {
	return &CreateFilePaymentsCreated{}
}

/* CreateFilePaymentsCreated describes a response with status code 201, with default header values.

File Payments Created
*/
type CreateFilePaymentsCreated struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteFileResponse3
}

func (o *CreateFilePaymentsCreated) Error() string {
	return fmt.Sprintf("[POST /file-payments][%d] createFilePaymentsCreated  %+v", 201, o.Payload)
}
func (o *CreateFilePaymentsCreated) GetPayload() *models.OBWriteFileResponse3 {
	return o.Payload
}

func (o *CreateFilePaymentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBWriteFileResponse3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFilePaymentsBadRequest creates a CreateFilePaymentsBadRequest with default headers values
func NewCreateFilePaymentsBadRequest() *CreateFilePaymentsBadRequest {
	return &CreateFilePaymentsBadRequest{}
}

/* CreateFilePaymentsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateFilePaymentsBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateFilePaymentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /file-payments][%d] createFilePaymentsBadRequest  %+v", 400, o.Payload)
}
func (o *CreateFilePaymentsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateFilePaymentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFilePaymentsUnauthorized creates a CreateFilePaymentsUnauthorized with default headers values
func NewCreateFilePaymentsUnauthorized() *CreateFilePaymentsUnauthorized {
	return &CreateFilePaymentsUnauthorized{}
}

/* CreateFilePaymentsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateFilePaymentsUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /file-payments][%d] createFilePaymentsUnauthorized ", 401)
}

func (o *CreateFilePaymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentsForbidden creates a CreateFilePaymentsForbidden with default headers values
func NewCreateFilePaymentsForbidden() *CreateFilePaymentsForbidden {
	return &CreateFilePaymentsForbidden{}
}

/* CreateFilePaymentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateFilePaymentsForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateFilePaymentsForbidden) Error() string {
	return fmt.Sprintf("[POST /file-payments][%d] createFilePaymentsForbidden  %+v", 403, o.Payload)
}
func (o *CreateFilePaymentsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateFilePaymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFilePaymentsNotFound creates a CreateFilePaymentsNotFound with default headers values
func NewCreateFilePaymentsNotFound() *CreateFilePaymentsNotFound {
	return &CreateFilePaymentsNotFound{}
}

/* CreateFilePaymentsNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateFilePaymentsNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentsNotFound) Error() string {
	return fmt.Sprintf("[POST /file-payments][%d] createFilePaymentsNotFound ", 404)
}

func (o *CreateFilePaymentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentsMethodNotAllowed creates a CreateFilePaymentsMethodNotAllowed with default headers values
func NewCreateFilePaymentsMethodNotAllowed() *CreateFilePaymentsMethodNotAllowed {
	return &CreateFilePaymentsMethodNotAllowed{}
}

/* CreateFilePaymentsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type CreateFilePaymentsMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /file-payments][%d] createFilePaymentsMethodNotAllowed ", 405)
}

func (o *CreateFilePaymentsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentsNotAcceptable creates a CreateFilePaymentsNotAcceptable with default headers values
func NewCreateFilePaymentsNotAcceptable() *CreateFilePaymentsNotAcceptable {
	return &CreateFilePaymentsNotAcceptable{}
}

/* CreateFilePaymentsNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type CreateFilePaymentsNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentsNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /file-payments][%d] createFilePaymentsNotAcceptable ", 406)
}

func (o *CreateFilePaymentsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentsUnsupportedMediaType creates a CreateFilePaymentsUnsupportedMediaType with default headers values
func NewCreateFilePaymentsUnsupportedMediaType() *CreateFilePaymentsUnsupportedMediaType {
	return &CreateFilePaymentsUnsupportedMediaType{}
}

/* CreateFilePaymentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type
*/
type CreateFilePaymentsUnsupportedMediaType struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /file-payments][%d] createFilePaymentsUnsupportedMediaType ", 415)
}

func (o *CreateFilePaymentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentsTooManyRequests creates a CreateFilePaymentsTooManyRequests with default headers values
func NewCreateFilePaymentsTooManyRequests() *CreateFilePaymentsTooManyRequests {
	return &CreateFilePaymentsTooManyRequests{}
}

/* CreateFilePaymentsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateFilePaymentsTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /file-payments][%d] createFilePaymentsTooManyRequests ", 429)
}

func (o *CreateFilePaymentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Retry-After
	hdrRetryAfter := response.GetHeader("Retry-After")

	if hdrRetryAfter != "" {
		valretryAfter, err := swag.ConvertInt64(hdrRetryAfter)
		if err != nil {
			return errors.InvalidType("Retry-After", "header", "int64", hdrRetryAfter)
		}
		o.RetryAfter = valretryAfter
	}

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentsInternalServerError creates a CreateFilePaymentsInternalServerError with default headers values
func NewCreateFilePaymentsInternalServerError() *CreateFilePaymentsInternalServerError {
	return &CreateFilePaymentsInternalServerError{}
}

/* CreateFilePaymentsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateFilePaymentsInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateFilePaymentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /file-payments][%d] createFilePaymentsInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateFilePaymentsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateFilePaymentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}