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
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostureCheckFailureProcessMulti posture check failure process multi
//
// swagger:model postureCheckFailureProcessMulti
type PostureCheckFailureProcessMulti struct {
	postureCheckIdField *string

	postureCheckNameField *string

	// actual value
	// Required: true
	ActualValue []*PostureCheckFailureProcessActual `json:"actualValue"`

	// expected value
	// Required: true
	ExpectedValue []*ProcessMulti `json:"expectedValue"`

	// semantic
	// Required: true
	Semantic Semantic `json:"semantic"`
}

// PostureCheckID gets the posture check Id of this subtype
func (m *PostureCheckFailureProcessMulti) PostureCheckID() *string {
	return m.postureCheckIdField
}

// SetPostureCheckID sets the posture check Id of this subtype
func (m *PostureCheckFailureProcessMulti) SetPostureCheckID(val *string) {
	m.postureCheckIdField = val
}

// PostureCheckName gets the posture check name of this subtype
func (m *PostureCheckFailureProcessMulti) PostureCheckName() *string {
	return m.postureCheckNameField
}

// SetPostureCheckName sets the posture check name of this subtype
func (m *PostureCheckFailureProcessMulti) SetPostureCheckName(val *string) {
	m.postureCheckNameField = val
}

// PostureCheckType gets the posture check type of this subtype
func (m *PostureCheckFailureProcessMulti) PostureCheckType() string {
	return "PROCESS_MULTI"
}

// SetPostureCheckType sets the posture check type of this subtype
func (m *PostureCheckFailureProcessMulti) SetPostureCheckType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PostureCheckFailureProcessMulti) UnmarshalJSON(raw []byte) error {
	var data struct {

		// actual value
		// Required: true
		ActualValue []*PostureCheckFailureProcessActual `json:"actualValue"`

		// expected value
		// Required: true
		ExpectedValue []*ProcessMulti `json:"expectedValue"`

		// semantic
		// Required: true
		Semantic Semantic `json:"semantic"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		PostureCheckID *string `json:"postureCheckId"`

		PostureCheckName *string `json:"postureCheckName"`

		PostureCheckType string `json:"postureCheckType"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result PostureCheckFailureProcessMulti

	result.postureCheckIdField = base.PostureCheckID

	result.postureCheckNameField = base.PostureCheckName

	if base.PostureCheckType != result.PostureCheckType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid postureCheckType value: %q", base.PostureCheckType)
	}

	result.ActualValue = data.ActualValue
	result.ExpectedValue = data.ExpectedValue
	result.Semantic = data.Semantic

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PostureCheckFailureProcessMulti) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// actual value
		// Required: true
		ActualValue []*PostureCheckFailureProcessActual `json:"actualValue"`

		// expected value
		// Required: true
		ExpectedValue []*ProcessMulti `json:"expectedValue"`

		// semantic
		// Required: true
		Semantic Semantic `json:"semantic"`
	}{

		ActualValue: m.ActualValue,

		ExpectedValue: m.ExpectedValue,

		Semantic: m.Semantic,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		PostureCheckID *string `json:"postureCheckId"`

		PostureCheckName *string `json:"postureCheckName"`

		PostureCheckType string `json:"postureCheckType"`
	}{

		PostureCheckID: m.PostureCheckID(),

		PostureCheckName: m.PostureCheckName(),

		PostureCheckType: m.PostureCheckType(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this posture check failure process multi
func (m *PostureCheckFailureProcessMulti) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePostureCheckID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostureCheckName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActualValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSemantic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostureCheckFailureProcessMulti) validatePostureCheckID(formats strfmt.Registry) error {

	if err := validate.Required("postureCheckId", "body", m.PostureCheckID()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckFailureProcessMulti) validatePostureCheckName(formats strfmt.Registry) error {

	if err := validate.Required("postureCheckName", "body", m.PostureCheckName()); err != nil {
		return err
	}

	return nil
}

func (m *PostureCheckFailureProcessMulti) validateActualValue(formats strfmt.Registry) error {

	if err := validate.Required("actualValue", "body", m.ActualValue); err != nil {
		return err
	}

	for i := 0; i < len(m.ActualValue); i++ {
		if swag.IsZero(m.ActualValue[i]) { // not required
			continue
		}

		if m.ActualValue[i] != nil {
			if err := m.ActualValue[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actualValue" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostureCheckFailureProcessMulti) validateExpectedValue(formats strfmt.Registry) error {

	if err := validate.Required("expectedValue", "body", m.ExpectedValue); err != nil {
		return err
	}

	for i := 0; i < len(m.ExpectedValue); i++ {
		if swag.IsZero(m.ExpectedValue[i]) { // not required
			continue
		}

		if m.ExpectedValue[i] != nil {
			if err := m.ExpectedValue[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("expectedValue" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostureCheckFailureProcessMulti) validateSemantic(formats strfmt.Registry) error {

	if err := m.Semantic.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("semantic")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostureCheckFailureProcessMulti) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostureCheckFailureProcessMulti) UnmarshalBinary(b []byte) error {
	var res PostureCheckFailureProcessMulti
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}