// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigurationBundleStreamsItemsStreamRulesItems configuration bundle streams items stream rules items
// swagger:model configurationBundleStreamsItemsStreamRulesItems
type ConfigurationBundleStreamsItemsStreamRulesItems struct {

	// description
	Description string `json:"description,omitempty"`

	// field
	Field string `json:"field,omitempty"`

	// inverted
	Inverted bool `json:"inverted,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this configuration bundle streams items stream rules items
func (m *ConfigurationBundleStreamsItemsStreamRulesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var configurationBundleStreamsItemsStreamRulesItemsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXACT","REGEX","GREATER","SMALLER","PRESENCE","CONTAINS","ALWAYS_MATCH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configurationBundleStreamsItemsStreamRulesItemsTypeTypePropEnum = append(configurationBundleStreamsItemsStreamRulesItemsTypeTypePropEnum, v)
	}
}

const (
	// ConfigurationBundleStreamsItemsStreamRulesItemsTypeEXACT captures enum value "EXACT"
	ConfigurationBundleStreamsItemsStreamRulesItemsTypeEXACT string = "EXACT"
	// ConfigurationBundleStreamsItemsStreamRulesItemsTypeREGEX captures enum value "REGEX"
	ConfigurationBundleStreamsItemsStreamRulesItemsTypeREGEX string = "REGEX"
	// ConfigurationBundleStreamsItemsStreamRulesItemsTypeGREATER captures enum value "GREATER"
	ConfigurationBundleStreamsItemsStreamRulesItemsTypeGREATER string = "GREATER"
	// ConfigurationBundleStreamsItemsStreamRulesItemsTypeSMALLER captures enum value "SMALLER"
	ConfigurationBundleStreamsItemsStreamRulesItemsTypeSMALLER string = "SMALLER"
	// ConfigurationBundleStreamsItemsStreamRulesItemsTypePRESENCE captures enum value "PRESENCE"
	ConfigurationBundleStreamsItemsStreamRulesItemsTypePRESENCE string = "PRESENCE"
	// ConfigurationBundleStreamsItemsStreamRulesItemsTypeCONTAINS captures enum value "CONTAINS"
	ConfigurationBundleStreamsItemsStreamRulesItemsTypeCONTAINS string = "CONTAINS"
	// ConfigurationBundleStreamsItemsStreamRulesItemsTypeALWAYSMATCH captures enum value "ALWAYS_MATCH"
	ConfigurationBundleStreamsItemsStreamRulesItemsTypeALWAYSMATCH string = "ALWAYS_MATCH"
)

// prop value enum
func (m *ConfigurationBundleStreamsItemsStreamRulesItems) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, configurationBundleStreamsItemsStreamRulesItemsTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ConfigurationBundleStreamsItemsStreamRulesItems) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigurationBundleStreamsItemsStreamRulesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigurationBundleStreamsItemsStreamRulesItems) UnmarshalBinary(b []byte) error {
	var res ConfigurationBundleStreamsItemsStreamRulesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
