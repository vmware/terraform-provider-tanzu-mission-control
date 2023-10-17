/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus Kubeconfig status types.
//
//   - STATUS_UNSPECIFIED: Default value for the enum.
//   - CREATING: CREATING indicates either the cluster or TMC resources on the cluster are not ready yet.
//   - PENDING: PENDING indicates kubeconfig data is not yet available from the cluster.
//   - READY: READY indicates kubeconfig is ready to use.
//   - UNAVAILABLE: UNAVAILABLE indicates kubeconfig cannot be provided for this cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.kubeconfig.GetKubeconfigResponse.Status
type VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus string

func NewVmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus(value VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus) *VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus.
func (m VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus) Pointer() *VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusSTATUSUNSPECIFIED captures enum value "STATUS_UNSPECIFIED"
	VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusSTATUSUNSPECIFIED VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus = "STATUS_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusCREATING captures enum value "CREATING"
	VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusCREATING VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus = "CREATING"

	// VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusPENDING captures enum value "PENDING"
	VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusPENDING VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus = "PENDING"

	// VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusREADY captures enum value "READY"
	VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusREADY VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus = "READY"

	// VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusUNAVAILABLE captures enum value "UNAVAILABLE"
	VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusUNAVAILABLE VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus = "UNAVAILABLE"
)

// for schema
var vmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus
	if err := json.Unmarshal([]byte(`["STATUS_UNSPECIFIED","CREATING","PENDING","READY","UNAVAILABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusEnum = append(vmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusEnum, v)
	}
}
