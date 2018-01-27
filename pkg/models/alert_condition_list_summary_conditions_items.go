// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertConditionListSummaryConditionsItems alert condition list summary conditions items
// swagger:model alertConditionListSummaryConditionsItems
type AlertConditionListSummaryConditionsItems struct {

	// created at
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// creator user id
	CreatorUserID string `json:"creator_user_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// in grace
	InGrace bool `json:"in_grace,omitempty"`

	// parameters
	Parameters interface{} `json:"parameters,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this alert condition list summary conditions items
func (m *AlertConditionListSummaryConditionsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertConditionListSummaryConditionsItems) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertConditionListSummaryConditionsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertConditionListSummaryConditionsItems) UnmarshalBinary(b []byte) error {
	var res AlertConditionListSummaryConditionsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}