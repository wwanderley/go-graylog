// Code generated by go-swagger; DO NOT EDIT.

package plugins_map_mapdata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new plugins org graylog plugins map mapdata API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for plugins org graylog plugins map mapdata API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CopyConfiguration copies a configuration

This is a stateless method which copies a collector configuration to one with another name
*/
func (a *Client) CopyConfiguration(params *CopyConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*CopyConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCopyConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "copyConfiguration",
		Method:             "POST",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CopyConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CopyConfigurationOK), nil

}

/*
CopyInput copies a configuration input

This is a stateless method which copies a collector input to one with another name
*/
func (a *Client) CopyInput(params *CopyInputParams, authInfo runtime.ClientAuthInfoWriter) (*CopyInputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCopyInputParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "copyInput",
		Method:             "POST",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{inputId}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CopyInputReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CopyInputOK), nil

}

/*
CopyOutput copies a configuration output

This is a stateless method which copies a collector output to one with another name
*/
func (a *Client) CopyOutput(params *CopyOutputParams, authInfo runtime.ClientAuthInfoWriter) (*CopyOutputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCopyOutputParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "copyOutput",
		Method:             "POST",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{outputId}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CopyOutputReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CopyOutputOK), nil

}

/*
CopySnippet copies a configuration snippet

This is a stateless method which copies a collector snippet to one with another name
*/
func (a *Client) CopySnippet(params *CopySnippetParams, authInfo runtime.ClientAuthInfoWriter) (*CopySnippetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCopySnippetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "copySnippet",
		Method:             "POST",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/snippets/{snippetId}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CopySnippetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CopySnippetOK), nil

}

/*
CreateConfiguration creates new collector configuration
*/
func (a *Client) CreateConfiguration(params *CreateConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createConfiguration",
		Method:             "POST",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateConfigurationOK), nil

}

/*
CreateInput creates a configuration input

This is a stateless method which inserts a collector input
*/
func (a *Client) CreateInput(params *CreateInputParams, authInfo runtime.ClientAuthInfoWriter) (*CreateInputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInputParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createInput",
		Method:             "POST",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/inputs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateInputReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateInputOK), nil

}

/*
CreateOutput creates a configuration output

This is a stateless method which inserts a collector output
*/
func (a *Client) CreateOutput(params *CreateOutputParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOutputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOutputParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOutput",
		Method:             "POST",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/outputs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateOutputReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateOutputOK), nil

}

/*
CreateSnippet creates a configuration snippet

This is a stateless method which inserts a collector configuration snippet
*/
func (a *Client) CreateSnippet(params *CreateSnippetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSnippetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSnippetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSnippet",
		Method:             "POST",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/snippets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSnippetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSnippetOK), nil

}

/*
DeleteConfiguration deletes a collector configuration
*/
func (a *Client) DeleteConfiguration(params *DeleteConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteConfiguration",
		Method:             "DELETE",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteConfigurationOK), nil

}

/*
DeleteInput deletes input form configuration
*/
func (a *Client) DeleteInput(params *DeleteInputParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteInputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInputParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteInput",
		Method:             "DELETE",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{inputId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteInputReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteInputOK), nil

}

/*
DeleteOutput deletes output from configuration
*/
func (a *Client) DeleteOutput(params *DeleteOutputParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOutputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOutputParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOutput",
		Method:             "DELETE",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{outputId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOutputReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOutputOK), nil

}

/*
DeleteSnippet deletes snippet from configuration
*/
func (a *Client) DeleteSnippet(params *DeleteSnippetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSnippetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSnippetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSnippet",
		Method:             "DELETE",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/snippets/{snippetId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSnippetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSnippetOK), nil

}

/*
GetConfiguration gets a single collector configuration
*/
func (a *Client) GetConfiguration(params *GetConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*GetConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConfiguration",
		Method:             "GET",
		PathPattern:        "/plugins/org.graylog.plugins.collector/{collectorId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfigurationOK), nil

}

/*
GetConfigurations shows collector configuration details
*/
func (a *Client) GetConfigurations(params *GetConfigurationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetConfigurationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigurationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConfigurations",
		Method:             "GET",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConfigurationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfigurationsOK), nil

}

/*
GetTags lists all used tags
*/
func (a *Client) GetTags(params *GetTagsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTags",
		Method:             "GET",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTagsOK), nil

}

/*
ListConfigurations lists all collector configurations
*/
func (a *Client) ListConfigurations(params *ListConfigurationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListConfigurationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListConfigurationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listConfigurations",
		Method:             "GET",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListConfigurationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListConfigurationsOK), nil

}

/*
UpdateConfigurationName updates a collector configuration name
*/
func (a *Client) UpdateConfigurationName(params *UpdateConfigurationNameParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateConfigurationNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateConfigurationNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateConfigurationName",
		Method:             "PUT",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateConfigurationNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateConfigurationNameOK), nil

}

/*
UpdateInput updates a configuration input

This is a stateless method which updates a collector input
*/
func (a *Client) UpdateInput(params *UpdateInputParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateInputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateInputParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateInput",
		Method:             "PUT",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{input_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateInputReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateInputOK), nil

}

/*
UpdateOutput updates a configuration output

This is a stateless method which updates a collector output
*/
func (a *Client) UpdateOutput(params *UpdateOutputParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOutputOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOutputParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOutput",
		Method:             "PUT",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{output_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateOutputReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateOutputOK), nil

}

/*
UpdateSnippet updates a configuration snippet

This is a stateless method which updates a collector snippet
*/
func (a *Client) UpdateSnippet(params *UpdateSnippetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSnippetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSnippetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSnippet",
		Method:             "PUT",
		PathPattern:        "/plugins/org.graylog.plugins.collector/configurations/{id}/snippets/{snippet_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSnippetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSnippetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}