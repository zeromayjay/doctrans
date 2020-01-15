// Code generated by go-swagger; DO NOT EDIT.

package d_t_a_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new d t a server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for d t a server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ListServices list services API
*/
func (a *Client) ListServices(params *ListServicesParams) (*ListServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListServices",
		Method:             "GET",
		PathPattern:        "/v1/dta/service/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListServicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListServices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TransformDocument requests to transform a plain text document
*/
func (a *Client) TransformDocument(params *TransformDocumentParams) (*TransformDocumentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransformDocumentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TransformDocument",
		Method:             "POST",
		PathPattern:        "/v1/dta/document/transform",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TransformDocumentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TransformDocumentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TransformDocument: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TransformPipe transform pipe API
*/
func (a *Client) TransformPipe(params *TransformPipeParams) (*TransformPipeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransformPipeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TransformPipe",
		Method:             "POST",
		PathPattern:        "/v1/document/transform-pipe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TransformPipeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TransformPipeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for TransformPipe: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}