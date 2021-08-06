// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OpenbankingBrasilStatus Status
//
// Retorna o estado do consentimento, o qual no momento de sua criao ser AWAITING_AUTHORISATION.
// Este estado ser alterado depois da autorizao do consentimento na detentora da conta do pagador (Debtor) para AUTHORISED ou REJECTED.
// O consentimento fica no estado CONSUMED aps ocorrer a iniciao do pagamento referente ao consentimento.
// Em caso de consentimento expirado a detentora dever retornar o status REJECTED.
// Estados possveis:
// AWAITING_AUTHORISATION - Aguardando autorizao
// AUTHORISED - Autorizado
// REJECTED - Rejeitado
// CONSUMED - Consumido
//
// swagger:model OpenbankingBrasilStatus
type OpenbankingBrasilStatus string

const (

	// OpenbankingBrasilStatusAWAITINGAUTHORISATION captures enum value "AWAITING_AUTHORISATION"
	OpenbankingBrasilStatusAWAITINGAUTHORISATION OpenbankingBrasilStatus = "AWAITING_AUTHORISATION"

	// OpenbankingBrasilStatusAUTHORISED captures enum value "AUTHORISED"
	OpenbankingBrasilStatusAUTHORISED OpenbankingBrasilStatus = "AUTHORISED"

	// OpenbankingBrasilStatusREJECTED captures enum value "REJECTED"
	OpenbankingBrasilStatusREJECTED OpenbankingBrasilStatus = "REJECTED"

	// OpenbankingBrasilStatusCONSUMED captures enum value "CONSUMED"
	OpenbankingBrasilStatusCONSUMED OpenbankingBrasilStatus = "CONSUMED"
)

// for schema
var openbankingBrasilStatusEnum []interface{}

func init() {
	var res []OpenbankingBrasilStatus
	if err := json.Unmarshal([]byte(`["AWAITING_AUTHORISATION","AUTHORISED","REJECTED","CONSUMED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingBrasilStatusEnum = append(openbankingBrasilStatusEnum, v)
	}
}

func (m OpenbankingBrasilStatus) validateOpenbankingBrasilStatusEnum(path, location string, value OpenbankingBrasilStatus) error {
	if err := validate.EnumCase(path, location, value, openbankingBrasilStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this openbanking brasil status
func (m OpenbankingBrasilStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenbankingBrasilStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this openbanking brasil status based on context it is used
func (m OpenbankingBrasilStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}