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

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// NewOrderParams creates a new OrderParams object
// with the default values initialized.
func NewOrderParams() *OrderParams {
	var ()
	return &OrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrderParamsWithTimeout creates a new OrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderParamsWithTimeout(timeout time.Duration) *OrderParams {
	var ()
	return &OrderParams{

		timeout: timeout,
	}
}

// NewOrderParamsWithContext creates a new OrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderParamsWithContext(ctx context.Context) *OrderParams {
	var ()
	return &OrderParams{

		Context: ctx,
	}
}

// NewOrderParamsWithHTTPClient creates a new OrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrderParamsWithHTTPClient(client *http.Client) *OrderParams {
	var ()
	return &OrderParams{
		HTTPClient: client,
	}
}

/*OrderParams contains all the parameters to send to the API endpoint
for the order operation typically these are written to a http.Request
*/
type OrderParams struct {

	/*JSONBody*/
	JSONBody *models.OrderExtractorsRequest
	/*InputID
	  Persist ID (!) of input.

	*/
	InputID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the order params
func (o *OrderParams) WithTimeout(timeout time.Duration) *OrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order params
func (o *OrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order params
func (o *OrderParams) WithContext(ctx context.Context) *OrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order params
func (o *OrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order params
func (o *OrderParams) WithHTTPClient(client *http.Client) *OrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order params
func (o *OrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the order params
func (o *OrderParams) WithJSONBody(jSONBody *models.OrderExtractorsRequest) *OrderParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the order params
func (o *OrderParams) SetJSONBody(jSONBody *models.OrderExtractorsRequest) {
	o.JSONBody = jSONBody
}

// WithInputID adds the inputID to the order params
func (o *OrderParams) WithInputID(inputID string) *OrderParams {
	o.SetInputID(inputID)
	return o
}

// SetInputID adds the inputId to the order params
func (o *OrderParams) SetInputID(inputID string) {
	o.InputID = inputID
}

// WriteToRequest writes these params to a swagger request
func (o *OrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
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
