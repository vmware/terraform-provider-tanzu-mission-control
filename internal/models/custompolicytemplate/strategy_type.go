/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplatemodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType PolicyUpdateStrategyType defines strategies for updating policies after a policy template update.
//
//   - POLICY_UPDATE_STRATEGY_TYPE_UNSPECIFIED: UNSPECIFIED policy update strategy (default).
//
// Updates will not be allowed when this strategy is selected.
//   - INPLACE_UPDATE: In-place policy update strategy.
//
// Existing Template will be forcibly updated without creating a new version.
// There will be no changes to the policies using the template.
// Warning: When using this strategy, make sure that the updated template does not
// adversely affect the existing policies.
//
// swagger:model vmware.tanzu.manage.v1alpha1.policy.template.PolicyUpdateStrategyType
type VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType string

const (

	// VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyTypePOLICYUPDATESTRATEGYTYPEUNSPECIFIED captures enum value "POLICY_UPDATE_STRATEGY_TYPE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyTypePOLICYUPDATESTRATEGYTYPEUNSPECIFIED VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType = "POLICY_UPDATE_STRATEGY_TYPE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyTypeINPLACEUPDATE captures enum value "INPLACE_UPDATE".
	VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyTypeINPLACEUPDATE VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType = "INPLACE_UPDATE"
)

// for schema.
var vmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyTypeEnum []interface{}

func NewVmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType(value VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType) *VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType.
func (m VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType) Pointer() *VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType {
	return &m
}

func init() {
	var res []VmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyType

	if err := json.Unmarshal([]byte(`["POLICY_UPDATE_STRATEGY_TYPE_UNSPECIFIED","INPLACE_UPDATE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyTypeEnum = append(vmwareTanzuManageV1alpha1PolicyTemplatePolicyUpdateStrategyTypeEnum, v)
	}
}
