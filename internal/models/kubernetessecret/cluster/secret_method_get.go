/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package secret

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretResponse Response from getting a Secret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.secret.GetSecretResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretResponse struct {

	// Secret returned.
	Secret *VmwareTanzuManageV1alpha1ClusterNamespaceSecret `json:"secret,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
