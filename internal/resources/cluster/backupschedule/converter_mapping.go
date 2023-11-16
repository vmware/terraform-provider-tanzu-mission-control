/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	backupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/dataprotection/cluster/backupschedule"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var (
	resourcesArrayField        = tfModelConverterHelper.BuildArrayField("resources")
	postHooksArrayField        = tfModelConverterHelper.BuildArrayField("postHooks")
	preHooksArrayField         = tfModelConverterHelper.BuildArrayField("preHooks")
	matchExpressionsArrayField = tfModelConverterHelper.BuildArrayField("matchExpressions")
	orLabelSelectorsArrayField = tfModelConverterHelper.BuildArrayField("orLabelSelectors")
)

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey:                  tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	ClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterName"),
	ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
	ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("fullName", "provisionerName"),
	common.MetaKey:           common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	SpecKey: &tfModelConverterHelper.BlockToStruct{
		PausedKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "paused"),
		ScheduleKey: &tfModelConverterHelper.BlockToStruct{
			RateKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "schedule", "rate"),
		},
		TemplateKey: &tfModelConverterHelper.BlockToStruct{
			CsiSnapshotTimeoutKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "csiSnapshotTimeout"),
			DefaultVolumesToFsBackupKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "defaultVolumesToFsBackup"),
			DefaultVolumesToResticKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "defaultVolumesToRestic"),
			ExcludedNamespacesKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "excludedNamespaces"),
			IncludedNamespacesKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "includedNamespaces"),
			ExcludedResourcesKey:        tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "excludedResources"),
			IncludedResourcesKey:        tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "includedResources"),
			IncludeClusterResourcesKey:  tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "includeClusterResources"),
			SnapshotVolumesKey:          tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "snapshotVolumes"),
			StorageLocationKey:          tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "storageLocation"),
			BackupTTLKey:                tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "ttl"),
			VolumeSnapshotLocationsKey:  tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "volumeSnapshotLocations"),
			HooksKey: &tfModelConverterHelper.BlockToStruct{
				ResourceKey: &tfModelConverterHelper.BlockSliceToStructSlice{
					{
						ExcludedNamespacesKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "excludedNamespaces"),
						IncludedNamespacesKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "includedNamespaces"),
						NameKey:               tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "name"),
						PostHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
							{
								ExecKey: &tfModelConverterHelper.BlockToStruct{
									CommandKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "command"),
									ContainerKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "container"),
									OnErrorKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "onError"),
									TimeoutKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, postHooksArrayField, "exec", "timeout"),
								},
							},
						},
						PreHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
							{
								ExecKey: &tfModelConverterHelper.BlockToStruct{
									CommandKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "command"),
									ContainerKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "container"),
									OnErrorKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "onError"),
									TimeoutKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, preHooksArrayField, "exec", "timeout"),
								},
							},
						},
						LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
							MatchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
								{
									MeKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "key"),
									MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "operator"),
									MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "labelSelector", matchExpressionsArrayField, "values"),
								},
							},
							MatchLabelsKey: &tfModelConverterHelper.Map{
								tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "hooks", resourcesArrayField, "labelSelector", "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
							},
						},
					},
				},
			},
			LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
				MatchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
					{
						MeKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "labelSelector", matchExpressionsArrayField, "key"),
						MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "labelSelector", matchExpressionsArrayField, "operator"),
						MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "labelSelector", matchExpressionsArrayField, "values"),
					},
				},
				MatchLabelsKey: &tfModelConverterHelper.Map{
					tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "labelSelector", "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
				},
			},
			OrLabelSelectorKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					MatchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
						{
							MeKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "template", orLabelSelectorsArrayField, matchExpressionsArrayField, "key"),
							MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", orLabelSelectorsArrayField, matchExpressionsArrayField, "operator"),
							MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "template", orLabelSelectorsArrayField, matchExpressionsArrayField, "values"),
						},
					},
					MatchLabelsKey: &tfModelConverterHelper.Map{
						tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", orLabelSelectorsArrayField, "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
					},
				},
			},
			OrderedResourcesKey: &tfModelConverterHelper.Map{
				tfModelConverterHelper.ArrayFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "template", "orderedResources", tfModelConverterHelper.AllMapKeysFieldMarker),
			},
		},
	},
}

var tfModelDataSourceRequestMap = &tfModelConverterHelper.BlockToStruct{
	SortByKey:            "sortBy",
	QueryKey:             "query",
	IncludeTotalCountKey: "includeTotal",
	ScopeKey: &tfModelConverterHelper.BlockToStruct{
		ClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("searchScope", "clusterName"),
		ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("searchScope", "managementClusterName"),
		ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("searchScope", "provisionerName"),
		NameKey:                  tfModelConverterHelper.BuildDefaultModelPath("searchScope", "name"),
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

var tfModelDataSourceRequestConverter = tfModelConverterHelper.TFSchemaModelConverter[*backupschedulemodels.ListBackupSchedulesRequest]{
	TFModelMap: tfModelDataSourceRequestMap,
}

var tfModelDataSourceResponseConverter = tfModelConverterHelper.TFSchemaModelConverter[*backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionScheduleListSchedulesResponse]{
	TFModelMap: tfModelDataSourceResponseMap,
}

func constructTFModelDataSourceResponseMap() {
	targetLocationDataSourceSchema := tfModelResourceConverter.UnpackSchema(tfModelConverterHelper.BuildArrayField("schedules"))

	*(*tfModelDataSourceResponseMap)[SchedulesKey].(*tfModelConverterHelper.BlockSliceToStructSlice) = append(
		*(*tfModelDataSourceResponseMap)[SchedulesKey].(*tfModelConverterHelper.BlockSliceToStructSlice),
		targetLocationDataSourceSchema,
	)
}
