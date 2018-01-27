// Code generated by go-swagger; DO NOT EDIT.

package system_fields

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

// NewCycleParams creates a new CycleParams object
// with the default values initialized.
func NewCycleParams() *CycleParams {
	var ()
	return &CycleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCycleParamsWithTimeout creates a new CycleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCycleParamsWithTimeout(timeout time.Duration) *CycleParams {
	var ()
	return &CycleParams{

		timeout: timeout,
	}
}

// NewCycleParamsWithContext creates a new CycleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCycleParamsWithContext(ctx context.Context) *CycleParams {
	var ()
	return &CycleParams{

		Context: ctx,
	}
}

// NewCycleParamsWithHTTPClient creates a new CycleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCycleParamsWithHTTPClient(client *http.Client) *CycleParams {
	var ()
	return &CycleParams{
		HTTPClient: client,
	}
}

/*CycleParams contains all the parameters to send to the API endpoint
for the cycle operation typically these are written to a http.Request
*/
type CycleParams struct {

	/*IndexSetID*/
	IndexSetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cycle params
func (o *CycleParams) WithTimeout(timeout time.Duration) *CycleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cycle params
func (o *CycleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cycle params
func (o *CycleParams) WithContext(ctx context.Context) *CycleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cycle params
func (o *CycleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cycle params
func (o *CycleParams) WithHTTPClient(client *http.Client) *CycleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cycle params
func (o *CycleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndexSetID adds the indexSetID to the cycle params
func (o *CycleParams) WithIndexSetID(indexSetID string) *CycleParams {
	o.SetIndexSetID(indexSetID)
	return o
}

// SetIndexSetID adds the indexSetId to the cycle params
func (o *CycleParams) SetIndexSetID(indexSetID string) {
	o.IndexSetID = indexSetID
}

// WriteToRequest writes these params to a swagger request
func (o *CycleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param indexSetId
	if err := r.SetPathParam("indexSetId", o.IndexSetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
