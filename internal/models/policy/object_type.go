/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policymodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuCoreV1alpha1ObjectType Holds general type metadatas.
//
// swagger:model vmware.tanzu.core.v1alpha1.object.Type
type VmwareTanzuCoreV1alpha1ObjectType struct {

	// Kind of the type.
	Kind string `json:"kind,omitempty"`

	// Package of the type.
	Package string `json:"package,omitempty"`

	// Version of the type.
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1ObjectType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1ObjectType) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1ObjectType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
