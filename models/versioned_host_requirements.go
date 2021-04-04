// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VersionedHostRequirements versioned host requirements
//
// swagger:model versioned-host-requirements
type VersionedHostRequirements struct {

	// Master node requirements
	MasterRequirements *ClusterHostRequirementsDetails `json:"master,omitempty"`

	// Worker node requirements
	WorkerRequirements *ClusterHostRequirementsDetails `json:"worker,omitempty"`

	// Version of the component for which requirements are defined
	Version string `json:"version,omitempty"`
}

// Validate validates this versioned host requirements
func (m *VersionedHostRequirements) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMasterRequirements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkerRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionedHostRequirements) validateMasterRequirements(formats strfmt.Registry) error {

	if swag.IsZero(m.MasterRequirements) { // not required
		return nil
	}

	if m.MasterRequirements != nil {
		if err := m.MasterRequirements.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master")
			}
			return err
		}
	}

	return nil
}

func (m *VersionedHostRequirements) validateWorkerRequirements(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkerRequirements) { // not required
		return nil
	}

	if m.WorkerRequirements != nil {
		if err := m.WorkerRequirements.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("worker")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionedHostRequirements) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedHostRequirements) UnmarshalBinary(b []byte) error {
	var res VersionedHostRequirements
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}