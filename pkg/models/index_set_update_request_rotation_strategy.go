// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IndexSetUpdateRequestRotationStrategy index set update request rotation strategy
// swagger:model indexSetUpdateRequestRotationStrategy
type IndexSetUpdateRequestRotationStrategy struct {

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this index set update request rotation strategy
func (m *IndexSetUpdateRequestRotationStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IndexSetUpdateRequestRotationStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndexSetUpdateRequestRotationStrategy) UnmarshalBinary(b []byte) error {
	var res IndexSetUpdateRequestRotationStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}