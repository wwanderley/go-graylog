// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateExtractorRequest create extractor request
// swagger:model CreateExtractorRequest
type CreateExtractorRequest struct {

	// condition type
	ConditionType string `json:"condition_type,omitempty"`

	// condition value
	ConditionValue string `json:"condition_value,omitempty"`

	// converters
	Converters interface{} `json:"converters,omitempty"`

	// cut or copy
	CutOrCopy string `json:"cut_or_copy,omitempty"`

	// extractor config
	ExtractorConfig interface{} `json:"extractor_config,omitempty"`

	// extractor type
	ExtractorType string `json:"extractor_type,omitempty"`

	// order
	Order int64 `json:"order,omitempty"`

	// source field
	SourceField string `json:"source_field,omitempty"`

	// target field
	TargetField string `json:"target_field,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this create extractor request
func (m *CreateExtractorRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CreateExtractorRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateExtractorRequest) UnmarshalBinary(b []byte) error {
	var res CreateExtractorRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}