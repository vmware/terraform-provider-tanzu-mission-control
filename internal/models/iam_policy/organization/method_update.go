// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package organizationiammodel

import (
	"github.com/go-openapi/swag"

	iammodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

// VmwareTanzuManageV1alpha1OrganizationUpdateOrganizationIAMPolicyResponse UpdateOrganizationIAMPolicy response message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.UpdateOrganizationIAMPolicyResponse
type VmwareTanzuManageV1alpha1OrganizationUpdateOrganizationIAMPolicyResponse struct {

	// Organization policy.
	Policy *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationUpdateOrganizationIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationUpdateOrganizationIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationUpdateOrganizationIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
