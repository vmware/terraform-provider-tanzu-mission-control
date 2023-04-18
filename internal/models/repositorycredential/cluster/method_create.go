/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package repositorycredentialclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterFluxcdSourceSecretRequest Request to create a credential.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.SourceSecret.CreateSourceSecretRequest
type VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretRequest struct {
	// SourceSecret to create.
	SourceSecret *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecret `json:"sourceSecret,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretResponse Response from creating a Credential.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.sourcesecret.CreateSourceSecretResponse// Response from creating a SourceSecret.
type VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretResponse struct {
	// SourceSecret created.
	SourceSecret *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecret `json:"sourceSecret,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
