// Code generated by go-swagger; DO NOT EDIT.

package system_cluster_traffic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system cluster traffic API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system cluster traffic API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ElasticsearchStats elasticsearches information

This resource returns information about the Elasticsearch Cluster.
*/
func (a *Client) ElasticsearchStats(params *ElasticsearchStatsParams, authInfo runtime.ClientAuthInfoWriter) (*ElasticsearchStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewElasticsearchStatsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "elasticsearchStats",
		Method:             "GET",
		PathPattern:        "/system/cluster/stats/elasticsearch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ElasticsearchStatsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ElasticsearchStatsOK), nil

}

/*
MongoStats mongos d b information

This resource returns information about MongoDB.
*/
func (a *Client) MongoStats(params *MongoStatsParams, authInfo runtime.ClientAuthInfoWriter) (*MongoStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMongoStatsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mongoStats",
		Method:             "GET",
		PathPattern:        "/system/cluster/stats/mongo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MongoStatsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MongoStatsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}