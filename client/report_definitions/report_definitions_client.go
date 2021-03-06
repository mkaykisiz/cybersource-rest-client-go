// Code generated by go-swagger; DO NOT EDIT.

package report_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new report definitions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for report definitions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetResourceInfoByReportDefinition gets report definition

View the attributes of an individual report type.
For a list of values for reportDefinitionName, see the
[Reporting Developer Guide](https://www.cybersource.com/developers/documentation/reporting_and_reconciliation/)

*/
func (a *Client) GetResourceInfoByReportDefinition(params *GetResourceInfoByReportDefinitionParams) (*GetResourceInfoByReportDefinitionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceInfoByReportDefinitionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResourceInfoByReportDefinition",
		Method:             "GET",
		PathPattern:        "/reporting/v3/report-definitions/{reportDefinitionName}",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceInfoByReportDefinitionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetResourceInfoByReportDefinitionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResourceInfoByReportDefinition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetResourceV2Info gets reporting resource information

View a list of supported reports and their attributes before subscribing to them.

*/
func (a *Client) GetResourceV2Info(params *GetResourceV2InfoParams) (*GetResourceV2InfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceV2InfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResourceV2Info",
		Method:             "GET",
		PathPattern:        "/reporting/v3/report-definitions",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceV2InfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetResourceV2InfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResourceV2Info: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
