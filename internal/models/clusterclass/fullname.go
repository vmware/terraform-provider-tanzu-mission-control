/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclass

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassFullName Full name of the cluster class. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.clusterclass.FullName
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassFullName struct {

	// Name of the management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of this cluster class.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Provisioner of the cluster class.
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
