// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// GrpcGatewayRuntimeError grpc gateway runtime error
//
// swagger:model grpc.gateway.runtime.Error
type GrpcGatewayRuntimeError struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*GoogleProtobufAny `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// MarshalBinary interface implementation.
func (m *GrpcGatewayRuntimeError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *GrpcGatewayRuntimeError) UnmarshalBinary(b []byte) error {
	var res GrpcGatewayRuntimeError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
