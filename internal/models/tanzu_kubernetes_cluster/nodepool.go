/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetescluster

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition Definition is the definition of nodepool for cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.nodepool.Definition
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition struct {

	// Info for the nodepool.
	Info *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolInfo `json:"info,omitempty"`

	// Spec for the nodepool.
	Spec *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec `json:"spec,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
