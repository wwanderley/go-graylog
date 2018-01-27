// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CachesPageCachesItemsConfig caches page caches items config
// swagger:model cachesPageCachesItemsConfig
type CachesPageCachesItemsConfig struct {

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this caches page caches items config
func (m *CachesPageCachesItemsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CachesPageCachesItemsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CachesPageCachesItemsConfig) UnmarshalBinary(b []byte) error {
	var res CachesPageCachesItemsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}