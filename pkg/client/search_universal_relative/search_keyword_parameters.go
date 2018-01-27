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

// NewSearchKeywordParams creates a new SearchKeywordParams object
// with the default values initialized.
func NewSearchKeywordParams() *SearchKeywordParams {
	var (
		decorateDefault = bool(true)
	)
	return &SearchKeywordParams{
		Decorate: &decorateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchKeywordParamsWithTimeout creates a new SearchKeywordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchKeywordParamsWithTimeout(timeout time.Duration) *SearchKeywordParams {
	var (
		decorateDefault = bool(true)
	)
	return &SearchKeywordParams{
		Decorate: &decorateDefault,

		timeout: timeout,
	}
}

// NewSearchKeywordParamsWithContext creates a new SearchKeywordParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchKeywordParamsWithContext(ctx context.Context) *SearchKeywordParams {
	var (
		decorateDefault = bool(true)
	)
	return &SearchKeywordParams{
		Decorate: &decorateDefault,

		Context: ctx,
	}
}

// NewSearchKeywordParamsWithHTTPClient creates a new SearchKeywordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchKeywordParamsWithHTTPClient(client *http.Client) *SearchKeywordParams {
	var (
		decorateDefault = bool(true)
	)
	return &SearchKeywordParams{
		Decorate:   &decorateDefault,
		HTTPClient: client,
	}
}

/*SearchKeywordParams contains all the parameters to send to the API endpoint
for the search keyword operation typically these are written to a http.Request
*/
type SearchKeywordParams struct {

	/*Decorate
	  Run decorators on search result

	*/
	Decorate *bool
	/*Fields
	  Comma separated list of fields to return

	*/
	Fields *string
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
	/*Sort
	  Sorting (field:asc / field:desc)

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search keyword params
func (o *SearchKeywordParams) WithTimeout(timeout time.Duration) *SearchKeywordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search keyword params
func (o *SearchKeywordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search keyword params
func (o *SearchKeywordParams) WithContext(ctx context.Context) *SearchKeywordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search keyword params
func (o *SearchKeywordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search keyword params
func (o *SearchKeywordParams) WithHTTPClient(client *http.Client) *SearchKeywordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search keyword params
func (o *SearchKeywordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDecorate adds the decorate to the search keyword params
func (o *SearchKeywordParams) WithDecorate(decorate *bool) *SearchKeywordParams {
	o.SetDecorate(decorate)
	return o
}

// SetDecorate adds the decorate to the search keyword params
func (o *SearchKeywordParams) SetDecorate(decorate *bool) {
	o.Decorate = decorate
}

// WithFields adds the fields to the search keyword params
func (o *SearchKeywordParams) WithFields(fields *string) *SearchKeywordParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search keyword params
func (o *SearchKeywordParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the search keyword params
func (o *SearchKeywordParams) WithFilter(filter *string) *SearchKeywordParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the search keyword params
func (o *SearchKeywordParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithKeyword adds the keyword to the search keyword params
func (o *SearchKeywordParams) WithKeyword(keyword string) *SearchKeywordParams {
	o.SetKeyword(keyword)
	return o
}

// SetKeyword adds the keyword to the search keyword params
func (o *SearchKeywordParams) SetKeyword(keyword string) {
	o.Keyword = keyword
}

// WithLimit adds the limit to the search keyword params
func (o *SearchKeywordParams) WithLimit(limit *int64) *SearchKeywordParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search keyword params
func (o *SearchKeywordParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search keyword params
func (o *SearchKeywordParams) WithOffset(offset *int64) *SearchKeywordParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search keyword params
func (o *SearchKeywordParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQuery adds the query to the search keyword params
func (o *SearchKeywordParams) WithQuery(query string) *SearchKeywordParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the search keyword params
func (o *SearchKeywordParams) SetQuery(query string) {
	o.Query = query
}

// WithSort adds the sort to the search keyword params
func (o *SearchKeywordParams) WithSort(sort *string) *SearchKeywordParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search keyword params
func (o *SearchKeywordParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *SearchKeywordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Decorate != nil {

		// query param decorate
		var qrDecorate bool
		if o.Decorate != nil {
			qrDecorate = *o.Decorate
		}
		qDecorate := swag.FormatBool(qrDecorate)
		if qDecorate != "" {
			if err := r.SetQueryParam("decorate", qDecorate); err != nil {
				return err
			}
		}

	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
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

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
