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
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RouterDetail router detail
//
// swagger:model routerDetail
type RouterDetail struct {
	BaseEntity

	// allow traversal
	// Required: true
	AllowTraversal *bool `json:"allowTraversal"`

	// cost
	// Required: true
	// Maximum: 65535
	// Minimum: 0
	Cost *int64 `json:"cost"`

	// enrollment created at
	// Format: date-time
	EnrollmentCreatedAt *strfmt.DateTime `json:"enrollmentCreatedAt,omitempty"`

	// enrollment expires at
	// Format: date-time
	EnrollmentExpiresAt *strfmt.DateTime `json:"enrollmentExpiresAt,omitempty"`

	// enrollment jwt
	EnrollmentJwt *string `json:"enrollmentJwt,omitempty"`

	// enrollment token
	EnrollmentToken *string `json:"enrollmentToken,omitempty"`

	// fingerprint
	// Required: true
	Fingerprint *string `json:"fingerprint"`

	// is online
	// Required: true
	IsOnline *bool `json:"isOnline"`

	// is verified
	// Required: true
	IsVerified *bool `json:"isVerified"`

	// name
	// Required: true
	Name *string `json:"name"`

	// unverified cert pem
	UnverifiedCertPem *string `json:"unverifiedCertPem"`

	// unverified fingerprint
	UnverifiedFingerprint *string `json:"unverifiedFingerprint"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RouterDetail) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseEntity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseEntity = aO0

	// AO1
	var dataAO1 struct {
		AllowTraversal *bool `json:"allowTraversal"`

		Cost *int64 `json:"cost"`

		EnrollmentCreatedAt *strfmt.DateTime `json:"enrollmentCreatedAt,omitempty"`

		EnrollmentExpiresAt *strfmt.DateTime `json:"enrollmentExpiresAt,omitempty"`

		EnrollmentJwt *string `json:"enrollmentJwt,omitempty"`

		EnrollmentToken *string `json:"enrollmentToken,omitempty"`

		Fingerprint *string `json:"fingerprint"`

		IsOnline *bool `json:"isOnline"`

		IsVerified *bool `json:"isVerified"`

		Name *string `json:"name"`

		UnverifiedCertPem *string `json:"unverifiedCertPem"`

		UnverifiedFingerprint *string `json:"unverifiedFingerprint"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AllowTraversal = dataAO1.AllowTraversal

	m.Cost = dataAO1.Cost

	m.EnrollmentCreatedAt = dataAO1.EnrollmentCreatedAt

	m.EnrollmentExpiresAt = dataAO1.EnrollmentExpiresAt

	m.EnrollmentJwt = dataAO1.EnrollmentJwt

	m.EnrollmentToken = dataAO1.EnrollmentToken

	m.Fingerprint = dataAO1.Fingerprint

	m.IsOnline = dataAO1.IsOnline

	m.IsVerified = dataAO1.IsVerified

	m.Name = dataAO1.Name

	m.UnverifiedCertPem = dataAO1.UnverifiedCertPem

	m.UnverifiedFingerprint = dataAO1.UnverifiedFingerprint

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RouterDetail) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AllowTraversal *bool `json:"allowTraversal"`

		Cost *int64 `json:"cost"`

		EnrollmentCreatedAt *strfmt.DateTime `json:"enrollmentCreatedAt,omitempty"`

		EnrollmentExpiresAt *strfmt.DateTime `json:"enrollmentExpiresAt,omitempty"`

		EnrollmentJwt *string `json:"enrollmentJwt,omitempty"`

		EnrollmentToken *string `json:"enrollmentToken,omitempty"`

		Fingerprint *string `json:"fingerprint"`

		IsOnline *bool `json:"isOnline"`

		IsVerified *bool `json:"isVerified"`

		Name *string `json:"name"`

		UnverifiedCertPem *string `json:"unverifiedCertPem"`

		UnverifiedFingerprint *string `json:"unverifiedFingerprint"`
	}

	dataAO1.AllowTraversal = m.AllowTraversal

	dataAO1.Cost = m.Cost

	dataAO1.EnrollmentCreatedAt = m.EnrollmentCreatedAt

	dataAO1.EnrollmentExpiresAt = m.EnrollmentExpiresAt

	dataAO1.EnrollmentJwt = m.EnrollmentJwt

	dataAO1.EnrollmentToken = m.EnrollmentToken

	dataAO1.Fingerprint = m.Fingerprint

	dataAO1.IsOnline = m.IsOnline

	dataAO1.IsVerified = m.IsVerified

	dataAO1.Name = m.Name

	dataAO1.UnverifiedCertPem = m.UnverifiedCertPem

	dataAO1.UnverifiedFingerprint = m.UnverifiedFingerprint

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this router detail
func (m *RouterDetail) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllowTraversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnrollmentCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnrollmentExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFingerprint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsOnline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsVerified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RouterDetail) validateAllowTraversal(formats strfmt.Registry) error {

	if err := validate.Required("allowTraversal", "body", m.AllowTraversal); err != nil {
		return err
	}

	return nil
}

func (m *RouterDetail) validateCost(formats strfmt.Registry) error {

	if err := validate.Required("cost", "body", m.Cost); err != nil {
		return err
	}

	if err := validate.MinimumInt("cost", "body", *m.Cost, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("cost", "body", *m.Cost, 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *RouterDetail) validateEnrollmentCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.EnrollmentCreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("enrollmentCreatedAt", "body", "date-time", m.EnrollmentCreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RouterDetail) validateEnrollmentExpiresAt(formats strfmt.Registry) error {

	if swag.IsZero(m.EnrollmentExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("enrollmentExpiresAt", "body", "date-time", m.EnrollmentExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RouterDetail) validateFingerprint(formats strfmt.Registry) error {

	if err := validate.Required("fingerprint", "body", m.Fingerprint); err != nil {
		return err
	}

	return nil
}

func (m *RouterDetail) validateIsOnline(formats strfmt.Registry) error {

	if err := validate.Required("isOnline", "body", m.IsOnline); err != nil {
		return err
	}

	return nil
}

func (m *RouterDetail) validateIsVerified(formats strfmt.Registry) error {

	if err := validate.Required("isVerified", "body", m.IsVerified); err != nil {
		return err
	}

	return nil
}

func (m *RouterDetail) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this router detail based on the context it is used
func (m *RouterDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseEntity
	if err := m.BaseEntity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RouterDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouterDetail) UnmarshalBinary(b []byte) error {
	var res RouterDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
