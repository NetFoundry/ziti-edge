// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IdentityCreate An identity to create
//
// swagger:model identityCreate
type IdentityCreate struct {

	// enrollment
	Enrollment *IdentityCreateEnrollment `json:"enrollment,omitempty"`

	// is admin
	IsAdmin bool `json:"isAdmin,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// tags
	Tags Tags `json:"tags,omitempty"`

	// type
	// Enum: [User Device Service]
	Type string `json:"type,omitempty"`
}

// Validate validates this identity create
func (m *IdentityCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnrollment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityCreate) validateEnrollment(formats strfmt.Registry) error {

	if swag.IsZero(m.Enrollment) { // not required
		return nil
	}

	if m.Enrollment != nil {
		if err := m.Enrollment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enrollment")
			}
			return err
		}
	}

	return nil
}

func (m *IdentityCreate) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := m.Tags.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tags")
		}
		return err
	}

	return nil
}

var identityCreateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["User","Device","Service"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		identityCreateTypeTypePropEnum = append(identityCreateTypeTypePropEnum, v)
	}
}

const (

	// IdentityCreateTypeUser captures enum value "User"
	IdentityCreateTypeUser string = "User"

	// IdentityCreateTypeDevice captures enum value "Device"
	IdentityCreateTypeDevice string = "Device"

	// IdentityCreateTypeService captures enum value "Service"
	IdentityCreateTypeService string = "Service"
)

// prop value enum
func (m *IdentityCreate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, identityCreateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IdentityCreate) validateType(formats strfmt.Registry) error {

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
func (m *IdentityCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityCreate) UnmarshalBinary(b []byte) error {
	var res IdentityCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IdentityCreateEnrollment identity create enrollment
//
// swagger:model IdentityCreateEnrollment
type IdentityCreateEnrollment struct {

	// ott
	Ott bool `json:"ott,omitempty"`

	// ottca
	// Format: uuid
	Ottca strfmt.UUID `json:"ottca,omitempty"`

	// updb
	Updb string `json:"updb,omitempty"`
}

// Validate validates this identity create enrollment
func (m *IdentityCreateEnrollment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOttca(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityCreateEnrollment) validateOttca(formats strfmt.Registry) error {

	if swag.IsZero(m.Ottca) { // not required
		return nil
	}

	if err := validate.FormatOf("enrollment"+"."+"ottca", "body", "uuid", m.Ottca.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdentityCreateEnrollment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityCreateEnrollment) UnmarshalBinary(b []byte) error {
	var res IdentityCreateEnrollment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
