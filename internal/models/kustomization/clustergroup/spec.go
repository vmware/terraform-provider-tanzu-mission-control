/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kustomizationclustergroupmodel

import (
	"github.com/go-openapi/swag"

	kustomizationclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/cluster"
)

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec Spec for the Kustomization.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.kustomization.Spec
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec struct {

	// Spec for the Kustomization as defined at atomic level.
	AtomicSpec *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec `json:"atomicSpec,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
