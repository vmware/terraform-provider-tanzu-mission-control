/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetescluster

import (
	"encoding/json"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const (
	ResourceName = "tanzu-mission-control_tanzu_kubernetes_cluster"

	// Common Keys.
	NameKey        = "name"
	SpecKey        = "spec"
	DescriptionKey = "description"
	VersionKey     = "version"
	ReplicasKey    = "replicas"
	OSImageKey     = "os_image"
	OSArchKey      = "arch"

	// Root Keys.
	ManagementClusterNameKey = "management_cluster_name"
	ProvisionerNameKey       = "provisioner_name"

	// Spec Directive Keys.
	ClusterGroupNameKey = "cluster_group_name"
	TMCManagedKey       = "tmc_managed"
	ProxyNameKey        = "proxy_name"
	ImageRegistryKey    = "image_registry"
	TopologyKey         = "topology"

	// Topology Directive Keys.
	ClusterClassKey     = "cluster_class"
	ControlPlaneKey     = "control_plane"
	NodePoolKey         = "nodepool"
	ClusterVariablesKey = "cluster_variables"
	NetworkKey          = "network"
	CoreAddonKey        = "core_addon"

	// Node Pool Directive Keys.
	WorkerClassKey   = "worker_class"
	FailureDomainKey = "failure_domain"
	OverridesKey     = "overrides"

	// Network Directive Keys.
	PodCIDRBlocksKey     = "pod_cidr_blocks"
	ServiceCIDRBlocksKey = "service_cidr_blocks"
	ServiceDomainKey     = "service_domain"

	// Core Addon Directive Keys.
	TypeKey     = "type"
	ProviderKey = "provider"
)

var tanzuKubernetesClusterSchema = map[string]*schema.Schema{
	NameKey:                  clusterNameSchema,
	ManagementClusterNameKey: managementClusterNameSchema,
	ProvisionerNameKey:       provisionerNameSchema,
	SpecKey:                  specSchema,
	common.MetaKey:           common.Meta,
}

var clusterNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Cluster name",
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

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec of tanzu kubernetes cluster (Unified TKG)",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ClusterGroupNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the cluster group to which this cluster belongs.",
				Default:     "default",
				Optional:    true,
			},
			TMCManagedKey: {
				Type:        schema.TypeBool,
				Description: "TMC-managed flag indicates if the cluster is managed by tmc.\n(Default: False)",
				Computed:    true,
			},
			ProxyNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the proxy configuration to use.",
				Optional:    true,
			},
			ImageRegistryKey: {
				Type:        schema.TypeString,
				Description: "Name of the image registry configuration to use.",
				Optional:    true,
			},
			TopologyKey: TopologySchema,
		},
	},
}

var TopologySchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The cluster topology.",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			VersionKey: {
				Type:        schema.TypeString,
				Description: "Kubernetes version of the cluster.",
				Required:    true,
			},
			ClusterClassKey: {
				Type:        schema.TypeString,
				Description: "The name of the cluster class for the cluster.",
				Optional:    true,
				Computed:    true,
			},
			ClusterVariablesKey: {
				Type:                  schema.TypeString,
				Description:           "Variables configuration for the cluster.",
				Required:              true,
				ValidateDiagFunc:      validateJSONString,
				DiffSuppressOnRefresh: true,
				DiffSuppressFunc:      checkVariablesValues,
			},
			CoreAddonKey: {
				Type:        schema.TypeList,
				Description: "(Repeatable Block) The core addons.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						TypeKey: {
							Type:        schema.TypeString,
							Description: "Type of core add on",
							Required:    true,
						},
						ProviderKey: {
							Type:        schema.TypeString,
							Description: "Provider of core add on",
							Required:    true,
						},
					},
				},
			},
			NetworkKey:      NetworkSchema,
			ControlPlaneKey: ControlPlaneSchema,
			NodePoolKey:     NodePoolSchema,
		},
	},
}

var NetworkSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Network specific configuration.",
	MaxItems:    1,
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			PodCIDRBlocksKey: {
				Type:        schema.TypeList,
				Description: "Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			ServiceCIDRBlocksKey: {
				Type:        schema.TypeList,
				Description: "Service CIDR for kubernetes services defaults to 10.96.0.0/12.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			ServiceDomainKey: {
				Type:        schema.TypeString,
				Description: "Domain name for services.",
				Optional:    true,
			},
		},
	},
}

var ControlPlaneSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Control plane specific configuration.",
	MaxItems:    1,
	Required:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ReplicasKey:    ReplicasSchema,
			OSImageKey:     OSImageSchema,
			common.MetaKey: common.Meta,
		},
	},
}

var NodePoolSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "(Repeatable Block) Node pool definition for the cluster.",
	Required:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			NameKey: {
				Type:        schema.TypeString,
				Description: "Name of the node pool.",
				Required:    true,
			},
			DescriptionKey: {
				Type:        schema.TypeString,
				Description: "Description for the node pool.",
				Optional:    true,
			},
			SpecKey: NodePoolSpecSchema,
		},
	},
}

var NodePoolSpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the node pool.",
	MaxItems:    1,
	Required:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			WorkerClassKey: {
				Type:        schema.TypeString,
				Description: "The name of the machine deployment class used to create the node pool.",
				Required:    true,
			},
			FailureDomainKey: {
				Type:        schema.TypeString,
				Description: "The failure domain the machines will be created in.",
				Optional:    true,
			},
			OverridesKey: {
				Type:                  schema.TypeString,
				Description:           "Overrides can be used to override cluster level variables.",
				Optional:              true,
				ValidateDiagFunc:      validateJSONString,
				DiffSuppressOnRefresh: true,
				DiffSuppressFunc:      checkVariablesValues,
			},
			ReplicasKey:    ReplicasSchema,
			OSImageKey:     OSImageSchema,
			common.MetaKey: common.Meta,
		},
	},
}

var ReplicasSchema = &schema.Schema{
	Type:        schema.TypeInt,
	Description: "Number of replicas",
	Required:    true,
}

var OSImageSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "OS image block",
	MaxItems:    1,
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			NameKey: {
				Type:        schema.TypeString,
				Description: "The name of the OS image.",
				Required:    true,
			},
			VersionKey: {
				Type:        schema.TypeString,
				Description: "The version of the OS image.",
				Required:    true,
			},
			OSArchKey: {
				Type:        schema.TypeString,
				Description: "The architecture of the OS image.",
				Required:    true,
			},
		},
	},
}

func validateJSONString(value interface{}, path cty.Path) diag.Diagnostics {
	valueJSON := make(map[string]interface{})
	err := json.Unmarshal([]byte(value.(string)), &valueJSON)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Value is not a valid JSON string."))
	}

	return nil
}

func checkVariablesValues(k, oldValue, newValue string, d *schema.ResourceData) bool {
	if oldValue == "" && newValue != "" {
		return false
	}

	oldValueJSON := make(map[string]interface{})
	newValuesJSON := make(map[string]interface{})
	_ = json.Unmarshal([]byte(oldValue), &oldValueJSON)
	_ = json.Unmarshal([]byte(newValue), &newValuesJSON)
	allMapKeys := getAllKeys(oldValueJSON, newValuesJSON)

	for k := range allMapKeys {
		isVariableEqual := isVariableEqual(oldValueJSON[k], newValuesJSON[k])

		if !isVariableEqual {
			return isVariableEqual
		}
	}

	return true
}

func isVariableEqual(oldVar interface{}, newVar interface{}) bool {
	if (oldVar == nil && newVar != nil) || (oldVar != nil && newVar == nil) {
		return false
	} else if oldVar != nil && newVar != nil {
		switch oldVar := oldVar.(type) {
		case []interface{}:
			oldVarLen := len(oldVar)

			if oldVarLen != len(newVar.([]interface{})) {
				return false
			}

			// List order is a mandatory requirement for deciding list equality
			for i := 0; i < oldVarLen; i++ {
				if !isVariableEqual(oldVar[i], newVar.([]interface{})[i]) {
					return false
				}
			}
		case map[string]interface{}:
			if len(oldVar) != len(newVar.(map[string]interface{})) {
				return false
			}

			allMapKeys := getAllKeys(oldVar, newVar.(map[string]interface{}))

			for k := range allMapKeys {
				if !isVariableEqual(oldVar[k], newVar.(map[string]interface{})[k]) {
					return false
				}
			}
		default:
			return oldVar == newVar
		}
	}

	return true
}

func getAllKeys(maps ...map[string]interface{}) map[string]bool {
	keys := make(map[string]bool)

	for _, m := range maps {
		for key := range m {
			keys[key] = true
		}
	}

	return keys
}
