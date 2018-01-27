// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterStatsMongoHostInfoOs cluster stats mongo host info os
// swagger:model clusterStatsMongoHostInfoOs
type ClusterStatsMongoHostInfoOs struct {

	// name
	Name string `json:"name,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this cluster stats mongo host info os
func (m *ClusterStatsMongoHostInfoOs) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatsMongoHostInfoOs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatsMongoHostInfoOs) UnmarshalBinary(b []byte) error {
	var res ClusterStatsMongoHostInfoOs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
