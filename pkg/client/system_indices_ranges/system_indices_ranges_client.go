// Code generated by go-swagger; DO NOT EDIT.

package system_indices_ranges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system indices ranges API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system indices ranges API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
BulkUpdatePatterns adds a list of new patterns
*/
func (a *Client) BulkUpdatePatterns(params *BulkUpdatePatternsParams, authInfo runtime.ClientAuthInfoWriter) (*BulkUpdatePatternsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkUpdatePatternsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "bulkUpdatePatterns",
		Method:             "PUT",
		PathPattern:        "/system/grok",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BulkUpdatePatternsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BulkUpdatePatternsOK), nil

}

/*
UpdatePattern updates an existing pattern
*/
func (a *Client) UpdatePattern(params *UpdatePatternParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePatternOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePatternParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePattern",
		Method:             "PUT",
		PathPattern:        "/system/grok/{patternId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePatternReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdatePatternOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
