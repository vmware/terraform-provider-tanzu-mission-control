/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupintegrationmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterGroupIntegrationSpec The integration configuration spec.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.integration.Spec
type VmwareTanzuManageV1alpha1ClusterGroupIntegrationSpec struct {

	// Configurations. The expected input schema can be found in v1alpha1/integration API.
	Configurations map[string]interface{} `json:"configurations,omitempty"`

	// Credential name is the name of the Organization's Account Credential to be used instead of secrets to add an integration on clusters within this cluster group.
	CredentialName string `json:"credentialName,omitempty"`

	// Secrets are for sensitive configurations. The values are write-only and will be masked when read.
	Secrets map[string]string `json:"secrets,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupIntegrationSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupIntegrationSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupIntegrationSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
