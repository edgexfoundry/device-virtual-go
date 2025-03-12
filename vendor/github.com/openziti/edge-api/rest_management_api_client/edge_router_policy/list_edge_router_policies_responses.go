// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package edge_router_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge-api/rest_model"
)

// ListEdgeRouterPoliciesReader is a Reader for the ListEdgeRouterPolicies structure.
type ListEdgeRouterPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEdgeRouterPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEdgeRouterPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListEdgeRouterPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListEdgeRouterPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListEdgeRouterPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListEdgeRouterPoliciesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListEdgeRouterPoliciesOK creates a ListEdgeRouterPoliciesOK with default headers values
func NewListEdgeRouterPoliciesOK() *ListEdgeRouterPoliciesOK {
	return &ListEdgeRouterPoliciesOK{}
}

/* ListEdgeRouterPoliciesOK describes a response with status code 200, with default header values.

A list of edge router policies
*/
type ListEdgeRouterPoliciesOK struct {
	Payload *rest_model.ListEdgeRouterPoliciesEnvelope
}

func (o *ListEdgeRouterPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies][%d] listEdgeRouterPoliciesOK  %+v", 200, o.Payload)
}
func (o *ListEdgeRouterPoliciesOK) GetPayload() *rest_model.ListEdgeRouterPoliciesEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListEdgeRouterPoliciesEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEdgeRouterPoliciesBadRequest creates a ListEdgeRouterPoliciesBadRequest with default headers values
func NewListEdgeRouterPoliciesBadRequest() *ListEdgeRouterPoliciesBadRequest {
	return &ListEdgeRouterPoliciesBadRequest{}
}

/* ListEdgeRouterPoliciesBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListEdgeRouterPoliciesBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListEdgeRouterPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies][%d] listEdgeRouterPoliciesBadRequest  %+v", 400, o.Payload)
}
func (o *ListEdgeRouterPoliciesBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEdgeRouterPoliciesUnauthorized creates a ListEdgeRouterPoliciesUnauthorized with default headers values
func NewListEdgeRouterPoliciesUnauthorized() *ListEdgeRouterPoliciesUnauthorized {
	return &ListEdgeRouterPoliciesUnauthorized{}
}

/* ListEdgeRouterPoliciesUnauthorized describes a response with status code 401, with default header values.

The supplied session does not have the correct access rights to request this resource
*/
type ListEdgeRouterPoliciesUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListEdgeRouterPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies][%d] listEdgeRouterPoliciesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListEdgeRouterPoliciesUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEdgeRouterPoliciesTooManyRequests creates a ListEdgeRouterPoliciesTooManyRequests with default headers values
func NewListEdgeRouterPoliciesTooManyRequests() *ListEdgeRouterPoliciesTooManyRequests {
	return &ListEdgeRouterPoliciesTooManyRequests{}
}

/* ListEdgeRouterPoliciesTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type ListEdgeRouterPoliciesTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListEdgeRouterPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies][%d] listEdgeRouterPoliciesTooManyRequests  %+v", 429, o.Payload)
}
func (o *ListEdgeRouterPoliciesTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEdgeRouterPoliciesServiceUnavailable creates a ListEdgeRouterPoliciesServiceUnavailable with default headers values
func NewListEdgeRouterPoliciesServiceUnavailable() *ListEdgeRouterPoliciesServiceUnavailable {
	return &ListEdgeRouterPoliciesServiceUnavailable{}
}

/* ListEdgeRouterPoliciesServiceUnavailable describes a response with status code 503, with default header values.

The request could not be completed due to the server being busy or in a temporarily bad state
*/
type ListEdgeRouterPoliciesServiceUnavailable struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListEdgeRouterPoliciesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /edge-router-policies][%d] listEdgeRouterPoliciesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ListEdgeRouterPoliciesServiceUnavailable) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListEdgeRouterPoliciesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
