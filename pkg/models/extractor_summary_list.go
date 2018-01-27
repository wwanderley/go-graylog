// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExtractorSummaryList extractor summary list
// swagger:model ExtractorSummaryList
type ExtractorSummaryList struct {

	// extractors
	Extractors ExtractorSummaryListExtractors `json:"extractors"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this extractor summary list
func (m *ExtractorSummaryList) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ExtractorSummaryList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtractorSummaryList) UnmarshalBinary(b []byte) error {
	var res ExtractorSummaryList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}