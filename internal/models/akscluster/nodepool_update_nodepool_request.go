package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolRequest Request to update (overwrite) a Nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.UpdateNodepoolRequest
type VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolRequest struct {

	// Update Nodepool.
	Nodepool *VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool `json:"nodepool,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolUpdateNodepoolRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
