// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MongoStatsServerStatusStorageEngine mongo stats server status storage engine
// swagger:model mongoStatsServerStatusStorageEngine
type MongoStatsServerStatusStorageEngine struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this mongo stats server status storage engine
func (m *MongoStatsServerStatusStorageEngine) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MongoStatsServerStatusStorageEngine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoStatsServerStatusStorageEngine) UnmarshalBinary(b []byte) error {
	var res MongoStatsServerStatusStorageEngine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
