// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenshiftVersion openshift version
//
// swagger:model openshift-version
type OpenshiftVersion struct {

	// Name of the version to be presented to the user.
	DisplayName string `json:"display_name,omitempty"`

	// The installation image of the OpenShift cluster.
	ReleaseImage string `json:"release_image,omitempty"`

	// The base RHCOS image used for the discovery iso.
	RhcosImage string `json:"rhcos_image,omitempty"`

	// Level of support of the version.
	// Enum: [beta production]
	SupportLevel string `json:"support_level,omitempty"`
}

// Validate validates this openshift version
func (m *OpenshiftVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSupportLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var openshiftVersionTypeSupportLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["beta","production"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openshiftVersionTypeSupportLevelPropEnum = append(openshiftVersionTypeSupportLevelPropEnum, v)
	}
}

const (

	// OpenshiftVersionSupportLevelBeta captures enum value "beta"
	OpenshiftVersionSupportLevelBeta string = "beta"

	// OpenshiftVersionSupportLevelProduction captures enum value "production"
	OpenshiftVersionSupportLevelProduction string = "production"
)

// prop value enum
func (m *OpenshiftVersion) validateSupportLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openshiftVersionTypeSupportLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenshiftVersion) validateSupportLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateSupportLevelEnum("support_level", "body", m.SupportLevel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenshiftVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenshiftVersion) UnmarshalBinary(b []byte) error {
	var res OpenshiftVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}