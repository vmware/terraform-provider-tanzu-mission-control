/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/clientcmd"
	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/manifest"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/tkgaws"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/tkgservicevsphere"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/tkgvsphere"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type (
	contextMethodKey struct{}
	updateCheck      []func(*schema.ResourceData, *clustermodel.VmwareTanzuManageV1alpha1ClusterCluster) bool
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
		Type:        schema.TypeString,
		Description: "Name of the management cluster",
		Default:     attachedValue,
		Optional:    true,
		ForceNew:    true,
	},
	ProvisionerNameKey: {
		Type:        schema.TypeString,
		Description: "Provisioner of the cluster",
		Default:     attachedValue,
		Optional:    true,
		ForceNew:    true,
	},
	NameKey: {
		Type:        schema.TypeString,
		Description: "Name of this cluster",
		Required:    true,
		ForceNew:    true,
	},
	attachClusterKey: attachCluster,
	common.MetaKey:   common.Meta,
	SpecKey:          clusterSpec,
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

func constructFullname(d *schema.ResourceData) (fullname *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) {
	fullname = &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{}

	if value, ok := d.GetOk(ManagementClusterNameKey); ok {
		fullname.ManagementClusterName = value.(string)
	}

	fullname.ManagementClusterName, _ = d.Get(ManagementClusterNameKey).(string)
	fullname.ProvisionerName, _ = d.Get(ProvisionerNameKey).(string)
	fullname.Name, _ = d.Get(NameKey).(string)

	return fullname
}

var attachCluster = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			attachClusterKubeConfigKey: {
				Type:        schema.TypeString,
				Description: "Attach cluster KUBECONFIG path",
				ForceNew:    true,
				Optional:    true,
			},
			attachClusterDescriptionKey: {
				Type:        schema.TypeString,
				Description: "Attach cluster description",
				Optional:    true,
			},
		},
	},
}

var clusterSpec = &schema.Schema{
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
			tkgAWSClusterKey:     tkgaws.TkgAWSClusterSpec,
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

	if v, ok := specData[proxyNameKey]; ok {
		spec.ProxyName = v.(string)
	}

	if v, ok := specData[tkgAWSClusterKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.TkgAws = tkgaws.ConstructTKGAWSClusterSpec(v1)
		}
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

	flattenSpecData[proxyNameKey] = spec.ProxyName

	if spec.TkgAws != nil {
		flattenSpecData[tkgAWSClusterKey] = tkgaws.FlattenTKGAWSClusterSpec(spec.TkgAws)
	}

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
				log.Println("[ERROR] error while creating kubernetes client: ", err.Error())
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
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control cluster entry, name : %s", d.Get(NameKey)))
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

		log.Printf("[INFO] Applying %s cluster's deployment link manifest objects on to kubernetes cluster", constructFullname(d).ToString())

		err = manifest.Create(k8sclient, deploymentManifests, true)
		if err != nil {
			return append(diags, diag.FromErr(err)...)
		}

		log.Printf("[INFO] Cluster attach successful. Tanzu Mission Control resources applied to the cluster(%s) successfully", constructFullname(d).ToString())
	}

	return append(diags, dataSourceTMCClusterRead(context.WithValue(ctx, contextMethodKey{}, "create"), d, m)...)
}

func resourceClusterDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceDelete(constructFullname(d), "false")
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control cluster entry, name : %s", d.Get(NameKey)))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	_ = schema.RemoveFromState(d, m)

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
		diag.FromErr(errors.Wrapf(err, "verify %s cluster resource clean up", d.Get(NameKey)))
	}

	return diags
}

func withTKGmVsphereVersionUpdate(d *schema.ResourceData, cluster *clustermodel.VmwareTanzuManageV1alpha1ClusterCluster) bool {
	if d.HasChange(helper.GetFirstElementOf(SpecKey, tkgVsphereClusterKey, distributionKey, versionKey)) {
		newVersion := d.Get(helper.GetFirstElementOf(SpecKey, tkgVsphereClusterKey, distributionKey, versionKey))
		if newVersion.(string) != "" {
			cluster.Spec.TkgVsphere.Distribution.Version = newVersion.(string)

			log.Printf("[INFO] updating TKGm vSphere workload cluster version to %s", newVersion.(string))

			return true
		}
	}

	return false
}

func withTKGsVsphereVersionUpdate(d *schema.ResourceData, cluster *clustermodel.VmwareTanzuManageV1alpha1ClusterCluster) bool {
	if d.HasChange(helper.GetFirstElementOf(SpecKey, tkgServiceVsphereKey, distributionKey, versionKey)) {
		newVersion := d.Get(helper.GetFirstElementOf(SpecKey, tkgServiceVsphereKey, distributionKey, versionKey))
		if newVersion.(string) != "" {
			cluster.Spec.TkgServiceVsphere.Distribution.Version = newVersion.(string)

			log.Printf("[INFO] updating TKGs workload cluster version to %s", newVersion.(string))

			return true
		}
	}

	return false
}

func withClusterGroupUpdate(d *schema.ResourceData, cluster *clustermodel.VmwareTanzuManageV1alpha1ClusterCluster) bool {
	if d.HasChange(helper.GetFirstElementOf(SpecKey, clusterGroupKey)) {
		newClusterGroupName := d.Get(helper.GetFirstElementOf(SpecKey, clusterGroupKey))
		if newClusterGroupName.(string) != "" {
			cluster.Spec.ClusterGroupName = newClusterGroupName.(string)

			log.Printf("[INFO] updating cluster group to %s", newClusterGroupName.(string))

			return true
		}
	}

	return false
}

func withMetaUpdate(d *schema.ResourceData, cluster *clustermodel.VmwareTanzuManageV1alpha1ClusterCluster) bool {
	if !common.HasMetaChanged(d) {
		return false
	}

	objectMeta := common.ConstructMeta(d)

	if value, ok := cluster.Meta.Labels[common.CreatorLabelKey]; ok {
		objectMeta.Labels[common.CreatorLabelKey] = value
	}

	cluster.Meta.Labels = objectMeta.Labels
	cluster.Meta.Description = objectMeta.Description

	log.Printf("[INFO] updating cluster meta data")

	return true
}

func resourceClusterInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	var updateAvailable bool

	// Get call to initialise the cluster struct
	getResp, err := config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceGet(constructFullname(d))
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster entry, name : %s", d.Get(NameKey)))
	}

	updates := updateCheck{withMetaUpdate, withClusterGroupUpdate, withTKGsVsphereVersionUpdate, withTKGmVsphereVersionUpdate}

	for _, update := range updates {
		if update(d, getResp.Cluster) {
			updateAvailable = true
		}
	}

	if updateAvailable {
		_, err = config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceUpdate(
			&clustermodel.VmwareTanzuManageV1alpha1ClusterRequest{
				Cluster: getResp.Cluster,
			},
		)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control cluster entry, name : %s", d.Get(NameKey)))
		}

		log.Printf("[INFO] cluster update successful")
	}

	return dataSourceTMCClusterRead(ctx, d, m)
}

func getK8sClient(kubeConfigFile string) (k8sClient.Client, error) {
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigFile)
	if err != nil {
		return nil, errors.WithMessagef(err, "Invalid kubeconfig file path provided, filepath : %s", kubeConfigFile)
	}

	restConfig.Timeout = 10 * time.Second

	k8sClient, err := k8sClient.New(restConfig, k8sClient.Options{})
	if err != nil {
		return nil, errors.WithMessagef(err, "Error in creating kubernetes client from kubeconfig file provided, filepath : %s", kubeConfigFile)
	}

	return k8sClient, nil
}
