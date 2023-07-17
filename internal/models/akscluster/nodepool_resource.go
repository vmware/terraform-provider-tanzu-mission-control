package models

import (
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool Nodepool associated with a AKS cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.Nodepool
type VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool struct {

	// Full name for the Nodepool.
	FullName *VmwareTanzuManageV1alpha1AksclusterNodepoolFullName `json:"fullName,omitempty"`

	// Metadata for the Nodepool object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Nodepool.
	Spec *VmwareTanzuManageV1alpha1AksclusterNodepoolSpec `json:"spec,omitempty"`

	// Status of the Nodepool.
	Status *VmwareTanzuManageV1alpha1AksclusterNodepoolStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
