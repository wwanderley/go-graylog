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

// NewDeleteInputParams creates a new DeleteInputParams object
// with the default values initialized.
func NewDeleteInputParams() *DeleteInputParams {
	var ()
	return &DeleteInputParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInputParamsWithTimeout creates a new DeleteInputParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInputParamsWithTimeout(timeout time.Duration) *DeleteInputParams {
	var ()
	return &DeleteInputParams{

		timeout: timeout,
	}
}

// NewDeleteInputParamsWithContext creates a new DeleteInputParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInputParamsWithContext(ctx context.Context) *DeleteInputParams {
	var ()
	return &DeleteInputParams{

		Context: ctx,
	}
}

// NewDeleteInputParamsWithHTTPClient creates a new DeleteInputParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInputParamsWithHTTPClient(client *http.Client) *DeleteInputParams {
	var ()
	return &DeleteInputParams{
		HTTPClient: client,
	}
}

/*DeleteInputParams contains all the parameters to send to the API endpoint
for the delete input operation typically these are written to a http.Request
*/
type DeleteInputParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete input params
func (o *DeleteInputParams) WithTimeout(timeout time.Duration) *DeleteInputParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete input params
func (o *DeleteInputParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete input params
func (o *DeleteInputParams) WithContext(ctx context.Context) *DeleteInputParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete input params
func (o *DeleteInputParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete input params
func (o *DeleteInputParams) WithHTTPClient(client *http.Client) *DeleteInputParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete input params
func (o *DeleteInputParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete input params
func (o *DeleteInputParams) WithID(id string) *DeleteInputParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete input params
func (o *DeleteInputParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInputParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
