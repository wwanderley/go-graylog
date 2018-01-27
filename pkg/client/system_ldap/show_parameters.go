// Code generated by go-swagger; DO NOT EDIT.

package system_ldap

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

// NewShowParams creates a new ShowParams object
// with the default values initialized.
func NewShowParams() *ShowParams {

	return &ShowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShowParamsWithTimeout creates a new ShowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShowParamsWithTimeout(timeout time.Duration) *ShowParams {

	return &ShowParams{

		timeout: timeout,
	}
}

// NewShowParamsWithContext creates a new ShowParams object
// with the default values initialized, and the ability to set a context for a request
func NewShowParamsWithContext(ctx context.Context) *ShowParams {

	return &ShowParams{

		Context: ctx,
	}
}

// NewShowParamsWithHTTPClient creates a new ShowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShowParamsWithHTTPClient(client *http.Client) *ShowParams {

	return &ShowParams{
		HTTPClient: client,
	}
}

/*ShowParams contains all the parameters to send to the API endpoint
for the show operation typically these are written to a http.Request
*/
type ShowParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the show params
func (o *ShowParams) WithTimeout(timeout time.Duration) *ShowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show params
func (o *ShowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show params
func (o *ShowParams) WithContext(ctx context.Context) *ShowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show params
func (o *ShowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show params
func (o *ShowParams) WithHTTPClient(client *http.Client) *ShowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show params
func (o *ShowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ShowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
