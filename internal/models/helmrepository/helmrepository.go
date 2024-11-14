// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmrepositoryclustermodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepository Represents Helm Repository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.repository.Repository
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepository struct {

	// Full name for the artifact metadata.
	FullName *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryFullName `json:"fullName,omitempty"`

	// Metadata for the helm repository object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the helm repository.
	Spec *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySpec `json:"spec,omitempty"`

	// Status for the helm repository.
	Status *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepository) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
