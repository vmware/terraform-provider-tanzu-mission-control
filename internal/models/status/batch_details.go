/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package statusmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1CommonBatchDetails Details contains information about a source resource being applied on its atomic targets.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.batch.Details
type VmwareTanzuManageV1alpha1CommonBatchDetails struct {

	// Number of atomic targets on which this source resource is successfully applied.
	Applied int32 `json:"applied,omitempty"`

	// Total number of targets available for this source resource.
	AvailableTargets int32 `json:"availableTargets,omitempty"`

	// Number of atomic targets on which this source resource failed to apply due to some error.
	Error int32 `json:"error,omitempty"`

	// Number of atomic targets on which this source resource is overridden by another resource.
	Overridden int32 `json:"overridden,omitempty"`

	// Number of atomic targets on which this source resource is still being applied.
	Pending int32 `json:"pending,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonBatchDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonBatchDetails) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonBatchDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
