/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcmodels

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1CommonClusterStatus Status is the common status of the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.Status
type VmwareTanzuManageV1alpha1CommonClusterStatus struct {

	// Conditions of the cluster resource.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Phase of the cluster resource.
	Phase *VmwareTanzuManageV1alpha1CommonClusterStatusPhase `json:"phase,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonClusterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
