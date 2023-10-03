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
	nameKey:                  "fullName.name",
	clusterNameKey:           "fullName.clusterName",
	managementClusterNameKey: "fullName.managementClusterName",
	provisionerNameKey:       "fullName.provisionerName",
	specKey: &tfModelConverterHelper.BlockToStruct{
		pausedKey: "spec.paused",
		scheduleKey: &tfModelConverterHelper.BlockToStruct{
			rateKey: "spec.schedule.rate",
		},
		templateKey: &tfModelConverterHelper.BlockToStruct{
			csiSnapshotTimeoutKey:       "spec.template.csiSnapshotTimeout",
			defaultVolumesToFsBackupKey: "spec.template.defaultVolumesToFsBackup",
			defaultVolumesToResticKey:   "spec.template.defaultVolumesToRestic",
			excludedNamespacesKey:       "spec.template.excludedNamespaces",
			includedNamespacesKey:       "spec.template.includedNamespaces",
			excludedResourcesKey:        "spec.template.excludedResources",
			includedResourcesKey:        "spec.template.includedResources",
			includeClusterResourcesKey:  "spec.template.includeClusterResources",
			snapshotVolumesKey:          "spec.template.snapshotVolumes",
			storageLocationKey:          "spec.template.storageLocation",
			backupTTLKey:                "spec.template.ttl",
			volumeSnapshotLocationsKey:  "spec.template.volumeSnapshotLocations",
			hooksKey: &tfModelConverterHelper.BlockToStruct{
				resourceKey: &tfModelConverterHelper.BlockSliceToStructSlice{
					{
						excludedNamespacesKey: "spec.template.hooks.resources[].excludedNamespaces",
						includedNamespacesKey: "spec.template.hooks.resources[].includedNamespaces",
						nameKey:               "spec.template.hooks.resources[].name",
						postHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
							{
								execKey: &tfModelConverterHelper.BlockToStruct{
									commandKey:   "spec.template.hooks.resources[].postHooks[].exec.command",
									containerKey: "spec.template.hooks.resources[].postHooks[].exec.container",
									onErrorKey:   "spec.template.hooks.resources[].postHooks[].exec.onError",
									timeoutKey:   "spec.template.hooks.resources[].postHooks[].exec.timeout",
								},
							},
						},
						preHookKey: &tfModelConverterHelper.BlockSliceToStructSlice{
							{
								execKey: &tfModelConverterHelper.BlockToStruct{
									commandKey:   "spec.template.hooks.resources[].preHooks[].exec.command",
									containerKey: "spec.template.hooks.resources[].preHooks[].exec.container",
									onErrorKey:   "spec.template.hooks.resources[].preHooks[].exec.onError",
									timeoutKey:   "spec.template.hooks.resources[].preHooks[].exec.timeout",
								},
							},
						},
						labelSelectorKey: &tfModelConverterHelper.BlockToStruct{
							matchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
								{
									meKey:         "spec.template.hooks.resources[].labelSelector.matchExpressions[].key",
									meOperatorKey: "spec.template.hooks.resources[].labelSelector.matchExpressions[].operator",
									meValuesKey:   "spec.template.hooks.resources[].labelSelector.matchExpressions[].values",
								},
							},
							matchLabelsKey: &tfModelConverterHelper.Map{
								"*": "spec.template.hooks.resources[].labelSelector.matchLabels.*",
							},
						},
					},
				},
			},
			labelSelectorKey: &tfModelConverterHelper.BlockToStruct{
				matchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
					{
						meKey:         "spec.template.labelSelector.matchExpressions[].key",
						meOperatorKey: "spec.template.labelSelector.matchExpressions[].operator",
						meValuesKey:   "spec.template.labelSelector.matchExpressions[].values",
					},
				},
				matchLabelsKey: &tfModelConverterHelper.Map{
					"*": "spec.template.labelSelector.matchLabels.*",
				},
			},
			orLabelSelectorKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					matchExrpessionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
						{
							meKey:         "spec.template.orLabelSelectors[].matchExpressions[].key",
							meOperatorKey: "spec.template.orLabelSelectors[].matchExpressions[].operator",
							meValuesKey:   "spec.template.orLabelSelectors[].matchExpressions[].values",
						},
					},
					matchLabelsKey: &tfModelConverterHelper.Map{
						"*": "spec.template.orLabelSelectors[].matchLabels.*",
					},
				},
			},
			orderedResourcesKey: &tfModelConverterHelper.Map{
				"*": "spec.template.orderedResources.*",
			},
		},
	},
}

var tfModelDataSourceRequestMap = &tfModelConverterHelper.BlockToStruct{
	sortByKey:            "sortBy",
	queryKey:             "query",
	includeTotalCountKey: "includeTotal",
	scopeKey: &tfModelConverterHelper.BlockToStruct{
		clusterNameKey:           "searchScope.clusterScope.clusterName",
		managementClusterNameKey: "searchScope.clusterScope.managementClusterName",
		provisionerNameKey:       "searchScope.clusterScope.provisionerName",
		nameKey:                  "searchScope.clusterScope.name",
	},
}

var tfModelDataSourceResponseMap = &tfModelConverterHelper.BlockToStruct{
	schedulesKey: &tfModelConverterHelper.BlockSliceToStructSlice{
		// UNPACK tfModelResourceMap HERE.
	},
	totalCountKey: "totalCount",
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

	*(*tfModelDataSourceResponseMap)[schedulesKey].(*tfModelConverterHelper.BlockSliceToStructSlice) = append(
		*(*tfModelDataSourceResponseMap)[schedulesKey].(*tfModelConverterHelper.BlockSliceToStructSlice),
		targetLocationDataSourceSchema,
	)
}
