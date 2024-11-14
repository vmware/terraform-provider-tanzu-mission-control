// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package integration

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterIntegrationUpdateIntegrationResponse Response from updating an Integration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.integration.UpdateIntegrationResponse
type VmwareTanzuManageV1alpha1ClusterIntegrationUpdateIntegrationResponse struct {
	// Integration updated.
	Integration *VmwareTanzuManageV1alpha1ClusterIntegrationIntegration `json:"integration,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationUpdateIntegrationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationUpdateIntegrationResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterIntegrationUpdateIntegrationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
