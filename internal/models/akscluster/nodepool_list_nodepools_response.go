// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse Response from listing Nodepools.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.ListNodepoolsResponse
type VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse struct {

	// List of nodepools.
	Nodepools []*VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool `json:"nodepools"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolListNodepoolsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
