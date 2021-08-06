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

// OBOtherCodeType15 Other fee rate type which is not in the standard rate type list
//
// swagger:model OB_OtherCodeType1_5
type OBOtherCodeType15 struct {

	// code
	Code OBCodeMnemonic `json:"Code,omitempty"`

	// description
	// Required: true
	Description *Description3 `json:"Description"`

	// name
	// Required: true
	Name *Name4 `json:"Name"`
}

// Validate validates this o b other code type1 5
func (m *OBOtherCodeType15) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBOtherCodeType15) validateCode(formats strfmt.Registry) error {
	if swag.IsZero(m.Code) { // not required
		return nil
	}

	if err := m.Code.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Code")
		}
		return err
	}

	return nil
}

func (m *OBOtherCodeType15) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	if m.Description != nil {
		if err := m.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Description")
			}
			return err
		}
	}

	return nil
}

func (m *OBOtherCodeType15) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Name")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this o b other code type1 5 based on the context it is used
func (m *OBOtherCodeType15) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBOtherCodeType15) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Code.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Code")
		}
		return err
	}

	return nil
}

func (m *OBOtherCodeType15) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if m.Description != nil {
		if err := m.Description.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Description")
			}
			return err
		}
	}

	return nil
}

func (m *OBOtherCodeType15) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if m.Name != nil {
		if err := m.Name.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Name")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBOtherCodeType15) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBOtherCodeType15) UnmarshalBinary(b []byte) error {
	var res OBOtherCodeType15
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}