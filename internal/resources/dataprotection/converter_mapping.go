// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package dataprotection

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	dataprotectionmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/dataprotection"
	dataprotectioncgmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup/dataprotection"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection/scope"
)

var (
	matchExpressionsArrayField = tfModelConverterHelper.BuildArrayField("matchExpressions")
)

func getTFModelConverterCluster() tfModelConverterHelper.TFSchemaModelConverter[*dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection] {
	return tfModelConverterHelper.TFSchemaModelConverter[*dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionDataProtection]{
		TFModelMap: getTFModelMapCommon(false),
	}
}

func getTFModelConverterClusterGroup() tfModelConverterHelper.TFSchemaModelConverter[*dataprotectioncgmodels.VmwareTanzuManageV1alpha1ClustergroupDataprotectionDataProtection] {
	return tfModelConverterHelper.TFSchemaModelConverter[*dataprotectioncgmodels.VmwareTanzuManageV1alpha1ClustergroupDataprotectionDataProtection]{
		TFModelMap: getTFModelMapCommon(true),
	}
}

func getTFModelMapCommon(forClusterGroup bool) *tfModelConverterHelper.BlockToStruct {
	var specBlock *tfModelConverterHelper.BlockToStruct
	if forClusterGroup {
		specBlock = &tfModelConverterHelper.BlockToStruct{
			DisableResticKey:                   tfModelConverterHelper.BuildDefaultModelPath("spec", "atomicSpec", "disableRestic"),
			EnableCSISnapshotsKey:              tfModelConverterHelper.BuildDefaultModelPath("spec", "atomicSpec", "enableCsiSnapshots"),
			EnableAllAPIGroupVersionsBackupKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "atomicSpec", "enableAllApiGroupVersionsBackup"),
			UseNodeAgentKey:                    tfModelConverterHelper.BuildDefaultModelPath("spec", "atomicSpec", "useNodeAgent"),
			SelectorKey: &tfModelConverterHelper.BlockToStruct{
				NamesKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "selector", "names"),
				ExcludedNamesKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "selector", "excludedNames"),
				LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
					MatchExpressionsKey: &tfModelConverterHelper.BlockSliceToStructSlice{
						{
							KeyKey:      tfModelConverterHelper.BuildDefaultModelPath("spec", "selector", "labelSelector", matchExpressionsArrayField, "key"),
							OperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "selector", "labelSelector", matchExpressionsArrayField, "operator"),
							ValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "selector", "labelSelector", matchExpressionsArrayField, "values"),
						},
					},
				},
			},
		}
	} else {
		specBlock = &tfModelConverterHelper.BlockToStruct{
			DisableResticKey:                   tfModelConverterHelper.BuildDefaultModelPath("spec", "disableRestic"),
			EnableCSISnapshotsKey:              tfModelConverterHelper.BuildDefaultModelPath("spec", "enableCsiSnapshots"),
			EnableAllAPIGroupVersionsBackupKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "enableAllApiGroupVersionsBackup"),
			UseNodeAgentKey:                    tfModelConverterHelper.BuildDefaultModelPath("spec", "useNodeAgent"),
		}
	}

	return &tfModelConverterHelper.BlockToStruct{
		scope.ScopeKey: &tfModelConverterHelper.BlockToStruct{
			scope.ClusterKey: &tfModelConverterHelper.BlockToStruct{
				scope.ClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterName"),
				scope.ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
				scope.ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("fullName", "provisionerName"),
			},
			scope.ClusterGroupKey: &tfModelConverterHelper.BlockToStruct{
				scope.ClusterGroupNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterGroupName"),
			},
		},
		common.MetaKey: common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
		SpecKey:        specBlock,
	}
}
