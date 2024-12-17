// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package kustomizationclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest Request to create a Kustomization.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.kustomization.CreateKustomizationRequest
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest struct {

	// Kustomization to create.
	Kustomization *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization `json:"kustomization,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse Response from creating a Kustomization.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.kustomization.CreateKustomizationResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse struct {

	// Kustomization created.
	Kustomization *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomization `json:"kustomization,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationKustomizationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
