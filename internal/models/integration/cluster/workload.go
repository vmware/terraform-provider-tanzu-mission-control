/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterintegrationmodels

import (
	"github.com/go-openapi/swag"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
)

// VmwareTanzuManageV1alpha1ClusterIntegrationWorkload Status of the Integration Partner's workloads.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.integration.Workload
type VmwareTanzuManageV1alpha1ClusterIntegrationWorkload struct {
	// Conditions of the workloads that supports readiness and health.
	Conditions map[string]nodepool.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Version of the application.
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationWorkload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegrationWorkload) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterIntegrationWorkload

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
