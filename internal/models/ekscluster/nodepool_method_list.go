package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterNodepoolListNodepoolsResponse Response from listing Nodepools.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.ListNodepoolsResponse
type VmwareTanzuManageV1alpha1EksclusterNodepoolListNodepoolsResponse struct {

	// List of nodepools.
	Nodepools []*VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool `json:"nodepools"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolListNodepoolsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolListNodepoolsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterNodepoolListNodepoolsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
