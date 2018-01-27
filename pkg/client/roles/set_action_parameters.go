// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// NewSetActionParams creates a new SetActionParams object
// with the default values initialized.
func NewSetActionParams() *SetActionParams {
	var ()
	return &SetActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetActionParamsWithTimeout creates a new SetActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetActionParamsWithTimeout(timeout time.Duration) *SetActionParams {
	var ()
	return &SetActionParams{

		timeout: timeout,
	}
}

// NewSetActionParamsWithContext creates a new SetActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetActionParamsWithContext(ctx context.Context) *SetActionParams {
	var ()
	return &SetActionParams{

		Context: ctx,
	}
}

// NewSetActionParamsWithHTTPClient creates a new SetActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetActionParamsWithHTTPClient(client *http.Client) *SetActionParams {
	var ()
	return &SetActionParams{
		HTTPClient: client,
	}
}

/*SetActionParams contains all the parameters to send to the API endpoint
for the set action operation typically these are written to a http.Request
*/
type SetActionParams struct {

	/*JSONBody*/
	JSONBody []interface{}
	/*CollectorID
	  The collector id this collector is registering as.

	*/
	CollectorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set action params
func (o *SetActionParams) WithTimeout(timeout time.Duration) *SetActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set action params
func (o *SetActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set action params
func (o *SetActionParams) WithContext(ctx context.Context) *SetActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set action params
func (o *SetActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set action params
func (o *SetActionParams) WithHTTPClient(client *http.Client) *SetActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set action params
func (o *SetActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the set action params
func (o *SetActionParams) WithJSONBody(jSONBody []interface{}) *SetActionParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the set action params
func (o *SetActionParams) SetJSONBody(jSONBody []interface{}) {
	o.JSONBody = jSONBody
}

// WithCollectorID adds the collectorID to the set action params
func (o *SetActionParams) WithCollectorID(collectorID string) *SetActionParams {
	o.SetCollectorID(collectorID)
	return o
}

// SetCollectorID adds the collectorId to the set action params
func (o *SetActionParams) SetCollectorID(collectorID string) {
	o.CollectorID = collectorID
}

// WriteToRequest writes these params to a swagger request
func (o *SetActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	// path param collectorId
	if err := r.SetPathParam("collectorId", o.CollectorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}