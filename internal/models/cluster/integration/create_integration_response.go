/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationResponse Response from creating an Integration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.integration.CreateIntegrationResponse
type VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationResponse struct {
	// Integration created.
	Integration *VmwareTanzuManageV1alpha1ClusterIntegrationIntegration `json:"integration,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterIntegrationCreateIntegrationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
