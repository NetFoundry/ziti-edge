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

package external_j_w_t_signer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteExternalJwtSignerParams creates a new DeleteExternalJwtSignerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteExternalJwtSignerParams() *DeleteExternalJwtSignerParams {
	return &DeleteExternalJwtSignerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteExternalJwtSignerParamsWithTimeout creates a new DeleteExternalJwtSignerParams object
// with the ability to set a timeout on a request.
func NewDeleteExternalJwtSignerParamsWithTimeout(timeout time.Duration) *DeleteExternalJwtSignerParams {
	return &DeleteExternalJwtSignerParams{
		timeout: timeout,
	}
}

// NewDeleteExternalJwtSignerParamsWithContext creates a new DeleteExternalJwtSignerParams object
// with the ability to set a context for a request.
func NewDeleteExternalJwtSignerParamsWithContext(ctx context.Context) *DeleteExternalJwtSignerParams {
	return &DeleteExternalJwtSignerParams{
		Context: ctx,
	}
}

// NewDeleteExternalJwtSignerParamsWithHTTPClient creates a new DeleteExternalJwtSignerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteExternalJwtSignerParamsWithHTTPClient(client *http.Client) *DeleteExternalJwtSignerParams {
	return &DeleteExternalJwtSignerParams{
		HTTPClient: client,
	}
}

/* DeleteExternalJwtSignerParams contains all the parameters to send to the API endpoint
   for the delete external jwt signer operation.

   Typically these are written to a http.Request.
*/
type DeleteExternalJwtSignerParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete external jwt signer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExternalJwtSignerParams) WithDefaults() *DeleteExternalJwtSignerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete external jwt signer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExternalJwtSignerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete external jwt signer params
func (o *DeleteExternalJwtSignerParams) WithTimeout(timeout time.Duration) *DeleteExternalJwtSignerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete external jwt signer params
func (o *DeleteExternalJwtSignerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete external jwt signer params
func (o *DeleteExternalJwtSignerParams) WithContext(ctx context.Context) *DeleteExternalJwtSignerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete external jwt signer params
func (o *DeleteExternalJwtSignerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete external jwt signer params
func (o *DeleteExternalJwtSignerParams) WithHTTPClient(client *http.Client) *DeleteExternalJwtSignerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete external jwt signer params
func (o *DeleteExternalJwtSignerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete external jwt signer params
func (o *DeleteExternalJwtSignerParams) WithID(id string) *DeleteExternalJwtSignerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete external jwt signer params
func (o *DeleteExternalJwtSignerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteExternalJwtSignerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}