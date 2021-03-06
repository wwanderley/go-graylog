// Code generated by go-swagger; DO NOT EDIT.

package system_cluster

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

// NewDeleteBundleParams creates a new DeleteBundleParams object
// with the default values initialized.
func NewDeleteBundleParams() *DeleteBundleParams {
	var ()
	return &DeleteBundleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBundleParamsWithTimeout creates a new DeleteBundleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBundleParamsWithTimeout(timeout time.Duration) *DeleteBundleParams {
	var ()
	return &DeleteBundleParams{

		timeout: timeout,
	}
}

// NewDeleteBundleParamsWithContext creates a new DeleteBundleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBundleParamsWithContext(ctx context.Context) *DeleteBundleParams {
	var ()
	return &DeleteBundleParams{

		Context: ctx,
	}
}

// NewDeleteBundleParamsWithHTTPClient creates a new DeleteBundleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBundleParamsWithHTTPClient(client *http.Client) *DeleteBundleParams {
	var ()
	return &DeleteBundleParams{
		HTTPClient: client,
	}
}

/*DeleteBundleParams contains all the parameters to send to the API endpoint
for the delete bundle operation typically these are written to a http.Request
*/
type DeleteBundleParams struct {

	/*BundleID
	  Content pack ID

	*/
	BundleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete bundle params
func (o *DeleteBundleParams) WithTimeout(timeout time.Duration) *DeleteBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete bundle params
func (o *DeleteBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete bundle params
func (o *DeleteBundleParams) WithContext(ctx context.Context) *DeleteBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete bundle params
func (o *DeleteBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete bundle params
func (o *DeleteBundleParams) WithHTTPClient(client *http.Client) *DeleteBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete bundle params
func (o *DeleteBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleID adds the bundleID to the delete bundle params
func (o *DeleteBundleParams) WithBundleID(bundleID string) *DeleteBundleParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the delete bundle params
func (o *DeleteBundleParams) SetBundleID(bundleID string) {
	o.BundleID = bundleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
