// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package integration

import (
	"github.com/go-openapi/swag"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/extension"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
)

// VmwareTanzuManageV1alpha1ClusterIntegrationOperator Status of the Integration Partner's operator deployed by TMC.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.integration.Operator
type VmwareTanzuManageV1alpha1ClusterIntegrationOperator struct {
	// The Conditions attached to an extension resource.
	Conditions map[string]nodepool.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Health of the deployed extension.
	Health *extension.VmwareTanzuManageV1alpha1ClusterExtensionHealth `json:"health,omitempty"`

	// Previous version of the extension.
	PreviousVersion string `json:"previousVersion,omitempty"`

	// Phase of the extension.
	State *extension.VmwareTanzuManageV1alpha1ClusterExtensionPhase `json:"state,omitempty"`

	// Version of the extension.
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationOperator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationOperator) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterIntegrationOperator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
