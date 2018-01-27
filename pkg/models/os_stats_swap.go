// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OsStatsSwap os stats swap
// swagger:model osStatsSwap
type OsStatsSwap struct {

	// free
	Free int64 `json:"free,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`

	// used
	Used int64 `json:"used,omitempty"`
}

// Validate validates this os stats swap
func (m *OsStatsSwap) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OsStatsSwap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsStatsSwap) UnmarshalBinary(b []byte) error {
	var res OsStatsSwap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
