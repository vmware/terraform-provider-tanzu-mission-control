// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustergroupsecret

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest Request to create a Secret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.secret.CreateSecretRequest
type VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest struct {

	// Secret to create.
	Secret *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret `json:"secret,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse Response from creating a Secret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.secret.CreateSecretResponse
type VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse struct {

	// Secret created.
	Secret *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret `json:"secret,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
