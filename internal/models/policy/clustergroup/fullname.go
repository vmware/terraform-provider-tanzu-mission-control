// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policyclustergroupmodel

import (
	"fmt"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClustergroupPolicyFullName Full name of the cluster group policy. This includes the object
// name along with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.policy.FullName
type VmwareTanzuManageV1alpha1ClustergroupPolicyFullName struct {

	// Name of the cluster group.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Name of the policy.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupPolicyFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

func (m *VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) ToString() string {
	if m == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s:%s", m.OrgID, m.ClusterGroupName, m.Name)
}
