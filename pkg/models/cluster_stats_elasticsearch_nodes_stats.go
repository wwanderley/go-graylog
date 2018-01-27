// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterStatsElasticsearchNodesStats cluster stats elasticsearch nodes stats
// swagger:model clusterStatsElasticsearchNodesStats
type ClusterStatsElasticsearchNodesStats struct {

	// client
	Client int64 `json:"client,omitempty"`

	// data only
	DataOnly int64 `json:"data_only,omitempty"`

	// master data
	MasterData int64 `json:"master_data,omitempty"`

	// master only
	MasterOnly int64 `json:"master_only,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this cluster stats elasticsearch nodes stats
func (m *ClusterStatsElasticsearchNodesStats) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatsElasticsearchNodesStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatsElasticsearchNodesStats) UnmarshalBinary(b []byte) error {
	var res ClusterStatsElasticsearchNodesStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
