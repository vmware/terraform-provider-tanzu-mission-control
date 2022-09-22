/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package extension

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterExtensionHealth Health status of the deployed extension.
//
//   - HEALTH_UNSPECIFIED: Unknown.
//   - HEALTHY: Healthy.
//   - UNHEALTHY: Unhealthy.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.extension.Health
type VmwareTanzuManageV1alpha1ClusterExtensionHealth string

func NewVmwareTanzuManageV1alpha1ClusterExtensionHealth(value VmwareTanzuManageV1alpha1ClusterExtensionHealth) *VmwareTanzuManageV1alpha1ClusterExtensionHealth {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterExtensionHealth.
func (m VmwareTanzuManageV1alpha1ClusterExtensionHealth) Pointer() *VmwareTanzuManageV1alpha1ClusterExtensionHealth {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterExtensionHealthHEALTHUNSPECIFIED captures enum value "HEALTH_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterExtensionHealthHEALTHUNSPECIFIED VmwareTanzuManageV1alpha1ClusterExtensionHealth = "HEALTH_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterExtensionHealthHEALTHY captures enum value "HEALTHY".
	VmwareTanzuManageV1alpha1ClusterExtensionHealthHEALTHY VmwareTanzuManageV1alpha1ClusterExtensionHealth = "HEALTHY"

	// VmwareTanzuManageV1alpha1ClusterExtensionHealthUNHEALTHY captures enum value "UNHEALTHY".
	VmwareTanzuManageV1alpha1ClusterExtensionHealthUNHEALTHY VmwareTanzuManageV1alpha1ClusterExtensionHealth = "UNHEALTHY"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterExtensionHealthEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterExtensionHealth
	if err := json.Unmarshal([]byte(`["HEALTH_UNSPECIFIED","HEALTHY","UNHEALTHY"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterExtensionHealthEnum = append(vmwareTanzuManageV1alpha1ClusterExtensionHealthEnum, v)
	}
}
