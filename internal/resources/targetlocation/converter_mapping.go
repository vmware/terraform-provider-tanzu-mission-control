/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocation

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	targetlocationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/targetlocation"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var assignedGroupsArrayField = tfModelConverterHelper.BuildArrayField("assignedGroups")

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey:        tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	common.MetaKey: common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	SpecKey: &tfModelConverterHelper.BlockToStruct{
		BucketKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "bucket"),
		CaCertKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "caCert"),
		TargetProviderKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "targetProvider"),
		RegionKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "region"),
		ConfigKey: &tfModelConverterHelper.BlockToStruct{
			AwsConfigKey: &tfModelConverterHelper.BlockToStruct{
				AwsS3PublicURLKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "config", "s3Config", "publicUrl"),
				AwsS3ForcePathKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "config", "s3Config", "s3ForcePathStyle"),
				AwsS3BucketURLKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "config", "s3Config", "s3Url"),
			},
			AzureConfigKey: &tfModelConverterHelper.BlockToStruct{
				AzureResourceGroupKey:  tfModelConverterHelper.BuildDefaultModelPath("spec", "config", "azureConfig", "resourceGroup"),
				AzureStorageAccountKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "config", "azureConfig", "storageAccount"),
				AzureSubscriptionIDKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "config", "azureConfig", "subscriptionId"),
			},
		},
		CredentialKey: &tfModelConverterHelper.Map{
			"name": tfModelConverterHelper.BuildDefaultModelPath("spec", "credential", "name"),
		},
		AssignedGroupsKey: &tfModelConverterHelper.BlockToStructSlice{
			{
				ClusterKey: &tfModelConverterHelper.BlockToStruct{
					ClustersManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("spec", assignedGroupsArrayField, "cluster", "managementClusterName"),
					ClustersProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", assignedGroupsArrayField, "cluster", "provisionerName"),
					ClustersNameKey:                  tfModelConverterHelper.BuildDefaultModelPath("spec", assignedGroupsArrayField, "cluster", "name"),
				},
			},
			{
				ClusterGroupsKey: &tfModelConverterHelper.ListToStruct{tfModelConverterHelper.BuildDefaultModelPath("spec", assignedGroupsArrayField, "clustergroup", "name")},
			},
		},
	},
}

var tfModelDataSourceRequestMap = &tfModelConverterHelper.BlockToStruct{
	SortByKey:            "sortBy",
	QueryKey:             "query",
	IncludeTotalCountKey: "includeTotal",
	ScopeKey: &tfModelConverterHelper.BlockToStruct{
		ProviderScopeKey: &tfModelConverterHelper.BlockToStruct{
			ProviderScopeNameKey:              tfModelConverterHelper.BuildDefaultModelPath("searchScope", "name"),
			ProviderScopeCredentialNameKey:    tfModelConverterHelper.BuildDefaultModelPath("searchScope", "credentialName"),
			ProviderScopeAssignedGroupNameKey: tfModelConverterHelper.BuildDefaultModelPath("searchScope", "assignedGroupName"),
		},
		ClusterScopeKey: &tfModelConverterHelper.BlockToStruct{
			ClusterScopeClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("searchScope", "clusterName"),
			ClusterScopeNameKey:                  tfModelConverterHelper.BuildDefaultModelPath("searchScope", "name"),
			ClusterScopeProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("searchScope", "provisionerName"),
			ClusterScopeManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("searchScope", "managementClusterName"),
		},
	},
}

var tfModelDataSourceResponseMap = &tfModelConverterHelper.BlockToStruct{
	TargetLocationsKey: &tfModelConverterHelper.BlockSliceToStructSlice{
		// UNPACK tfModelResourceMap HERE.
	},
	TotalCountKey: "totalCount",
}

var tfModelResourceConverter = tfModelConverterHelper.TFSchemaModelConverter[*targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationBackupLocation]{
	TFModelMap: tfModelResourceMap,
}

var tfModelDataSourceRequestConverter = tfModelConverterHelper.TFSchemaModelConverter[*targetlocationmodels.ListBackupLocationsRequest]{
	TFModelMap: tfModelDataSourceRequestMap,
}

var tfModelDataSourceResponseConverter = tfModelConverterHelper.TFSchemaModelConverter[*targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationListBackupLocationsResponse]{
	TFModelMap: tfModelDataSourceResponseMap,
}

func constructTFModelDataSourceResponseMap() {
	targetLocationDataSourceSchema := tfModelResourceConverter.UnpackSchema(tfModelConverterHelper.BuildArrayField("backupLocations"))

	*(*tfModelDataSourceResponseMap)[TargetLocationsKey].(*tfModelConverterHelper.BlockSliceToStructSlice) = append(
		*(*tfModelDataSourceResponseMap)[TargetLocationsKey].(*tfModelConverterHelper.BlockSliceToStructSlice),
		targetLocationDataSourceSchema,
	)
}
