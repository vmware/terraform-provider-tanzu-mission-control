/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrolemodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1IamRoleAggregationRule AggregationRule for a role.
//
// swagger:model vmware.tanzu.manage.v1alpha1.iam.role.AggregationRule
type VmwareTanzuManageV1alpha1IamRoleAggregationRule struct {

	// Label based Cluster Role Selector.
	ClusterRoleSelectors []*K8sIoApimachineryPkgApisMetaV1LabelSelector `json:"clusterRoleSelectors"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRoleAggregationRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1IamRoleAggregationRule) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1IamRoleAggregationRule

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
