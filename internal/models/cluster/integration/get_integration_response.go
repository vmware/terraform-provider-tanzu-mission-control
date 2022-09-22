/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterIntegrationGetIntegrationResponse Response from getting an Integration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.integration.GetIntegrationResponse
type VmwareTanzuManageV1alpha1ClusterIntegrationGetIntegrationResponse struct {
	// Integration returned.
	Integration *VmwareTanzuManageV1alpha1ClusterIntegrationIntegration `json:"integration,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationGetIntegrationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationGetIntegrationResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterIntegrationGetIntegrationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
