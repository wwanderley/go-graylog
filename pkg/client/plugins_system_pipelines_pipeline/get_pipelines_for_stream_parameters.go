// Code generated by go-swagger; DO NOT EDIT.

package plugins_system_pipelines_pipeline

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

// NewGetPipelinesForStreamParams creates a new GetPipelinesForStreamParams object
// with the default values initialized.
func NewGetPipelinesForStreamParams() *GetPipelinesForStreamParams {
	var ()
	return &GetPipelinesForStreamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPipelinesForStreamParamsWithTimeout creates a new GetPipelinesForStreamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPipelinesForStreamParamsWithTimeout(timeout time.Duration) *GetPipelinesForStreamParams {
	var ()
	return &GetPipelinesForStreamParams{

		timeout: timeout,
	}
}

// NewGetPipelinesForStreamParamsWithContext creates a new GetPipelinesForStreamParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPipelinesForStreamParamsWithContext(ctx context.Context) *GetPipelinesForStreamParams {
	var ()
	return &GetPipelinesForStreamParams{

		Context: ctx,
	}
}

// NewGetPipelinesForStreamParamsWithHTTPClient creates a new GetPipelinesForStreamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPipelinesForStreamParamsWithHTTPClient(client *http.Client) *GetPipelinesForStreamParams {
	var ()
	return &GetPipelinesForStreamParams{
		HTTPClient: client,
	}
}

/*GetPipelinesForStreamParams contains all the parameters to send to the API endpoint
for the get pipelines for stream operation typically these are written to a http.Request
*/
type GetPipelinesForStreamParams struct {

	/*StreamID*/
	StreamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pipelines for stream params
func (o *GetPipelinesForStreamParams) WithTimeout(timeout time.Duration) *GetPipelinesForStreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pipelines for stream params
func (o *GetPipelinesForStreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pipelines for stream params
func (o *GetPipelinesForStreamParams) WithContext(ctx context.Context) *GetPipelinesForStreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pipelines for stream params
func (o *GetPipelinesForStreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pipelines for stream params
func (o *GetPipelinesForStreamParams) WithHTTPClient(client *http.Client) *GetPipelinesForStreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pipelines for stream params
func (o *GetPipelinesForStreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStreamID adds the streamID to the get pipelines for stream params
func (o *GetPipelinesForStreamParams) WithStreamID(streamID string) *GetPipelinesForStreamParams {
	o.SetStreamID(streamID)
	return o
}

// SetStreamID adds the streamId to the get pipelines for stream params
func (o *GetPipelinesForStreamParams) SetStreamID(streamID string) {
	o.StreamID = streamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPipelinesForStreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param streamId
	if err := r.SetPathParam("streamId", o.StreamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}