// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IndexSetResponse index set response
// swagger:model IndexSetResponse
type IndexSetResponse struct {

	// index sets
	IndexSets IndexSetResponseIndexSets `json:"index_sets"`

	// stats
	Stats interface{} `json:"stats,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this index set response
func (m *IndexSetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IndexSetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndexSetResponse) UnmarshalBinary(b []byte) error {
	var res IndexSetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
