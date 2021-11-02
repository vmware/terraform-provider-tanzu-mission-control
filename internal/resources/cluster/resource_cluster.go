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

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/cluster/manifest"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/cluster/tkgservicevsphere"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/cluster/tkgvsphere"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/common"
)

func ResourceTMCCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceTMCClusterRead,
		CreateContext: resourceClusterCreate,
		UpdateContext: resourceClusterInPlaceUpdate,
		DeleteContext: resourceClusterDelete,
		Schema:        clusterSchema,
	}
}

var clusterSchema = map[string]*schema.Schema{
	ManagementClusterNameKey: {
		Type:     schema.TypeString,
		Default:  "attached",
		Optional: true,
		ForceNew: true,
	},
	ProvisionerNameKey: {
		Type:     schema.TypeString,
		Default:  "attached",
		Optional: true,
		ForceNew: true,
	},
	clusterNameKey: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	attachClusterKey: attachCluster,
	common.MetaKey:   common.Meta,
	SpecKey:          clusterSpec,
	StatusKey: {
		Type:     schema.TypeMap,
		Computed: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
	waitKey: {
		Type:     schema.TypeBool,
		Default:  false,
		Optional: true,
	},
}

func constructFullname(d *schema.ResourceData) (fullname *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) {
	fullname = &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{}

	if value, ok := d.GetOk(ManagementClusterNameKey); ok {
		fullname.ManagementClusterName = value.(string)
	}

	fullname.ManagementClusterName, _ = d.Get(ManagementClusterNameKey).(string)
	fullname.ProvisionerName, _ = d.Get(ProvisionerNameKey).(string)
	fullname.Name, _ = d.Get(clusterNameKey).(string)

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
				ForceNew: true,
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
			tkgServiceVsphereKey: tkgservicevsphere.TkgServiceVsphere,
			tkgVsphereClusterKey: tkgvsphere.TkgVsphereClusterSpec,
		},
	},
}

func constructSpec(d *schema.ResourceData) (spec *clustermodel.VmwareTanzuManageV1alpha1ClusterSpec) {
	spec = &clustermodel.VmwareTanzuManageV1alpha1ClusterSpec{
		ClusterGroupName: clusterGroupDefaultValue,
	}

	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	if v, ok := specData[clusterGroupKey]; ok {
		spec.ClusterGroupName = v.(string)
	}

	if v, ok := specData[tkgServiceVsphereKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.TkgServiceVsphere = tkgservicevsphere.ConstructTKGSSpec(v1)
		}
	}

	if v, ok := specData[tkgVsphereClusterKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.TkgVsphere = tkgvsphere.ConstructTKGVsphereClusterSpec(v1)
		}
	}

	return spec
}

func flattenSpec(spec *clustermodel.VmwareTanzuManageV1alpha1ClusterSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	flattenSpecData[clusterGroupKey] = spec.ClusterGroupName

	if spec.TkgServiceVsphere != nil {
		flattenSpecData[tkgServiceVsphereKey] = tkgservicevsphere.FlattenTKGSSpec(spec.TkgServiceVsphere)
	}

	if spec.TkgVsphere != nil {
		flattenSpecData[tkgVsphereClusterKey] = tkgvsphere.FlattenTKGVsphereClusterSpec(spec.TkgVsphere)
	}

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
		kubeconfigfile, _ = d.Get(helper.GetFirstElementOf(attachClusterKey, attachClusterKubeConfigKey)).(string)

		if kubeconfigfile != "" {
			k8sclient, err = getK8sClient(kubeconfigfile)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	clusterReq := &clustermodel.VmwareTanzuManageV1alpha1ClusterRequest{
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
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete tanzu TMC cluster entry, name : %s", d.Get(clusterNameKey)))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	getClusterResourceRetryableFn := func() (retry bool, err error) {
		_, err = config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceGet(constructFullname(d))
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
		diag.FromErr(errors.Wrapf(err, "verify %s cluster resource clean up", d.Get(clusterNameKey)))
	}

	return diags
}

func resourceClusterInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	updateRequired := false

	switch {
	case common.HasMetaChanged(d):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, clusterGroupKey)):
		updateRequired = true
	}

	if !updateRequired {
		return diags
	}

	getResp, err := config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceGet(constructFullname(d))
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get tanzu TMC cluster entry, name : %s", d.Get(clusterNameKey)))
	}

	if common.HasMetaChanged(d) {
		meta := common.ConstructMeta(d)

		if value, ok := getResp.Cluster.Meta.Labels[common.CreatorLabelKey]; ok {
			meta.Labels[common.CreatorLabelKey] = value
		}

		getResp.Cluster.Meta.Labels = meta.Labels
		getResp.Cluster.Meta.Description = meta.Description
	}

	incomingCGName := d.Get(helper.GetFirstElementOf(SpecKey, clusterGroupKey))

	if incomingCGName.(string) != "" {
		getResp.Cluster.Spec.ClusterGroupName = incomingCGName.(string)
	}

	_, err = config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceUpdate(
		&clustermodel.VmwareTanzuManageV1alpha1ClusterRequest{
			Cluster: getResp.Cluster,
		},
	)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to update tanzu TMC cluster entry, name : %s", d.Get(clusterNameKey)))
	}

	return dataSourceTMCClusterRead(ctx, d, m)
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
