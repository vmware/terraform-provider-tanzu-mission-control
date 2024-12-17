// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kustomizationclustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationGetKustomizationResponse Response from getting a Kustomization.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.kustomization.GetKustomizationResponse
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationGetKustomizationResponse struct {

	// Kustomization returned.
	Kustomization *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization `json:"kustomization,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationGetKustomizationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationGetKustomizationResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationGetKustomizationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
