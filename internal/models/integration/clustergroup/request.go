/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupintegrationmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterGroupIntegrationCreateIntegrationRequest Request to create an Integration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.integration.CreateIntegrationRequest
type VmwareTanzuManageV1alpha1ClusterGroupIntegrationData struct {

	// Integration to create.
	Integration *VmwareTanzuManageV1alpha1ClusterGroupIntegrationIntegration `json:"integration,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupIntegrationData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupIntegrationData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupIntegrationData

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
