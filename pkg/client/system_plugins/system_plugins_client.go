// Code generated by go-swagger; DO NOT EDIT.

package system_plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system plugins API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system plugins API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Permissions gets all available user permissions
*/
func (a *Client) Permissions(params *PermissionsParams, authInfo runtime.ClientAuthInfoWriter) (*PermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPermissionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "permissions",
		Method:             "GET",
		PathPattern:        "/system/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PermissionsOK), nil

}

/*
ReaderPermissions gets the initial permissions assigned to a reader account
*/
func (a *Client) ReaderPermissions(params *ReaderPermissionsParams, authInfo runtime.ClientAuthInfoWriter) (*ReaderPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReaderPermissionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readerPermissions",
		Method:             "GET",
		PathPattern:        "/system/permissions/reader/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReaderPermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReaderPermissionsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
