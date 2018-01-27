// Code generated by go-swagger; DO NOT EDIT.

package system_journal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system journal API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system journal API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Trigger triggers new job
*/
func (a *Client) Trigger(params *TriggerParams, authInfo runtime.ClientAuthInfoWriter) (*TriggerOK, *TriggerAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTriggerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "trigger",
		Method:             "POST",
		PathPattern:        "/system/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TriggerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *TriggerOK:
		return value, nil, nil
	case *TriggerAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
