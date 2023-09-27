/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkc

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
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	tkcmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type (
	contextMethodKey struct{}
)

const defaultTimeout = 30 * time.Minute

func ResourceTMCTKCCluster() *schema.Resource {
	return &schema.Resource{
		Schema:        clusterSchema,
		CreateContext: resourceClusterCreate,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceTMCTKCClusterRead(helper.GetContextWithCaller(ctx, helper.RefreshState), d, m)
		},
		UpdateContext: resourceClusterInPlaceUpdate,
		DeleteContext: resourceClusterDelete,
		// Importer: &schema.ResourceImporter{
		// 	StateContext: resourceClusterImporter,
		// },
		Description: "Tanzu Mission Control EKS Cluster Resource",
	}
}

var clusterSchema = map[string]*schema.Schema{
	NameKey: {
		Type:        schema.TypeString,
		Description: "Name of this cluster",
		Required:    true,
		ForceNew:    true,
	},
	ManagementClusterNameKey: {
		Type:        schema.TypeString,
		Description: "Name of the management cluster",
		Required:    true,
		ForceNew:    true,
	},
	ProvisionerNameKey: {
		Type:        schema.TypeString,
		Description: "Provisioner of the cluster",
		Required:    true,
		ForceNew:    true,
	},
	common.MetaKey: common.Meta,
	specKey: {
		Type:        schema.TypeList,
		Description: "Spec for the cluster",
		Required:    true,
		MaxItems:    1,
		Elem:        clusterSpecSchema,
	},
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
		DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
			return true
		},
	},
}

var clusterSpecSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		clusterGroupKey: {
			Type:        schema.TypeString,
			Description: "Name of the cluster group to which this cluster belongs",
			Required:    true,
		},
		tmcManagedKey: {
			Type:        schema.TypeBool,
			Description: "TMC-managed flag indicates if the cluster is managed by tmc",
			Required:    true,
		},
		proxyNameKey: {
			Type:        schema.TypeString,
			Description: "Optional proxy name is the name of the Proxy Config to be used for the cluster",
			Optional:    true,
		},
		imageRegistryKey: {
			Type:        schema.TypeString,
			Description: "Name of the image registry configuration to use",
			Optional:    true,
		},
		topologyKey: {
			Type:        schema.TypeList,
			Description: "The cluster topology",
			Required:    true,
			MaxItems:    1,
			Elem:        topologySchema,
		},
	},
}

var topologySchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		versionKey: {
			Type:        schema.TypeString,
			Description: "Kubernetes version of the cluster",
			Required:    true,
		},
		clusterClassKey: {
			Type:        schema.TypeString,
			Description: "The name of the cluster class for the cluster",
			Required:    true,
			// ForceNew:    true,
		},
		controlPlaneKey: {
			Type:        schema.TypeList,
			Description: "The cluster specific control plane configuration",
			Required:    true,
			MinItems:    1,
			Elem:        controlPlaneSchema,
		},
		nodepoolsKey: {
			Type:        schema.TypeList,
			Description: "Nodepool definition for the cluster",
			Required:    true,
			MinItems:    1,
			Elem:        nodepoolsDefinitionSchema,
		},
		clusterVariablesKey: {
			Type:        schema.TypeList,
			Description: "Variables configuration for the cluster",
			Required:    true,
			Elem:        commonClusterClusterVariableSchema,
		},
		networkKey: {
			Type:        schema.TypeList,
			Description: "Network related settings for the cluster",
			Required:    true,
			ForceNew:    true,
			Elem:        networkSchema,
		},
		coreAddonsKey: {
			Type:        schema.TypeList,
			Description: "The core addons",
			Required:    true,
			Elem:        coreAddonsSchema,
		},
	},
}

var controlPlaneSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		replicasKey: {
			Type:        schema.TypeInt,
			Description: "The replicas of the control plane",
			Required:    true,
			// ForceNew:    false,
		},
		metadataKey: {
			Type:        schema.TypeList,
			Description: "The labels and annotations configurations of the control plane",
			Optional:    true,
			Elem:        commonClusterMetadataSchema,
		},
		osImageKey: {
			Type:        schema.TypeList,
			Description: "The OS image configuration of the control plane",
			Required:    true,
			Elem:        commonClusterOsImageSchema,
		},
	},
}

var networkSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		podCidrBlocksKey: {
			Type:        schema.TypeList,
			Description: "NetworkRanges describes a collection of IP addresses as a list of ranges",
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		serviceCidrBlocksKey: {
			Type:        schema.TypeList,
			Description: "NetworkRanges describes a collection of IP addresses as a list of ranges",
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		serviceDomainKey: {
			Type:        schema.TypeString,
			Description: "Domain name for services",
			Optional:    true,
		},
	},
}

var coreAddonsSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			typeKey: {
				Type:        schema.TypeString,
				Description: "Type of core addon, e.g. 'cni'",
				Required:    true,
			},
			providerKey: {
				Type:        schema.TypeString,
				Description: "Provider of core addon, e.g. 'antrea', 'calico'",
				Required:    true,
			},
		},
	},
}

func constructTkcClusterSpec(d *schema.ResourceData) (spec *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterSpec) {
	spec = &tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterSpec{}

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

	if v, ok := specData[tmcManagedKey]; ok {
		helper.SetPrimitiveValue(v, &spec.TmcManaged, tmcManagedKey)
	}

	if v, ok := specData[proxyNameKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ProxyName, proxyNameKey)
	}

	if v, ok := specData[imageRegistryKey]; ok {
		helper.SetPrimitiveValue(v, &spec.ImageRegistry, imageRegistryKey)
	}

	if v, ok := specData[topologyKey]; ok {
		topologyData, _ := v.([]interface{})
		spec.Topology = constructTopology(topologyData)
	}

	return spec
}

func constructTopology(data []interface{}) *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology {
	topology := &tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology{}

	if len(data) == 0 || data[0] == nil {
		return topology
	}

	topologyData, _ := data[0].(map[string]interface{})

	if v, ok := topologyData[versionKey]; ok {
		helper.SetPrimitiveValue(v, &topology.Version, versionKey)
	}

	if v, ok := topologyData[clusterClassKey]; ok {
		helper.SetPrimitiveValue(v, &topology.ClusterClass, clusterClassKey)
	}

	if v, ok := topologyData[controlPlaneKey]; ok {
		data, _ := v.([]interface{})
		topology.ControlPlane = constructControlPlane(data)
	}

	if v, ok := topologyData[nodepoolsKey]; ok {
		data, _ := v.([]interface{})
		topology.NodePools = constructNodePools(data)
	}

	if v, ok := topologyData[clusterVariablesKey]; ok {
		clusterVariablesData, _ := v.([]interface{})
		variablesData, _ := clusterVariablesData[0].(map[string]interface{})
		if v2, ok := variablesData[variablesKey]; ok {
			data, _ := v2.([]interface{})
			topology.Variables = constructCommonClusterClusterVariables(data)
		}
	}

	if v, ok := topologyData[networkKey]; ok {
		data, _ := v.([]interface{})
		topology.Network = constructNetwork(data)
	}

	if v, ok := topologyData[coreAddonsKey]; ok {
		data, _ := v.([]interface{})
		topology.CoreAddons = constructCoreAddons(data)
	}

	return topology
}

func constructControlPlane(data []interface{}) *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterControlPlane {
	cp := &tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterControlPlane{}
	if len(data) == 0 || data[0] == nil {
		return cp
	}

	controlPlaneData, _ := data[0].(map[string]interface{})

	if v, ok := controlPlaneData[replicasKey]; ok {
		helper.SetPrimitiveValue(v, &cp.Replicas, replicasKey)
	}

	if v, ok := controlPlaneData[metadataKey]; ok {
		data, _ := v.([]interface{})
		cp.Metadata = constructCommonClusterMetadata(data)
	}

	if v, ok := controlPlaneData[osImageKey]; ok {
		data, _ := v.([]interface{})
		cp.OsImage = constructCommonClusterOsImage(data)
	}

	return cp
}

func constructNetwork(data []interface{}) *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings {
	netSettings := &tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings{}

	if len(data) == 0 || data[0] == nil {
		return netSettings
	}

	netSettingsData, _ := data[0].(map[string]interface{})

	if v, ok := netSettingsData[podCidrBlocksKey]; ok {
		if data, ok := v.(*schema.Set); ok {
			netSettings.Pods.CidrBlocks = constructStringList(data.List())
		}
	}

	if v, ok := netSettingsData[serviceCidrBlocksKey]; ok {
		if data, ok := v.(*schema.Set); ok {
			netSettings.Services.CidrBlocks = constructStringList(data.List())
		}
	}

	if v, ok := netSettingsData[serviceDomainKey]; ok {
		helper.SetPrimitiveValue(v, &netSettings.ServiceDomain, serviceDomainKey)
	}

	return netSettings
}

func constructCoreAddons(data []interface{}) []*tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCoreAddon {
	coreAddons := []*tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCoreAddon{}

	if len(data) == 0 || data[0] == nil {
		return coreAddons
	}

	for _, v := range data {
		ca := constructCoreAddon(v)
		coreAddons = append(coreAddons, ca)
	}

	return coreAddons
}

