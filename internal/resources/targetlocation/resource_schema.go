// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package targetlocation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	targetlocationmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/targetlocation"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const (
	ResourceName    = "tanzu-mission-control_target_location"
	TMCProviderName = "tmc"

	// Root Keys.
	NameKey = "name"
	SpecKey = "spec"

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

var TargetProviderValidValues = []string{
	string(targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderAWS),
	string(targetlocationmodels.VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderAZURE),
}

var backupTargetLocationResourceSchema = map[string]*schema.Schema{
	NameKey:        nameSchema,
	SpecKey:        specSchema,
	common.MetaKey: common.Meta,
}

var nameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "The name of the target location",
	Required:    true,
	ForceNew:    true,
}

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec block of backup target location",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			BucketKey: {
				Type:        schema.TypeString,
				Description: "The bucket to use for object storage.",
				ForceNew:    true,
				Optional:    true,
			},
			SysBucketKey: {
				Type:        schema.TypeString,
				Description: "System bucket to use for object storage.\n(Only used for Managed TMC)",
				Computed:    true,
			},
			RegionKey: {
				Type:        schema.TypeString,
				Description: "The region of the bucket origin.\nRequired only when target location is AWS Self Managed.",
				Optional:    true,
			},
			SysRegionKey: {
				Type:        schema.TypeString,
				Description: "System bucket region (Only used for Managed TMC)",
				Computed:    true,
			},
			TargetProviderKey: {
				Type:             schema.TypeString,
				Description:      fmt.Sprintf("The target provider of the backup storage.\nValid values are (%s)", strings.Join(TargetProviderValidValues, ", ")),
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(TargetProviderValidValues, false)),
			},
			CaCertKey: {
				Type:        schema.TypeString,
				Description: "A PEM-encoded certificate bundle to trust while connecting to the storage backend.",
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
		Type:        schema.TypeString,
		Description: "The name of credential to be used to access the bucket.",
	},
}

var assignedGroupsSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Assigned groups block for the target location.",
	MaxItems:    1,
	MinItems:    1,
	Optional:    true,
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
				Description: "(Repeatable Block) Cluster block.",
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
	Type:        schema.TypeList,
	Description: "Target location config block.\nRequired only when target location is Self Managed and should contain either AWS or Azure blocks but not both.",
	MaxItems:    1,
	MinItems:    1,
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			AwsConfigKey: {
				Type:        schema.TypeList,
				Description: "AWS S3 and S3-compatible target location config block.",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						AwsS3ForcePathKey: {
							Type:        schema.TypeBool,
							Description: "A flag for whether to force path style URLs for S3 objects.\nIt is default to false and set it to true when using local storage service like Minio.",
							Optional:    true,
							Default:     false,
						},
						AwsS3BucketURLKey: {
							Type:        schema.TypeString,
							Description: "The service endpoint for non-AWS S3 storage solution.",
							Optional:    true,
						},
						AwsS3PublicURLKey: {
							Type:        schema.TypeString,
							Description: "The service endpoint used for generating download URLs. This field is primarily for local storage services like Minio.",
							Optional:    true,
						},
					},
				},
			},
			AzureConfigKey: {
				Type:        schema.TypeList,
				Description: "Azure target location config block.",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						AzureResourceGroupKey: {
							Type:        schema.TypeString,
							Description: "Name of the resource group containing the storage account for this backup storage location.",
							Optional:    true,
						},
						AzureSubscriptionIDKey: {
							Type:        schema.TypeString,
							Description: "Name of the storage account for this backup storage location.",
							Optional:    true,
						},
						AzureStorageAccountKey: {
							Type:        schema.TypeString,
							Description: "Subscription ID under which all the resources are being managed in azure.",
							Optional:    true,
						},
					},
				},
			},
		},
	},
}
