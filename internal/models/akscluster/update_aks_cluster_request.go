package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest Request to update (overwrite) an AksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.UpdateAksClusterRequest
type VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest struct {

	// Update AksCluster.
	AksCluster *VmwareTanzuManageV1alpha1AksclusterAksCluster `json:"aksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
