// Code generated by go-swagger; DO NOT EDIT.

package system_inputs

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

// NewStopParams creates a new StopParams object
// with the default values initialized.
func NewStopParams() *StopParams {
	var ()
	return &StopParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopParamsWithTimeout creates a new StopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopParamsWithTimeout(timeout time.Duration) *StopParams {
	var ()
	return &StopParams{

		timeout: timeout,
	}
}

// NewStopParamsWithContext creates a new StopParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopParamsWithContext(ctx context.Context) *StopParams {
	var ()
	return &StopParams{

		Context: ctx,
	}
}

// NewStopParamsWithHTTPClient creates a new StopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopParamsWithHTTPClient(client *http.Client) *StopParams {
	var ()
	return &StopParams{
		HTTPClient: client,
	}
}

/*StopParams contains all the parameters to send to the API endpoint
for the stop operation typically these are written to a http.Request
*/
type StopParams struct {

	/*InputID*/
	InputID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop params
func (o *StopParams) WithTimeout(timeout time.Duration) *StopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop params
func (o *StopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop params
func (o *StopParams) WithContext(ctx context.Context) *StopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop params
func (o *StopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop params
func (o *StopParams) WithHTTPClient(client *http.Client) *StopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop params
func (o *StopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInputID adds the inputID to the stop params
func (o *StopParams) WithInputID(inputID string) *StopParams {
	o.SetInputID(inputID)
	return o
}

// SetInputID adds the inputId to the stop params
func (o *StopParams) SetInputID(inputID string) {
	o.InputID = inputID
}

// WriteToRequest writes these params to a swagger request
func (o *StopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inputId
	if err := r.SetPathParam("inputId", o.InputID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
