package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GoogleProtobufAny google protobuf any
//
// swagger:model google.protobuf.Any
type GoogleProtobufAny struct {

	// type Url
	TypeURL string `json:"typeUrl,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// MarshalBinary interface implementation.
func (m *GoogleProtobufAny) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *GoogleProtobufAny) UnmarshalBinary(b []byte) error {
	var res GoogleProtobufAny
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
