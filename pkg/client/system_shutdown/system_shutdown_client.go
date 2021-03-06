// Code generated by go-swagger; DO NOT EDIT.

package system_shutdown

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system shutdown API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system shutdown API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TerminateSession terminates an existing session

Destroys the session with the given ID: the equivalent of logging out.
*/
func (a *Client) TerminateSession(params *TerminateSessionParams, authInfo runtime.ClientAuthInfoWriter) (*TerminateSessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTerminateSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "terminateSession",
		Method:             "DELETE",
		PathPattern:        "/system/sessions/{sessionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TerminateSessionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TerminateSessionOK), nil

}

/*
ValidateSession validates an existing session

Checks the session with the given ID: returns http status 204 (No Content) if session is valid.
*/
func (a *Client) ValidateSession(params *ValidateSessionParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateSessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateSession",
		Method:             "GET",
		PathPattern:        "/system/sessions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidateSessionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValidateSessionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
