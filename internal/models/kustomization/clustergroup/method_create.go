/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kustomizationclustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest Request to create a Kustomization.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.kustomization.CreateKustomizationRequest
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest struct {

	// Kustomization to create.
	Kustomization *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization `json:"kustomization,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse Response from creating a Kustomization.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.kustomization.CreateKustomizationResponse
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse struct {

	// Kustomization created.
	Kustomization *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomization `json:"kustomization,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationKustomizationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
