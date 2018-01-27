// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RolesResponseRolesItems roles response roles items
// swagger:model rolesResponseRolesItems
type RolesResponseRolesItems struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// permissions
	Permissions []string `json:"permissions"`

	// read only
	ReadOnly bool `json:"read_only,omitempty"`
}

// Validate validates this roles response roles items
func (m *RolesResponseRolesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RolesResponseRolesItems) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RolesResponseRolesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RolesResponseRolesItems) UnmarshalBinary(b []byte) error {
	var res RolesResponseRolesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}