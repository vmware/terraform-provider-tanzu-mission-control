// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgawsmodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings Network and security settings for the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgaws.Settings
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings struct {

	// NetworkSettings specifies network-related settings for the cluster.
	Network *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkSettings `json:"network,omitempty"`

	// SecuritySettings specifies security-related settings for the cluster.
	Security *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSecuritySettings `json:"security,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
