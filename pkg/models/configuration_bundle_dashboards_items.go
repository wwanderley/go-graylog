// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ConfigurationBundleDashboardsItems configuration bundle dashboards items
// swagger:model configurationBundleDashboardsItems
type ConfigurationBundleDashboardsItems struct {

	// dashboard widgets
	DashboardWidgets ConfigurationBundleDashboardsItemsDashboardWidgets `json:"dashboard_widgets"`

	// description
	Description string `json:"description,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this configuration bundle dashboards items
func (m *ConfigurationBundleDashboardsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigurationBundleDashboardsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigurationBundleDashboardsItems) UnmarshalBinary(b []byte) error {
	var res ConfigurationBundleDashboardsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
