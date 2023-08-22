/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcnodepool

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName Full name of the nodepool. This includes the object name along with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.nodepool.FullName
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName struct {
	// Name of the management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of this nodepool.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgId string `json:"orgId,omitempty"`

	// Provisioner of the cluster.
	ProvisionerName string `json:"provisionerName,omitempty"`

	// Name of the cluster.
	TanzuKubernetesClusterName string `json:"tanzuKubernetesClusterName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
