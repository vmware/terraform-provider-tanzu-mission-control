/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	nodepoolmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"

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
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceClusterRead(helper.GetContextWithCaller(ctx, helper.RefreshState), d, m)
		},
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
		Description: "Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero. Should be set to 0 in case of simple attach cluster where kubeconfig input is not provided.",
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

func constructNpFullName(d *schema.ResourceData) (fullname *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolFullName) {
	fullname = &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolFullName{}

	if value, ok := d.GetOk(ManagementClusterNameKey); ok {
		fullname.ManagementClusterName = value.(string)
	}

	if value, ok := d.GetOk(ProvisionerNameKey); ok {
		fullname.ProvisionerName = value.(string)
	}

	if value, ok := d.GetOk(NameKey); ok {
		fullname.ClusterName = value.(string)
	}

	if value, ok := d.Get(helper.GetFirstElementOf(SpecKey, tkgServiceVsphereKey, topologyKey, nodePoolKey)).([]interface{})[0].(map[string]interface{}); ok {
		poolInfo := value["info"].([]interface{})[0].(map[string]interface{})
		fullname.Name = poolInfo["name"].(string)
	}

	return fullname
}

var (
	attachCluster = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				attachClusterKubeConfigPathKey: {
					Type:        schema.TypeString,
					Description: "Attach cluster KUBECONFIG path",
					ForceNew:    true,
					Optional:    true,
				},
				attachClusterKubeConfigRawKey: {
					Type:        schema.TypeString,
					Description: "Attach cluster KUBECONFIG",
					Optional:    true,
					ForceNew:    true,
					Sensitive:   true,
				},
				attachClusterDescriptionKey: {
					Type:         schema.TypeString,
					Description:  "Attach cluster description",
					Optional:     true,
					ValidateFunc: validation.StringIsNotWhiteSpace,
				},
			},
		},
	}
	KubeConfigWayAllowed = [...]string{attachClusterKubeConfigPathKey, attachClusterKubeConfigRawKey}
)

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
			imageRegistryNameKey: {
				Type:        schema.TypeString,
				Description: "Optional image registry name is the name of the image registry to be used for the cluster",
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

	if v, ok := specData[imageRegistryNameKey]; ok {
		spec.ImageRegistry = v.(string)
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

	flattenSpecData[imageRegistryNameKey] = spec.ImageRegistry

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

func validateKubeConfig(value interface{}) error {
	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("attach cluster data: %v is not valid: minimum one valid kube config type is required among: %v", data, strings.Join(KubeConfigWayAllowed[:], `, `))
	}

	kubeConfigData := data[0].(map[string]interface{})

	kubeConfigTypeFound := make([]string, 0)

	if v, ok := kubeConfigData[attachClusterKubeConfigPathKey]; ok {
		if v1, ok := v.(string); ok && len(v1) != 0 {
			kubeConfigTypeFound = append(kubeConfigTypeFound, attachClusterKubeConfigPathKey)
		}
	}

	if v, ok := kubeConfigData[attachClusterKubeConfigRawKey]; ok {
		if v1, ok := v.(string); ok && len(v1) != 0 {
			kubeConfigTypeFound = append(kubeConfigTypeFound, attachClusterKubeConfigRawKey)
		}
	}

	if len(kubeConfigTypeFound) == 0 {
		return fmt.Errorf("no valid kube config type found: minimum one valid kube config type is required among: %v", strings.Join(KubeConfigWayAllowed[:], `, `))
	} else if len(kubeConfigTypeFound) > 1 {
		return fmt.Errorf("found kube config types: %v are not valid: maximum one valid kube config type is allowed", strings.Join(kubeConfigTypeFound, `, `))
	}

	return nil
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	var (
		k8sclient  *k8sClient.Client
		err        error
		kubeConfig interface{}
		manifests  string
	)

	if v, ok := d.GetOk(attachClusterKey); ok {
		if v == nil {
			return diag.Errorf("data for attach cluster block not found: %v", v)
		}

		err = validateKubeConfig(v)
		if err != nil {
			return diag.FromErr(err)
		}

		isKubeConfigPresent := func(typeKey string) bool {
			if value, ok := d.GetOk(helper.GetFirstElementOf(attachClusterKey, typeKey)); ok {
				if value != nil {
					kubeConfig = value
					return true
				}
			}

			return false
		}

		switch {
		case isKubeConfigPresent(attachClusterKubeConfigPathKey):
			kubeConfigFile, _ := kubeConfig.(string)
			if strings.TrimSpace(kubeConfigFile) == "" {
				return diag.FromErr(fmt.Errorf("expected kubeconfig file path to not be an empty string or whitespace"))
			}

			k8sclient, err = getK8sClient(withPath(kubeConfigFile))

		case isKubeConfigPresent(attachClusterKubeConfigRawKey):
			rawKubeConfig, _ := kubeConfig.(string)
			if strings.TrimSpace(rawKubeConfig) == "" {
				return diag.FromErr(fmt.Errorf("expected raw kubeconfig to not be an empty string or whitespace"))
			}

			k8sclient, err = getK8sClient(withRaw(rawKubeConfig))
		}

		if err != nil {
			log.Println("[ERROR] error while creating kubernetes client: ", err.Error())
			return diag.FromErr(err)
		}

		if k8sclient == nil {
			err = errors.New("error while obtaining k8s client from REST config")
			log.Println("[ERROR] error while creating kubernetes client: ", err.Error())

			return diag.FromErr(err)
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

	if _, ok := d.GetOk(attachClusterKey); ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Kubernetes cluster's kubeconfig provided. Proceeding to attach the cluster TMC",
		})

		if clusterResponse.Cluster.Spec.ImageRegistry != "" || clusterResponse.Cluster.Spec.ProxyName != "" {
			clusterManifest, err := config.TMCConnection.ManifestResourceService.ClusterManifestHelperGetManifest(constructFullname(d))
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to get manifest (%s) for cluster entry, name : %s", clusterManifest.Manifest, err))
			}

			manifests = clusterManifest.Manifest
		} else {
			deploymentManifest, err := manifest.GetK8sManifest(clusterResponse.Cluster.Status.InstallerLink)
			if err != nil {
				return append(diags, diag.FromErr(err)...)
			}

			manifests = string(deploymentManifest)
		}

		log.Printf("[INFO] Applying %s cluster's deployment link manifest objects on to kubernetes cluster", constructFullname(d).ToString())

		err = manifest.Create(k8sclient, manifests, true)
		if err != nil {
			return append(diags, diag.FromErr(err)...)
		}

		log.Printf("[INFO] Cluster attach successful. Tanzu Mission Control resources applied to the cluster(%s) successfully", constructFullname(d).ToString())
	}

	return append(diags, dataSourceClusterRead(context.WithValue(ctx, contextMethodKey{}, "create"), d, m)...)
}

