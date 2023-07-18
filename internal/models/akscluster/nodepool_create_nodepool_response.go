package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse Response from creating a Nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.CreateNodepoolResponse
type VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse struct {

	// Nodepool created.
	Nodepool *VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolCreateNodepoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
