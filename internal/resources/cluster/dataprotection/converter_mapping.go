/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotection

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	dataprotectionmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/dataprotection"
)

var tfModelMap = &tfModelConverterHelper.BlockToStruct{
	ClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterName"),
	ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
	ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("fullName", "provisionerName"),
	SpecKey: &tfModelConverterHelper.BlockToStruct{
		DisableResticKey:                   tfModelConverterHelper.BuildDefaultModelPath("spec", "disableRestic"),
		EnableCSISnapshotsKey:              tfModelConverterHelper.BuildDefaultModelPath("spec", "enableCsiSnapshots"),
		EnableAllAPIGroupVersionsBackupKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "enableAllApiGroupVersionsBackup"),
	},
}

var tfModelConverter = tfModelConverterHelper.TFSchemaModelConverter[*dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection]{
	TFModelMap: tfModelMap,
}
