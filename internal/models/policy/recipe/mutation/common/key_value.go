/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyrecipemutationcommonmodel

import "github.com/go-openapi/swag"

type KeyValue struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (m *KeyValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

func (m *KeyValue) UnmarshalBinary(b []byte) error {
	var res KeyValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
