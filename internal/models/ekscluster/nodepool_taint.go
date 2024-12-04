// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterNodepoolTaint The node this Taint is attached to has the "effect" on
// any pod that does not tolerate the Taint.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.Taint
type VmwareTanzuManageV1alpha1EksclusterNodepoolTaint struct {

	// Current effect state of the node pool.
	Effect *VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect `json:"effect,omitempty"`

	// The taint key to be applied to a node.
	Key string `json:"key,omitempty"`

	// The taint value corresponding to the taint key.
	Value string `json:"value,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolTaint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolTaint) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterNodepoolTaint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
