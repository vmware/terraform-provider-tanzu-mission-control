/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package provisioner

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	provisionermodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/provisioner"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var tfModelMap = &tfModelConverterHelper.BlockToStruct{
	nameKey:                  tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	managementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
	common.MetaKey:           common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
}

var tfModelConverter = tfModelConverterHelper.TFSchemaModelConverter[*provisionermodels.VmwareTanzuManageV1alpha1ManagementclusterProvisionerProvisioner]{
	TFModelMap: tfModelMap,
}
