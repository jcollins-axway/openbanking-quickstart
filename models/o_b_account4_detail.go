// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OBAccount4Detail Unambiguous identification of the account to which credit and debit entries are made.
//
// swagger:model OBAccount4Detail
type OBAccount4Detail struct {

	// account
	// Required: true
	Account []*OBAccount4DetailAccountItems0 `json:"Account"`

	// account Id
	// Required: true
	AccountID *AccountID `json:"AccountId"`

	// account sub type
	// Required: true
	AccountSubType *OBExternalAccountSubType1Code `json:"AccountSubType"`

	// account type
	// Required: true
	AccountType *OBExternalAccountType1Code `json:"AccountType"`

	// currency
	// Required: true
	Currency *ActiveOrHistoricCurrencyCode0 `json:"Currency"`

	// description
	Description Description0 `json:"Description,omitempty"`

	// nickname
	Nickname Nickname `json:"Nickname,omitempty"`

	// servicer
	Servicer *OBBranchAndFinancialInstitutionIdentification50 `json:"Servicer,omitempty"`

	// status
	Status OBAccountStatus1Code `json:"Status,omitempty"`

	// status update date time
	// Format: date-time
	StatusUpdateDateTime StatusUpdateDateTime `json:"StatusUpdateDateTime,omitempty"`
}

// Validate validates this o b account4 detail
func (m *OBAccount4Detail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountSubType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNickname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusUpdateDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBAccount4Detail) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("Account", "body", m.Account); err != nil {
		return err
	}

	for i := 0; i < len(m.Account); i++ {
		if swag.IsZero(m.Account[i]) { // not required
			continue
		}

		if m.Account[i] != nil {
			if err := m.Account[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Account" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OBAccount4Detail) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("AccountId", "body", m.AccountID); err != nil {
		return err
	}

	if err := validate.Required("AccountId", "body", m.AccountID); err != nil {
		return err
	}

	if m.AccountID != nil {
		if err := m.AccountID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AccountId")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4Detail) validateAccountSubType(formats strfmt.Registry) error {

	if err := validate.Required("AccountSubType", "body", m.AccountSubType); err != nil {
		return err
	}

	if err := validate.Required("AccountSubType", "body", m.AccountSubType); err != nil {
		return err
	}

	if m.AccountSubType != nil {
		if err := m.AccountSubType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AccountSubType")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4Detail) validateAccountType(formats strfmt.Registry) error {

	if err := validate.Required("AccountType", "body", m.AccountType); err != nil {
		return err
	}

	if err := validate.Required("AccountType", "body", m.AccountType); err != nil {
		return err
	}

	if m.AccountType != nil {
		if err := m.AccountType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AccountType")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4Detail) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("Currency", "body", m.Currency); err != nil {
		return err
	}

	if err := validate.Required("Currency", "body", m.Currency); err != nil {
		return err
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Currency")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4Detail) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Description")
		}
		return err
	}

	return nil
}

func (m *OBAccount4Detail) validateNickname(formats strfmt.Registry) error {
	if swag.IsZero(m.Nickname) { // not required
		return nil
	}

	if err := m.Nickname.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Nickname")
		}
		return err
	}

	return nil
}

func (m *OBAccount4Detail) validateServicer(formats strfmt.Registry) error {
	if swag.IsZero(m.Servicer) { // not required
		return nil
	}

	if m.Servicer != nil {
		if err := m.Servicer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Servicer")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4Detail) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Status")
		}
		return err
	}

	return nil
}

func (m *OBAccount4Detail) validateStatusUpdateDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusUpdateDateTime) { // not required
		return nil
	}

	if err := m.StatusUpdateDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StatusUpdateDateTime")
		}
		return err
	}

	return nil
}

