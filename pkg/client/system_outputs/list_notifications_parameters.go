// Code generated by go-swagger; DO NOT EDIT.

package system_outputs

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

// NewListNotificationsParams creates a new ListNotificationsParams object
// with the default values initialized.
func NewListNotificationsParams() *ListNotificationsParams {

	return &ListNotificationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListNotificationsParamsWithTimeout creates a new ListNotificationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListNotificationsParamsWithTimeout(timeout time.Duration) *ListNotificationsParams {

	return &ListNotificationsParams{

		timeout: timeout,
	}
}

// NewListNotificationsParamsWithContext creates a new ListNotificationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListNotificationsParamsWithContext(ctx context.Context) *ListNotificationsParams {

	return &ListNotificationsParams{

		Context: ctx,
	}
}

// NewListNotificationsParamsWithHTTPClient creates a new ListNotificationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListNotificationsParamsWithHTTPClient(client *http.Client) *ListNotificationsParams {

	return &ListNotificationsParams{
		HTTPClient: client,
	}
}

/*ListNotificationsParams contains all the parameters to send to the API endpoint
for the list notifications operation typically these are written to a http.Request
*/
type ListNotificationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list notifications params
func (o *ListNotificationsParams) WithTimeout(timeout time.Duration) *ListNotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list notifications params
func (o *ListNotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list notifications params
func (o *ListNotificationsParams) WithContext(ctx context.Context) *ListNotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list notifications params
func (o *ListNotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list notifications params
func (o *ListNotificationsParams) WithHTTPClient(client *http.Client) *ListNotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list notifications params
func (o *ListNotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListNotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
