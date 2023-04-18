/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package repositorycredentialclustermodel

import "github.com/go-openapi/swag"

type VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretGetSourceSecretResponse struct {
	// SourceSecret returned.
	SourceSecret *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSourceSecret `json:"sourceSecret,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretGetSourceSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretGetSourceSecretResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretGetSourceSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
