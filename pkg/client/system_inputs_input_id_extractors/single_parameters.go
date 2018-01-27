// Code generated by go-swagger; DO NOT EDIT.

package system_inputs_input_id_extractors

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

// NewSingleParams creates a new SingleParams object
// with the default values initialized.
func NewSingleParams() *SingleParams {
	var ()
	return &SingleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSingleParamsWithTimeout creates a new SingleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSingleParamsWithTimeout(timeout time.Duration) *SingleParams {
	var ()
	return &SingleParams{

		timeout: timeout,
	}
}

// NewSingleParamsWithContext creates a new SingleParams object
// with the default values initialized, and the ability to set a context for a request
func NewSingleParamsWithContext(ctx context.Context) *SingleParams {
	var ()
	return &SingleParams{

		Context: ctx,
	}
}

// NewSingleParamsWithHTTPClient creates a new SingleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSingleParamsWithHTTPClient(client *http.Client) *SingleParams {
	var ()
	return &SingleParams{
		HTTPClient: client,
	}
}

/*SingleParams contains all the parameters to send to the API endpoint
for the single operation typically these are written to a http.Request
*/
type SingleParams struct {

	/*ExtractorID*/
	ExtractorID string
	/*InputID*/
	InputID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the single params
func (o *SingleParams) WithTimeout(timeout time.Duration) *SingleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the single params
func (o *SingleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the single params
func (o *SingleParams) WithContext(ctx context.Context) *SingleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the single params
func (o *SingleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the single params
func (o *SingleParams) WithHTTPClient(client *http.Client) *SingleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the single params
func (o *SingleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtractorID adds the extractorID to the single params
func (o *SingleParams) WithExtractorID(extractorID string) *SingleParams {
	o.SetExtractorID(extractorID)
	return o
}

// SetExtractorID adds the extractorId to the single params
func (o *SingleParams) SetExtractorID(extractorID string) {
	o.ExtractorID = extractorID
}

// WithInputID adds the inputID to the single params
func (o *SingleParams) WithInputID(inputID string) *SingleParams {
	o.SetInputID(inputID)
	return o
}

// SetInputID adds the inputId to the single params
func (o *SingleParams) SetInputID(inputID string) {
	o.InputID = inputID
}

// WriteToRequest writes these params to a swagger request
func (o *SingleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param extractorId
	if err := r.SetPathParam("extractorId", o.ExtractorID); err != nil {
		return err
	}

	// path param inputId
	if err := r.SetPathParam("inputId", o.InputID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
