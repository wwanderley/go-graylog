// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterStatsMongoDatabaseStats cluster stats mongo database stats
// swagger:model clusterStatsMongoDatabaseStats
type ClusterStatsMongoDatabaseStats struct {

	// avg obj size
	AvgObjSize float64 `json:"avg_obj_size,omitempty"`

	// collections
	Collections int64 `json:"collections,omitempty"`

	// data file version
	DataFileVersion *ClusterStatsMongoDatabaseStatsDataFileVersion `json:"data_file_version,omitempty"`

	// data size
	DataSize int64 `json:"data_size,omitempty"`

	// db
	Db string `json:"db,omitempty"`

	// extent free list
	ExtentFreeList *ClusterStatsMongoDatabaseStatsExtentFreeList `json:"extent_free_list,omitempty"`

	// file size
	FileSize int64 `json:"file_size,omitempty"`

	// index size
	IndexSize int64 `json:"index_size,omitempty"`

	// indexes
	Indexes int64 `json:"indexes,omitempty"`

	// ns size mb
	NsSizeMb int64 `json:"ns_size_mb,omitempty"`

	// num extents
	NumExtents int64 `json:"num_extents,omitempty"`

	// objects
	Objects int64 `json:"objects,omitempty"`

	// storage size
	StorageSize int64 `json:"storage_size,omitempty"`
}

// Validate validates this cluster stats mongo database stats
func (m *ClusterStatsMongoDatabaseStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataFileVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExtentFreeList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterStatsMongoDatabaseStats) validateDataFileVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.DataFileVersion) { // not required
		return nil
	}

	if m.DataFileVersion != nil {

		if err := m.DataFileVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data_file_version")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterStatsMongoDatabaseStats) validateExtentFreeList(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtentFreeList) { // not required
		return nil
	}

	if m.ExtentFreeList != nil {

		if err := m.ExtentFreeList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extent_free_list")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatsMongoDatabaseStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatsMongoDatabaseStats) UnmarshalBinary(b []byte) error {
	var res ClusterStatsMongoDatabaseStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
