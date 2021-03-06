// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MongoStatsServerStatusNetwork mongo stats server status network
// swagger:model mongoStatsServerStatusNetwork
type MongoStatsServerStatusNetwork struct {

	// bytes in
	BytesIn int64 `json:"bytes_in,omitempty"`

	// bytes out
	BytesOut int64 `json:"bytes_out,omitempty"`

	// num requests
	NumRequests int64 `json:"num_requests,omitempty"`
}

// Validate validates this mongo stats server status network
func (m *MongoStatsServerStatusNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MongoStatsServerStatusNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoStatsServerStatusNetwork) UnmarshalBinary(b []byte) error {
	var res MongoStatsServerStatusNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
