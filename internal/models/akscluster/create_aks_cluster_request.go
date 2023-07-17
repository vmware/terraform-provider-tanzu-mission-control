package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest Request to create an AksCluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.CreateAksClusterRequest
type VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest struct {

	// AksCluster to create.
	AksCluster *VmwareTanzuManageV1alpha1AksclusterAksCluster `json:"aksCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
