// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package integration

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterIntegrationDeleteIntegrationResponse Response from deleting an Integration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.integration.DeleteIntegrationResponse
type VmwareTanzuManageV1alpha1ClusterIntegrationDeleteIntegrationResponse struct {
	// Message regarding deletion.
	Message string `json:"message,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationDeleteIntegrationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationDeleteIntegrationResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterIntegrationDeleteIntegrationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
