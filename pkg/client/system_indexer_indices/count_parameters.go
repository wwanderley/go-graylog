// Code generated by go-swagger; DO NOT EDIT.

package system_indexer_indices

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

// NewCountParams creates a new CountParams object
// with the default values initialized.
func NewCountParams() *CountParams {
	var ()
	return &CountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCountParamsWithTimeout creates a new CountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCountParamsWithTimeout(timeout time.Duration) *CountParams {
	var ()
	return &CountParams{

		timeout: timeout,
	}
}

// NewCountParamsWithContext creates a new CountParams object
// with the default values initialized, and the ability to set a context for a request
func NewCountParamsWithContext(ctx context.Context) *CountParams {
	var ()
	return &CountParams{

		Context: ctx,
	}
}

// NewCountParamsWithHTTPClient creates a new CountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCountParamsWithHTTPClient(client *http.Client) *CountParams {
	var ()
	return &CountParams{
		HTTPClient: client,
	}
}

/*CountParams contains all the parameters to send to the API endpoint
for the count operation typically these are written to a http.Request
*/
type CountParams struct {

	/*Since
	  ISO8601 date

	*/
	Since string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the count params
func (o *CountParams) WithTimeout(timeout time.Duration) *CountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the count params
func (o *CountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the count params
func (o *CountParams) WithContext(ctx context.Context) *CountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the count params
func (o *CountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the count params
func (o *CountParams) WithHTTPClient(client *http.Client) *CountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the count params
func (o *CountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSince adds the since to the count params
func (o *CountParams) WithSince(since string) *CountParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the count params
func (o *CountParams) SetSince(since string) {
	o.Since = since
}

// WriteToRequest writes these params to a swagger request
func (o *CountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param since
	qrSince := o.Since
	qSince := qrSince
	if qSince != "" {
		if err := r.SetQueryParam("since", qSince); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
