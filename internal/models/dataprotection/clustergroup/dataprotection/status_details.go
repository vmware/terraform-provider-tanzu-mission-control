/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustergroupmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusDetails Status of the DataProtection configure resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.Status.details
type VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusDetails struct {
	// Total number of targets available for this source resource.
	AvailableTargets int32 `json:"available_targets,omitempty"`
	// Number of atomic targets on which this source resource is successfully applied.
	Applied int32 `json:"applied,omitempty"`
	// Number of atomic targets on which this source resource is overridden by another resource.
	Overridden int32 `json:"overridden,omitempty"`
	// Number of atomic targets on which this source resource is still being applied.
	Pending int32 `json:"pending,omitempty"`
	// Number of atomic targets on which this source resource failed to apply due to some error.
	Error int32 `json:"error,omitempty"`
	// Number of atomic targets on which this source resource is currently being deleted (only applicable on some source resource types).
	Deleting int32 `json:"deleting,omitempty"`
	// Number of atomic targets on which this source resource is not applied because they don't match the provided selectors.
	Skipped int32 `json:"skipped,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusDetails) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusDetails

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
