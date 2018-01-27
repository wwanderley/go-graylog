// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ThreadDump gets a thread dump of the given node
*/
func (a *Client) ThreadDump(params *ThreadDumpParams, authInfo runtime.ClientAuthInfoWriter) (*ThreadDumpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewThreadDumpParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "threadDump",
		Method:             "GET",
		PathPattern:        "/cluster/{nodeId}/threaddump",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ThreadDumpReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ThreadDumpOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
