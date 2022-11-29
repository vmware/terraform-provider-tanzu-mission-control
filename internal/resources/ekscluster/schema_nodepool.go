/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

// nodepoolDefinitionSchema defines the info and nodepool spec for EKS clusters.
//
// Note: ForceNew is not used in any of the elements because this is a part of
// EKS cluster and we don't want to replace full clusters because of Nodepool
// change.
var nodepoolDefinitionSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		infoKey: {
			Type:        schema.TypeList,
			Description: "Info for the nodepool",
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					NameKey: {
						Type:        schema.TypeString,
						Description: "Name of the nodepool",
						Required:    true,
					},
					common.DescriptionKey: {
						Type:        schema.TypeString,
						Description: "Description for the nodepool",
						Optional:    true,
					},
				},
			},
		},
		specKey: nodepoolSpecSchema,
	},
}

var nodepoolSpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			roleArnKey: {
				Type:        schema.TypeString,
				Description: "ARN of the IAM role that provides permissions for the Kubernetes nodepool to make calls to AWS API operations",
				Required:    true,
			},
			amiTypeKey: {
				Type:        schema.TypeString,
				Description: "AMI Type",
				Optional:    true,
			},
			capacityTypeKey: {
				Type:        schema.TypeString,
				Description: "Capacity Type",
				Optional:    true,
			},
			rootDiskSizeKey: {
				Type:        schema.TypeInt,
				Description: "Root disk size in GiB",
				Optional:    true,
			},
			tagsKey: {
				Type:        schema.TypeMap,
				Description: "EKS specific tags",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			nodeLabelsKey: {
				Type:        schema.TypeMap,
				Description: "Kubernetes node labels",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			subnetIdsKey: {
				Type:        schema.TypeList,
				Description: "Subnets required for the nodepool",
				Required:    true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					MinItems: 2,
				},
			},
			remoteAccessKey: {
				Type:        schema.TypeList,
				Description: "Remote access to worker nodes",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						sshKeyKey: {
							Type:        schema.TypeString,
							Description: "SSH key for the nodepool VMs",
							Optional:    true,
						},
						securityGroupsKey: {
							Type:        schema.TypeList,
							Description: "Security groups for the VMs",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    5, // TODO: check about this
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			launchTemplateKey: {
				Type:        schema.TypeList,
				Description: "Launch template for the nodepool",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						idKey: {
							Type:        schema.TypeString,
							Description: "The ID of the launch template",
							Optional:    true,
						},
						nameKey: {
							Type:        schema.TypeString,
							Description: "The name of the launch template",
							Optional:    true,
						},
						versionKey: {
							Type:        schema.TypeString,
							Description: "The version of the launch template to use",
							Optional:    true,
						},
					},
				},
			},
			scalingConfigKey: {
				Type:        schema.TypeList,
				Description: "Nodepool scaling config",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						desiredSizeKey: {
							Type:        schema.TypeInt,
							Description: "Desired size of nodepool",
							Optional:    true,
						},
						maxSizeKey: {
							Type:        schema.TypeInt,
							Description: "Maximum size of nodepool",
							Optional:    true,
						},
						minSizeKey: {
							Type:        schema.TypeInt,
							Description: "Minimum size of nodepool",
							Optional:    true,
						},
					},
				},
			},
			updateConfigKey: {
				Type:        schema.TypeList,
				Description: "Update config for the nodepool",
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						maxUnavailableNodesKey: {
							Type:        schema.TypeString,
							Description: "Maximum number of nodes unavailable at once during a version update",
							Optional:    true,
						},
						maxUnavailablePercentageKey: {
							Type:        schema.TypeString,
							Description: "Maximum percentage of nodes unavailable during a version update",
							Optional:    true,
						},
					},
				},
			},
			taintsKey: {
				Type:        schema.TypeList,
				Description: "If specified, the node's taints",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						effectKey: {
							Type:        schema.TypeString,
							Description: "Current effect state of the node pool",
							Optional:    true,
						},
						keyKey: {
							Type:        schema.TypeString,
							Description: "The taint key to be applied to a node",
							Optional:    true,
						},
						valueKey: {
							Type:        schema.TypeString,
							Description: "The taint value corresponding to the taint key",
							Optional:    true,
						},
					},
				},
			},
			instanceTypesKey: {
				Type:        schema.TypeList,
				Description: "Nodepool instance types",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	},
}
