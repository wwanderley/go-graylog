// Code generated by go-swagger; DO NOT EDIT.

package system_messageprocessors

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

// NewPerformLookupParams creates a new PerformLookupParams object
// with the default values initialized.
func NewPerformLookupParams() *PerformLookupParams {
	var ()
	return &PerformLookupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPerformLookupParamsWithTimeout creates a new PerformLookupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPerformLookupParamsWithTimeout(timeout time.Duration) *PerformLookupParams {
	var ()
	return &PerformLookupParams{

		timeout: timeout,
	}
}

// NewPerformLookupParamsWithContext creates a new PerformLookupParams object
// with the default values initialized, and the ability to set a context for a request
func NewPerformLookupParamsWithContext(ctx context.Context) *PerformLookupParams {
	var ()
	return &PerformLookupParams{

		Context: ctx,
	}
}

// NewPerformLookupParamsWithHTTPClient creates a new PerformLookupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPerformLookupParamsWithHTTPClient(client *http.Client) *PerformLookupParams {
	var ()
	return &PerformLookupParams{
		HTTPClient: client,
	}
}

/*PerformLookupParams contains all the parameters to send to the API endpoint
for the perform lookup operation typically these are written to a http.Request
*/
type PerformLookupParams struct {

	/*Key*/
	Key *string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the perform lookup params
func (o *PerformLookupParams) WithTimeout(timeout time.Duration) *PerformLookupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform lookup params
func (o *PerformLookupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform lookup params
func (o *PerformLookupParams) WithContext(ctx context.Context) *PerformLookupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform lookup params
func (o *PerformLookupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform lookup params
func (o *PerformLookupParams) WithHTTPClient(client *http.Client) *PerformLookupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform lookup params
func (o *PerformLookupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the perform lookup params
func (o *PerformLookupParams) WithKey(key *string) *PerformLookupParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the perform lookup params
func (o *PerformLookupParams) SetKey(key *string) {
	o.Key = key
}

// WithName adds the name to the perform lookup params
func (o *PerformLookupParams) WithName(name string) *PerformLookupParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the perform lookup params
func (o *PerformLookupParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PerformLookupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Key != nil {

		// query param key
		var qrKey string
		if o.Key != nil {
			qrKey = *o.Key
		}
		qKey := qrKey
		if qKey != "" {
			if err := r.SetQueryParam("key", qKey); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}