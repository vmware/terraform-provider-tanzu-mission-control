/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyclustermodel

import (
	"fmt"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterPolicyFullName Full name of the cluster policy. This includes the object
// name along with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.policy.FullName
type VmwareTanzuManageV1alpha1ClusterPolicyFullName struct {

	// Name of the cluster.
	ClusterName string `json:"clusterName,omitempty"`

	// Name of the management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of the policy.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of the cluster provisioner.
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPolicyFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPolicyFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterPolicyFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

func (m *VmwareTanzuManageV1alpha1ClusterPolicyFullName) ToString() string {
	if m == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s:%s:%s:%s", m.OrgID, m.ManagementClusterName, m.ProvisionerName, m.ClusterName, m.Name)
}
