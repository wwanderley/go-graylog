// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RotationStrategyDescriptionDefaultConfig rotation strategy description default config
// swagger:model rotationStrategyDescriptionDefaultConfig
type RotationStrategyDescriptionDefaultConfig struct {

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this rotation strategy description default config
func (m *RotationStrategyDescriptionDefaultConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RotationStrategyDescriptionDefaultConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RotationStrategyDescriptionDefaultConfig) UnmarshalBinary(b []byte) error {
	var res RotationStrategyDescriptionDefaultConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
