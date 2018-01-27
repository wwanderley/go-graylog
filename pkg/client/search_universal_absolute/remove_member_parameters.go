// Code generated by go-swagger; DO NOT EDIT.

package search_universal_absolute

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

// NewRemoveMemberParams creates a new RemoveMemberParams object
// with the default values initialized.
func NewRemoveMemberParams() *RemoveMemberParams {
	var ()
	return &RemoveMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveMemberParamsWithTimeout creates a new RemoveMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveMemberParamsWithTimeout(timeout time.Duration) *RemoveMemberParams {
	var ()
	return &RemoveMemberParams{

		timeout: timeout,
	}
}

// NewRemoveMemberParamsWithContext creates a new RemoveMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveMemberParamsWithContext(ctx context.Context) *RemoveMemberParams {
	var ()
	return &RemoveMemberParams{

		Context: ctx,
	}
}

// NewRemoveMemberParamsWithHTTPClient creates a new RemoveMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveMemberParamsWithHTTPClient(client *http.Client) *RemoveMemberParams {
	var ()
	return &RemoveMemberParams{
		HTTPClient: client,
	}
}

/*RemoveMemberParams contains all the parameters to send to the API endpoint
for the remove member operation typically these are written to a http.Request
*/
type RemoveMemberParams struct {

	/*Rolename*/
	Rolename string
	/*Username*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove member params
func (o *RemoveMemberParams) WithTimeout(timeout time.Duration) *RemoveMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove member params
func (o *RemoveMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove member params
func (o *RemoveMemberParams) WithContext(ctx context.Context) *RemoveMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove member params
func (o *RemoveMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove member params
func (o *RemoveMemberParams) WithHTTPClient(client *http.Client) *RemoveMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove member params
func (o *RemoveMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRolename adds the rolename to the remove member params
func (o *RemoveMemberParams) WithRolename(rolename string) *RemoveMemberParams {
	o.SetRolename(rolename)
	return o
}

// SetRolename adds the rolename to the remove member params
func (o *RemoveMemberParams) SetRolename(rolename string) {
	o.Rolename = rolename
}

// WithUsername adds the username to the remove member params
func (o *RemoveMemberParams) WithUsername(username string) *RemoveMemberParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the remove member params
func (o *RemoveMemberParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rolename
	if err := r.SetPathParam("rolename", o.Rolename); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}