// Code generated by go-swagger; DO NOT EDIT.

package system_lookup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewMessagesParams creates a new MessagesParams object
// with the default values initialized.
func NewMessagesParams() *MessagesParams {
	var (
		levelDefault = string("ALL")
		limitDefault = int64(500)
	)
	return &MessagesParams{
		Level: &levelDefault,
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMessagesParamsWithTimeout creates a new MessagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMessagesParamsWithTimeout(timeout time.Duration) *MessagesParams {
	var (
		levelDefault = string("ALL")
		limitDefault = int64(500)
	)
	return &MessagesParams{
		Level: &levelDefault,
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewMessagesParamsWithContext creates a new MessagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewMessagesParamsWithContext(ctx context.Context) *MessagesParams {
	var (
		levelDefault = string("ALL")
		limitDefault = int64(500)
	)
	return &MessagesParams{
		Level: &levelDefault,
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewMessagesParamsWithHTTPClient creates a new MessagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMessagesParamsWithHTTPClient(client *http.Client) *MessagesParams {
	var (
		levelDefault = string("ALL")
		limitDefault = int64(500)
	)
	return &MessagesParams{
		Level:      &levelDefault,
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*MessagesParams contains all the parameters to send to the API endpoint
for the messages operation typically these are written to a http.Request
*/
type MessagesParams struct {

	/*Level
	  Which log level (or higher) should the messages have

	*/
	Level *string
	/*Limit
	  How many log messages should be returned

	*/
	Limit *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the messages params
func (o *MessagesParams) WithTimeout(timeout time.Duration) *MessagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the messages params
func (o *MessagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the messages params
func (o *MessagesParams) WithContext(ctx context.Context) *MessagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the messages params
func (o *MessagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the messages params
func (o *MessagesParams) WithHTTPClient(client *http.Client) *MessagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the messages params
func (o *MessagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLevel adds the level to the messages params
func (o *MessagesParams) WithLevel(level *string) *MessagesParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the messages params
func (o *MessagesParams) SetLevel(level *string) {
	o.Level = level
}

// WithLimit adds the limit to the messages params
func (o *MessagesParams) WithLimit(limit *int64) *MessagesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the messages params
func (o *MessagesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *MessagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Level != nil {

		// query param level
		var qrLevel string
		if o.Level != nil {
			qrLevel = *o.Level
		}
		qLevel := qrLevel
		if qLevel != "" {
			if err := r.SetQueryParam("level", qLevel); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
