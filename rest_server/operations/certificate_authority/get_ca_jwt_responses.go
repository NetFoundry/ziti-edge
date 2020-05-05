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

package certificate_authority

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/netfoundry/ziti-edge/rest_model"
)

// GetCaJwtOKCode is the HTTP code returned for type GetCaJwtOK
const GetCaJwtOKCode int = 200

/*GetCaJwtOK The result is the JWT text to validate the CA

swagger:response getCaJwtOK
*/
type GetCaJwtOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetCaJwtOK creates GetCaJwtOK with default headers values
func NewGetCaJwtOK() *GetCaJwtOK {

	return &GetCaJwtOK{}
}

// WithPayload adds the payload to the get ca jwt o k response
func (o *GetCaJwtOK) WithPayload(payload string) *GetCaJwtOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ca jwt o k response
func (o *GetCaJwtOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCaJwtOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCaJwtUnauthorizedCode is the HTTP code returned for type GetCaJwtUnauthorized
const GetCaJwtUnauthorizedCode int = 401

/*GetCaJwtUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response getCaJwtUnauthorized
*/
type GetCaJwtUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetCaJwtUnauthorized creates GetCaJwtUnauthorized with default headers values
func NewGetCaJwtUnauthorized() *GetCaJwtUnauthorized {

	return &GetCaJwtUnauthorized{}
}

// WithPayload adds the payload to the get ca jwt unauthorized response
func (o *GetCaJwtUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *GetCaJwtUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ca jwt unauthorized response
func (o *GetCaJwtUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCaJwtUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCaJwtNotFoundCode is the HTTP code returned for type GetCaJwtNotFound
const GetCaJwtNotFoundCode int = 404

/*GetCaJwtNotFound The requested resource does not exist

swagger:response getCaJwtNotFound
*/
type GetCaJwtNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewGetCaJwtNotFound creates GetCaJwtNotFound with default headers values
func NewGetCaJwtNotFound() *GetCaJwtNotFound {

	return &GetCaJwtNotFound{}
}

// WithPayload adds the payload to the get ca jwt not found response
func (o *GetCaJwtNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *GetCaJwtNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ca jwt not found response
func (o *GetCaJwtNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCaJwtNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
