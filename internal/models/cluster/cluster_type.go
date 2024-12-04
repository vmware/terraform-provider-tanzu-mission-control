// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustermodel

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterType Type describe type of cluster.
/*
  - TYPE_UNSPECIFIED: TYPE_UNSPECIFIED cluster type.
  - ATTACHED: ATTACHED cluster type.
  - PROVISIONED: PROVISIONED cluster type.
  - TANZU_KUBERNETES_GRID_SERVICE: Tanzu Kubernetes Grid Service cluster type.
  - TANZU_KUBERNETES_GRID: Tanzu Kubernetes Grid cluster type.

 swagger:model vmware.tanzu.manage.v1alpha1.cluster.Type
*/
type VmwareTanzuManageV1alpha1ClusterType string

func NewVmwareTanzuManageV1alpha1ClusterType(value VmwareTanzuManageV1alpha1ClusterType) *VmwareTanzuManageV1alpha1ClusterType {
	v := value
	return &v
}

const (

	// VmwareTanzuManageV1alpha1ClusterTypeTYPEUNSPECIFIED captures enum value "TYPE_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterTypeTYPEUNSPECIFIED VmwareTanzuManageV1alpha1ClusterType = "TYPE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterTypeATTACHED captures enum value "ATTACHED".
	VmwareTanzuManageV1alpha1ClusterTypeATTACHED VmwareTanzuManageV1alpha1ClusterType = "ATTACHED"

	// VmwareTanzuManageV1alpha1ClusterTypePROVISIONED captures enum value "PROVISIONED".
	VmwareTanzuManageV1alpha1ClusterTypePROVISIONED VmwareTanzuManageV1alpha1ClusterType = "PROVISIONED"

	// VmwareTanzuManageV1alpha1ClusterTypeTANZUKUBERNETESGRIDSERVICE captures enum value "TANZU_KUBERNETES_GRID_SERVICE".
	VmwareTanzuManageV1alpha1ClusterTypeTANZUKUBERNETESGRIDSERVICE VmwareTanzuManageV1alpha1ClusterType = "TANZU_KUBERNETES_GRID_SERVICE"

	// VmwareTanzuManageV1alpha1ClusterTypeTANZUKUBERNETESGRID captures enum value "TANZU_KUBERNETES_GRID".
	VmwareTanzuManageV1alpha1ClusterTypeTANZUKUBERNETESGRID VmwareTanzuManageV1alpha1ClusterType = "TANZU_KUBERNETES_GRID"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterType
	if err := json.Unmarshal([]byte(`["TYPE_UNSPECIFIED","ATTACHED","PROVISIONED","TANZU_KUBERNETES_GRID_SERVICE","TANZU_KUBERNETES_GRID"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterTypeEnum = append(vmwareTanzuManageV1alpha1ClusterTypeEnum, v)
	}
}
