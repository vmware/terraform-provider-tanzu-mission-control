/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrolemodels

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1IamRoleRole Manage permissions on resources.
//
// swagger:model vmware.tanzu.manage.v1alpha1.iam.role.Role
type VmwareTanzuManageV1alpha1IamRole struct {

	// Full name for the role.
	FullName *VmwareTanzuManageV1alpha1IamRoleFullName `json:"fullName,omitempty"`

	// Metadata for the role object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the role.
	Spec *VmwareTanzuManageV1alpha1IamRoleSpec `json:"spec,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRole) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1IamRole

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
