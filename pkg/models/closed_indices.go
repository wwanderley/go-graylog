// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClosedIndices closed indices
// swagger:model ClosedIndices
type ClosedIndices struct {

	// indices
	Indices []string `json:"indices"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this closed indices
func (m *ClosedIndices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClosedIndices) validateIndices(formats strfmt.Registry) error {

	if swag.IsZero(m.Indices) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClosedIndices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClosedIndices) UnmarshalBinary(b []byte) error {
	var res ClosedIndices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
