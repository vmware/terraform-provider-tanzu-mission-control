// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package secret

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest Request to create a Secret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.secret.CreateSecretRequest
type VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest struct {

	// Secret to create.
	Secret *VmwareTanzuManageV1alpha1ClusterNamespaceSecret `json:"secret,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceSecretRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse Response from creating a Secret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.secret.CreateSecretResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse struct {

	// Secret created.
	Secret *VmwareTanzuManageV1alpha1ClusterNamespaceSecret `json:"secret,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
