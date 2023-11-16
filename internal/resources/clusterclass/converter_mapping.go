/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclass

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	clusterclassmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clusterclass"
)

var tfModelDataSourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey:                  tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
	ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("fullName", "provisionerName"),
	WorkerClassesKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "workersClasses"),
}

var tfModelDataSourceConverter = tfModelConverterHelper.TFSchemaModelConverter[*clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClass]{
	TFModelMap: tfModelDataSourceMap,
}
