/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type BackupScope string

const (
	fullClusterBackupScope   BackupScope = "FULL_CLUSTER"
	namespacesBackupScope    BackupScope = "SET_NAMESPACES"
	labelSelectorBackupScope BackupScope = "LABEL_SELECTOR"
)

const (
	ResourceName = "tanzu-mission-control_backup_schedule"

	// Root Keys.
	nameKey                  = "name"
	clusterNameKey           = "cluster_name"
	specKey                  = "spec"
	provisionerNameKey       = "provisioner_name"
	managementClusterNameKey = "management_cluster_name"
	scopeKey                 = "scope"

	// Spec Directive Keys.
	pausedKey   = "paused"
	scheduleKey = "schedule"
	templateKey = "template"

	// Schedule Directive Keys.
	rateKey = "rate"

	// Template Directive Keys.
	backupTTLKey                = "backup_ttl"
	systemExcludedNamespacesKey = "sys_excluded_namespaces"
	excludedNamespacesKey       = "excluded_namespaces"
	includedNamespacesKey       = "included_namespaces"
	excludedResourcesKey        = "excluded_resources"
	includedResourcesKey        = "included_resources"
	includeClusterResourcesKey  = "include_cluster_resources"
	defaultVolumesToResticKey   = "default_volumes_to_restic"
	snapshotVolumesKey          = "snapshot_volumes"
	csiSnapshotTimeoutKey       = "csi_snapshot_timeout"
	defaultVolumesToFsBackupKey = "default_volumes_to_fs_backup"
	storageLocationKey          = "storage_location"
	volumeSnapshotLocationsKey  = "volume_snapshot_locations"
	orderedResourcesKey         = "ordered_resources"
	hooksKey                    = "hooks"
	labelSelectorKey            = "label_selector"
	orLabelSelectorKey          = "or_label_selector"

	// Hooks Directive Keys.
	resourceKey = "resource"

	// Resource Directive Keys.
	preHookKey  = "pre_hook"
	postHookKey = "post_hook"

	// Pre/Post Hook Directive Keys.
	execKey = "exec"

	// Exec Directive Keys.
	commandKey   = "command"
	containerKey = "container"
	onErrorKey   = "on_error"
	timeoutKey   = "timeout"

	// (Or)Label Selector Directive Keys.
	matchLabelsKey     = "match_labels"
	matchExrpessionKey = "match_expression"

	// Match Expressions Directive Keys.
	meKey         = "key"
	meOperatorKey = "operator"
	meValuesKey   = "values"
)

var backupScheduleResourceSchema = map[string]*schema.Schema{
	nameKey:                  nameSchema,
	managementClusterNameKey: managementClusterNameSchema,
	provisionerNameKey:       provisionerNameSchema,
	clusterNameKey:           clusterNameSchema,
	scopeKey:                 backupScopeSchema,
	specKey:                  specSchema,
	common.MetaKey:           common.Meta,
}

var nameSchema = &schema.Schema{
	Type:     schema.TypeString,
	Required: true,
	ForceNew: true,
}

var managementClusterNameSchema = &schema.Schema{
	Type:     schema.TypeString,
	Required: true,
	ForceNew: true,
}

var provisionerNameSchema = &schema.Schema{
	Type:     schema.TypeString,
	Required: true,
	ForceNew: true,
}

var clusterNameSchema = &schema.Schema{
	Type:     schema.TypeString,
	Required: true,
	ForceNew: true,
}

var backupScopeSchema = &schema.Schema{
	Type:     schema.TypeString,
	Required: true,
	ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
		string(fullClusterBackupScope),
		string(namespacesBackupScope),
		string(labelSelectorBackupScope),
	}, false)),
}

var specSchema = &schema.Schema{
	Type:     schema.TypeList,
	Required: true,
	MaxItems: 1,
	MinItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			pausedKey: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			scheduleKey: {
				Type:     schema.TypeList,
				MaxItems: 1,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						rateKey: {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			templateKey: templateSchema,
		},
	},
}

var templateSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	MinItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			csiSnapshotTimeoutKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			defaultVolumesToFsBackupKey: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			defaultVolumesToResticKey: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			excludedNamespacesKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			systemExcludedNamespacesKey: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			excludedResourcesKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			includedNamespacesKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			includedResourcesKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			includeClusterResourcesKey: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			orderedResourcesKey: {
				Type:     schema.TypeMap,
				Optional: true,
			},
			orLabelSelectorKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: labelSelectorResource,
				},
			},
			labelSelectorKey: labelSelectorSchema,
			hooksKey:         hooksSchema,
			snapshotVolumesKey: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			storageLocationKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			backupTTLKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			volumeSnapshotLocationsKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	},
}

var hooksSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			resourceKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						nameKey: nameSchema,
						excludedNamespacesKey: {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						includedNamespacesKey: {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						labelSelectorKey: labelSelectorSchema,
						postHookKey:      pHookSchema,
						preHookKey:       pHookSchema,
					},
				},
			},
		},
	},
}

var labelSelectorSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: labelSelectorResource,
	},
}

var labelSelectorResource = map[string]*schema.Schema{
	matchExrpessionKey: {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				meKey: {
					Type:     schema.TypeString,
					Required: true,
				},
				meOperatorKey: {
					Type:     schema.TypeString,
					Required: true,
				},
				meValuesKey: {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	},
	matchLabelsKey: {
		Type:     schema.TypeMap,
		Optional: true,
	},
}

var pHookSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			execKey: {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						commandKey: {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						containerKey: {
							Type:     schema.TypeString,
							Required: true,
						},
						onErrorKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
						timeoutKey: {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	},
}
