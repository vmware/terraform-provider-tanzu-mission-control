/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackagerepository

import (
	"github.com/go-openapi/swag"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository Represents Tanzu Carvel Package repository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.repository.Repository
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository struct {

	// Full name for the Package Repository.
	FullName *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryFullName `json:"fullName,omitempty"`

	// Metadata for the Package Repository object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Package Repository.
	Spec *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec `json:"spec,omitempty"`

	// Status for the Package Repository.
	Status *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
