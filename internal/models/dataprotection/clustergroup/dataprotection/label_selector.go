/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustergroupmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelectorLabelSelector Label based Selector.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.spec.selector.labelselector
type VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelectorLabelSelector struct {
	// key is the label key that the selector applies to.
	// +patchMergeKey=key
	// +patchStrategy=merge
	Key string `json:"key,omitempty"`

	// operator represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator string `json:"operator,omitempty"`

	// values is an array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. This array is replaced during a strategic
	// merge patch.
	// +optional
	Values []string `json:"values"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelectorLabelSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelectorLabelSelector) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupDataprotectionSpecSelectorLabelSelector

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
