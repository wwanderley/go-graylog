// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PipelineConnections pipeline connections
// swagger:model PipelineConnections
type PipelineConnections struct {

	// id
	ID string `json:"id,omitempty"`

	// pipeline ids
	PipelineIds []string `json:"pipeline_ids"`

	// stream id
	StreamID string `json:"stream_id,omitempty"`
}

// Validate validates this pipeline connections
func (m *PipelineConnections) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePipelineIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PipelineConnections) validatePipelineIds(formats strfmt.Registry) error {

	if swag.IsZero(m.PipelineIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PipelineConnections) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineConnections) UnmarshalBinary(b []byte) error {
	var res PipelineConnections
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
