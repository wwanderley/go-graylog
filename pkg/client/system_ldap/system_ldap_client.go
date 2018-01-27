// Code generated by go-swagger; DO NOT EDIT.

package system_ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system ldap API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system ldap API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Show gets current state of the journal on this node
*/
func (a *Client) Show(params *ShowParams, authInfo runtime.ClientAuthInfoWriter) (*ShowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "show",
		Method:             "GET",
		PathPattern:        "/system/journal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ShowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ShowOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}