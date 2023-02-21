/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package secret

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretExportResponse Response from getting a SecretExport.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.secretexport.GetSecretExportResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretExportResponse struct {

	// SecretExport returned.
	SecretExport *VmwareTanzuManageV1alpha1ClusterNamespaceSecretExport `json:"secretExport,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretExportResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretExportResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceGetSecretExportResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
