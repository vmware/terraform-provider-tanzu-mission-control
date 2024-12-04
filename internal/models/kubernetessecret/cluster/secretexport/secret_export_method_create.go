// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package secret

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportRequest Request to create a SecretExport.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.secretexport.CreateSecretExportRequest
type VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportRequest struct {

	// SecretExport to create.
	SecretExport *VmwareTanzuManageV1alpha1ClusterNamespaceSecretExport `json:"secretExport,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceSecretexportCreateSecretExportResponse Response from creating a SecretExport.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.secretexport.CreateSecretExportResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportResponse struct {

	// SecretExport created.
	SecretExport *VmwareTanzuManageV1alpha1ClusterNamespaceSecretExport `json:"secretExport,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceSecretExportResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
