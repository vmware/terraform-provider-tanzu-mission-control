// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgservicevspheremodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSettings Settings is the tkg service specific cluster setting.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgservicevsphere.Settings
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSettings struct {

	// NetworkSettings specifies network-related settings for the cluster.
	Network *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkSettings `json:"network,omitempty"`

	// StorageSettings specifies storage-related settings for the cluster.
	Storage *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereStorageSettings `json:"storage,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSettings) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
