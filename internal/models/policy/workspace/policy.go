// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policyworkspacemodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
)

// VmwareTanzuManageV1alpha1WorkspacePolicyPolicy A Policy to apply on a group of managed Kubernetes namespaces.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.policy.Policy
type VmwareTanzuManageV1alpha1WorkspacePolicyPolicy struct {

	// Full name for the Workspace policy.
	FullName *VmwareTanzuManageV1alpha1WorkspacePolicyFullName `json:"fullName,omitempty"`

	// Metadata for the Workspace policy.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Workspace policy.
	Spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec `json:"spec,omitempty"`

	// Metadata describing the type of the resource.
	Type *policymodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePolicyPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePolicyPolicy) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspacePolicyPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
