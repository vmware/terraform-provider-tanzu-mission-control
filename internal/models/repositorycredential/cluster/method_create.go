/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package repositorycredentialclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRequest Request to create a credential.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.repositorycredential.CreateRepositoryCredentialRequest
type VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRequest struct {

	// Credential to create.
	Repositorycredential *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredential `json:"respositorycredential,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialResponse Response from creating a Credential.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.repositorycredential.CreateRepositoryCredentialResponse
type VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialResponse struct {

	// Secret created.
	Repositorycredential *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialRepositorycredential `json:"respositorycredential,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdRepositorycredentialResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
