// Code generated by go-swagger; DO NOT EDIT.

package system_metrics_metric_name_history

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

// NewMetricNamesParams creates a new MetricNamesParams object
// with the default values initialized.
func NewMetricNamesParams() *MetricNamesParams {

	return &MetricNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMetricNamesParamsWithTimeout creates a new MetricNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMetricNamesParamsWithTimeout(timeout time.Duration) *MetricNamesParams {

	return &MetricNamesParams{

		timeout: timeout,
	}
}

// NewMetricNamesParamsWithContext creates a new MetricNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewMetricNamesParamsWithContext(ctx context.Context) *MetricNamesParams {

	return &MetricNamesParams{

		Context: ctx,
	}
}

// NewMetricNamesParamsWithHTTPClient creates a new MetricNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMetricNamesParamsWithHTTPClient(client *http.Client) *MetricNamesParams {

	return &MetricNamesParams{
		HTTPClient: client,
	}
}

/*MetricNamesParams contains all the parameters to send to the API endpoint
for the metric names operation typically these are written to a http.Request
*/
type MetricNamesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the metric names params
func (o *MetricNamesParams) WithTimeout(timeout time.Duration) *MetricNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metric names params
func (o *MetricNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metric names params
func (o *MetricNamesParams) WithContext(ctx context.Context) *MetricNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metric names params
func (o *MetricNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metric names params
func (o *MetricNamesParams) WithHTTPClient(client *http.Client) *MetricNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metric names params
func (o *MetricNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MetricNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
