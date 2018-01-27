// Code generated by go-swagger; DO NOT EDIT.

package system_deflector

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

// NewGenerateClusterDebugEventParams creates a new GenerateClusterDebugEventParams object
// with the default values initialized.
func NewGenerateClusterDebugEventParams() *GenerateClusterDebugEventParams {
	var ()
	return &GenerateClusterDebugEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateClusterDebugEventParamsWithTimeout creates a new GenerateClusterDebugEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGenerateClusterDebugEventParamsWithTimeout(timeout time.Duration) *GenerateClusterDebugEventParams {
	var ()
	return &GenerateClusterDebugEventParams{

		timeout: timeout,
	}
}

// NewGenerateClusterDebugEventParamsWithContext creates a new GenerateClusterDebugEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewGenerateClusterDebugEventParamsWithContext(ctx context.Context) *GenerateClusterDebugEventParams {
	var ()
	return &GenerateClusterDebugEventParams{

		Context: ctx,
	}
}

// NewGenerateClusterDebugEventParamsWithHTTPClient creates a new GenerateClusterDebugEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGenerateClusterDebugEventParamsWithHTTPClient(client *http.Client) *GenerateClusterDebugEventParams {
	var ()
	return &GenerateClusterDebugEventParams{
		HTTPClient: client,
	}
}

/*GenerateClusterDebugEventParams contains all the parameters to send to the API endpoint
for the generate cluster debug event operation typically these are written to a http.Request
*/
type GenerateClusterDebugEventParams struct {

	/*Text*/
	Text *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the generate cluster debug event params
func (o *GenerateClusterDebugEventParams) WithTimeout(timeout time.Duration) *GenerateClusterDebugEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate cluster debug event params
func (o *GenerateClusterDebugEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate cluster debug event params
func (o *GenerateClusterDebugEventParams) WithContext(ctx context.Context) *GenerateClusterDebugEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate cluster debug event params
func (o *GenerateClusterDebugEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate cluster debug event params
func (o *GenerateClusterDebugEventParams) WithHTTPClient(client *http.Client) *GenerateClusterDebugEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate cluster debug event params
func (o *GenerateClusterDebugEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithText adds the text to the generate cluster debug event params
func (o *GenerateClusterDebugEventParams) WithText(text *string) *GenerateClusterDebugEventParams {
	o.SetText(text)
	return o
}

// SetText adds the text to the generate cluster debug event params
func (o *GenerateClusterDebugEventParams) SetText(text *string) {
	o.Text = text
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateClusterDebugEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Text != nil {
		if err := r.SetBodyParam(o.Text); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
