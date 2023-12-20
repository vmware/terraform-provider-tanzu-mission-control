/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	backupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backupschedule/cluster"
	cgbackupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backupschedule/clustergroup"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var (
	resourcesArrayField        = tfModelConverterHelper.BuildArrayField("resources")
	postHooksArrayField        = tfModelConverterHelper.BuildArrayField("postHooks")
	preHooksArrayField         = tfModelConverterHelper.BuildArrayField("preHooks")
	matchExpressionsArrayField = tfModelConverterHelper.BuildArrayField("matchExpressions")
	orLabelSelectorsArrayField = tfModelConverterHelper.BuildArrayField("orLabelSelectors")
)

func getTfModelResourceSpecMap(rootPath string) *tfModelConverterHelper.BlockToStruct {
	return &tfModelConverterHelper.BlockToStruct{
		PausedKey: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "paused"),
		ScheduleKey: &tfModelConverterHelper.BlockToStruct{
			RateKey: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "schedule", "rate"),
		},
		TemplateKey: &tfModelConverterHelper.BlockToStruct{
			CsiSnapshotTimeoutKey:               tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "csiSnapshotTimeout"),
			DefaultVolumesToFsBackupKey:         tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "defaultVolumesToFsBackup"),
			DefaultVolumesToResticKey:           tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "defaultVolumesToRestic"),
			ExcludedNamespacesKey:               tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "excludedNamespaces"),
			IncludedNamespacesKey:               tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "includedNamespaces"),
			ExcludedResourcesKey:                tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "excludedResources"),
			IncludedResourcesKey:                tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "includedResources"),
			IncludeClusterResourcesKey:          tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "includeClusterResources"),
			SnapshotVolumesKey:                  tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "snapshotVolumes"),
			StorageLocationKey:                  tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "storageLocation"),
			BackupTTLKey:                        tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "ttl"),
			VolumeSnapshotLocationsKey:          tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "volumeSnapshotLocations"),
			IncludedClusterScopedResourcesKey:   tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "includedClusterScopedResources"),
			ExcludedClusterScopedResourcesKey:   tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "excludedClusterScopedResources"),
			IncludedNamespaceScopedResourcesKey: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "includedNamespaceScopedResources"),
			ExcludedNamespaceScopedResourcesKey: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "excludedNamespaceScopedResources"),
			SnapshotMoveDataKey:                 tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "snapshotMoveData"),
			HooksKey: &tfModelConverterHelper.BlockToStruct{
				ResourceKey: &tfModelConverterHelper.BlockSliceToStructSlice{
					{
						ExcludedNamespacesKey: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, "excludedNamespaces"),
						IncludedNamespacesKey: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, "includedNamespaces"),
						NameKey:               tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, "name"),
						PostHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
							{
								ExecKey: &tfModelConverterHelper.BlockToStruct{
									CommandKey:   tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "command"),
									ContainerKey: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "container"),
									OnErrorKey:   tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "onError"),
									TimeoutKey:   tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "timeout"),
								},
							},
						},
						PreHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
							{
								ExecKey: &tfModelConverterHelper.BlockToStruct{
									CommandKey:   tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "command"),
									ContainerKey: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "container"),
									OnErrorKey:   tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "onError"),
									TimeoutKey:   tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "timeout"),
								},
							},
						},
						LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
							MatchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
								{
									MeKey:         tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "key"),
									MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "operator"),
									MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "values"),
								},
							},
							MatchLabelsKey: &tfModelConverterHelper.Map{
								tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "hooks", resourcesArrayField, "labelSelector", "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
							},
						},
					},
				},
			},
			LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
				MatchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
					{
						MeKey:         tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "labelSelector", matchExpressionsArrayField, "key"),
						MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "labelSelector", matchExpressionsArrayField, "operator"),
						MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "labelSelector", matchExpressionsArrayField, "values"),
					},
				},
				MatchLabelsKey: &tfModelConverterHelper.Map{
					tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "labelSelector", "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
				},
			},
			OrLabelSelectorKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					MatchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
						{
							MeKey:         tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", orLabelSelectorsArrayField, matchExpressionsArrayField, "key"),
							MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", orLabelSelectorsArrayField, matchExpressionsArrayField, "operator"),
							MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", orLabelSelectorsArrayField, matchExpressionsArrayField, "values"),
						},
					},
					MatchLabelsKey: &tfModelConverterHelper.Map{
						tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", orLabelSelectorsArrayField, "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
					},
				},
			},
			OrderedResourcesKey: &tfModelConverterHelper.Map{
				tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath(rootPath, "template", "orderedResources", tfModelConverterHelper.AllMapKeysFieldMarker),
			},
		},
	}
}

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	ScopeKey: &tfModelConverterHelper.BlockToStruct{
		ClusterGroupScopeKey: &tfModelConverterHelper.BlockToStruct{
			ClusterGroupNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterGroupName"),
		},
		ClusterScopeKey: &tfModelConverterHelper.BlockToStruct{
			ClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterName"),
			ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
			ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("fullName", "provisionerName"),
		},
	},
	common.MetaKey: common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	SpecKey:        getTfModelResourceSpecMap("spec"),
}

var tfModelCGResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	ScopeKey: &tfModelConverterHelper.BlockToStruct{
		ClusterGroupScopeKey: &tfModelConverterHelper.BlockToStruct{
			ClusterGroupNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterGroupName"),
		},
		ClusterScopeKey: &tfModelConverterHelper.BlockToStruct{
			ClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterName"),
			ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
			ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("fullName", "provisionerName"),
		},
	},
	common.MetaKey: common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	SpecKey:        getTfModelResourceSpecMap("spec" + tfModelConverterHelper.DefaultModelPathSeparator + "atomicSpec"),
	SelectorKey: &tfModelConverterHelper.BlockToStruct{
		NamesKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "selector", "names"),
		ExcludedNamesKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "selector", "excludedNames"),
		LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
			MatchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					MeKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "selector", "labelSelector", matchExpressionsArrayField, "key"),
					MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "selector", "labelSelector", matchExpressionsArrayField, "operator"),
					MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "selector", "labelSelector", matchExpressionsArrayField, "values"),
				},
			},
		},
	},
}

var tfModelDataSourceRequestMap = &tfModelConverterHelper.BlockToStruct{
	SortByKey:            "sortBy",
	QueryKey:             "query",
	IncludeTotalCountKey: "includeTotal",
	NameKey:              tfModelConverterHelper.BuildDefaultModelPath("searchScope", "name"),
	ScopeKey: &tfModelConverterHelper.BlockToStruct{
		ClusterGroupScopeKey: &tfModelConverterHelper.BlockToStruct{
			ClusterGroupNameKey: tfModelConverterHelper.BuildDefaultModelPath("searchScope", "clusterGroupName"),
		},
		ClusterScopeKey: &tfModelConverterHelper.BlockToStruct{
			ClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("searchScope", "clusterName"),
			ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("searchScope", "managementClusterName"),
			ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("searchScope", "provisionerName"),
		},
	},
}

var tfModelDataSourceResponseMap = &tfModelConverterHelper.BlockToStruct{
	SchedulesKey: &tfModelConverterHelper.BlockSliceToStructSlice{
		// UNPACK tfModelResourceMap HERE.
	},
	TotalCountKey: "totalCount",
}

var tfModelResourceConverter = tfModelConverterHelper.TFSchemaModelConverter[*backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleSchedule]{
	TFModelMap: tfModelResourceMap,
}

var tfModelCGResourceConverter = tfModelConverterHelper.TFSchemaModelConverter[*cgbackupschedulemodels.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleSchedule]{
	TFModelMap: tfModelCGResourceMap,
}

var tfModelDataSourceRequestConverter = tfModelConverterHelper.TFSchemaModelConverter[*backupschedulemodels.ListBackupSchedulesRequest]{
	TFModelMap: tfModelDataSourceRequestMap,
}

var tfModelDataSourceResponseConverter = tfModelConverterHelper.TFSchemaModelConverter[*backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse]{
	TFModelMap: tfModelDataSourceResponseMap,
}

var tfModelCGDataSourceRequestConverter = tfModelConverterHelper.TFSchemaModelConverter[*cgbackupschedulemodels.ListClusterGroupBackupSchedulesRequest]{
	TFModelMap: tfModelDataSourceRequestMap,
}

var tfModelCGDataSourceResponseConverter = tfModelConverterHelper.TFSchemaModelConverter[*cgbackupschedulemodels.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleListSchedulesResponse]{
	TFModelMap: tfModelDataSourceResponseMap,
}

func constructTFModelDataSourceResponseMap() {
	targetLocationDataSourceSchema := tfModelResourceConverter.UnpackSchema(tfModelConverterHelper.BuildArrayField("schedules"))

	*(*tfModelDataSourceResponseMap)[SchedulesKey].(*tfModelConverterHelper.BlockSliceToStructSlice) = append(
		*(*tfModelDataSourceResponseMap)[SchedulesKey].(*tfModelConverterHelper.BlockSliceToStructSlice),
		targetLocationDataSourceSchema,
	)
}
