// Code generated by go-swagger; DO NOT EDIT.

package system_indices_retention

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

// NewIndexSetStatisticsParams creates a new IndexSetStatisticsParams object
// with the default values initialized.
func NewIndexSetStatisticsParams() *IndexSetStatisticsParams {
	var ()
	return &IndexSetStatisticsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexSetStatisticsParamsWithTimeout creates a new IndexSetStatisticsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexSetStatisticsParamsWithTimeout(timeout time.Duration) *IndexSetStatisticsParams {
	var ()
	return &IndexSetStatisticsParams{

		timeout: timeout,
	}
}

// NewIndexSetStatisticsParamsWithContext creates a new IndexSetStatisticsParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexSetStatisticsParamsWithContext(ctx context.Context) *IndexSetStatisticsParams {
	var ()
	return &IndexSetStatisticsParams{

		Context: ctx,
	}
}

// NewIndexSetStatisticsParamsWithHTTPClient creates a new IndexSetStatisticsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexSetStatisticsParamsWithHTTPClient(client *http.Client) *IndexSetStatisticsParams {
	var ()
	return &IndexSetStatisticsParams{
		HTTPClient: client,
	}
}

/*IndexSetStatisticsParams contains all the parameters to send to the API endpoint
for the index set statistics operation typically these are written to a http.Request
*/
type IndexSetStatisticsParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the index set statistics params
func (o *IndexSetStatisticsParams) WithTimeout(timeout time.Duration) *IndexSetStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index set statistics params
func (o *IndexSetStatisticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index set statistics params
func (o *IndexSetStatisticsParams) WithContext(ctx context.Context) *IndexSetStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index set statistics params
func (o *IndexSetStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index set statistics params
func (o *IndexSetStatisticsParams) WithHTTPClient(client *http.Client) *IndexSetStatisticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index set statistics params
func (o *IndexSetStatisticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the index set statistics params
func (o *IndexSetStatisticsParams) WithID(id string) *IndexSetStatisticsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the index set statistics params
func (o *IndexSetStatisticsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IndexSetStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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