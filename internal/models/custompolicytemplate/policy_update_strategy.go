/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplatemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategy PolicyUpdateStrategy on how to handle policies after a policy template update.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.template.PolicyUpdateStrategy
type VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategy struct {

	// The strategy to use for policy updates.
	Type *VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategy) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategy

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
