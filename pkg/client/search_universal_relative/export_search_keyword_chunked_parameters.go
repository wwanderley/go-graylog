// Code generated by go-swagger; DO NOT EDIT.

package search_universal_relative

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

// NewExportSearchKeywordChunkedParams creates a new ExportSearchKeywordChunkedParams object
// with the default values initialized.
func NewExportSearchKeywordChunkedParams() *ExportSearchKeywordChunkedParams {
	var ()
	return &ExportSearchKeywordChunkedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExportSearchKeywordChunkedParamsWithTimeout creates a new ExportSearchKeywordChunkedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExportSearchKeywordChunkedParamsWithTimeout(timeout time.Duration) *ExportSearchKeywordChunkedParams {
	var ()
	return &ExportSearchKeywordChunkedParams{

		timeout: timeout,
	}
}

// NewExportSearchKeywordChunkedParamsWithContext creates a new ExportSearchKeywordChunkedParams object
// with the default values initialized, and the ability to set a context for a request
func NewExportSearchKeywordChunkedParamsWithContext(ctx context.Context) *ExportSearchKeywordChunkedParams {
	var ()
	return &ExportSearchKeywordChunkedParams{

		Context: ctx,
	}
}

// NewExportSearchKeywordChunkedParamsWithHTTPClient creates a new ExportSearchKeywordChunkedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExportSearchKeywordChunkedParamsWithHTTPClient(client *http.Client) *ExportSearchKeywordChunkedParams {
	var ()
	return &ExportSearchKeywordChunkedParams{
		HTTPClient: client,
	}
}

/*ExportSearchKeywordChunkedParams contains all the parameters to send to the API endpoint
for the export search keyword chunked operation typically these are written to a http.Request
*/
type ExportSearchKeywordChunkedParams struct {

	/*Fields
	  Comma separated list of fields to return

	*/
	Fields string
	/*Filter
	  Filter

	*/
	Filter *string
	/*Keyword
	  Range keyword

	*/
	Keyword string
	/*Limit
	  Maximum number of messages to return.

	*/
	Limit *int64
	/*Offset
	  Offset

	*/
	Offset *int64
	/*Query
	  Query (Lucene syntax)

	*/
	Query string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) WithTimeout(timeout time.Duration) *ExportSearchKeywordChunkedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) WithContext(ctx context.Context) *ExportSearchKeywordChunkedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) WithHTTPClient(client *http.Client) *ExportSearchKeywordChunkedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) WithFields(fields string) *ExportSearchKeywordChunkedParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) SetFields(fields string) {
	o.Fields = fields
}

// WithFilter adds the filter to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) WithFilter(filter *string) *ExportSearchKeywordChunkedParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithKeyword adds the keyword to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) WithKeyword(keyword string) *ExportSearchKeywordChunkedParams {
	o.SetKeyword(keyword)
	return o
}

// SetKeyword adds the keyword to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) SetKeyword(keyword string) {
	o.Keyword = keyword
}

// WithLimit adds the limit to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) WithLimit(limit *int64) *ExportSearchKeywordChunkedParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) WithOffset(offset *int64) *ExportSearchKeywordChunkedParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQuery adds the query to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) WithQuery(query string) *ExportSearchKeywordChunkedParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the export search keyword chunked params
func (o *ExportSearchKeywordChunkedParams) SetQuery(query string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *ExportSearchKeywordChunkedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param fields
	qrFields := o.Fields
	qFields := qrFields
	if qFields != "" {
		if err := r.SetQueryParam("fields", qFields); err != nil {
			return err
		}
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	// query param keyword
	qrKeyword := o.Keyword
	qKeyword := qrKeyword
	if qKeyword != "" {
		if err := r.SetQueryParam("keyword", qKeyword); err != nil {
			return err
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

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// query param query
	qrQuery := o.Query
	qQuery := qrQuery
	if qQuery != "" {
		if err := r.SetQueryParam("query", qQuery); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
