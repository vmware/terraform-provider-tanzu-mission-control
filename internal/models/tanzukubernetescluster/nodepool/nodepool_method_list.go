/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcnodepool

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolListNodepoolsResponse Response from listing Nodepools.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.nodepool.ListNodepoolsResponse
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolListNodepoolsResponse struct {

	// List of nodepools.
	Nodepools []*VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolNodepool `json:"nodepools"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolListNodepoolsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolListNodepoolsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolListNodepoolsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
