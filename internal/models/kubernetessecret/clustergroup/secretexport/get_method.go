// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustergroupsecretexport

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportGetSecretExportResponse Response from getting a SecretExport.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.secretexport.GetSecretExportResponse
type VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportGetSecretExportResponse struct {

	// SecretExport returned.
	SecretExport *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportSecretExport `json:"secretExport,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportGetSecretExportResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportGetSecretExportResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportGetSecretExportResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
