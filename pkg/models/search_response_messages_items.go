// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SearchResponseMessagesItems search response messages items
// swagger:model searchResponseMessagesItems
type SearchResponseMessagesItems struct {

	// decoration stats
	DecorationStats *SearchResponseMessagesItemsDecorationStats `json:"decoration_stats,omitempty"`

	// highlight ranges
	HighlightRanges interface{} `json:"highlight_ranges,omitempty"`

	// index
	Index string `json:"index,omitempty"`

	// message
	Message interface{} `json:"message,omitempty"`
}

// Validate validates this search response messages items
func (m *SearchResponseMessagesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecorationStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchResponseMessagesItems) validateDecorationStats(formats strfmt.Registry) error {

	if swag.IsZero(m.DecorationStats) { // not required
		return nil
	}

	if m.DecorationStats != nil {

		if err := m.DecorationStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decoration_stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchResponseMessagesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchResponseMessagesItems) UnmarshalBinary(b []byte) error {
	var res SearchResponseMessagesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
