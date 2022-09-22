/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"github.com/go-openapi/swag"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
)

// VmwareTanzuManageV1alpha1ClusterIntegrationStatus Status of the integration configuration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.integration.Status
type VmwareTanzuManageV1alpha1ClusterIntegrationStatus struct {
	// Deep link to integration service that shows details for this cluster.
	ClusterViewURL string `json:"clusterViewUrl,omitempty"`

	// Conditions that help identify the phase.
	Conditions map[string]nodepool.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Integration Workload backed indicator abstracts workload in Status. This indicator indicates the state of
	// the workload deployed by the Integration Partner team. In case of issues with this indicator, Integration Partner
	// team should be able to help resolve the issues related to this indicator.
	IntegrationWorkload *VmwareTanzuManageV1alpha1ClusterIntegrationIndicator `json:"integrationWorkload,omitempty"`

	// Status of the Integration Partner's Operator deployed by TMC.
	Operator *VmwareTanzuManageV1alpha1ClusterIntegrationOperator `json:"operator,omitempty"`

	// Phase of the integration.
	Phase *VmwareTanzuManageV1alpha1ClusterIntegrationPhase `json:"phase,omitempty"`

	// System indicator abstracts Phase and Operator in Status.
	TmcAdapter *VmwareTanzuManageV1alpha1ClusterIntegrationIndicator `json:"tmcAdapter,omitempty"`

	// Existing version of the integration.
	Version string `json:"version,omitempty"`

	// Status of the Integration Partner's workloads.
	Workload *VmwareTanzuManageV1alpha1ClusterIntegrationWorkload `json:"workload,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterIntegrationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
