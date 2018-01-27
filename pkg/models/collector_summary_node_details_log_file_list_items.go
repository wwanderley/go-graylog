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

// CollectorSummaryNodeDetailsLogFileListItems collector summary node details log file list items
// swagger:model collectorSummaryNodeDetailsLogFileListItems
type CollectorSummaryNodeDetailsLogFileListItems struct {

	// is dir
	IsDir bool `json:"is_dir,omitempty"`

	// mod time
	ModTime strfmt.DateTime `json:"mod_time,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this collector summary node details log file list items
func (m *CollectorSummaryNodeDetailsLogFileListItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectorSummaryNodeDetailsLogFileListItems) validateModTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ModTime) { // not required
		return nil
	}

	if err := validate.FormatOf("mod_time", "body", "date-time", m.ModTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CollectorSummaryNodeDetailsLogFileListItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectorSummaryNodeDetailsLogFileListItems) UnmarshalBinary(b []byte) error {
	var res CollectorSummaryNodeDetailsLogFileListItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
