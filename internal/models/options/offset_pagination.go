/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package optionsmodel

import "github.com/go-openapi/swag"

// VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions Options to paginate a response using offsets.
//
// swagger:model vmware.tanzu.core.v1alpha1.options.OffsetPaginationOptions
type VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions struct {

	// Offset at which to start returning records.
	Offset string `json:"offset,omitempty"`

	// Number of records to return.
	Size string `json:"size,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
