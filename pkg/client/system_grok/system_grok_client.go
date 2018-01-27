// Code generated by go-swagger; DO NOT EDIT.

package system_grok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system grok API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system grok API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DismissGettingStarted dismisses auto showing getting started guide for this version
*/
func (a *Client) DismissGettingStarted(params *DismissGettingStartedParams, authInfo runtime.ClientAuthInfoWriter) (*DismissGettingStartedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDismissGettingStartedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dismissGettingStarted",
		Method:             "POST",
		PathPattern:        "/system/gettingstarted/dismiss",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DismissGettingStartedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DismissGettingStartedOK), nil

}

/*
DisplayGettingStarted checks whether to display the getting started guide for this version
*/
func (a *Client) DisplayGettingStarted(params *DisplayGettingStartedParams, authInfo runtime.ClientAuthInfoWriter) (*DisplayGettingStartedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisplayGettingStartedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "displayGettingStarted",
		Method:             "GET",
		PathPattern:        "/system/gettingstarted",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisplayGettingStartedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DisplayGettingStartedOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
