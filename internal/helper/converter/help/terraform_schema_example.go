//go:build ignore
// +build ignore

package help

/*
######### Suggested Terraform Schema #########
var backupTargetLocationSchema = map[string]*schema.Schema{
	nameKey: &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	providerNameKey: &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	specKey: &schema.Schema{
		Type:        schema.TypeList,
		Description: "Spec of enable backup target location",
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				bucketKey: {
					Type:        schema.TypeString,
					Description: "Bucket as target location for backups",
					Required:    true,
				},
				regionKey: {
					Type:        schema.TypeString,
					Description: "CA cert",
					Required:    true,
				},
				targetProviderKey: {
					Type:        schema.TypeString,
					Description: "CA cert",
					Required:    true,
				},
				caCertKey: {
					Type:        schema.TypeString,
					Description: "CA cert",
					Optional:    true,
				},
				credentialKey: &schema.Schema{
					Type:        schema.TypeMap,
					Description: "Credentials block",
					Required:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				assignedGroupsKey: &schema.Schema{
					Type:     schema.TypeList,
					MaxItems: 1,
					MinItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							clusterGroupsKey: {
								Type:        schema.TypeList,
								Description: "Cluster group names",
								Optional:    true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
							clusterKey: {
								Type:        schema.TypeList,
								Optional:    true,
								Description: "Cluster objects",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										clustersNameKey: {
											Type:        schema.TypeString,
											Description: "Cluster name",
											Optional:    true,
										},
										clustersManagementClusterNameKey: {
											Type:        schema.TypeString,
											Description: "Management cluster name",
											Optional:    true,
										},
										clustersProvisionerNameKey: {
											Type:        schema.TypeString,
											Description: "Cluster provisioner name",
											Optional:    true,
										},
									},
								},
							},
						},
					},
				},
				configKey: {
					Type:     schema.TypeList,
					MaxItems: 1,
					MinItems: 1,
					Required: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							awsConfigKey: {
								Type:     schema.TypeList,
								Optional: true,
								MaxItems: 1,
								MinItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										awsS3ForcePathKey: {
											Type:     schema.TypeBool,
											Optional: true,
										},
										awsS3BucketUrlKey: {
											Type:     schema.TypeString,
											Optional: true,
										},
										awsS3PublicUrlKey: {
											Type:     schema.TypeString,
											Optional: true,
										},
									},
								},
							},
							azureConfigKey: {
								Type:     schema.TypeList,
								Optional: true,
								MaxItems: 1,
								MinItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										azureResourceGroupKey: {
											Type:     schema.TypeString,
											Optional: true,
										},
										azureSubscriptionIdKey: {
											Type:     schema.TypeString,
											Optional: true,
										},
										azureStorageAccountKey: {
											Type:     schema.TypeString,
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
	common.MetaKey: common.Meta,
}
*/
