// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustergroupsecret

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretGetSecretResponse Response from getting a Secret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.secret.GetSecretResponse
type VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretGetSecretResponse struct {

	// Secret returned.
	Secret *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret `json:"secret,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretGetSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretGetSecretResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretGetSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
