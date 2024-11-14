// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package workspaceiammodel

import (
	"github.com/go-openapi/swag"

	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
	workspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/workspace"
)

// VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyRequest PatchWorkspaceIAMPolicy request message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.PatchWorkspaceIAMPolicyRequest
type VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyRequest struct {

	// Binding delta to be applied.
	BindingDeltaList []*iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta `json:"bindingDeltaList"`

	// Workspace full_name.
	FullName *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName `json:"fullName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyResponse PatchWorkspaceIAMPolicy response message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.PatchWorkspaceIAMPolicyResponse
type VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyResponse struct {

	// New policy object.
	Policy *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspacePatchWorkspaceIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
