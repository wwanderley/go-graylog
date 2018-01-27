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

// NewUpdateSnippetParams creates a new UpdateSnippetParams object
// with the default values initialized.
func NewUpdateSnippetParams() *UpdateSnippetParams {
	var ()
	return &UpdateSnippetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSnippetParamsWithTimeout creates a new UpdateSnippetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSnippetParamsWithTimeout(timeout time.Duration) *UpdateSnippetParams {
	var ()
	return &UpdateSnippetParams{

		timeout: timeout,
	}
}

// NewUpdateSnippetParamsWithContext creates a new UpdateSnippetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSnippetParamsWithContext(ctx context.Context) *UpdateSnippetParams {
	var ()
	return &UpdateSnippetParams{

		Context: ctx,
	}
}

// NewUpdateSnippetParamsWithHTTPClient creates a new UpdateSnippetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSnippetParamsWithHTTPClient(client *http.Client) *UpdateSnippetParams {
	var ()
	return &UpdateSnippetParams{
		HTTPClient: client,
	}
}

/*UpdateSnippetParams contains all the parameters to send to the API endpoint
for the update snippet operation typically these are written to a http.Request
*/
type UpdateSnippetParams struct {

	/*JSONBody*/
	JSONBody interface{}
	/*ID*/
	ID string
	/*SnippetID*/
	SnippetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update snippet params
func (o *UpdateSnippetParams) WithTimeout(timeout time.Duration) *UpdateSnippetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update snippet params
func (o *UpdateSnippetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update snippet params
func (o *UpdateSnippetParams) WithContext(ctx context.Context) *UpdateSnippetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update snippet params
func (o *UpdateSnippetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update snippet params
func (o *UpdateSnippetParams) WithHTTPClient(client *http.Client) *UpdateSnippetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update snippet params
func (o *UpdateSnippetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the update snippet params
func (o *UpdateSnippetParams) WithJSONBody(jSONBody interface{}) *UpdateSnippetParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the update snippet params
func (o *UpdateSnippetParams) SetJSONBody(jSONBody interface{}) {
	o.JSONBody = jSONBody
}

// WithID adds the id to the update snippet params
func (o *UpdateSnippetParams) WithID(id string) *UpdateSnippetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update snippet params
func (o *UpdateSnippetParams) SetID(id string) {
	o.ID = id
}

// WithSnippetID adds the snippetID to the update snippet params
func (o *UpdateSnippetParams) WithSnippetID(snippetID string) *UpdateSnippetParams {
	o.SetSnippetID(snippetID)
	return o
}

// SetSnippetID adds the snippetId to the update snippet params
func (o *UpdateSnippetParams) SetSnippetID(snippetID string) {
	o.SnippetID = snippetID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSnippetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param snippet_id
	if err := r.SetPathParam("snippet_id", o.SnippetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
