/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package permissontemplatemodels

import (
	"github.com/go-openapi/swag"

	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
)

// VmwareTanzuManageV1alpha1AccountCredentialGeneratePermissionTemplateRequest Request to generate a permission template for a non-existent credential.
//
// swagger:model vmware.tanzu.manage.v1alpha1.account.credential.GeneratePermissionTemplateRequest
type VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateRequest struct {

	// The Tanzu capability for which the credential shall be used.
	Capability string `json:"capability,omitempty"`

	// The full name of the credential that will be created.
	FullName *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialFullName `json:"fullName,omitempty"`

	// The infrastructure provider that the permission template is for.
	Provider *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProvider `json:"provider,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateRequest

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse Response containing the generated permission template.
//
// swagger:model vmware.tanzu.manage.v1alpha1.account.credential.GeneratePermissionTemplateResponse
type VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse struct {

	// The Tanzu capability for which the credential shall be used.
	Capability string `json:"capability,omitempty"`

	// The full name of the credential that this template is for.
	FullName *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialFullName `json:"fullName,omitempty"`

	// Base64 encoded permission template.
	PermissionTemplate string `json:"permissionTemplate,omitempty"`

	// The infrastructure provider that the permission template is for.
	Provider *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProvider `json:"provider,omitempty"`

	// URL for permission template.
	TemplateURL string `json:"templateUrl,omitempty"`

	// Values which helps in forming the CLI command output and displaying values on UI.
	TemplateValues map[string]string `json:"templateValues,omitempty"`

	// Values which are not defined in the template parameters definition
	UndefinedTemplateValues map[string]string `json:"undefinedTemplateValues,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
