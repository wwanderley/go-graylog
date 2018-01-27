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

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// NewUpdateBundleParams creates a new UpdateBundleParams object
// with the default values initialized.
func NewUpdateBundleParams() *UpdateBundleParams {
	var ()
	return &UpdateBundleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBundleParamsWithTimeout creates a new UpdateBundleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateBundleParamsWithTimeout(timeout time.Duration) *UpdateBundleParams {
	var ()
	return &UpdateBundleParams{

		timeout: timeout,
	}
}

// NewUpdateBundleParamsWithContext creates a new UpdateBundleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateBundleParamsWithContext(ctx context.Context) *UpdateBundleParams {
	var ()
	return &UpdateBundleParams{

		Context: ctx,
	}
}

// NewUpdateBundleParamsWithHTTPClient creates a new UpdateBundleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateBundleParamsWithHTTPClient(client *http.Client) *UpdateBundleParams {
	var ()
	return &UpdateBundleParams{
		HTTPClient: client,
	}
}

/*UpdateBundleParams contains all the parameters to send to the API endpoint
for the update bundle operation typically these are written to a http.Request
*/
type UpdateBundleParams struct {

	/*RequestBody
	  Content pack

	*/
	RequestBody *models.ConfigurationBundle
	/*BundleID
	  Content pack ID

	*/
	BundleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update bundle params
func (o *UpdateBundleParams) WithTimeout(timeout time.Duration) *UpdateBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update bundle params
func (o *UpdateBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update bundle params
func (o *UpdateBundleParams) WithContext(ctx context.Context) *UpdateBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update bundle params
func (o *UpdateBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update bundle params
func (o *UpdateBundleParams) WithHTTPClient(client *http.Client) *UpdateBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update bundle params
func (o *UpdateBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update bundle params
func (o *UpdateBundleParams) WithRequestBody(requestBody *models.ConfigurationBundle) *UpdateBundleParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update bundle params
func (o *UpdateBundleParams) SetRequestBody(requestBody *models.ConfigurationBundle) {
	o.RequestBody = requestBody
}

// WithBundleID adds the bundleID to the update bundle params
func (o *UpdateBundleParams) WithBundleID(bundleID string) *UpdateBundleParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the update bundle params
func (o *UpdateBundleParams) SetBundleID(bundleID string) {
	o.BundleID = bundleID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RequestBody != nil {
		if err := r.SetBodyParam(o.RequestBody); err != nil {
			return err
		}
	}

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
