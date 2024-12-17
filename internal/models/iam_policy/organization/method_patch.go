// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package organizationiammodel

import (
	"github.com/go-openapi/swag"

	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
	organizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/organization"
)

// VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyRequest PatchOrganizationIAMPolicy request message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.PatchOrganizationIAMPolicyRequest
type VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyRequest struct {

	// Policy delta to be applied.
	BindingDeltaList []*iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta `json:"bindingDeltaList"`

	// Organization fullname.
	FullName *organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName `json:"fullName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyResponse PatchOrganizationIAMPolicy response message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.PatchOrganizationIAMPolicyResponse
type VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyResponse struct {

	// New Organization policy.
	Policy *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationPatchOrganizationIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
