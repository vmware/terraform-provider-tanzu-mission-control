/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kustomizationclustergroupmodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization Represents configuration that needs to be applied to cluster group.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.kustomization.Kustomization
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization struct {

	// Full name for the Kustomization.
	FullName *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName `json:"fullName,omitempty"`

	// Metadata for the Kustomization object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Kustomization.
	Spec *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec `json:"spec,omitempty"`

	// Status for the Kustomization.
	Status *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
