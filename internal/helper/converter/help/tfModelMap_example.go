//go:build ignore
// +build ignore

package help

/*
######### tfModelMap Example #########
var tfModelMap = &tfModelConverterHelper.BlockToStruct{
	nameKey:         "fullName.name",
	providerNameKey: "fullName.providerName",
	specKey: &tfModelConverterHelper.BlockToStruct{
		bucketKey:         "spec.bucket",
		caCertKey:         "spec.caCert",
		targetProviderKey: "spec.targetProvider",
		regionKey:         "spec.region",
		configKey: &tfModelConverterHelper.BlockToStruct{
			awsConfigKey: &tfModelConverterHelper.BlockToStruct{
				awsS3PublicUrlKey: "spec.config.s3Config.publicUrl",
				awsS3ForcePathKey: "spec.config.s3Config.s3ForcePathStyle",
				awsS3BucketUrlKey: "spec.config.s3Config.s3Url",
			},
			azureConfigKey: &tfModelConverterHelper.BlockToStruct{
				azureResourceGroupKey:  "spec.config.azureConfig.resourceGroup",
				azureStorageAccountKey: "spec.config.azureConfig.storageAccount",
				azureSubscriptionIdKey: "spec.config.azureConfig.subscriptionId",
			},
		},
		credentialKey: &tfModelConverterHelper.Map{
			"name": "spec.credential.name",
		},
		assignedGroupsKey: &tfModelConverterHelper.BlockToStructSlice{
			{
				clusterKey: &tfModelConverterHelper.BlockToStruct{
					clustersManagementClusterNameKey: "spec.assignedGroups[].cluster.managementClusterName",
					clustersProvisionerNameKey:       "spec.assignedGroups[].cluster.provisionerName",
					clustersNameKey:                  "spec.assignedGroups[].cluster.name",
				},
			},
			{
				clusterGroupsKey: &tfModelConverterHelper.ListToStruct{"spec.assignedGroups[].clustergroup.name"},
			},
		},
	},
}
*/
