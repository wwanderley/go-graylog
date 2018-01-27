// Code generated by go-swagger; DO NOT EDIT.

package streams_streamid_alerts_alert_id_history

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

// NewGetForAlertParams creates a new GetForAlertParams object
// with the default values initialized.
func NewGetForAlertParams() *GetForAlertParams {
	var ()
	return &GetForAlertParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetForAlertParamsWithTimeout creates a new GetForAlertParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetForAlertParamsWithTimeout(timeout time.Duration) *GetForAlertParams {
	var ()
	return &GetForAlertParams{

		timeout: timeout,
	}
}

// NewGetForAlertParamsWithContext creates a new GetForAlertParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetForAlertParamsWithContext(ctx context.Context) *GetForAlertParams {
	var ()
	return &GetForAlertParams{

		Context: ctx,
	}
}

// NewGetForAlertParamsWithHTTPClient creates a new GetForAlertParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetForAlertParamsWithHTTPClient(client *http.Client) *GetForAlertParams {
	var ()
	return &GetForAlertParams{
		HTTPClient: client,
	}
}

/*GetForAlertParams contains all the parameters to send to the API endpoint
for the get for alert operation typically these are written to a http.Request
*/
type GetForAlertParams struct {

	/*AlertID
	  The id of the alert whose callback history we want.

	*/
	AlertID string
	/*Streamid
	  The id of the stream whose alarm callbacks history we want.

	*/
	Streamid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get for alert params
func (o *GetForAlertParams) WithTimeout(timeout time.Duration) *GetForAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get for alert params
func (o *GetForAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get for alert params
func (o *GetForAlertParams) WithContext(ctx context.Context) *GetForAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get for alert params
func (o *GetForAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get for alert params
func (o *GetForAlertParams) WithHTTPClient(client *http.Client) *GetForAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get for alert params
func (o *GetForAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertID adds the alertID to the get for alert params
func (o *GetForAlertParams) WithAlertID(alertID string) *GetForAlertParams {
	o.SetAlertID(alertID)
	return o
}

// SetAlertID adds the alertId to the get for alert params
func (o *GetForAlertParams) SetAlertID(alertID string) {
	o.AlertID = alertID
}

// WithStreamid adds the streamid to the get for alert params
func (o *GetForAlertParams) WithStreamid(streamid string) *GetForAlertParams {
	o.SetStreamid(streamid)
	return o
}

// SetStreamid adds the streamid to the get for alert params
func (o *GetForAlertParams) SetStreamid(streamid string) {
	o.Streamid = streamid
}

// WriteToRequest writes these params to a swagger request
func (o *GetForAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertId
	if err := r.SetPathParam("alertId", o.AlertID); err != nil {
		return err
	}

	// path param streamid
	if err := r.SetPathParam("streamid", o.Streamid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
