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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netfoundry/ziti-edge/rest_model"
)

// CreateServiceReader is a Reader for the CreateService structure.
type CreateServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateServiceOK creates a CreateServiceOK with default headers values
func NewCreateServiceOK() *CreateServiceOK {
	return &CreateServiceOK{}
}

/*CreateServiceOK handles this case with default header values.

The create request was successful and the resource has been added at the following location
*/
type CreateServiceOK struct {
	Payload *rest_model.CreateEnvelope
}

func (o *CreateServiceOK) Error() string {
	return fmt.Sprintf("[POST /services][%d] createServiceOK  %+v", 200, o.Payload)
}

func (o *CreateServiceOK) GetPayload() *rest_model.CreateEnvelope {
	return o.Payload
}

func (o *CreateServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.CreateEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceBadRequest creates a CreateServiceBadRequest with default headers values
func NewCreateServiceBadRequest() *CreateServiceBadRequest {
	return &CreateServiceBadRequest{}
}

/*CreateServiceBadRequest handles this case with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type CreateServiceBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateServiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /services][%d] createServiceBadRequest  %+v", 400, o.Payload)
}

func (o *CreateServiceBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceUnauthorized creates a CreateServiceUnauthorized with default headers values
func NewCreateServiceUnauthorized() *CreateServiceUnauthorized {
	return &CreateServiceUnauthorized{}
}

/*CreateServiceUnauthorized handles this case with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type CreateServiceUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *CreateServiceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /services][%d] createServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateServiceUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *CreateServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
