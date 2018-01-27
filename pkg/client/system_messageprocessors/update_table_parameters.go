// Code generated by go-swagger; DO NOT EDIT.

package system_messageprocessors

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

// NewUpdateTableParams creates a new UpdateTableParams object
// with the default values initialized.
func NewUpdateTableParams() *UpdateTableParams {
	var ()
	return &UpdateTableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTableParamsWithTimeout creates a new UpdateTableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTableParamsWithTimeout(timeout time.Duration) *UpdateTableParams {
	var ()
	return &UpdateTableParams{

		timeout: timeout,
	}
}

// NewUpdateTableParamsWithContext creates a new UpdateTableParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTableParamsWithContext(ctx context.Context) *UpdateTableParams {
	var ()
	return &UpdateTableParams{

		Context: ctx,
	}
}

// NewUpdateTableParamsWithHTTPClient creates a new UpdateTableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTableParamsWithHTTPClient(client *http.Client) *UpdateTableParams {
	var ()
	return &UpdateTableParams{
		HTTPClient: client,
	}
}

/*UpdateTableParams contains all the parameters to send to the API endpoint
for the update table operation typically these are written to a http.Request
*/
type UpdateTableParams struct {

	/*Body*/
	Body *models.LookupTableAPI
	/*IDOrName*/
	IDOrName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update table params
func (o *UpdateTableParams) WithTimeout(timeout time.Duration) *UpdateTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update table params
func (o *UpdateTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update table params
func (o *UpdateTableParams) WithContext(ctx context.Context) *UpdateTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update table params
func (o *UpdateTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update table params
func (o *UpdateTableParams) WithHTTPClient(client *http.Client) *UpdateTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update table params
func (o *UpdateTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update table params
func (o *UpdateTableParams) WithBody(body *models.LookupTableAPI) *UpdateTableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update table params
func (o *UpdateTableParams) SetBody(body *models.LookupTableAPI) {
	o.Body = body
}

// WithIDOrName adds the iDOrName to the update table params
func (o *UpdateTableParams) WithIDOrName(iDOrName string) *UpdateTableParams {
	o.SetIDOrName(iDOrName)
	return o
}

// SetIDOrName adds the idOrName to the update table params
func (o *UpdateTableParams) SetIDOrName(iDOrName string) {
	o.IDOrName = iDOrName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param idOrName
	if err := r.SetPathParam("idOrName", o.IDOrName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
