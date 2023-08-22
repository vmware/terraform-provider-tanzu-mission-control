/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcnodepool

import (
	"github.com/go-openapi/swag"

	tkcstatus "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/status"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolStatus Status of node pool resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.nodepool.Status

type VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolStatus struct {
	// Phase of the nodepool resource.
	Phase *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolStatusPhase `json:"phase,omitempty"`

	// Conditions for the nodepool resource.
	Conditions map[string]tkcstatus.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
