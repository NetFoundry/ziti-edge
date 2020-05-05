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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TerminatorUpdate terminator update
//
// swagger:model terminatorUpdate
type TerminatorUpdate struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// binding
	// Required: true
	Binding *string `json:"binding"`

	// cost
	Cost TerminatorCost `json:"cost,omitempty"`

	// precedence
	Precedence TerminatorPrecedence `json:"precedence,omitempty"`

	// router
	// Required: true
	Router *string `json:"router"`

	// service
	// Required: true
	Service *string `json:"service"`

	// tags
	// Required: true
	Tags Tags `json:"tags"`
}

// Validate validates this terminator update
func (m *TerminatorUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBinding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrecedence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerminatorUpdate) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *TerminatorUpdate) validateBinding(formats strfmt.Registry) error {

	if err := validate.Required("binding", "body", m.Binding); err != nil {
		return err
	}

	return nil
}

func (m *TerminatorUpdate) validateCost(formats strfmt.Registry) error {

	if swag.IsZero(m.Cost) { // not required
		return nil
	}

	if err := m.Cost.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cost")
		}
		return err
	}

	return nil
}

func (m *TerminatorUpdate) validatePrecedence(formats strfmt.Registry) error {

	if swag.IsZero(m.Precedence) { // not required
		return nil
	}

	if err := m.Precedence.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("precedence")
		}
		return err
	}

	return nil
}

func (m *TerminatorUpdate) validateRouter(formats strfmt.Registry) error {

	if err := validate.Required("router", "body", m.Router); err != nil {
		return err
	}

	return nil
}

func (m *TerminatorUpdate) validateService(formats strfmt.Registry) error {

	if err := validate.Required("service", "body", m.Service); err != nil {
		return err
	}

	return nil
}

func (m *TerminatorUpdate) validateTags(formats strfmt.Registry) error {

	if err := m.Tags.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tags")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerminatorUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerminatorUpdate) UnmarshalBinary(b []byte) error {
	var res TerminatorUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
