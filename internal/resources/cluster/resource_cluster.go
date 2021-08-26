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
	clusterNameKey: {
		Type:     schema.TypeString,
		Required: true,
	},
	attachClusterKey: attachCluster,
	common.MetaKey:   common.Meta,
	specKey:          clusterSpec,
	statusKey: {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
}

func constructFullname(d *schema.ResourceData) (fullname *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) {
	fullname = &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{}

	if value, ok := d.GetOk(managementClusterNameKey); ok {
		fullname.ManagementClusterName = value.(string)
	}

	fullname.ManagementClusterName = d.Get(managementClusterNameKey).(string)
	fullname.ProvisionerName = d.Get(provisionerNameKey).(string)
	fullname.Name = d.Get(clusterNameKey).(string)

	return fullname
}

var attachCluster = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			attachClusterKubeConfigKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			attachClusterDescriptionKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	},
}

var clusterSpec = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			clusterGroupKey: {
				Type:     schema.TypeString,
				Default:  clusterGroupDefaultValue,
				Optional: true,
			},
		},
	},
}

func constructSpec(d *schema.ResourceData) (spec *clustermodel.VmwareTanzuManageV1alpha1ClusterSpec) {
	spec = &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
		ClusterGroupName: clusterGroupDefaultValue,
	}

	value, ok := d.GetOk(specKey)
	if !ok {
		return spec
	}

	data := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	if v, ok := specData[clusterGroupKey]; ok {
		spec.ClusterGroupName = v.(string)
	}

	return spec
}

func flattenSpec(spec *clustermodel.VmwareTanzuManageV1alpha1ClusterSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	flattenSpecData[clusterGroupKey] = spec.ClusterGroupName

	return []interface{}{flattenSpecData}
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	var (
		kubeconfigfile string
		k8sclient      k8sClient.Client
		err            error
	)

	if _, ok := d.GetOk(attachClusterKey); ok {
		kubeconfigfile = d.Get(getKubeConfigFileKeyFromRoot).(string)

		if kubeconfigfile != "" {
			k8sclient, err = getK8sClient(kubeconfigfile)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	clusterReq := &clustermodel.VmwareTanzuManageV1alpha1ClusterCreateClusterRequest{
		Cluster: &clustermodel.VmwareTanzuManageV1alpha1ClusterCluster{
			FullName: constructFullname(d),
			Meta:     common.ConstructMeta(d),
			Spec:     constructSpec(d),
		},
	}

	clusterResponse, err := config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceCreate(clusterReq)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create tanzu TMC cluster entry, name : %s", d.Get(clusterNameKey)))
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

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceDelete(constructFullname(d), "false")
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete tanzu TMC cluster entry, name : %s", d.Get(clusterNameKey)))
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
