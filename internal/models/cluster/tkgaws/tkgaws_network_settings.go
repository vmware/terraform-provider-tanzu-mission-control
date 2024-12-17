// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgawsmodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkSettings Network and provider information for the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgaws.NetworkSettings
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkSettings struct {

	// Kubernetes network information for the cluster.
	Cluster *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork `json:"cluster,omitempty"`

	// Provider specific network information for the cluster.
	Provider *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork `json:"provider,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkSettings) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
