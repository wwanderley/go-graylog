// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TrafficHistogram traffic histogram
// swagger:model TrafficHistogram
type TrafficHistogram struct {

	// decoded
	Decoded interface{} `json:"decoded,omitempty"`

	// from
	From strfmt.DateTime `json:"from,omitempty"`

	// input
	Input interface{} `json:"input,omitempty"`

	// output
	Output interface{} `json:"output,omitempty"`

	// to
	To strfmt.DateTime `json:"to,omitempty"`
}

// Validate validates this traffic histogram
func (m *TrafficHistogram) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrafficHistogram) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if err := validate.FormatOf("from", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrafficHistogram) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	if err := validate.FormatOf("to", "body", "date-time", m.To.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrafficHistogram) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrafficHistogram) UnmarshalBinary(b []byte) error {
	var res TrafficHistogram
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}