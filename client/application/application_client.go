// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new application API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for application API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AppConfigDatastoreShow gets bitrise yml of a specific app

Show details of a specific app
*/
func (a *Client) AppConfigDatastoreShow(params *AppConfigDatastoreShowParams, authInfo runtime.ClientAuthInfoWriter) (*AppConfigDatastoreShowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppConfigDatastoreShowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "app-config-datastore-show",
		Method:             "GET",
		PathPattern:        "/apps/{app-slug}/bitrise.yml",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AppConfigDatastoreShowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AppConfigDatastoreShowOK), nil

}

/*
AppList gets list of the apps

List all the available apps
*/
func (a *Client) AppList(params *AppListParams, authInfo runtime.ClientAuthInfoWriter) (*AppListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "app-list",
		Method:             "GET",
		PathPattern:        "/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AppListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AppListOK), nil

}

/*
AppListByOrganization gets list of the apps for an organization

List all the available apps for a specific organization
*/
func (a *Client) AppListByOrganization(params *AppListByOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*AppListByOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppListByOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "app-list-by-organization",
		Method:             "GET",
		PathPattern:        "/organizations/{org-slug}/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AppListByOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AppListByOrganizationOK), nil

}

/*
AppListByUser gets list of the apps for a user

List all the available apps for a specific user
*/
func (a *Client) AppListByUser(params *AppListByUserParams, authInfo runtime.ClientAuthInfoWriter) (*AppListByUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppListByUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "app-list-by-user",
		Method:             "GET",
		PathPattern:        "/users/{user-slug}/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AppListByUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AppListByUserOK), nil

}

/*
AppShow gets a specific app

Show details of a specific app
*/
func (a *Client) AppShow(params *AppShowParams, authInfo runtime.ClientAuthInfoWriter) (*AppShowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAppShowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "app-show",
		Method:             "GET",
		PathPattern:        "/apps/{app-slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AppShowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AppShowOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}