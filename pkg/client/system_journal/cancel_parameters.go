// Code generated by go-swagger; DO NOT EDIT.

package system_journal

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

// NewCancelParams creates a new CancelParams object
// with the default values initialized.
func NewCancelParams() *CancelParams {
	var ()
	return &CancelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelParamsWithTimeout creates a new CancelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelParamsWithTimeout(timeout time.Duration) *CancelParams {
	var ()
	return &CancelParams{

		timeout: timeout,
	}
}

// NewCancelParamsWithContext creates a new CancelParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelParamsWithContext(ctx context.Context) *CancelParams {
	var ()
	return &CancelParams{

		Context: ctx,
	}
}

// NewCancelParamsWithHTTPClient creates a new CancelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelParamsWithHTTPClient(client *http.Client) *CancelParams {
	var ()
	return &CancelParams{
		HTTPClient: client,
	}
}

/*CancelParams contains all the parameters to send to the API endpoint
for the cancel operation typically these are written to a http.Request
*/
type CancelParams struct {

	/*JobID*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel params
func (o *CancelParams) WithTimeout(timeout time.Duration) *CancelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel params
func (o *CancelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel params
func (o *CancelParams) WithContext(ctx context.Context) *CancelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel params
func (o *CancelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel params
func (o *CancelParams) WithHTTPClient(client *http.Client) *CancelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel params
func (o *CancelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the cancel params
func (o *CancelParams) WithJobID(jobID string) *CancelParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the cancel params
func (o *CancelParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param jobId
	if err := r.SetPathParam("jobId", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}