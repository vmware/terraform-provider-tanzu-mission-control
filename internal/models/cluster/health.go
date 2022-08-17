/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustermodel

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1CommonClusterHealth Health describes the health of a resource.
/*
  - HEALTH_UNSPECIFIED: Unspecified health.
  - HEALTHY: Resource is healthy.
  - WARNING: Resource is in warning state.
  - UNHEALTHY: Resource is unhealthy.
  - DISCONNECTED: Resource is disconnected.

 swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.Health
*/
type VmwareTanzuManageV1alpha1CommonClusterHealth string

func NewVmwareTanzuManageV1alpha1CommonClusterHealth(value VmwareTanzuManageV1alpha1CommonClusterHealth) *VmwareTanzuManageV1alpha1CommonClusterHealth {
	v := value
	return &v
}

const (

	// VmwareTanzuManageV1alpha1CommonClusterHealthHEALTHUNSPECIFIED captures enum value "HEALTH_UNSPECIFIED".
	VmwareTanzuManageV1alpha1CommonClusterHealthHEALTHUNSPECIFIED VmwareTanzuManageV1alpha1CommonClusterHealth = "HEALTH_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1CommonClusterHealthHEALTHY captures enum value "HEALTHY".
	VmwareTanzuManageV1alpha1CommonClusterHealthHEALTHY VmwareTanzuManageV1alpha1CommonClusterHealth = "HEALTHY"

	// VmwareTanzuManageV1alpha1CommonClusterHealthWARNING captures enum value "WARNING".
	VmwareTanzuManageV1alpha1CommonClusterHealthWARNING VmwareTanzuManageV1alpha1CommonClusterHealth = "WARNING"

	// VmwareTanzuManageV1alpha1CommonClusterHealthUNHEALTHY captures enum value "UNHEALTHY".
	VmwareTanzuManageV1alpha1CommonClusterHealthUNHEALTHY VmwareTanzuManageV1alpha1CommonClusterHealth = "UNHEALTHY"

	// VmwareTanzuManageV1alpha1CommonClusterHealthDISCONNECTED captures enum value "DISCONNECTED".
	VmwareTanzuManageV1alpha1CommonClusterHealthDISCONNECTED VmwareTanzuManageV1alpha1CommonClusterHealth = "DISCONNECTED"
)

// for schema.
var vmwareTanzuManageV1alpha1CommonClusterHealthEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1CommonClusterHealth
	if err := json.Unmarshal([]byte(`["HEALTH_UNSPECIFIED","HEALTHY","WARNING","UNHEALTHY","DISCONNECTED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1CommonClusterHealthEnum = append(vmwareTanzuManageV1alpha1CommonClusterHealthEnum, v)
	}
}