func constructCoreAddon(data interface{}) *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCoreAddon {
	coreAddon := &tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCoreAddon{}

	coreAddonData := data.(map[string]interface{})

	if v, ok := coreAddonData[typeKey]; ok {
		helper.SetPrimitiveValue(v, &coreAddon.Type, typeKey)
	}

	if v, ok := coreAddonData[providerKey]; ok {
		helper.SetPrimitiveValue(v, &coreAddon.Provider, providerKey)
	}

	return coreAddon
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config, ok := m.(authctx.TanzuContext)
	if !ok {
		log.Println("[ERROR] error while retrieving Tanzu auth config")
		return diag.Errorf("error while retrieving Tanzu auth config")
	}

	clusterFn := constructFullname(d)
	clusterSpec := constructTkcClusterSpec(d)

	clusterReq := &tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterRequest{
		TanzuKubernetesCluster: &tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster{
			FullName: clusterFn,
			Meta:     common.ConstructMeta(d),
			Spec:     clusterSpec,
		},
	}

	var tkcCluster *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster

	clusterResponse, err := config.TMCConnection.TKCClusterResourceService.TanzuKubernetesClusterResourceServiceCreate(clusterReq)
	if err != nil {
		if !clienterrors.IsAlreadyExistsError(err) {
			return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
		}

		clusterResponse, err := config.TMCConnection.TKCClusterResourceService.TanzuKubernetesClusterResourceServiceGet(clusterFn)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
		}

		tkcCluster = clusterResponse.TanzuKubernetesCluster
	} else {
		tkcCluster = clusterResponse.TanzuKubernetesCluster
	}

	d.SetId(tkcCluster.Meta.UID)

	return dataSourceTMCTKCClusterRead(context.WithValue(ctx, contextMethodKey{}, "create"), d, m)
}

func resourceClusterDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := config.TMCConnection.TKCClusterResourceService.TanzuKubernetesClusterResourceServiceDelete(constructFullname(d), "false")
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	clusterFn := constructFullname(d)
	getClusterResourceRetryableFn := func() (retry bool, err error) {
		_, err = config.TMCConnection.TKCClusterResourceService.TanzuKubernetesClusterResourceServiceGet(clusterFn)
		if err == nil {
			log.Printf("[DEBUG] cluster(%s) deletion in progress", clusterFn.ToString())
			return true, errors.New("cluster deletion in progress")
		}

		if !clienterrors.IsNotFoundError(err) {
			return true, err
		}

		return false, nil
	}

	timeoutDuration := getRetryTimeout(d)

	_, err = helper.RetryUntilTimeout(getClusterResourceRetryableFn, 10*time.Second, timeoutDuration)
	if err != nil {
		diag.FromErr(errors.Wrapf(err, "verify %s EKS cluster resource clean up", d.Get(NameKey)))
	}

	return diags
}

func resourceClusterInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	// Get call to initialise the cluster struct
	getResp, err := config.TMCConnection.TKCClusterResourceService.TanzuKubernetesClusterResourceServiceGet(constructFullname(d))
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	clusterSpec := constructTkcClusterSpec(d)

	errcl := handleClusterDiff(config, getResp.TanzuKubernetesCluster, common.ConstructMeta(d), clusterSpec)
	if errcl != nil {
		return diag.FromErr(errors.Wrapf(errcl, "Unable to update Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	log.Printf("[INFO] cluster update successful")

	return dataSourceTMCTKCClusterRead(ctx, d, m)
}

// func resourceClusterImporter(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
// 	config := m.(authctx.TanzuContext)

// 	// id := d.Id()

// 	// if id == "" {
// 	// 	return nil, errors.New("ID is needed to import an TMC EKS cluster")
// 	// }

// 	// resp, err := config.TMCConnection.TKCClusterResourceService.TanzuKubernetesClusterResourceServiceGetByID(id)
// 	// if err != nil {
// 	// 	return nil, errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS cluster entry for id %s", id)
// 	// }

// 	// if err = d.Set(NameKey, resp.TanzuKubernetesCluster.FullName.Name); err != nil {
// 	// 	return nil, errors.Wrapf(err, "Failed to set name for the cluster %s", resp.TanzuKubernetesCluster.FullName.Name)
// 	// }

// 	// if err = d.Set(ManagementClusterNameKey, resp.TanzuKubernetesCluster.FullName.ManagementClusterName); err != nil {
// 	// 	return nil, errors.Wrapf(err, "Failed to set management cluster name for the cluster %s", resp.TanzuKubernetesCluster.FullName.Name)
// 	// }

// 	// if err = d.Set(ProvisionerNameKey, resp.TanzuKubernetesCluster.FullName.ProvisionerName); err != nil {
// 	// 	return nil, errors.Wrapf(err, "Failed to set provisioner name for the cluster %s", resp.TanzuKubernetesCluster.FullName.Name)
// 	// }

// 	// err = setResourceData(d, resp.TanzuKubernetesCluster)
// 	// if err != nil {
// 	// 	return nil, errors.Wrapf(err, "Failed to set resource data during import for %s", resp.TanzuKubernetesCluster.FullName.Name)
// 	// }

// 	// return []*schema.ResourceData{d}, nil

// 	return dataSourceTMCTKCClusterRead(context.WithValue(ctx, contextMethodKey{}, "import"), d, m)
// }

func handleClusterDiff(config authctx.TanzuContext, tmcCluster *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster, meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta, clusterSpec *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterSpec) error {
	// updateCluster := false

	// if meta.Description != tmcCluster.Meta.Description ||
	// 	!mapEqual(meta.Labels, tmcCluster.Meta.Labels) {
	// 	updateCluster = true
	// 	tmcCluster.Meta.Description = meta.Description
	// 	tmcCluster.Meta.Labels = meta.Labels
	// }

	// if !clusterSpecEqual(clusterSpec, tmcCluster.Spec) {
	// 	updateCluster = true
	// }

	// if !updateCluster {
	// 	return nil
	// }

	newCluster := &tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTanzuKubernetesCluster{
		FullName: tmcCluster.FullName,
		Meta:     tmcCluster.Meta,
		Spec:     clusterSpec,
	}

	_, err := config.TMCConnection.TKCClusterResourceService.TanzuKubernetesClusterResourceServiceUpdate(
		&tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCreateUpdateTanzuKubernetesClusterRequest{
			TanzuKubernetesCluster: newCluster,
		},
	)

	if err != nil {
		return errors.Wrapf(err, "Unable to update Tanzu Mission Control EKS cluster entry, name : %s", tmcCluster.FullName.Name)
	}

	return nil
}

func constructFullname(d *schema.ResourceData) (fullname *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName) {
	fullname = &tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterFullName{}

	fullname.ManagementClusterName, _ = d.Get(ManagementClusterNameKey).(string)
	fullname.Name, _ = d.Get(NameKey).(string)
	// fullname.OrgID, _ = d.Get(OrgIDKey).(string)
	fullname.ProvisionerName, _ = d.Get(ProvisionerNameKey).(string)

	return fullname
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

func flattenClusterSpec(item *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterSpec) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	data[clusterGroupKey] = item.ClusterGroupName

	if item.ImageRegistry != "" {
		data[imageRegistryKey] = item.ImageRegistry
	}

	if item.ProxyName != "" {
		data[proxyNameKey] = item.ProxyName
	}

	data[tmcManagedKey] = item.TmcManaged

	if item.Topology != nil {
		data[topologyKey] = flattenTopology(item.Topology)
	}

	return []interface{}{data}
}

func flattenTopology(item *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	if item.ClusterClass != "" {
		data[clusterClassKey] = item.ClusterClass
	}

	if item.ControlPlane != nil {
		data[controlPlaneKey] = flattenControlPlane(item.ControlPlane)
	}

	if item.CoreAddons != nil {
		data[coreAddonsKey] = flattenCoreAddons(item.CoreAddons)
	}

	if item.Network != nil {
		data[networkKey] = flattenNetwork(item.Network)
	}

	if item.NodePools != nil {
		data[nodepoolsKey] = flattenNodePools(item.NodePools)
	}

	if item.Variables != nil {
		data[clusterVariablesKey] = flattenCommonClusterVariables(item.Variables)
	}

	if item.Version != "" {
		data[versionKey] = item.Version
	}

	return []interface{}{data}
}

func flattenControlPlane(item *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterControlPlane) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	if item.Metadata != nil {
		data[metadataKey] = flattenCommonClusterMetadata(item.Metadata)
	}

	if item.OsImage != nil {
		data[osImageKey] = flattenCommonClusterOsImage(item.OsImage)
	}

	data[replicasKey] = item.Replicas

	return []interface{}{data}
}

func flattenCoreAddons(arr []*tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCoreAddon) []interface{} {
	data := make([]interface{}, 0, len(arr))

	if len(arr) == 0 {
		return data
	}

	for _, item := range arr {
		data = append(data, flattenCoreAddon(item))
	}

	return data
}

func flattenCoreAddon(item *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCoreAddon) []interface{} {
	data := make(map[string]interface{})

	if item.Provider != "" {
		data[providerKey] = item.Provider
	}

	if item.Type != "" {
		data[typeKey] = item.Type
	}

	return []interface{}{data}
}

func flattenNetwork(item *tkcmodel.VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings) []interface{} {
	if item == nil {
		return []interface{}{}
	}

	data := make(map[string]interface{})

	if len(item.Pods.CidrBlocks) > 0 {
		data[podCidrBlocksKey] = item.Pods.CidrBlocks
	}

	if item.ServiceDomain != "" {
		data[serviceDomainKey] = item.ServiceDomain
	}

	if len(item.Services.CidrBlocks) > 0 {
		data[serviceCidrBlocksKey] = item.Services.CidrBlocks
	}

	return []interface{}{data}
}
