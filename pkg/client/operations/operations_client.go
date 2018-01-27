// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChangePassword updates the password for a user
*/
func (a *Client) ChangePassword(params *ChangePasswordParams, authInfo runtime.ClientAuthInfoWriter) (*ChangePasswordOK, *ChangePasswordNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changePassword",
		Method:             "PUT",
		PathPattern:        "/users/{username}/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChangePasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ChangePasswordOK:
		return value, nil, nil
	case *ChangePasswordNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
ChangeUser modifies user details
*/
func (a *Client) ChangeUser(params *ChangeUserParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeUser",
		Method:             "PUT",
		PathPattern:        "/users/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChangeUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeUserOK), nil

}

/*
EditPermissions updates a user s permission set
*/
func (a *Client) EditPermissions(params *EditPermissionsParams, authInfo runtime.ClientAuthInfoWriter) (*EditPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditPermissionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editPermissions",
		Method:             "PUT",
		PathPattern:        "/users/{username}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EditPermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditPermissionsOK), nil

}

/*
GenerateNewToken generates a new access token for a user
*/
func (a *Client) GenerateNewToken(params *GenerateNewTokenParams, authInfo runtime.ClientAuthInfoWriter) (*GenerateNewTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateNewTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "generateNewToken",
		Method:             "POST",
		PathPattern:        "/users/{username}/tokens/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GenerateNewTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenerateNewTokenOK), nil

}

/*
ListTokens retrieves the list of access tokens for a user
*/
func (a *Client) ListTokens(params *ListTokensParams, authInfo runtime.ClientAuthInfoWriter) (*ListTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTokensParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listTokens",
		Method:             "GET",
		PathPattern:        "/users/{username}/tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListTokensReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListTokensOK), nil

}

/*
ListUsers lists all users

The permissions assigned to the users are always included.
*/
func (a *Client) ListUsers(params *ListUsersParams, authInfo runtime.ClientAuthInfoWriter) (*ListUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListUsersOK), nil

}

/*
RevokeToken removes a token for a user
*/
func (a *Client) RevokeToken(params *RevokeTokenParams, authInfo runtime.ClientAuthInfoWriter) (*RevokeTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "revokeToken",
		Method:             "DELETE",
		PathPattern:        "/users/{username}/tokens/{token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RevokeTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RevokeTokenOK), nil

}

/*
SavePreferences updates a user s preferences set
*/
func (a *Client) SavePreferences(params *SavePreferencesParams, authInfo runtime.ClientAuthInfoWriter) (*SavePreferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSavePreferencesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "savePreferences",
		Method:             "PUT",
		PathPattern:        "/users/{username}/preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SavePreferencesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SavePreferencesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
