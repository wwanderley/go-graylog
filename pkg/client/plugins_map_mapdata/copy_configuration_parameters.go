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

// NewCopyConfigurationParams creates a new CopyConfigurationParams object
// with the default values initialized.
func NewCopyConfigurationParams() *CopyConfigurationParams {
	var ()
	return &CopyConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCopyConfigurationParamsWithTimeout creates a new CopyConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCopyConfigurationParamsWithTimeout(timeout time.Duration) *CopyConfigurationParams {
	var ()
	return &CopyConfigurationParams{

		timeout: timeout,
	}
}

// NewCopyConfigurationParamsWithContext creates a new CopyConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCopyConfigurationParamsWithContext(ctx context.Context) *CopyConfigurationParams {
	var ()
	return &CopyConfigurationParams{

		Context: ctx,
	}
}

// NewCopyConfigurationParamsWithHTTPClient creates a new CopyConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCopyConfigurationParamsWithHTTPClient(client *http.Client) *CopyConfigurationParams {
	var ()
	return &CopyConfigurationParams{
		HTTPClient: client,
	}
}

/*CopyConfigurationParams contains all the parameters to send to the API endpoint
for the copy configuration operation typically these are written to a http.Request
*/
type CopyConfigurationParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the copy configuration params
func (o *CopyConfigurationParams) WithTimeout(timeout time.Duration) *CopyConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the copy configuration params
func (o *CopyConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the copy configuration params
func (o *CopyConfigurationParams) WithContext(ctx context.Context) *CopyConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the copy configuration params
func (o *CopyConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the copy configuration params
func (o *CopyConfigurationParams) WithHTTPClient(client *http.Client) *CopyConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the copy configuration params
func (o *CopyConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the copy configuration params
func (o *CopyConfigurationParams) WithID(id string) *CopyConfigurationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the copy configuration params
func (o *CopyConfigurationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CopyConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
