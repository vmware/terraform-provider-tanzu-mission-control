package models

import (
	"github.com/go-openapi/swag"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1AksclusterAksCluster AksCluster is an AKS Kubernetes Cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.AksCluster
type VmwareTanzuManageV1alpha1AksclusterAksCluster struct {

	// Full name for the cluster.
	FullName *VmwareTanzuManageV1alpha1AksclusterFullName `json:"fullName,omitempty"`

	// Metadata for the cluster object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the cluster.
	Spec *VmwareTanzuManageV1alpha1AksclusterSpec `json:"spec,omitempty"`

	// Status for the cluster.
	Status *VmwareTanzuManageV1alpha1AksclusterStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1AksclusterAksCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAksCluster) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterAksCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
