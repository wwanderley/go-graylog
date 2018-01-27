// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExposedConfiguration exposed configuration
// swagger:model ExposedConfiguration
type ExposedConfiguration struct {

	// allow highlighting
	AllowHighlighting bool `json:"allow_highlighting,omitempty"`

	// allow leading wildcard searches
	AllowLeadingWildcardSearches bool `json:"allow_leading_wildcard_searches,omitempty"`

	// gc warning threshold
	GcWarningThreshold string `json:"gc_warning_threshold,omitempty"`

	// inputbuffer processors
	InputbufferProcessors int64 `json:"inputbuffer_processors,omitempty"`

	// inputbuffer ring size
	InputbufferRingSize int64 `json:"inputbuffer_ring_size,omitempty"`

	// inputbuffer wait strategy
	InputbufferWaitStrategy string `json:"inputbuffer_wait_strategy,omitempty"`

	// node id file
	NodeIDFile string `json:"node_id_file,omitempty"`

	// output module timeout
	OutputModuleTimeout int64 `json:"output_module_timeout,omitempty"`

	// outputbuffer processors
	OutputbufferProcessors int64 `json:"outputbuffer_processors,omitempty"`

	// plugin dir
	PluginDir string `json:"plugin_dir,omitempty"`

	// processbuffer processors
	ProcessbufferProcessors int64 `json:"processbuffer_processors,omitempty"`

	// processor wait strategy
	ProcessorWaitStrategy string `json:"processor_wait_strategy,omitempty"`

	// ring size
	RingSize int64 `json:"ring_size,omitempty"`

	// stale master timeout
	StaleMasterTimeout int64 `json:"stale_master_timeout,omitempty"`

	// stream processing max faults
	StreamProcessingMaxFaults int64 `json:"stream_processing_max_faults,omitempty"`

	// stream processing timeout
	StreamProcessingTimeout int64 `json:"stream_processing_timeout,omitempty"`
}

// Validate validates this exposed configuration
func (m *ExposedConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ExposedConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExposedConfiguration) UnmarshalBinary(b []byte) error {
	var res ExposedConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}