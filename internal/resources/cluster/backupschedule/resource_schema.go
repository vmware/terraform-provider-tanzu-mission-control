/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	backupschedulemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/dataprotection/cluster/backupschedule"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type BackupScope string

const (
	FullClusterBackupScope   BackupScope = "FULL_CLUSTER"
	NamespacesBackupScope    BackupScope = "SET_NAMESPACES"
	LabelSelectorBackupScope BackupScope = "LABEL_SELECTOR"
)

const (
	ResourceName = "tanzu-mission-control_backup_schedule"

	// Root Keys.
	NameKey                  = "name"
	ClusterNameKey           = "cluster_name"
	SpecKey                  = "spec"
	ProvisionerNameKey       = "provisioner_name"
	ManagementClusterNameKey = "management_cluster_name"
	ScopeKey                 = "scope"

	// Spec Directive Keys.
	PausedKey   = "paused"
	ScheduleKey = "schedule"
	TemplateKey = "template"

	// Schedule Directive Keys.
	RateKey = "rate"

	// Template Directive Keys.
	BackupTTLKey                = "backup_ttl"
	SystemExcludedNamespacesKey = "sys_excluded_namespaces"
	ExcludedNamespacesKey       = "excluded_namespaces"
	IncludedNamespacesKey       = "included_namespaces"
	ExcludedResourcesKey        = "excluded_resources"
	IncludedResourcesKey        = "included_resources"
	IncludeClusterResourcesKey  = "include_cluster_resources"
	DefaultVolumesToResticKey   = "default_volumes_to_restic"
	SnapshotVolumesKey          = "snapshot_volumes"
	CsiSnapshotTimeoutKey       = "csi_snapshot_timeout"
	DefaultVolumesToFsBackupKey = "default_volumes_to_fs_backup"
	StorageLocationKey          = "storage_location"
	VolumeSnapshotLocationsKey  = "volume_snapshot_locations"
	OrderedResourcesKey         = "ordered_resources"
	HooksKey                    = "hooks"
	LabelSelectorKey            = "label_selector"
	OrLabelSelectorKey          = "or_label_selector"

	// Hooks Directive Keys.
	ResourceKey = "resource"

	// Resource Directive Keys.
	PreHookKey  = "pre_hook"
	PostHookKey = "post_hook"

	// Pre/Post Hook Directive Keys.
	ExecKey = "exec"

	// Exec Directive Keys.
	CommandKey   = "command"
	ContainerKey = "container"
	OnErrorKey   = "on_error"
	TimeoutKey   = "timeout"

	// (Or)Label Selector Directive Keys.
	MatchLabelsKey     = "match_labels"
	MatchExrpessionKey = "match_expression"

	// Match Expressions Directive Keys.
	MeKey         = "key"
	MeOperatorKey = "operator"
	MeValuesKey   = "values"
)

var (
	ScopeValidValues   = []string{string(FullClusterBackupScope), string(NamespacesBackupScope), string(LabelSelectorBackupScope)}
	OnErrorValidValues = []string{
		string(backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorModeFAIL),
		string(backupschedulemodels.VmwareTanzuManageV1alpha1ClusterDataprotectionBackupHookErrorModeCONTINUE),
	}
)

var backupScheduleResourceSchema = map[string]*schema.Schema{
	NameKey:                  nameSchema,
	ManagementClusterNameKey: managementClusterNameSchema,
	ProvisionerNameKey:       provisionerNameSchema,
	ClusterNameKey:           clusterNameSchema,
	ScopeKey:                 backupScopeSchema,
	SpecKey:                  specSchema,
	common.MetaKey:           common.Meta,
}

var nameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "The name of the backup schedule",
	Required:    true,
	ForceNew:    true,
}

var managementClusterNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Management cluster name",
	Required:    true,
	ForceNew:    true,
}

var provisionerNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Cluster provisioner name",
	Required:    true,
	ForceNew:    true,
}

var clusterNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Cluster name",
	Required:    true,
	ForceNew:    true,
}