// ContextValidate validate this o b account4 detail based on the context it is used
func (m *OBAccount4Detail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAccountID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAccountSubType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAccountType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCurrency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNickname(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServicer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusUpdateDateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBAccount4Detail) contextValidateAccount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Account); i++ {

		if m.Account[i] != nil {
			if err := m.Account[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Account" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OBAccount4Detail) contextValidateAccountID(ctx context.Context, formats strfmt.Registry) error {

	if m.AccountID != nil {
		if err := m.AccountID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AccountId")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4Detail) contextValidateAccountSubType(ctx context.Context, formats strfmt.Registry) error {

	if m.AccountSubType != nil {
		if err := m.AccountSubType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AccountSubType")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4Detail) contextValidateAccountType(ctx context.Context, formats strfmt.Registry) error {

	if m.AccountType != nil {
		if err := m.AccountType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AccountType")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4Detail) contextValidateCurrency(ctx context.Context, formats strfmt.Registry) error {

	if m.Currency != nil {
		if err := m.Currency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Currency")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4Detail) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Description.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Description")
		}
		return err
	}

	return nil
}

func (m *OBAccount4Detail) contextValidateNickname(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Nickname.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Nickname")
		}
		return err
	}

	return nil
}

func (m *OBAccount4Detail) contextValidateServicer(ctx context.Context, formats strfmt.Registry) error {

	if m.Servicer != nil {
		if err := m.Servicer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Servicer")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4Detail) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Status")
		}
		return err
	}

	return nil
}

func (m *OBAccount4Detail) contextValidateStatusUpdateDateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StatusUpdateDateTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StatusUpdateDateTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBAccount4Detail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBAccount4Detail) UnmarshalBinary(b []byte) error {
	var res OBAccount4Detail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBAccount4DetailAccountItems0 Provides the details to identify an account.
//
// swagger:model OBAccount4DetailAccountItems0
type OBAccount4DetailAccountItems0 struct {

	// identification
	// Required: true
	Identification *Identification0 `json:"Identification"`

	// name
	Name Name0 `json:"Name,omitempty"`

	// scheme name
	// Required: true
	SchemeName *OBExternalAccountIdentification4Code `json:"SchemeName"`

	// secondary identification
	SecondaryIdentification SecondaryIdentification `json:"SecondaryIdentification,omitempty"`
}

// Validate validates this o b account4 detail account items0
func (m *OBAccount4DetailAccountItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBAccount4DetailAccountItems0) validateIdentification(formats strfmt.Registry) error {

	if err := validate.Required("Identification", "body", m.Identification); err != nil {
		return err
	}

	if err := validate.Required("Identification", "body", m.Identification); err != nil {
		return err
	}

	if m.Identification != nil {
		if err := m.Identification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Identification")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4DetailAccountItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Name")
		}
		return err
	}

	return nil
}

func (m *OBAccount4DetailAccountItems0) validateSchemeName(formats strfmt.Registry) error {

	if err := validate.Required("SchemeName", "body", m.SchemeName); err != nil {
		return err
	}

	if err := validate.Required("SchemeName", "body", m.SchemeName); err != nil {
		return err
	}

	if m.SchemeName != nil {
		if err := m.SchemeName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SchemeName")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4DetailAccountItems0) validateSecondaryIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryIdentification) { // not required
		return nil
	}

	if err := m.SecondaryIdentification.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SecondaryIdentification")
		}
		return err
	}

	return nil
}

// ContextValidate validate this o b account4 detail account items0 based on the context it is used
func (m *OBAccount4DetailAccountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIdentification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchemeName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondaryIdentification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBAccount4DetailAccountItems0) contextValidateIdentification(ctx context.Context, formats strfmt.Registry) error {

	if m.Identification != nil {
		if err := m.Identification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Identification")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4DetailAccountItems0) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Name.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Name")
		}
		return err
	}

	return nil
}

func (m *OBAccount4DetailAccountItems0) contextValidateSchemeName(ctx context.Context, formats strfmt.Registry) error {

	if m.SchemeName != nil {
		if err := m.SchemeName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SchemeName")
			}
			return err
		}
	}

	return nil
}

func (m *OBAccount4DetailAccountItems0) contextValidateSecondaryIdentification(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SecondaryIdentification.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SecondaryIdentification")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBAccount4DetailAccountItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBAccount4DetailAccountItems0) UnmarshalBinary(b []byte) error {
	var res OBAccount4DetailAccountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}