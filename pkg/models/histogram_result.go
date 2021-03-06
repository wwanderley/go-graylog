// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HistogramResult histogram result
// swagger:model HistogramResult
type HistogramResult struct {

	// built query
	BuiltQuery string `json:"built_query,omitempty"`

	// interval
	Interval string `json:"interval,omitempty"`

	// queried timerange
	QueriedTimerange *HistogramResultQueriedTimerange `json:"queried_timerange,omitempty"`

	// results
	Results interface{} `json:"results,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`
}

// Validate validates this histogram result
func (m *HistogramResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueriedTimerange(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistogramResult) validateQueriedTimerange(formats strfmt.Registry) error {

	if swag.IsZero(m.QueriedTimerange) { // not required
		return nil
	}

	if m.QueriedTimerange != nil {

		if err := m.QueriedTimerange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queried_timerange")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistogramResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistogramResult) UnmarshalBinary(b []byte) error {
	var res HistogramResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
