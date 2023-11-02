/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterKubeconfigCliType CliType will help identify the appropriate kubeconfig to be returned for CLI application.
//
//   - TMC_CLI: TMC_CLI platform.
//   - TANZU_CLI: TANZU_CLI platform.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.kubeconfig.CliType
type VmwareTanzuManageV1alpha1ClusterKubeconfigCliType string

func NewVmwareTanzuManageV1alpha1ClusterKubeconfigCliType(value VmwareTanzuManageV1alpha1ClusterKubeconfigCliType) *VmwareTanzuManageV1alpha1ClusterKubeconfigCliType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterKubeconfigCliType.
func (m VmwareTanzuManageV1alpha1ClusterKubeconfigCliType) Pointer() *VmwareTanzuManageV1alpha1ClusterKubeconfigCliType {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterKubeconfigCliTypeTMCCLI captures enum value "TMC_CLI".
	VmwareTanzuManageV1alpha1ClusterKubeconfigCliTypeTMCCLI VmwareTanzuManageV1alpha1ClusterKubeconfigCliType = "TMC_CLI"

	// VmwareTanzuManageV1alpha1ClusterKubeconfigCliTypeTANZUCLI captures enum value "TANZU_CLI".
	VmwareTanzuManageV1alpha1ClusterKubeconfigCliTypeTANZUCLI VmwareTanzuManageV1alpha1ClusterKubeconfigCliType = "TANZU_CLI"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterKubeconfigCliTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterKubeconfigCliType
	if err := json.Unmarshal([]byte(`["TMC_CLI","TANZU_CLI"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterKubeconfigCliTypeEnum = append(vmwareTanzuManageV1alpha1ClusterKubeconfigCliTypeEnum, v)
	}
}
