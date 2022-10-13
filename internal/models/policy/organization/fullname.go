/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyorganizationmodel

import (
	"fmt"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1OrganizationPolicyFullName Full name of the organization policy. This includes the object
// name along with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.policy.FullName
type VmwareTanzuManageV1alpha1OrganizationPolicyFullName struct {

	// Name of the policy.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPolicyFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPolicyFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationPolicyFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

func (m *VmwareTanzuManageV1alpha1OrganizationPolicyFullName) ToString() string {
	if m == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s", m.OrgID, m.Name)
}
