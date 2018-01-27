// Code generated by go-swagger; DO NOT EDIT.

package search_universal_absolute

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

// NewListAllParams creates a new ListAllParams object
// with the default values initialized.
func NewListAllParams() *ListAllParams {

	return &ListAllParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllParamsWithTimeout creates a new ListAllParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllParamsWithTimeout(timeout time.Duration) *ListAllParams {

	return &ListAllParams{

		timeout: timeout,
	}
}

// NewListAllParamsWithContext creates a new ListAllParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllParamsWithContext(ctx context.Context) *ListAllParams {

	return &ListAllParams{

		Context: ctx,
	}
}

// NewListAllParamsWithHTTPClient creates a new ListAllParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllParamsWithHTTPClient(client *http.Client) *ListAllParams {

	return &ListAllParams{
		HTTPClient: client,
	}
}

/*ListAllParams contains all the parameters to send to the API endpoint
for the list all operation typically these are written to a http.Request
*/
type ListAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all params
func (o *ListAllParams) WithTimeout(timeout time.Duration) *ListAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all params
func (o *ListAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all params
func (o *ListAllParams) WithContext(ctx context.Context) *ListAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all params
func (o *ListAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all params
func (o *ListAllParams) WithHTTPClient(client *http.Client) *ListAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all params
func (o *ListAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
