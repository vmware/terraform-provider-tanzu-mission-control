/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterRequest Request to create a TanzuKubernetesCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.CreateTanzuKubernetesClusterRequest
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterRequest struct {

	// TanzuKubernetesCluster to create.
	TanzuKubernetesCluster *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster `json:"tanzuKubernetesCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterResponse Response from creating a TanzuKubernetesCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.CreateTanzuKubernetesClusterResponse
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterResponse struct {

	// TanzuKubernetesCluster created.
	TanzuKubernetesCluster *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster `json:"tanzuKubernetesCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
