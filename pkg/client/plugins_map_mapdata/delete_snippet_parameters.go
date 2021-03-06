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

// NewDeleteSnippetParams creates a new DeleteSnippetParams object
// with the default values initialized.
func NewDeleteSnippetParams() *DeleteSnippetParams {
	var ()
	return &DeleteSnippetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSnippetParamsWithTimeout creates a new DeleteSnippetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSnippetParamsWithTimeout(timeout time.Duration) *DeleteSnippetParams {
	var ()
	return &DeleteSnippetParams{

		timeout: timeout,
	}
}

// NewDeleteSnippetParamsWithContext creates a new DeleteSnippetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSnippetParamsWithContext(ctx context.Context) *DeleteSnippetParams {
	var ()
	return &DeleteSnippetParams{

		Context: ctx,
	}
}

// NewDeleteSnippetParamsWithHTTPClient creates a new DeleteSnippetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSnippetParamsWithHTTPClient(client *http.Client) *DeleteSnippetParams {
	var ()
	return &DeleteSnippetParams{
		HTTPClient: client,
	}
}

/*DeleteSnippetParams contains all the parameters to send to the API endpoint
for the delete snippet operation typically these are written to a http.Request
*/
type DeleteSnippetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete snippet params
func (o *DeleteSnippetParams) WithTimeout(timeout time.Duration) *DeleteSnippetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete snippet params
func (o *DeleteSnippetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete snippet params
func (o *DeleteSnippetParams) WithContext(ctx context.Context) *DeleteSnippetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete snippet params
func (o *DeleteSnippetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete snippet params
func (o *DeleteSnippetParams) WithHTTPClient(client *http.Client) *DeleteSnippetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete snippet params
func (o *DeleteSnippetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete snippet params
func (o *DeleteSnippetParams) WithID(id string) *DeleteSnippetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete snippet params
func (o *DeleteSnippetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSnippetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
