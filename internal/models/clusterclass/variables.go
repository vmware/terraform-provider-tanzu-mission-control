/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclass

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassClusterClassVariable ClusterClassVariable defines a variable which can be configured in the Cluster topology.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.clusterclass.ClusterClassVariable
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassClusterClassVariable struct {

	// Name of the cluster class variable.
	Name string `json:"name,omitempty"`

	// Required specifies if the variable is required.
	Required bool `json:"required,omitempty"`

	// Schema defines the schema of the variable.
	Schema *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassVariableSchema `json:"schema,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassClusterClassVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassClusterClassVariable) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassClusterClassVariable

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
