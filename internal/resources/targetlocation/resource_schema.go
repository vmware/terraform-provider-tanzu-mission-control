/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocation

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const (
	ResourceName = "tanzu-mission-control_target_location"

	// Root Keys.
	NameKey         = "name"
	ProviderNameKey = "provider_name"
	SpecKey         = "spec"

	// Spec Directive Keys.
	BucketKey         = "bucket"
	RegionKey         = "region"
	CaCertKey         = "ca_cert"
	AssignedGroupsKey = "assigned_groups"
	ConfigKey         = "config"
	CredentialKey     = "credential"
	TargetProviderKey = "target_provider"
	SysBucketKey      = "sys_bucket_key"
	SysRegionKey      = "sys_region_key"

	// Assigned Groups Directive Keys.
	ClusterGroupsKey                 = "cluster_groups"
	ClusterKey                       = "cluster"
	ClustersManagementClusterNameKey = "management_cluster_name"
	ClustersProvisionerNameKey       = "provisioner_name"
	ClustersNameKey                  = "name"

	// Config Directive Keys.
	AwsConfigKey           = "aws"
	AwsS3ForcePathKey      = "s3_force_path_style"
	AwsS3BucketURLKey      = "s3_bucket_url"
	AwsS3PublicURLKey      = "s3_public_url"
	AzureConfigKey         = "azure"
	AzureResourceGroupKey  = "resource_group"
	AzureStorageAccountKey = "storage_account"
	AzureSubscriptionIDKey = "subscription_id"
)

var backupTargetLocationResourceSchema = map[string]*schema.Schema{
	NameKey:         nameSchema,
	ProviderNameKey: providerNameSchema,
	SpecKey:         specSchema,
	common.MetaKey:  common.Meta,
}

var nameSchema = &schema.Schema{
	Type:     schema.TypeString,
	Required: true,
	ForceNew: true,
}

var providerNameSchema = &schema.Schema{
	Type:     schema.TypeString,
	Required: true,
	ForceNew: true,
}

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec of enable backup target location",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			BucketKey: {
				Type:        schema.TypeString,
				Description: "Bucket as target location for backups",
				ForceNew:    true,
				Optional:    true,
			},
			SysBucketKey: {
				Type:        schema.TypeString,
				Description: "Bucket as target location for backups",
				Computed:    true,
			},
			RegionKey: {
				Type:        schema.TypeString,
				Description: "Bucket region",
				Optional:    true,
			},
			SysRegionKey: {
				Type:        schema.TypeString,
				Description: "Bucket as target location for backups",
				Computed:    true,
			},
			TargetProviderKey: {
				Type:        schema.TypeString,
				Description: "Target provider",
				Required:    true,
			},
			CaCertKey: {
				Type:        schema.TypeString,
				Description: "CA cert",
				Optional:    true,
			},
			CredentialKey:     credentialsSchema,
			AssignedGroupsKey: assignedGroupsSchema,
			ConfigKey:         configSchema,
		},
	},
}

var credentialsSchema = &schema.Schema{
	Type:        schema.TypeMap,
	Description: "Credentials block",
	Required:    true,
	ForceNew:    true,
	Elem: &schema.Schema{
		Type: schema.TypeString,
	},
}

var assignedGroupsSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	MinItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ClusterGroupsKey: {
				Type:        schema.TypeList,
				Description: "Cluster group names",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			ClusterKey: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Cluster objects",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ClustersNameKey: {
							Type:        schema.TypeString,
							Description: "Cluster name",
							Required:    true,
						},
						ClustersManagementClusterNameKey: {
							Type:        schema.TypeString,
							Description: "Management cluster name",
							Required:    true,
						},
						ClustersProvisionerNameKey: {
							Type:        schema.TypeString,
							Description: "Cluster provisioner name",
							Required:    true,
						},
					},
				},
			},
		},
	},
}

var configSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	MinItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			AwsConfigKey: {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						AwsS3ForcePathKey: {
							Type:     schema.TypeBool,
							Optional: true,
						},
						AwsS3BucketURLKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						AwsS3PublicURLKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			AzureConfigKey: {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						AzureResourceGroupKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						AzureSubscriptionIDKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						AzureStorageAccountKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	},
}
