/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	backupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/backupschedule"
)

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey:                  "fullName.name",
	ClusterNameKey:           "fullName.clusterName",
	ManagementClusterNameKey: "fullName.managementClusterName",
	ProvisionerNameKey:       "fullName.provisionerName",
	SpecKey: &tfModelConverterHelper.BlockToStruct{
		PausedKey: "spec.paused",
		ScheduleKey: &tfModelConverterHelper.BlockToStruct{
			RateKey: "spec.schedule.rate",
		},
		TemplateKey: &tfModelConverterHelper.BlockToStruct{
			CsiSnapshotTimeoutKey:       "spec.template.csiSnapshotTimeout",
			DefaultVolumesToFsBackupKey: "spec.template.defaultVolumesToFsBackup",
			DefaultVolumesToResticKey:   "spec.template.defaultVolumesToRestic",
			ExcludedNamespacesKey:       "spec.template.excludedNamespaces",
			IncludedNamespacesKey:       "spec.template.includedNamespaces",
			ExcludedResourcesKey:        "spec.template.excludedResources",
			IncludedResourcesKey:        "spec.template.includedResources",
			IncludeClusterResourcesKey:  "spec.template.includeClusterResources",
			SnapshotVolumesKey:          "spec.template.snapshotVolumes",
			StorageLocationKey:          "spec.template.storageLocation",
			BackupTTLKey:                "spec.template.ttl",
			VolumeSnapshotLocationsKey:  "spec.template.volumeSnapshotLocations",
			HooksKey: &tfModelConverterHelper.BlockToStruct{
				ResourceKey: &tfModelConverterHelper.BlockSliceToStructSlice{
					{
						ExcludedNamespacesKey: "spec.template.hooks.resources[].excludedNamespaces",
						IncludedNamespacesKey: "spec.template.hooks.resources[].includedNamespaces",
						NameKey:               "spec.template.hooks.resources[].name",
						PostHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
							{
								ExecKey: &tfModelConverterHelper.BlockToStruct{
									CommandKey:   "spec.template.hooks.resources[].postHooks[].exec.command",
									ContainerKey: "spec.template.hooks.resources[].postHooks[].exec.container",
									OnErrorKey:   "spec.template.hooks.resources[].postHooks[].exec.onError",
									TimeoutKey:   "spec.template.hooks.resources[].postHooks[].exec.timeout",
								},
							},
						},
						PreHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
							{
								ExecKey: &tfModelConverterHelper.BlockToStruct{
									CommandKey:   "spec.template.hooks.resources[].preHooks[].exec.command",
									ContainerKey: "spec.template.hooks.resources[].preHooks[].exec.container",
									OnErrorKey:   "spec.template.hooks.resources[].preHooks[].exec.onError",
									TimeoutKey:   "spec.template.hooks.resources[].preHooks[].exec.timeout",
								},
							},
						},
						LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
							MatchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
								{
									MeKey:         "spec.template.hooks.resources[].labelSelector.matchExpressions[].key",
									MeOperatorKey: "spec.template.hooks.resources[].labelSelector.matchExpressions[].operator",
									MeValuesKey:   "spec.template.hooks.resources[].labelSelector.matchExpressions[].values",
								},
							},
							MatchLabelsKey: &tfModelConverterHelper.Map{
								"*": "spec.template.hooks.resources[].labelSelector.matchLabels.*",
							},
						},
					},
				},
			},
			LabelSelectorKey: &tfModelConverterHelper.BlockToStruct{
				MatchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
					{
						MeKey:         "spec.template.labelSelector.matchExpressions[].key",
						MeOperatorKey: "spec.template.labelSelector.matchExpressions[].operator",
						MeValuesKey:   "spec.template.labelSelector.matchExpressions[].values",
					},
				},
				MatchLabelsKey: &tfModelConverterHelper.Map{
					"*": "spec.template.labelSelector.matchLabels.*",
				},
			},
			OrLabelSelectorKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					MatchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
						{
							MeKey:         "spec.template.orLabelSelectors[].matchExpressions[].key",
							MeOperatorKey: "spec.template.orLabelSelectors[].matchExpressions[].operator",
							MeValuesKey:   "spec.template.orLabelSelectors[].matchExpressions[].values",
						},
					},
					MatchLabelsKey: &tfModelConverterHelper.Map{
						"*": "spec.template.orLabelSelectors[].matchLabels.*",
					},
				},
			},
			OrderedResourcesKey: &tfModelConverterHelper.Map{
				"*": "spec.template.orderedResources.*",
			},
		},
	},
}

var tfModelDataSourceRequestMap = &tfModelConverterHelper.BlockToStruct{
	SortByKey:            "sortBy",
	QueryKey:             "query",
	IncludeTotalCountKey: "includeTotal",
	ScopeKey: &tfModelConverterHelper.BlockToStruct{
		ClusterNameKey:           "searchScope.clusterName",
		ManagementClusterNameKey: "searchScope.managementClusterName",
		ProvisionerNameKey:       "searchScope.provisionerName",
		NameKey:                  "searchScope.name",
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
	targetLocationDataSourceSchema := tfModelResourceConverter.UnpackSchema("schedules[]")

	*(*tfModelDataSourceResponseMap)[SchedulesKey].(*tfModelConverterHelper.BlockSliceToStructSlice) = append(
		*(*tfModelDataSourceResponseMap)[SchedulesKey].(*tfModelConverterHelper.BlockSliceToStructSlice),
		targetLocationDataSourceSchema,
	)
}
