// Code generated by go-swagger; DO NOT EDIT.

package streams_streamid_alerts_alert_id_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new streams streamid alerts alert id history API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for streams streamid alerts alert id history API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetForAlert gets a list of all alarm callbacks for this stream
*/
func (a *Client) GetForAlert(params *GetForAlertParams, authInfo runtime.ClientAuthInfoWriter) (*GetForAlertOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetForAlertParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getForAlert",
		Method:             "GET",
		PathPattern:        "/streams/{streamid}/alerts/{alertId}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetForAlertReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetForAlertOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
