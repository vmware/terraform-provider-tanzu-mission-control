/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotection

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const (
	ResourceName         = "tanzu-mission-control_enable_data_protection"
	ScopeKey             = "scope"
	ClusterScopeKey      = "cluster"
	ClusterGroupScopeKey = "cluster_group"

	// Root Keys.
	ClusterNameKey           = "cluster_name"
	ClusterGroupNameKey      = "cluster_group_name"
	ManagementClusterNameKey = "management_cluster_name"
	ProvisionerNameKey       = "provisioner_name"
	SpecKey                  = "spec"
	DeletionPolicyKey        = "deletion_policy"

	// Spec Directive Keys.
	EnableCSISnapshotsKey              = "enable_csi_snapshots"
	DisableResticKey                   = "disable_restic"
	EnableAllAPIGroupVersionsBackupKey = "enable_all_api_group_versions_backup"

	// Deletion Policy Directive Keys.
	DeleteBackupsKey = "delete_backups"
)

var enableDataProtectionSchema = map[string]*schema.Schema{
	ScopeKey:          scopeSchema,
	SpecKey:           specSchema,
	common.MetaKey:    common.Meta,
	DeletionPolicyKey: deletionPolicySchema,
}

var scopeSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Scope block for Data Protection (cluster/cluster group)",
	Required:    true,
	MaxItems:    1,
	Optional:    false,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ClusterGroupScopeKey: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Cluster group scope block",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ClusterGroupNameKey: {
							Type:        schema.TypeString,
							Description: "Cluster group name",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			ClusterScopeKey: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Cluster scope block",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ClusterNameKey: {
							Type:        schema.TypeString,
							Description: "Cluster name",
							Required:    true,
							ForceNew:    true,
						},
						ManagementClusterNameKey: {
							Type:        schema.TypeString,
							Description: "Management cluster name",
							Required:    true,
							ForceNew:    true,
						},
						ProvisionerNameKey: {
							Type:        schema.TypeString,
							Description: "Cluster provisioner name",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	},
}

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec of enable data protection",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			EnableCSISnapshotsKey: {
				Type:        schema.TypeBool,
				Description: "A flag to indicate whether to install CSI snapshotting related capabilities.\n(Default: False)",
				Default:     false,
				Optional:    true,
			},
			DisableResticKey: {
				Type:        schema.TypeBool,
				Description: "A flag to indicate whether to skip installation of restic server (https://github.com/restic/restic).\nOtherwise, restic would be enabled by default as part of Data Protection installation.\n(Default: False)",
				Default:     false,
				Optional:    true,
			},
			EnableAllAPIGroupVersionsBackupKey: {
				Type:        schema.TypeBool,
				Description: "A flag to indicate whether to backup all the supported API Group versions of a resource on the cluster.\n(Default: False)",
				Default:     false,
				Optional:    true,
			},
		},
	},
}

var deletionPolicySchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Deletion policy block of data protection",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			DeleteBackupsKey: {
				Type:        schema.TypeBool,
				Description: "Destroy backups upon deleting data protection.\n(default: false)",
				Default:     false,
				Optional:    true,
			},
		},
	},
}
