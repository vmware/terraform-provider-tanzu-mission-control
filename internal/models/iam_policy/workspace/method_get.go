/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package workspaceiammodel

import (
	"github.com/go-openapi/swag"

	iammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

// VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceIAMPolicyResponse GetWorkspaceIAMPolicy response message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.workspace.GetWorkspaceIAMPolicyResponse
type VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceIAMPolicyResponse struct {

	// Workspace policy.
	PolicyList []*iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policyList"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1WorkspaceGetWorkspaceIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
