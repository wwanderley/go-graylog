// Code generated by go-swagger; DO NOT EDIT.

package dashboards_dashboard_id_widgets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new dashboards dashboard id widgets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dashboards dashboard id widgets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddWidget adds a widget to a dashboard
*/
func (a *Client) AddWidget(params *AddWidgetParams, authInfo runtime.ClientAuthInfoWriter) (*AddWidgetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddWidgetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addWidget",
		Method:             "POST",
		PathPattern:        "/dashboards/{dashboardId}/widgets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddWidgetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddWidgetOK), nil

}

/*
UpdateCacheTime updates cache time of a widget
*/
func (a *Client) UpdateCacheTime(params *UpdateCacheTimeParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCacheTimeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCacheTimeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCacheTime",
		Method:             "PUT",
		PathPattern:        "/dashboards/{dashboardId}/widgets/{widgetId}/cachetime",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCacheTimeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateCacheTimeOK), nil

}

/*
UpdateDescription updates description of a widget
*/
func (a *Client) UpdateDescription(params *UpdateDescriptionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDescriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDescriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDescription",
		Method:             "PUT",
		PathPattern:        "/dashboards/{dashboardId}/widgets/{widgetId}/description",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDescriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateDescriptionOK), nil

}

/*
UpdateWidget updates a widget
*/
func (a *Client) UpdateWidget(params *UpdateWidgetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateWidgetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWidgetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateWidget",
		Method:             "PUT",
		PathPattern:        "/dashboards/{dashboardId}/widgets/{widgetId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateWidgetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateWidgetOK), nil

}

/*
WidgetValue gets a single widget value
*/
func (a *Client) WidgetValue(params *WidgetValueParams, authInfo runtime.ClientAuthInfoWriter) (*WidgetValueOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWidgetValueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "widgetValue",
		Method:             "GET",
		PathPattern:        "/dashboards/{dashboardId}/widgets/{widgetId}/value",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WidgetValueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WidgetValueOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
