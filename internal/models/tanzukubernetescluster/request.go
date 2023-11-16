/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData Request to create a TanzuKubernetesCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.CreateTanzuKubernetesClusterRequest
type VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData struct {

	// TanzuKubernetesCluster to create.
	TanzuKubernetesCluster *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster `json:"tanzuKubernetesCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
