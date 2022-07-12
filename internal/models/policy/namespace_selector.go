/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policymodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1CommonPolicyLabelSelector Label Selector.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.policy.LabelSelector
type VmwareTanzuManageV1alpha1CommonPolicyLabelSelector struct {

	// Match expressions is a list of label selector requirements, the requirements are ANDed.
	// Label selector requirements support 4 operators for matching labels - in, notin, exists and doesnotexist.
	MatchExpressions []*K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement `json:"matchExpressions"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicyLabelSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicyLabelSelector) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicyLabelSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
