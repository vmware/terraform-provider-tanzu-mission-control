/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplatemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1PolicyTemplateFullName Full name of the policy template. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.template.FullName
type VmwareTanzuManageV1alpha1PolicyTemplateFullName struct {

	// Name of policy template.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplateFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplateFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1PolicyTemplateFullName

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
