/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type (
	contextMethodKey struct{}
)

var ignoredTagsPrefix = "tmc.cloud.vmware.com/"

func ResourceTMCEKSCluster() *schema.Resource {
	return &schema.Resource{
		Schema:        clusterSchema,
		CreateContext: resourceClusterCreate,
		ReadContext:   dataSourceTMCEKSClusterRead,
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			panic("not implemented")
		},
		DeleteContext: resourceClusterDelete,
		Description:   "Tanzu Mission Control EKS Cluster Resource",
	}
}

var clusterSchema = map[string]*schema.Schema{
	CredentialNameKey: {
		Type:        schema.TypeString,
		Description: "Name of the AWS Crendential in Tanzu Mission Control",
		Required:    true,
		ForceNew:    true,
	},
	RegionKey: {
		Type:        schema.TypeString,
		Description: "AWS Region of the this cluster",
		Required:    true,
		ForceNew:    true,
	},
	NameKey: {
		Type:        schema.TypeString,
		Description: "Name of this cluster",
		Required:    true,
		ForceNew:    true,
	},
	common.MetaKey: common.Meta,
	specKey:        clusterSpecSchema,
	StatusKey: {
		Type:        schema.TypeMap,
		Description: "Status of the cluster",
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
	waitKey: {
		Type:        schema.TypeString,
		Description: "Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero",
		Default:     "default",
		Optional:    true,
	},
}

var clusterSpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the cluster",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			clusterGroupKey: {
				Type:        schema.TypeString,
				Description: "Name of the cluster group to which this cluster belongs",
				Default:     clusterGroupDefaultValue,
				Optional:    true,
			},
			proxyNameKey: {
				Type:        schema.TypeString,
				Description: "Optional proxy name is the name of the Proxy Config to be used for the cluster",
				Optional:    true,
			},
			configKey: configSchema,
			nodepoolKey: {
				Type:        schema.TypeList,
				Description: "Nodepool definitions for the cluster",
				Required:    true,
				MinItems:    1,
				Elem:        nodepoolDefinitionSchema,
			},
		},
	},
}

var configSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "EKS config for the cluster control plane",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			roleArnKey: {
				Type:        schema.TypeString,
				Description: "ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations",
				Required:    true,
				ForceNew:    true,
			},
			kubernetesVersionKey: {
				Type:        schema.TypeString,
				Description: "Kubernetes version of the cluster",
				Required:    true,
			},
			tagsKey: {
				Type:        schema.TypeMap,
				Description: "The metadata to apply to the cluster to assist with categorization and organization",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.Contains(k, ignoredTagsPrefix)
				},
			},
			kubernetesNetworkConfigKey: {
				Type:        schema.TypeList,
				Description: "Kubernetes Network Config",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						serviceCidrKey: {
							Type:        schema.TypeString,
							Description: "Service CIDR for Kubernetes services",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			loggingKey: {
				Type:        schema.TypeList,
				Description: "EKS logging configuration",
				Optional:    true,
				ForceNew:    false,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						apiServerKey: {
							Type:        schema.TypeBool,
							Description: "Enable API server logs",
							Optional:    true,
							ForceNew:    false,
						},
						auditKey: {
							Type:        schema.TypeBool,
							Description: "Enable audit logs",
							Optional:    true,
							ForceNew:    false,
						},
						authenticatorKey: {
							Type:        schema.TypeBool,
							Description: "Enable authenticator logs",
							Optional:    true,
							ForceNew:    false,
						},
						controllerManagerKey: {
							Type:        schema.TypeBool,
							Description: "Enable controller manager logs",
							Optional:    true,
							ForceNew:    false,
						},
						schedulerKey: {
							Type:        schema.TypeBool,
							Description: "Enable scheduler logs",
							Optional:    true,
							ForceNew:    false,
						},
					},
				},
			},
			vpcKey: vpcSchema,
		},
	},
}

var vpcSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "VPC config",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			enablePrivateAccessKey: {
				Type:        schema.TypeBool,
				Description: "Enable private access on the cluster",
				Optional:    true,
				ForceNew:    false,
			},
			enablePublicAccessKey: {
				Type:        schema.TypeBool,
				Description: "Enable public access on the cluster",
				Optional:    true,
				ForceNew:    false,
			},
			publicAccessCidrsKey: {
				Type:        schema.TypeList,
				Description: "Public access cidrs",
				Optional:    true,
				ForceNew:    false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			securityGroupsKey: {
				Type:        schema.TypeList,
				Description: "Security groups for the cluster VMs",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    5,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			subnetIdsKey: {
				Type:        schema.TypeList,
				Description: "Subnet ids used by the cluster",
				Required:    true,
				ForceNew:    true,
				MinItems:    2,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	},
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config, ok := m.(authctx.TanzuContext)
	if !ok {
		log.Println("[ERROR] error while retrieving Tanzu auth config")
		return diag.Errorf("error while retrieving Tanzu auth config")
	}

	clusterReq := &eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest{
		EksCluster: &eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster{
			FullName: constructFullname(d),
			Meta:     common.ConstructMeta(d),
			Spec:     constructSpec(d),
		},
	}

	clusterResponse, err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceCreate(clusterReq)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster entry, name : %s", d.Get(NameKey)))
	}

	d.SetId(clusterResponse.EksCluster.Meta.UID)

	return dataSourceTMCEKSClusterRead(context.WithValue(ctx, contextMethodKey{}, "create"), d, m)
}

func resourceClusterDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceDelete(constructFullname(d), "false")
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	getClusterResourceRetryableFn := func() (retry bool, err error) {
		_, err = config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceGet(constructFullname(d))
		if err == nil {
			return true, errors.New("cluster deletion in progress")
		}

		if !clienterrors.IsNotFoundError(err) {
			return true, err
		}

		return false, nil
	}

	_, err = helper.Retry(getClusterResourceRetryableFn, 10*time.Second, 18)
	if err != nil {
		diag.FromErr(errors.Wrapf(err, "verify %s EKS cluster resource clean up", d.Get(NameKey)))
	}

	return diags
}

func constructFullname(d *schema.ResourceData) (fullname *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName) {
	fullname = &eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName{}

	fullname.CredentialName, _ = d.Get(CredentialNameKey).(string)
	fullname.Region, _ = d.Get(RegionKey).(string)
	fullname.Name, _ = d.Get(NameKey).(string)

	return fullname
}
