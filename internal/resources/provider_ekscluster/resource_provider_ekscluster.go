/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package providerekscluster

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/provider_ekscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type (
	contextMethodKey struct{}
)

const defaultTimeout = 3 * time.Minute

func ResourceTMCProviderEKSCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceProviderClusterRead(helper.GetContextWithCaller(ctx, helper.RefreshState), d, m)
		},
		CreateContext: resourceProviderClusterUpdate,
		UpdateContext: resourceProviderClusterUpdate,
		DeleteContext: resourceProviderClusterDelete,
		Schema:        providerClusterSchema,
	}
}

var providerClusterSchema = map[string]*schema.Schema{
	credentialNameKey: {
		Type:        schema.TypeString,
		Description: "Name of the AWS Crendential in Tanzu Mission Control",
		Required:    true,
		ForceNew:    true,
	},
	regionKey: {
		Type:        schema.TypeString,
		Description: "AWS Region of the this cluster",
		Required:    true,
		ForceNew:    true,
	},
	nameKey: {
		Type:        schema.TypeString,
		Description: "Name of this cluster in EKS",
		Required:    true,
		ForceNew:    true,
	},
	common.MetaKey: common.Meta,
	specKey:        providerClusterSpecSchema,
	statusKey: {
		Type:        schema.TypeMap,
		Description: "Status of the cluster",
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
	waitKey: {
		Type:        schema.TypeString,
		Description: "Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero. Should be set to 0 in case of simple attach cluster where kubeconfig input is not provided.",
		Default:     "default",
		Optional:    true,
	},
}

var providerClusterSpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the cluster",
	Optional:    true,
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
			agentNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the cluster in TMC",
				Required:    true,
				ForceNew:    true,
			},
			eksARNKey: {
				Type:        schema.TypeString,
				Description: "ARN of the EKS cluster",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

func resourceProviderClusterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cluster, err := providerClusterUpdate(ctx, d, m, true)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to manage AWS EKS cluster with Tanzu Mission Control"))
	}

	d.SetId(cluster.Meta.UID)

	return dataSourceProviderClusterRead(context.WithValue(ctx, contextMethodKey{}, "create"), d, m)
}

func resourceProviderClusterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := providerClusterUpdate(ctx, d, m, false)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to manage AWS EKS cluster with Tanzu Mission Control"))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	clusterFn := constructFullname(d)
	getClusterResourceRetryableFn := func() (retry bool, err error) {
		resp, err := config.TMCConnection.ProviderEKSClusterResourceService.ProviderEksClusterResourceServiceGet(clusterFn)
		if err != nil {
			if clienterrors.IsNotFoundError(err) {
				return false, nil
			}

			return true, err
		}

		if err == nil &&
			resp.ProviderEksCluster.Status.Phase != nil &&
			*resp.ProviderEksCluster.Status.Phase == models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhasePENDINGUNMANAGE {
			log.Printf("[DEBUG] cluster(%s) deletion in progress", clusterFn.ToString())
			return true, errors.New("cluster deletion in progress")
		}

		return false, errors.New("cluster is in unexpected phase")
	}

	timeoutDuration := getRetryTimeout(d)

	_, err = helper.RetryUntilTimeout(getClusterResourceRetryableFn, 10*time.Second, timeoutDuration)
	if err != nil {
		diag.FromErr(errors.Wrapf(err, "verify %s EKS cluster resource clean up", d.Get(nameKey)))
	}

	return diags
}

func providerClusterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}, tmcManaged bool) (*models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster, error) {
	config, ok := m.(authctx.TanzuContext)
	if !ok {
		log.Println("[ERROR] error while retrieving Tanzu auth config")
		return nil, errors.New("error while retrieving Tanzu auth config")
	}

	clusterFn := constructFullname(d)

	clusterReq := &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterUpdateProviderEksClusterRequest{
		ProviderEksCluster: &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster{
			FullName: clusterFn,
			Meta:     common.ConstructMeta(d),
			Spec:     constructClusterSpec(d),
		},
	}

	clusterReq.ProviderEksCluster.Spec.TmcManaged = tmcManaged

	clusterResponse, err := config.TMCConnection.ProviderEKSClusterResourceService.ProviderEksClusterResourceServiceUpdate(clusterReq)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to update management of AWS EKS Cluster %s with Tanzu Mission Control", d.Get(nameKey))
	}

	return clusterResponse.ProviderEksCluster, nil
}

func constructFullname(d *schema.ResourceData) (fullname *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName) {
	fullname = &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName{}

	fullname.CredentialName, _ = d.Get(credentialNameKey).(string)
	fullname.Region, _ = d.Get(regionKey).(string)
	fullname.Name, _ = d.Get(nameKey).(string)

	return fullname
}

func constructClusterSpec(d *schema.ResourceData) *models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec {
	spec := &models.VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec{
		ClusterGroupName: clusterGroupDefaultValue,
	}

	value, ok := d.GetOk(specKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})
	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData, _ := data[0].(map[string]interface{})

	if v, ok := specData[clusterGroupKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ClusterGroupName, clusterGroupKey)
	}

	if v, ok := specData[proxyNameKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ProxyName, proxyNameKey)
	}

	if v, ok := specData[agentNameKey]; ok {
		helper.SetPrimitiveValue(v, &spec.AgentName, proxyNameKey)
	}

	if v, ok := specData[eksARNKey]; ok {
		helper.SetPrimitiveValue(v, &spec.Arn, proxyNameKey)
	}

	return spec
}

func getRetryTimeout(d *schema.ResourceData) time.Duration {
	timeoutValueData, _ := d.Get(waitKey).(string)
	if timeoutValueData != "default" {
		providedDuration, parseErr := time.ParseDuration(timeoutValueData)
		if parseErr == nil {
			return providedDuration
		}
	}

	return defaultTimeout
}
