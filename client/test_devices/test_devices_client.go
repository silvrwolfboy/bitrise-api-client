// Code generated by go-swagger; DO NOT EDIT.

package test_devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new test devices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for test devices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TestDeviceList lists the test devices for an app

List registered test devices of all members of a specified Bitrise app
*/
func (a *Client) TestDeviceList(params *TestDeviceListParams, authInfo runtime.ClientAuthInfoWriter) (*TestDeviceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestDeviceListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "test-device-list",
		Method:             "GET",
		PathPattern:        "/apps/{app-slug}/test-devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestDeviceListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TestDeviceListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}