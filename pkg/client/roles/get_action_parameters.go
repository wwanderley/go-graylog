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

// NewGetActionParams creates a new GetActionParams object
// with the default values initialized.
func NewGetActionParams() *GetActionParams {
	var ()
	return &GetActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetActionParamsWithTimeout creates a new GetActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetActionParamsWithTimeout(timeout time.Duration) *GetActionParams {
	var ()
	return &GetActionParams{

		timeout: timeout,
	}
}

// NewGetActionParamsWithContext creates a new GetActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetActionParamsWithContext(ctx context.Context) *GetActionParams {
	var ()
	return &GetActionParams{

		Context: ctx,
	}
}

// NewGetActionParamsWithHTTPClient creates a new GetActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetActionParamsWithHTTPClient(client *http.Client) *GetActionParams {
	var ()
	return &GetActionParams{
		HTTPClient: client,
	}
}

/*GetActionParams contains all the parameters to send to the API endpoint
for the get action operation typically these are written to a http.Request
*/
type GetActionParams struct {

	/*CollectorID*/
	CollectorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get action params
func (o *GetActionParams) WithTimeout(timeout time.Duration) *GetActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get action params
func (o *GetActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get action params
func (o *GetActionParams) WithContext(ctx context.Context) *GetActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get action params
func (o *GetActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get action params
func (o *GetActionParams) WithHTTPClient(client *http.Client) *GetActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get action params
func (o *GetActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectorID adds the collectorID to the get action params
func (o *GetActionParams) WithCollectorID(collectorID string) *GetActionParams {
	o.SetCollectorID(collectorID)
	return o
}

// SetCollectorID adds the collectorId to the get action params
func (o *GetActionParams) SetCollectorID(collectorID string) {
	o.CollectorID = collectorID
}

// WriteToRequest writes these params to a swagger request
func (o *GetActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collectorId
	if err := r.SetPathParam("collectorId", o.CollectorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
