/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgawsmodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSecuritySettings Security settings for the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgaws.SecuritySettings
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSecuritySettings struct {

	// SSH key for the cluster VMs.
	SSHKey string `json:"sshKey,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSecuritySettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSecuritySettings) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSecuritySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
