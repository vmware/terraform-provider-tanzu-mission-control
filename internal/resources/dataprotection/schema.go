// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package dataprotection

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/dataprotection/scope"
)

const (
	ResourceName = "tanzu-mission-control_enable_data_protection"

	// Root Keys.
	SpecKey           = "spec"
	DeletionPolicyKey = "deletion_policy"

	// Spec Directive Keys.
	EnableCSISnapshotsKey              = "enable_csi_snapshots"
	DisableResticKey                   = "disable_restic"
	EnableAllAPIGroupVersionsBackupKey = "enable_all_api_group_versions_backup"
	UseNodeAgentKey                    = "use_node_agent"
	SelectorKey                        = "selector"

	// Selector keys.
	ExcludedNamesKey    = "excludednames"
	NamesKey            = "names"
	LabelSelectorKey    = "labelselector"
	MatchExpressionsKey = "matchexpressions"
	KeyKey              = "key"
	OperatorKey         = "operator"
	ValuesKey           = "values"

	// Deletion Policy Directive Keys.
	DeleteBackupsKey = "delete_backups"
	ForceDeleteKey   = "force"
)

var enableDataProtectionSchema = map[string]*schema.Schema{
	scope.ScopeKey:    scopeSchema,
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
			scope.ClusterGroupKey: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Cluster group scope block",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						scope.ClusterGroupNameKey: {
							Type:        schema.TypeString,
							Description: "Cluster group name",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			scope.ClusterKey: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Cluster scope block",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						scope.ClusterNameKey: {
							Type:        schema.TypeString,
							Description: "Cluster name",
							Required:    true,
							ForceNew:    true,
						},
						scope.ManagementClusterNameKey: {
							Type:        schema.TypeString,
							Description: "Management cluster name",
							Required:    true,
							ForceNew:    true,
						},
						scope.ProvisionerNameKey: {
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
				Optional:    true,
				Computed:    true,
			},
			DisableResticKey: {
				Type:        schema.TypeBool,
				Description: "A flag to indicate whether to skip installation of restic server (https://github.com/restic/restic).\nOtherwise, restic would be enabled by default as part of Data Protection installation.\n(Default: False)",
				Optional:    true,
				Computed:    true,
			},
			EnableAllAPIGroupVersionsBackupKey: {
				Type:        schema.TypeBool,
				Description: "A flag to indicate whether to backup all the supported API Group versions of a resource on the cluster.\n(Default: False)",
				Optional:    true,
				Computed:    true,
			},
			UseNodeAgentKey: {
				Type:        schema.TypeBool,
				Description: "A flag to indicate whether to install the node agent daemonset which is responsible for volume data transfer to the target location.",
				Optional:    true,
				Computed:    true,
			},
			SelectorKey: {
				Type:        schema.TypeList,
				Description: "A selector to include/exclude specific clusters in a cluster group (optional)",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ExcludedNamesKey: {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						LabelSelectorKey: {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									MatchExpressionsKey: {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												KeyKey: {
													Type:     schema.TypeString,
													Optional: true,
												},
												OperatorKey: {
													Type:     schema.TypeString,
													Optional: true,
												},
												ValuesKey: {
													Type:     schema.TypeList,
													Required: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
								},
							},
						},
						NamesKey: {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
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
			ForceDeleteKey: {
				Type:        schema.TypeBool,
				Description: "Disable data protection on all clusters in the cluster group even if cluster level schedules present.",
				Optional:    true,
				Computed:    true,
			},
		},
	},
}
