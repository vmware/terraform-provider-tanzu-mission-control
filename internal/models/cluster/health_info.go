// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustermodel

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1CommonClusterHealthInfo Health information about cluster components healths.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.HealthInfo
type VmwareTanzuManageV1alpha1CommonClusterHealthInfo struct {

	// Controller manager's health status.
	ControllerManagerHealth *VmwareTanzuManageV1alpha1CommonClusterComponentHealth `json:"controllerManagerHealth,omitempty"`

	// etcd's health status.
	EtcdHealth []*VmwareTanzuManageV1alpha1CommonClusterComponentHealth `json:"etcdHealth"`

	// Message providing overall health details.
	Message string `json:"message,omitempty"`

	// Scheduler's health status.
	SchedulerHealth *VmwareTanzuManageV1alpha1CommonClusterComponentHealth `json:"schedulerHealth,omitempty"`

	// Timestamp of the record.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterHealthInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterHealthInfo) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonClusterHealthInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

type VmwareTanzuManageV1alpha1CommonClusterComponentHealth struct {

	// Health of the component.
	Health *VmwareTanzuManageV1alpha1CommonClusterHealth `json:"health,omitempty"`

	// Message providing details.
	Message string `json:"message,omitempty"`

	// Name of the component.
	Name string `json:"name,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterComponentHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterComponentHealth) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonClusterComponentHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
