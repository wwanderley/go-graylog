// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AddOutputRequest add output request
// swagger:model AddOutputRequest
type AddOutputRequest struct {

	// outputs
	Outputs []string `json:"outputs"`
}

// Validate validates this add output request
func (m *AddOutputRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutputs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddOutputRequest) validateOutputs(formats strfmt.Registry) error {

	if swag.IsZero(m.Outputs) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddOutputRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddOutputRequest) UnmarshalBinary(b []byte) error {
	var res AddOutputRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
