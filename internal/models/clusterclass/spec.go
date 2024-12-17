// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clusterclass

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassSpec Spec of the cluster class.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.clusterclass.Spec
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassSpec struct {

	// Variables defines the variables which can be configured
	// in the Cluster topology and are then used in patches.
	Variables []*VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassVariable `json:"variables"`

	// Workers classes is a collection of node types which can be used to create
	// the worker nodes of the cluster.
	WorkersClasses []string `json:"workersClasses"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
