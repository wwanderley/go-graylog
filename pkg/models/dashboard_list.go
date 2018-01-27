// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DashboardList dashboard list
// swagger:model DashboardList
type DashboardList struct {

	// dashboards
	Dashboards []interface{} `json:"dashboards"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this dashboard list
func (m *DashboardList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboards(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardList) validateDashboards(formats strfmt.Registry) error {

	if swag.IsZero(m.Dashboards) { // not required
		return nil
	}

	for i := 0; i < len(m.Dashboards); i++ {

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardList) UnmarshalBinary(b []byte) error {
	var res DashboardList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
