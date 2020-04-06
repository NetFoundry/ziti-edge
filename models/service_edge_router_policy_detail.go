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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceEdgeRouterPolicyDetail service edge router policy detail
//
// swagger:model serviceEdgeRouterPolicyDetail
type ServiceEdgeRouterPolicyDetail struct {

	// links
	Links Links `json:"_links,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// edge router roles
	EdgeRouterRoles Roles `json:"edgeRouterRoles,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// semantic
	Semantic Semantic `json:"semantic,omitempty"`

	// service roles
	ServiceRoles Roles `json:"serviceRoles,omitempty"`

	// tags
	Tags Tags `json:"tags,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this service edge router policy detail
func (m *ServiceEdgeRouterPolicyDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeRouterRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSemantic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if err := m.Links.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("_links")
		}
		return err
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) validateEdgeRouterRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeRouterRoles) { // not required
		return nil
	}

	if err := m.EdgeRouterRoles.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("edgeRouterRoles")
		}
		return err
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) validateSemantic(formats strfmt.Registry) error {

	if swag.IsZero(m.Semantic) { // not required
		return nil
	}

	if err := m.Semantic.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("semantic")
		}
		return err
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) validateServiceRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceRoles) { // not required
		return nil
	}

	if err := m.ServiceRoles.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceRoles")
		}
		return err
	}

	return nil
}

func (m *ServiceEdgeRouterPolicyDetail) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceEdgeRouterPolicyDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceEdgeRouterPolicyDetail) UnmarshalBinary(b []byte) error {
	var res ServiceEdgeRouterPolicyDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