var backupScopeSchema = &schema.Schema{
	Type:             schema.TypeString,
	Description:      fmt.Sprintf("Scope for backup schedule.\nValid values are (%s)", strings.Join(ScopeValidValues, ", ")),
	Required:         true,
	ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(ScopeValidValues, false)),
}

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Backup schedule spec block",
	Required:    true,
	MaxItems:    1,
	MinItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			PausedKey: {
				Type:        schema.TypeBool,
				Description: "Paused specifies whether the schedule is paused or not. (Default: False)",
				Optional:    true,
				Default:     false,
			},
			ScheduleKey: {
				Type:        schema.TypeList,
				Description: "Schedule block",
				MaxItems:    1,
				MinItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						RateKey: {
							Type:        schema.TypeString,
							Description: "Cron expression of backup schedule rate/interval",
							Required:    true,
						},
					},
				},
			},
			TemplateKey: templateSchema,
		},
	},
}

var templateSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Backup schedule template block, backup definition to be run on the provided schedule",
	MaxItems:    1,
	MinItems:    1,
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			CsiSnapshotTimeoutKey: {
				Description: "Specifies the time used to wait for CSI VolumeSnapshot status turns to ReadyToUse during creation, before returning error as timeout.\nThe default value is 10 minute.\nFormat is the time number and time sign, example: \"50s\" (50 seconds)",
				Type:        schema.TypeString,
				Optional:    true,
			},
			DefaultVolumesToFsBackupKey: {
				Type:        schema.TypeBool,
				Description: "Specifies whether all pod volumes should be backed up via file system backup by default.\n(Default: True)",
				Optional:    true,
				Default:     true,
			},
			DefaultVolumesToResticKey: {
				Type:        schema.TypeBool,
				Description: "Specifies whether restic should be used to take a backup of all pod volumes by default.\n(Default: False)",
				Optional:    true,
				Default:     false,
			},
			ExcludedNamespacesKey: {
				Type:        schema.TypeList,
				Description: "The namespaces to be excluded in the backup.\nCan't be used if scope is SET_NAMESPACES.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			SystemExcludedNamespacesKey: {
				Type:        schema.TypeList,
				Description: "System excluded namespaces for state.",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			ExcludedResourcesKey: {
				Type:        schema.TypeList,
				Description: "The name list for the resources to be excluded in backup.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			IncludedNamespacesKey: {
				Type:        schema.TypeList,
				Description: "The namespace to be included for backup from.\nIf empty, all namespaces are included.\nCan't be used if scope is FULL_CLUSTER.\nRequired if scope is SET_NAMESPACES.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			IncludedResourcesKey: {
				Type:        schema.TypeList,
				Description: "The name list for the resources to be included into backup. If empty, all resources are included.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			IncludeClusterResourcesKey: {
				Type:        schema.TypeBool,
				Description: "A flag which specifies whether cluster-scoped resources should be included for consideration in the backup.\nIf set to true, all cluster-scoped resources will be backed up. If set to false, all cluster-scoped resources will be excluded from the backup.\nIf unset, all cluster-scoped resources are included if and only if all namespaces are included and there are no excluded namespaces.\nOtherwise, only cluster-scoped resources associated with namespace-scoped resources included in the backup spec are backed up.\nFor example, if a PersistentVolumeClaim is included in the backup, its associated PersistentVolume (which is cluster-scoped) would also be backed up.\n(Default: False)",
				Optional:    true,
				Default:     false,
			},
			OrderedResourcesKey: {
				Type:        schema.TypeMap,
				Description: "Specifies the backup order of resources of specific Kind. The map key is the Kind name and value is a list of resource names separated by commas.\nEach resource name has format \"namespace/resourcename\".\nFor cluster resources, simply use \"resourcename\".",
				Optional:    true,
			},
			SnapshotVolumesKey: {
				Type:        schema.TypeBool,
				Description: "A flag which specifies whether to take cloud snapshots of any PV's referenced in the set of objects included in the Backup.\nIf set to true, snapshots will be taken, otherwise, snapshots will be skipped.\nIf left unset, snapshots will be attempted if volume snapshots are configured for the cluster.",
				Optional:    true,
				Default:     false,
			},
			StorageLocationKey: {
				Type:        schema.TypeString,
				Description: "The name of a BackupStorageLocation where the backup should be stored.",
				Optional:    true,
			},
			BackupTTLKey: {
				Type:        schema.TypeString,
				Description: "The backup retention period.",
				Optional:    true,
			},
			VolumeSnapshotLocationsKey: {
				Type:        schema.TypeList,
				Description: "A list containing names of VolumeSnapshotLocations associated with this backup.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			LabelSelectorKey: {
				Type:        schema.TypeList,
				Description: "The label selector to selectively adding individual objects to the backup schedule.\nIf not specified, all objects are included.\nCan't be used if scope is FULL_CLUSTER or SET_NAMESPACES.\nRequired if scope is LABEL_SELECTOR and Or Label Selectors are not defined",
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: labelSelectorResource,
				},
			},
			OrLabelSelectorKey: {
				Type:        schema.TypeList,
				Description: "(Repeatable Block) A list of label selectors to filter with when adding individual objects to the backup.\nIf multiple provided they will be joined by the OR operator.\nLabelSelector as well as OrLabelSelectors cannot co-exist in backup request, only one of them can be used.\nCan't be used if scope is FULL_CLUSTER or SET_NAMESPACES.\nRequired if scope is LABEL_SELECTOR and Label Selector is not defined",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: labelSelectorResource,
				},
			},
			HooksKey: hooksSchema,
		},
	},
}

var hooksSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Hooks block represent custom actions that should be executed at different phases of the backup.",
	MaxItems:    1,
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ResourceKey: {
				Type:        schema.TypeList,
				Description: "(Repeatable Block) Resources are hooks that should be executed when backing up individual instances of a resource.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						NameKey: {
							Type:        schema.TypeString,
							Description: "The name of the hook resource.",
							Required:    true,
						},
						ExcludedNamespacesKey: {
							Type:        schema.TypeList,
							Description: "Specifies the namespaces to which this hook spec does not apply.",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						IncludedNamespacesKey: {
							Type:        schema.TypeList,
							Description: "Specifies the namespaces to which this hook spec applies.\nIf empty, it applies to all namespaces.",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						LabelSelectorKey: {
							Type:        schema.TypeList,
							Description: "The label selector to selectively adding individual objects to the hook resource.\nIf not specified, all objects are included.",
							MaxItems:    1,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: labelSelectorResource,
							},
						},
						PostHookKey: pHookSchema,
						PreHookKey:  pHookSchema,
					},
				},
			},
		},
	},
}

var labelSelectorResource = map[string]*schema.Schema{
	MatchExrpessionKey: {
		Type:        schema.TypeList,
		Description: "(Repeatable Block) A list of label selector requirements. The requirements are ANDed.",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				MeKey: {
					Type:        schema.TypeString,
					Description: "Key is the label key that the selector applies to.",
					Required:    true,
				},
				MeOperatorKey: {
					Type:        schema.TypeString,
					Description: "Operator represents a key's relationship to a set of values.\nValid operators are \"In\", \"NotIn\", \"Exists\" and \"DoesNotExist\".",
					Required:    true,
				},
				MeValuesKey: {
					Type:        schema.TypeList,
					Description: "Values is an array of string values.\nIf the operator is \"In\" or \"NotIn\", the values array must be non-empty.\nIf the operator is \"Exists\" or \"DoesNotExist\", the values array must be empty.\nThis array is replaced during a strategic merge patch.",
					Optional:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	},
	MatchLabelsKey: {
		Type:        schema.TypeMap,
		Description: "A map of {key,value} pairs. A single {key,value} in the map is equivalent to an element of match_expressions, whose key field is \"key\", the operator is \"In\" and the values array contains only \"value\".\nThe requirements are ANDed.",
		Optional:    true,
	},
}

var pHookSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "(Repeatable Block) A list of backup hooks to execute after storing the item in the backup.\nThese are executed after all \"additional items\" from item actions are processed.",
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ExecKey: {
				Type:        schema.TypeList,
				Description: "Exec block defines an exec hook.",
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						CommandKey: {
							Type:        schema.TypeList,
							Description: "The command and arguments to execute.",
							Required:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						ContainerKey: {
							Type:        schema.TypeString,
							Description: "The container in the pod where the command should be executed.\nIf not specified, the pod's first container is used.",
							Required:    true,
						},
						OnErrorKey: {
							Type:             schema.TypeString,
							Description:      fmt.Sprintf("Specifies how Velero should behave if it encounters an error executing this hook.\nValid values are (%s)", strings.Join(OnErrorValidValues, ", ")),
							Optional:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(OnErrorValidValues, false)),
						},
						TimeoutKey: {
							Type:        schema.TypeString,
							Description: "Defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.",
							Optional:    true,
						},
					},
				},
			},
		},
	},
}
