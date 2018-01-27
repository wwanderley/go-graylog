// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CollectorListCollectorsItemsNodeDetailsStatus collector list collectors items node details status
// swagger:model collectorListCollectorsItemsNodeDetailsStatus
type CollectorListCollectorsItemsNodeDetailsStatus struct {

	// backends
	Backends interface{} `json:"backends,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`
}

// Validate validates this collector list collectors items node details status
func (m *CollectorListCollectorsItemsNodeDetailsStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CollectorListCollectorsItemsNodeDetailsStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectorListCollectorsItemsNodeDetailsStatus) UnmarshalBinary(b []byte) error {
	var res CollectorListCollectorsItemsNodeDetailsStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
