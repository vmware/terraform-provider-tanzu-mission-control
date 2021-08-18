/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/clientcmd"
	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/authctx"
	clustermodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/cluster"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/cluster/manifest"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/common"
)

func ResourceTMCCluster() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceClusterCreate,
		ReadContext:   dataSourceTMCClusterRead,
		UpdateContext: schema.NoopContext,
		DeleteContext: resourceClusterDelete,
		Schema:        clusterSchema,
	}
}

var clusterSchema = map[string]*schema.Schema{
	managementClusterNameKey: {
		Type:     schema.TypeString,
		Default:  "attached",
		Optional: true,
	},
	provisionerNameKey: {
		Type:     schema.TypeString,
		Default:  "attached",
		Optional: true,
	},
	clusterGroupKey: {
		Type:     schema.TypeString,
		Default:  "default",
		Optional: true,
	},
	clusterNameKey: {
		Type:     schema.TypeString,
		Required: true,
	},
	common.MetaKey:   common.Meta,
	attachClusterKey: attachCluster,
	statusKey: {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
}

var attachCluster = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kube_config_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"execution_cmd": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	},
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	var (
		kubeconfigfile string
		k8sclient      k8sClient.Client
		err            error
	)

	var spec = &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
		ClusterGroupName: d.Get(clusterGroupKey).(string),
	}

	if _, ok := d.GetOk(attachClusterKey); ok {
		kubeconfigfile = d.Get("attach.0.kube_config_path").(string)

		if kubeconfigfile != "" {
			k8sclient, err = getK8sClient(kubeconfigfile)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	managementClusterName := d.Get(managementClusterNameKey).(string)
	provisionerName := d.Get(provisionerNameKey).(string)
	clusterName := d.Get(clusterNameKey).(string)

	clusterReq := &clustermodel.VmwareTanzuManageV1alpha1ClusterCreateClusterRequest{
		Cluster: &clustermodel.VmwareTanzuManageV1alpha1ClusterCluster{
			FullName: &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{
				ManagementClusterName: managementClusterName,
				ProvisionerName:       provisionerName,
				Name:                  clusterName,
			},
			Meta: common.ConstructMeta(d),
			Spec: spec,
		},
	}

	clusterResponse, err := config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceCreate(clusterReq)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create tanzu TMC cluster entry, name : %s", clusterName))
	}

	// always run
	d.SetId(clusterResponse.Cluster.Meta.UID)

	if kubeconfigfile != "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Kubernetes cluster's kubeconfig file provided. Proceeding to attach the cluster TMC",
		})

		deploymentManifests, err := manifest.GetK8sManifest(clusterResponse.Cluster.Status.InstallerLink)
		if err != nil {
			return append(diags, diag.FromErr(err)...)
		}

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Creating deployment link manifests objects on to kubernetes cluster",
		})

		err = manifest.Create(k8sclient, deploymentManifests, true)
		if err != nil {
			return append(diags, diag.FromErr(err)...)
		}

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "TMC resources applied to the cluster successfully",
		})
	}

	return append(diags, dataSourceTMCClusterRead(ctx, d, m)...)
}

func resourceClusterDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	managementClusterName := d.Get(managementClusterNameKey).(string)
	provisionerName := d.Get(provisionerNameKey).(string)
	clusterName := d.Get(clusterNameKey).(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	fn := &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{
		ManagementClusterName: managementClusterName,
		ProvisionerName:       provisionerName,
		Name:                  clusterName,
	}

	err := config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceDelete(fn, "false")
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete tanzu TMC cluster entry, name : %s", clusterName))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func getK8sClient(kubeconfigfile string) (k8sClient.Client, error) {
	restconfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigfile)
	if err != nil {
		return nil, errors.WithMessagef(err, "Invalid kubeconfig file path provided, filepath : %s", kubeconfigfile)
	}

	restconfig.Timeout = 10 * time.Second

	k8sclient, err := k8sClient.New(restconfig, k8sClient.Options{})
	if err != nil {
		return nil, errors.WithMessagef(err, "Error in creating kubernetes client from kubeconfig file provided, filepath : %s", kubeconfigfile)
	}

	return k8sclient, nil
}
