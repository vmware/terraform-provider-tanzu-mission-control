// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package integration

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationRequest Request to create an Integration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.integration.CreateIntegrationRequest
type VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationRequest struct {
	// Integration to create.
	Integration *VmwareTanzuManageV1alpha1ClusterIntegrationIntegration `json:"integration,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
