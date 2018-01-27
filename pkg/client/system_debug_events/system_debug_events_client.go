// Code generated by go-swagger; DO NOT EDIT.

package system_debug_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system debug events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system debug events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetRelevant gets relevant configuration settings and their values
*/
func (a *Client) GetRelevant(params *GetRelevantParams, authInfo runtime.ClientAuthInfoWriter) (*GetRelevantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRelevantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRelevant",
		Method:             "GET",
		PathPattern:        "/system/configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRelevantReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRelevantOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}