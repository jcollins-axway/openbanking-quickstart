// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatusUpdateDateTime Date and time at which the resource status was updated.All dates in the JSON payloads are represented in ISO 8601 date-time format.
// All date-time fields in responses must include the timezone. An example is below:
// 2017-04-05T10:43:07+00:00
//
// swagger:model StatusUpdateDateTime
type StatusUpdateDateTime strfmt.DateTime

// UnmarshalJSON sets a StatusUpdateDateTime value from JSON input
func (m *StatusUpdateDateTime) UnmarshalJSON(b []byte) error {
	return ((*strfmt.DateTime)(m)).UnmarshalJSON(b)
}

// MarshalJSON retrieves a StatusUpdateDateTime value as JSON output
func (m StatusUpdateDateTime) MarshalJSON() ([]byte, error) {
	return (strfmt.DateTime(m)).MarshalJSON()
}

// Validate validates this status update date time
func (m StatusUpdateDateTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.FormatOf("", "body", "date-time", strfmt.DateTime(m).String(), formats); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this status update date time based on context it is used
func (m StatusUpdateDateTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatusUpdateDateTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusUpdateDateTime) UnmarshalBinary(b []byte) error {
	var res StatusUpdateDateTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}