type (
	kubeConfigOption func(*kubeConfig)

	kubeConfig struct {
		filePath string
		raw      string
	}
)

func withPath(p string) kubeConfigOption {
	return func(config *kubeConfig) {
		config.filePath = p
	}
}

func withRaw(r string) kubeConfigOption {
	return func(config *kubeConfig) {
		config.raw = r
	}
}

func getK8sClient(opts ...kubeConfigOption) (*k8sClient.Client, error) {
	cfg := &kubeConfig{}

	for _, o := range opts {
		o(cfg)
	}

	var (
		restConfig *rest.Config
		err        error
	)

	switch {
	case cfg.filePath != "":
		restConfig, err = clientcmd.BuildConfigFromFlags("", cfg.filePath)
		if err != nil {
			return nil, errors.WithMessagef(err, "Invalid kubeconfig file path provided, filepath : %s", cfg.filePath)
		}
	case cfg.raw != "":
		restConfig, err = clientcmd.RESTConfigFromKubeConfig([]byte(cfg.raw))
		if err != nil {
			return nil, errors.WithMessagef(err, "Invalid raw kubeconfig provided.")
		}
	}

	if restConfig == nil {
		return nil, errors.WithMessagef(err, "Kubeconfig not provided.")
	}

	restConfig.Timeout = 10 * time.Second

	client, err := k8sClient.New(restConfig, k8sClient.Options{})
	if err != nil {
		return nil, errors.WithMessagef(err, "Error in creating kubernetes client from kubeconfig file provided, filepath : %s", cfg.filePath)
	}

	return &client, nil
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

	_, err = helper.Retry(getClusterResourceRetryableFn, 10*time.Second, 25)
	if err == nil {
		return diags
	}

	// if the cluster is still not removed then invoke force delete of the cluster.
	_, err = config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceGet(constructFullname(d))
	if err == nil {
		_ = config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceDelete(constructFullname(d), "true")

		log.Printf("[INFO] Cluster deletion in progress. Initiating force detach of the cluster entry as k8s cluster might not be responsive %s", constructFullname(d).ToString())

		diags = diag.FromErr(errors.Wrapf(err, "Initiating force detach for %s cluster."+
			"Ideally clean up of tmc agents and vmware-system-tmc namespace should have happened if not please remove them manually following "+
			"https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-3061A796-CA3D-4354-A0B7-19F50F2617CE.html", d.Get(NameKey)))
	}

	if err != nil {
		diags = diag.FromErr(errors.Wrapf(err, "verify %s cluster resource clean up", d.Get(NameKey)))
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

func withTKGNodePoolUpdate(d *schema.ResourceData, nodepool *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolNodepool) bool {
	switch {
	case nodepool.Spec.TkgServiceVsphere != nil:
		nodepools := d.Get(helper.GetFirstElementOf(SpecKey, tkgServiceVsphereKey, topologyKey, nodePoolKey)).([]interface{})[0].(map[string]interface{})
		poolSpec := nodepools[SpecKey].([]interface{})[0].(map[string]interface{})
		// poolInfo := nodepools["info"].([]interface{})[0].(map[string]interface{})
		tkgsSpec := poolSpec[tkgServiceVsphereKey].([]interface{})[0].(map[string]interface{})
		if d.HasChange(helper.GetFirstElementOf(SpecKey, tkgServiceVsphereKey, topologyKey, nodePoolKey)) {

			incomingWorkerNodeCount := poolSpec[workerNodeCountKey].(string)

			if incomingWorkerNodeCount != "" {
				nodepool.Spec.WorkerNodeCount = incomingWorkerNodeCount
			}

			incomingTkgServiceVsphereClass := tkgsSpec[classKey].(string)

			if incomingTkgServiceVsphereClass != "" {
				nodepool.Spec.TkgServiceVsphere.Class = incomingTkgServiceVsphereClass
			}

			incomingTkgServiceVsphereStorageClass := tkgsSpec[storageClassKey].(string)

			if incomingTkgServiceVsphereStorageClass != "" {
				nodepool.Spec.TkgServiceVsphere.StorageClass = incomingTkgServiceVsphereStorageClass
			}

			log.Printf("[INFO] updating TKGs workload cluster nodepools")
			return true
		}
	case nodepool.Spec.TkgVsphere != nil:
		nodepools := d.Get(helper.GetFirstElementOf(SpecKey, tkgVsphereClusterKey, topologyKey, nodePoolKey)).([]interface{})[0].(map[string]interface{})
		poolSpec := nodepools[SpecKey].([]interface{})[0].(map[string]interface{})
		if d.HasChange(helper.GetFirstElementOf(SpecKey, tkgVsphereClusterKey, topologyKey, nodePoolKey)) {
			incomingWorkerNodeCount := poolSpec[workerNodeCountKey].(string)

			if incomingWorkerNodeCount != "" {
				nodepool.Spec.WorkerNodeCount = incomingWorkerNodeCount
			}
		}
	}

	return false
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
	// check default nodepool configuration update
	npFullName := constructNpFullName(d)
	npResp, err := config.TMCConnection.NodePoolResourceService.ManageV1alpha1ClusterNodePoolResourceServiceGet(npFullName)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control cluster node pool entry, name : %s", npFullName.Name))
	}

	if withTKGNodePoolUpdate(d, npResp.Nodepool) {
		_, err = config.TMCConnection.NodePoolResourceService.ManageV1alpha1ClusterNodePoolResourceServiceUpdate(
			&nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest{
				Nodepool: npResp.Nodepool,
			},
		)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "Unable to update tanzu cluster node pool entry"))
		}

		log.Printf("[INFO] node pool update successful")
	}

	return dataSourceClusterRead(ctx, d, m)
}
