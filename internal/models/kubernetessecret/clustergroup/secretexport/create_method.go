// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustergroupsecretexport

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportRequest Request to create a SecretExport.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.secretexport.CreateSecretExportRequest
type VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportRequest struct {

	// SecretExport to create.
	SecretExport *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportSecretExport `json:"secretExport,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportResponse Response from creating a SecretExport.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.secretexport.CreateSecretExportResponse
type VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportResponse struct {

	// SecretExport created.
	SecretExport *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportSecretExport `json:"secretExport,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
