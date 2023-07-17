package models

import (
	"encoding/json"
)

// GoogleProtobufNullValue `NullValue` is a singleton enumeration to represent the null value for the
// `Value` type union.
//
//	The JSON representation for `NullValue` is JSON `null`.
//
//	- NULL_VALUE: Null value.
//
// swagger:model google.protobuf.NullValue
type GoogleProtobufNullValue string

func NewGoogleProtobufNullValue(value GoogleProtobufNullValue) *GoogleProtobufNullValue {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GoogleProtobufNullValue.
func (m GoogleProtobufNullValue) Pointer() *GoogleProtobufNullValue {
	return &m
}

const (

	// GoogleProtobufNullValueNULLVALUE captures enum value "NULL_VALUE".
	GoogleProtobufNullValueNULLVALUE GoogleProtobufNullValue = "NULL_VALUE"
)

// for schema.
var googleProtobufNullValueEnum []interface{}

func init() {
	var res []GoogleProtobufNullValue
	if err := json.Unmarshal([]byte(`["NULL_VALUE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		googleProtobufNullValueEnum = append(googleProtobufNullValueEnum, v)
	}
}
