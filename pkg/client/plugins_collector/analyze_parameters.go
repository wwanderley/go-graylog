// Code generated by go-swagger; DO NOT EDIT.

package plugins_collector

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

// NewAnalyzeParams creates a new AnalyzeParams object
// with the default values initialized.
func NewAnalyzeParams() *AnalyzeParams {
	var ()
	return &AnalyzeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAnalyzeParamsWithTimeout creates a new AnalyzeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAnalyzeParamsWithTimeout(timeout time.Duration) *AnalyzeParams {
	var ()
	return &AnalyzeParams{

		timeout: timeout,
	}
}

// NewAnalyzeParamsWithContext creates a new AnalyzeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAnalyzeParamsWithContext(ctx context.Context) *AnalyzeParams {
	var ()
	return &AnalyzeParams{

		Context: ctx,
	}
}

// NewAnalyzeParamsWithHTTPClient creates a new AnalyzeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAnalyzeParamsWithHTTPClient(client *http.Client) *AnalyzeParams {
	var ()
	return &AnalyzeParams{
		HTTPClient: client,
	}
}

/*AnalyzeParams contains all the parameters to send to the API endpoint
for the analyze operation typically these are written to a http.Request
*/
type AnalyzeParams struct {

	/*Analyzer
	  The analyzer to use.

	*/
	Analyzer *string
	/*Index
	  The index the message containing the string is stored in.

	*/
	Index string
	/*String
	  The string to analyze.

	*/
	String string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the analyze params
func (o *AnalyzeParams) WithTimeout(timeout time.Duration) *AnalyzeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the analyze params
func (o *AnalyzeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the analyze params
func (o *AnalyzeParams) WithContext(ctx context.Context) *AnalyzeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the analyze params
func (o *AnalyzeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the analyze params
func (o *AnalyzeParams) WithHTTPClient(client *http.Client) *AnalyzeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the analyze params
func (o *AnalyzeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAnalyzer adds the analyzer to the analyze params
func (o *AnalyzeParams) WithAnalyzer(analyzer *string) *AnalyzeParams {
	o.SetAnalyzer(analyzer)
	return o
}

// SetAnalyzer adds the analyzer to the analyze params
func (o *AnalyzeParams) SetAnalyzer(analyzer *string) {
	o.Analyzer = analyzer
}

// WithIndex adds the index to the analyze params
func (o *AnalyzeParams) WithIndex(index string) *AnalyzeParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the analyze params
func (o *AnalyzeParams) SetIndex(index string) {
	o.Index = index
}

// WithString adds the string to the analyze params
func (o *AnalyzeParams) WithString(string string) *AnalyzeParams {
	o.SetString(string)
	return o
}

// SetString adds the string to the analyze params
func (o *AnalyzeParams) SetString(string string) {
	o.String = string
}

// WriteToRequest writes these params to a swagger request
func (o *AnalyzeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Analyzer != nil {

		// query param analyzer
		var qrAnalyzer string
		if o.Analyzer != nil {
			qrAnalyzer = *o.Analyzer
		}
		qAnalyzer := qrAnalyzer
		if qAnalyzer != "" {
			if err := r.SetQueryParam("analyzer", qAnalyzer); err != nil {
				return err
			}
		}

	}

	// path param index
	if err := r.SetPathParam("index", o.Index); err != nil {
		return err
	}

	// query param string
	qrString := o.String
	qString := qrString
	if qString != "" {
		if err := r.SetQueryParam("string", qString); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}