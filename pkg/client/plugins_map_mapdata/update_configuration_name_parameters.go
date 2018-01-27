// Code generated by go-swagger; DO NOT EDIT.

package plugins_map_mapdata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateConfigurationNameParams creates a new UpdateConfigurationNameParams object
// with the default values initialized.
func NewUpdateConfigurationNameParams() *UpdateConfigurationNameParams {
	var ()
	return &UpdateConfigurationNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConfigurationNameParamsWithTimeout creates a new UpdateConfigurationNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateConfigurationNameParamsWithTimeout(timeout time.Duration) *UpdateConfigurationNameParams {
	var ()
	return &UpdateConfigurationNameParams{

		timeout: timeout,
	}
}

// NewUpdateConfigurationNameParamsWithContext creates a new UpdateConfigurationNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateConfigurationNameParamsWithContext(ctx context.Context) *UpdateConfigurationNameParams {
	var ()
	return &UpdateConfigurationNameParams{

		Context: ctx,
	}
}

// NewUpdateConfigurationNameParamsWithHTTPClient creates a new UpdateConfigurationNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateConfigurationNameParamsWithHTTPClient(client *http.Client) *UpdateConfigurationNameParams {
	var ()
	return &UpdateConfigurationNameParams{
		HTTPClient: client,
	}
}

/*UpdateConfigurationNameParams contains all the parameters to send to the API endpoint
for the update configuration name operation typically these are written to a http.Request
*/
type UpdateConfigurationNameParams struct {

	/*JSONBody*/
	JSONBody interface{}
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update configuration name params
func (o *UpdateConfigurationNameParams) WithTimeout(timeout time.Duration) *UpdateConfigurationNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update configuration name params
func (o *UpdateConfigurationNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update configuration name params
func (o *UpdateConfigurationNameParams) WithContext(ctx context.Context) *UpdateConfigurationNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update configuration name params
func (o *UpdateConfigurationNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update configuration name params
func (o *UpdateConfigurationNameParams) WithHTTPClient(client *http.Client) *UpdateConfigurationNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update configuration name params
func (o *UpdateConfigurationNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the update configuration name params
func (o *UpdateConfigurationNameParams) WithJSONBody(jSONBody interface{}) *UpdateConfigurationNameParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the update configuration name params
func (o *UpdateConfigurationNameParams) SetJSONBody(jSONBody interface{}) {
	o.JSONBody = jSONBody
}

// WithID adds the id to the update configuration name params
func (o *UpdateConfigurationNameParams) WithID(id string) *UpdateConfigurationNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update configuration name params
func (o *UpdateConfigurationNameParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConfigurationNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
