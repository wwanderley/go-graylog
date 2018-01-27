// Code generated by go-swagger; DO NOT EDIT.

package plugins_system_pipelines_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new plugins org graylog plugins pipelineprocessor system pipelines connections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for plugins org graylog plugins pipelineprocessor system pipelines connections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
MapData gets map data
*/
func (a *Client) MapData(params *MapDataParams, authInfo runtime.ClientAuthInfoWriter) (*MapDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMapDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mapData",
		Method:             "POST",
		PathPattern:        "/plugins/org.graylog.plugins.map/mapdata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MapDataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MapDataOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
