// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package nodepool

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNodepoolListNodepoolsResponse Response from listing Nodepools.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.ListNodepoolsResponse
type VmwareTanzuManageV1alpha1ClusterNodepoolListNodepoolsResponse struct {

	// List of nodepools.
	Nodepools []*VmwareTanzuManageV1alpha1ClusterNodepoolNodepool `json:"nodepools"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolListNodepoolsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolListNodepoolsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolListNodepoolsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
