/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package organizationmodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1OrganizationFullName Full name of the organization. This includes the org_id.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.FullName
type VmwareTanzuManageV1alpha1OrganizationFullName struct {

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
