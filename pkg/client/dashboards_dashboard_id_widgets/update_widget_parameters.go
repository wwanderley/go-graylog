// Code generated by go-swagger; DO NOT EDIT.

package dashboards_dashboard_id_widgets

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

// NewUpdateWidgetParams creates a new UpdateWidgetParams object
// with the default values initialized.
func NewUpdateWidgetParams() *UpdateWidgetParams {
	var ()
	return &UpdateWidgetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWidgetParamsWithTimeout creates a new UpdateWidgetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateWidgetParamsWithTimeout(timeout time.Duration) *UpdateWidgetParams {
	var ()
	return &UpdateWidgetParams{

		timeout: timeout,
	}
}

// NewUpdateWidgetParamsWithContext creates a new UpdateWidgetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateWidgetParamsWithContext(ctx context.Context) *UpdateWidgetParams {
	var ()
	return &UpdateWidgetParams{

		Context: ctx,
	}
}

// NewUpdateWidgetParamsWithHTTPClient creates a new UpdateWidgetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateWidgetParamsWithHTTPClient(client *http.Client) *UpdateWidgetParams {
	var ()
	return &UpdateWidgetParams{
		HTTPClient: client,
	}
}

/*UpdateWidgetParams contains all the parameters to send to the API endpoint
for the update widget operation typically these are written to a http.Request
*/
type UpdateWidgetParams struct {

	/*JSONBody*/
	JSONBody *models.AddWidgetRequest
	/*DashboardID*/
	DashboardID string
	/*WidgetID*/
	WidgetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update widget params
func (o *UpdateWidgetParams) WithTimeout(timeout time.Duration) *UpdateWidgetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update widget params
func (o *UpdateWidgetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update widget params
func (o *UpdateWidgetParams) WithContext(ctx context.Context) *UpdateWidgetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update widget params
func (o *UpdateWidgetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update widget params
func (o *UpdateWidgetParams) WithHTTPClient(client *http.Client) *UpdateWidgetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update widget params
func (o *UpdateWidgetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the update widget params
func (o *UpdateWidgetParams) WithJSONBody(jSONBody *models.AddWidgetRequest) *UpdateWidgetParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the update widget params
func (o *UpdateWidgetParams) SetJSONBody(jSONBody *models.AddWidgetRequest) {
	o.JSONBody = jSONBody
}

// WithDashboardID adds the dashboardID to the update widget params
func (o *UpdateWidgetParams) WithDashboardID(dashboardID string) *UpdateWidgetParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the update widget params
func (o *UpdateWidgetParams) SetDashboardID(dashboardID string) {
	o.DashboardID = dashboardID
}

// WithWidgetID adds the widgetID to the update widget params
func (o *UpdateWidgetParams) WithWidgetID(widgetID string) *UpdateWidgetParams {
	o.SetWidgetID(widgetID)
	return o
}

// SetWidgetID adds the widgetId to the update widget params
func (o *UpdateWidgetParams) SetWidgetID(widgetID string) {
	o.WidgetID = widgetID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWidgetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	// path param dashboardId
	if err := r.SetPathParam("dashboardId", o.DashboardID); err != nil {
		return err
	}

	// path param widgetId
	if err := r.SetPathParam("widgetId", o.WidgetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}