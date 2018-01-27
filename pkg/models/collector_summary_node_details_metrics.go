// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CollectorSummaryNodeDetailsMetrics collector summary node details metrics
// swagger:model collectorSummaryNodeDetailsMetrics
type CollectorSummaryNodeDetailsMetrics struct {

	// cpu idle
	CPUIDLE float64 `json:"cpu_idle,omitempty"`

	// disks 75
	Disks75 []string `json:"disks_75"`

	// load 1
	Load1 float64 `json:"load_1,omitempty"`
}

// Validate validates this collector summary node details metrics
func (m *CollectorSummaryNodeDetailsMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisks75(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectorSummaryNodeDetailsMetrics) validateDisks75(formats strfmt.Registry) error {

	if swag.IsZero(m.Disks75) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CollectorSummaryNodeDetailsMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectorSummaryNodeDetailsMetrics) UnmarshalBinary(b []byte) error {
	var res CollectorSummaryNodeDetailsMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
