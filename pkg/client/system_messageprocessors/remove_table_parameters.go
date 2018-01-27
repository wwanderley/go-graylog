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
)

// NewRemoveTableParams creates a new RemoveTableParams object
// with the default values initialized.
func NewRemoveTableParams() *RemoveTableParams {
	var ()
	return &RemoveTableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveTableParamsWithTimeout creates a new RemoveTableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveTableParamsWithTimeout(timeout time.Duration) *RemoveTableParams {
	var ()
	return &RemoveTableParams{

		timeout: timeout,
	}
}

// NewRemoveTableParamsWithContext creates a new RemoveTableParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveTableParamsWithContext(ctx context.Context) *RemoveTableParams {
	var ()
	return &RemoveTableParams{

		Context: ctx,
	}
}

// NewRemoveTableParamsWithHTTPClient creates a new RemoveTableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveTableParamsWithHTTPClient(client *http.Client) *RemoveTableParams {
	var ()
	return &RemoveTableParams{
		HTTPClient: client,
	}
}

/*RemoveTableParams contains all the parameters to send to the API endpoint
for the remove table operation typically these are written to a http.Request
*/
type RemoveTableParams struct {

	/*IDOrName*/
	IDOrName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove table params
func (o *RemoveTableParams) WithTimeout(timeout time.Duration) *RemoveTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove table params
func (o *RemoveTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove table params
func (o *RemoveTableParams) WithContext(ctx context.Context) *RemoveTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove table params
func (o *RemoveTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove table params
func (o *RemoveTableParams) WithHTTPClient(client *http.Client) *RemoveTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove table params
func (o *RemoveTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDOrName adds the iDOrName to the remove table params
func (o *RemoveTableParams) WithIDOrName(iDOrName string) *RemoveTableParams {
	o.SetIDOrName(iDOrName)
	return o
}

// SetIDOrName adds the idOrName to the remove table params
func (o *RemoveTableParams) SetIDOrName(iDOrName string) {
	o.IDOrName = iDOrName
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param idOrName
	if err := r.SetPathParam("idOrName", o.IDOrName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
