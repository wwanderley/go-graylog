// Code generated by go-swagger; DO NOT EDIT.

package plugins_system_pipelines_simulate

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

// NewCreateFromParserParams creates a new CreateFromParserParams object
// with the default values initialized.
func NewCreateFromParserParams() *CreateFromParserParams {
	var ()
	return &CreateFromParserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFromParserParamsWithTimeout creates a new CreateFromParserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateFromParserParamsWithTimeout(timeout time.Duration) *CreateFromParserParams {
	var ()
	return &CreateFromParserParams{

		timeout: timeout,
	}
}

// NewCreateFromParserParamsWithContext creates a new CreateFromParserParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateFromParserParamsWithContext(ctx context.Context) *CreateFromParserParams {
	var ()
	return &CreateFromParserParams{

		Context: ctx,
	}
}

// NewCreateFromParserParamsWithHTTPClient creates a new CreateFromParserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateFromParserParamsWithHTTPClient(client *http.Client) *CreateFromParserParams {
	var ()
	return &CreateFromParserParams{
		HTTPClient: client,
	}
}

/*CreateFromParserParams contains all the parameters to send to the API endpoint
for the create from parser operation typically these are written to a http.Request
*/
type CreateFromParserParams struct {

	/*Rule*/
	Rule *models.RuleSource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create from parser params
func (o *CreateFromParserParams) WithTimeout(timeout time.Duration) *CreateFromParserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create from parser params
func (o *CreateFromParserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create from parser params
func (o *CreateFromParserParams) WithContext(ctx context.Context) *CreateFromParserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create from parser params
func (o *CreateFromParserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create from parser params
func (o *CreateFromParserParams) WithHTTPClient(client *http.Client) *CreateFromParserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create from parser params
func (o *CreateFromParserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRule adds the rule to the create from parser params
func (o *CreateFromParserParams) WithRule(rule *models.RuleSource) *CreateFromParserParams {
	o.SetRule(rule)
	return o
}

// SetRule adds the rule to the create from parser params
func (o *CreateFromParserParams) SetRule(rule *models.RuleSource) {
	o.Rule = rule
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFromParserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Rule != nil {
		if err := r.SetBodyParam(o.Rule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
