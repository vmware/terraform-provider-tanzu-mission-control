// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgvspheremodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSettings VSphere related settings for workload cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgvsphere.Settings
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSettings struct {

	// NetworkSettings specifies network-related settings for the cluster.
	Network *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkSettings `json:"network,omitempty"`

	// SecuritySettings specifies security-related settings for the cluster.
	Security *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSecuritySettings `json:"security,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSettings) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
