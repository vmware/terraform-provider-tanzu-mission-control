// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kustomizationclustermodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization Represents configuration that needs to be applied to cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.kustomization.Kustomization
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization struct {

	// Full name for the Kustomization.
	FullName *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationFullName `json:"fullName,omitempty"`

	// Metadata for the Kustomization object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Kustomization.
	Spec *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec `json:"spec,omitempty"`

	// Status for the Kustomization.
	Status *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
