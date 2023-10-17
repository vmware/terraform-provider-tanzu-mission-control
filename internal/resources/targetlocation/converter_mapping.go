/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocation

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	targetlocationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/targetlocation"
)

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey: "fullName.name",
	SpecKey: &tfModelConverterHelper.BlockToStruct{
		BucketKey:         "spec.bucket",
		CaCertKey:         "spec.caCert",
		TargetProviderKey: "spec.targetProvider",
		RegionKey:         "spec.region",
		ConfigKey: &tfModelConverterHelper.BlockToStruct{
			AwsConfigKey: &tfModelConverterHelper.BlockToStruct{
				AwsS3PublicURLKey: "spec.config.s3Config.publicUrl",
				AwsS3ForcePathKey: "spec.config.s3Config.s3ForcePathStyle",
				AwsS3BucketURLKey: "spec.config.s3Config.s3Url",
			},
			AzureConfigKey: &tfModelConverterHelper.BlockToStruct{
				AzureResourceGroupKey:  "spec.config.azureConfig.resourceGroup",
				AzureStorageAccountKey: "spec.config.azureConfig.storageAccount",
				AzureSubscriptionIDKey: "spec.config.azureConfig.subscriptionId",
			},
		},
		CredentialKey: &tfModelConverterHelper.Map{
			"name": "spec.credential.name",
		},
		AssignedGroupsKey: &tfModelConverterHelper.BlockToStructSlice{
			{
				ClusterKey: &tfModelConverterHelper.BlockToStruct{
					ClustersManagementClusterNameKey: "spec.assignedGroups[].cluster.managementClusterName",
					ClustersProvisionerNameKey:       "spec.assignedGroups[].cluster.provisionerName",
					ClustersNameKey:                  "spec.assignedGroups[].cluster.name",
				},
			},
			{
				ClusterGroupsKey: &tfModelConverterHelper.ListToStruct{"spec.assignedGroups[].clustergroup.name"},
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
			ProviderScopeNameKey:              "searchScope.name",
			ProviderScopeCredentialNameKey:    "searchScope.credentialName",
			ProviderScopeAssignedGroupNameKey: "searchScope.assignedGroupName",
		},
		ClusterScopeKey: &tfModelConverterHelper.BlockToStruct{
			ClusterScopeClusterNameKey:           "searchScope.clusterName",
			ClusterScopeNameKey:                  "searchScope.name",
			ClusterScopeProvisionerNameKey:       "searchScope.provisionerName",
			ClusterScopeManagementClusterNameKey: "searchScope.managementClusterName",
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
	targetLocationDataSourceSchema := tfModelResourceConverter.UnpackSchema("backupLocations[]")

	*(*tfModelDataSourceResponseMap)[TargetLocationsKey].(*tfModelConverterHelper.BlockSliceToStructSlice) = append(
		*(*tfModelDataSourceResponseMap)[TargetLocationsKey].(*tfModelConverterHelper.BlockSliceToStructSlice),
		targetLocationDataSourceSchema,
	)
}
