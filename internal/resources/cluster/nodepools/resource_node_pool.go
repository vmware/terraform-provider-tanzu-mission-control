/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepools

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/helper"
	nodepoolmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
	"github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/resources/common"
)

const (
	ResourceName = "tanzu-mission-control_cluster_node_pool"
)

func ResourceNodePool() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceClusterNodePoolRead,
		CreateContext: resourceNodePoolCreate,
		UpdateContext: resourceClusterNodePoolInPlaceUpdate,
		DeleteContext: resourceClusterNodePoolDelete,
		Schema:        nodePoolSchema,
	}
}

var nodePoolSchema = map[string]*schema.Schema{
	managementClusterNameKey: {
		Type:        schema.TypeString,
		Description: "Name of the management cluster",
		Required:    true,
		ForceNew:    true,
	},
	provisionerNameKey: {
		Type:        schema.TypeString,
		Description: "Provisioner of the cluster",
		Required:    true,
		ForceNew:    true,
	},
	clusterNameKey: {
		Type:        schema.TypeString,
		Description: "Name of the cluster",
		Required:    true,
		ForceNew:    true,
	},
	nodePoolNameKey: {
		Type:        schema.TypeString,
		Description: "Name of this nodepool",
		Required:    true,
		ForceNew:    true,
	},
	common.MetaKey: common.Meta,
	specKey:        nodePoolSpec,
	statusKey: {
		Type:        schema.TypeMap,
		Description: "Status of node pool resource",
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
}

var nodePoolSpec = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the cluster nodepool",
	Optional:    true,
	ForceNew:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			workerNodeCountKey: {
				Type:        schema.TypeString,
				Description: "Count is the number of nodes",
				Default:     "1",
				Optional:    true,
			},
			cloudLabelsKey: {
				Type:        schema.TypeMap,
				Description: "Cloud labels",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			nodeLabelsKey: {
				Type:        schema.TypeMap,
				Description: "Node labels",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			tkgServiceVsphereKey: NodePoolTkgServiceVsphere,
		},
	},
}

var NodePoolTkgServiceVsphere = &schema.Schema{
	Type:        schema.TypeList,
	Description: "TKGServiceVsphereNodepool is the nodepool spec for TKG service vsphere cluster",
	Optional:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			classKey: {
				Type:        schema.TypeString,
				Description: "Nodepool instance type",
				Optional:    true,
			},
			storageClassKey: {
				Type:        schema.TypeString,
				Description: "Storage Class to be used for storage of the disks which store the root filesystem of the nodes",
				Optional:    true,
			},
		},
	},
}

func constructFullName(d *schema.ResourceData) (fullname *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolFullName) {
	fullname = &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolFullName{}

	if value, ok := d.GetOk(managementClusterNameKey); ok {
		fullname.ManagementClusterName = value.(string)
	}

	if value, ok := d.GetOk(provisionerNameKey); ok {
		fullname.ProvisionerName = value.(string)
	}

	if value, ok := d.GetOk(clusterNameKey); ok {
		fullname.ClusterName = value.(string)
	}

	if value, ok := d.GetOk(nodePoolNameKey); ok {
		fullname.Name = value.(string)
	}

	return fullname
}

func constructNodePoolSpec(d *schema.ResourceData) (spec *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec) {
	spec = &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec{
		WorkerNodeCount:   "1",
		CloudLabels:       make(map[string]string),
		NodeLabels:        make(map[string]string),
		TkgServiceVsphere: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{},
		TkgVsphere:        &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool{},
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

	if v, ok := specData[workerNodeCountKey]; ok {
		spec.WorkerNodeCount, _ = v.(string)
	}

	if v, ok := specData[cloudLabelsKey]; ok {
		if v1, ok := v.(map[string]interface{}); ok {
			spec.CloudLabels = common.GetTypeMapData(v1)
		}
	}

	if v, ok := specData[nodeLabelsKey]; ok {
		if v1, ok := v.(map[string]interface{}); ok {
			spec.NodeLabels = common.GetTypeMapData(v1)
		}
	}

	if v, ok := specData[tkgServiceVsphereKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.TkgServiceVsphere = constructTkgServiceVsphere(v1)
		}
	}

	return spec
}

func flattenSpec(spec *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	flattenSpecData[cloudLabelsKey] = spec.CloudLabels
	flattenSpecData[nodeLabelsKey] = spec.NodeLabels
	flattenSpecData[workerNodeCountKey] = spec.WorkerNodeCount

	if spec.TkgServiceVsphere != nil {
		flattenSpecData[tkgServiceVsphereKey] = flattenTkgServiceVsphere(spec.TkgServiceVsphere)
	}

	if spec.TkgVsphere != nil {
		flattenSpecData[tkgVsphereKey] = flattenNodePoolTKGVsphere(spec.TkgVsphere)
	}

	return []interface{}{flattenSpecData}
}

func flattenNodePoolTKGVsphere(tkgVsphere *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGVsphereNodepool) (data []interface{}) {
	flattenTKGVsphere := make(map[string]interface{})

	if tkgVsphere.VMConfig != nil {
		flattenTKGVsphere[vmConfigKey] = flattenTKGVsphereVMConfig(tkgVsphere.VMConfig)
	}

	return []interface{}{flattenTKGVsphere}
}

func flattenTKGVsphereVMConfig(vmConfig *nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig) (data []interface{}) {
	flattenVMConfig := make(map[string]interface{})

	flattenVMConfig[cpuKey] = vmConfig.CPU
	flattenVMConfig[diskKey] = vmConfig.DiskGib
	flattenVMConfig[memoryKey] = vmConfig.MemoryMib

	return []interface{}{flattenVMConfig}
}

func flattenTkgServiceVsphere(tkgServiceVsphere *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool) (data []interface{}) {
	flattenTkgServiceVsphereData := make(map[string]interface{})

	flattenTkgServiceVsphereData[classKey] = tkgServiceVsphere.Class
	flattenTkgServiceVsphereData[storageClassKey] = tkgServiceVsphere.StorageClass

	return []interface{}{flattenTkgServiceVsphereData}
}

func constructTkgServiceVsphere(data []interface{}) (tkgServiceVsphere *nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool) {
	if len(data) == 0 || data[0] == nil {
		return tkgServiceVsphere
	}

	lookUpTkgServiceVsphere, _ := data[0].(map[string]interface{})
	tkgServiceVsphere = &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool{}

	if v, ok := lookUpTkgServiceVsphere[classKey]; ok {
		tkgServiceVsphere.Class, _ = v.(string)
	}

	if v, ok := lookUpTkgServiceVsphere[storageClassKey]; ok {
		tkgServiceVsphere.StorageClass, _ = v.(string)
	}

	return tkgServiceVsphere
}

func resourceNodePoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	nodePoolRequest := &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest{
		Nodepool: &nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolNodepool{
			FullName: constructFullName(d),
			Meta:     common.ConstructMeta(d),
			Spec:     constructNodePoolSpec(d),
		},
	}

	if nodePoolRequest.Nodepool.Spec.TkgServiceVsphere == nil {
		return diag.FromErr(fmt.Errorf("TKGs vsphere nodepool spec has to be provided"))
	}

	nodePoolResponse, err := config.TMCConnection.NodePoolResourceService.ManageV1alpha1ClusterNodePoolResourceServiceCreate(nodePoolRequest)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "unable to create tanzu node pool entry"))
	}

	d.SetId(nodePoolResponse.Nodepool.FullName.Name + ":" + nodePoolResponse.Nodepool.FullName.ClusterName)

	return dataSourceClusterNodePoolRead(ctx, d, m)
}

func resourceClusterNodePoolInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	updateRequired := false

	getResp, err := config.TMCConnection.NodePoolResourceService.ManageV1alpha1ClusterNodePoolResourceServiceGet(constructFullName(d))
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get tanzu cluster node pool entry"))
	}

	switch {
	case getResp.Nodepool.Spec.TkgServiceVsphere != nil:
		if d.HasChange(helper.GetFirstElementOf(specKey, workerNodeCountKey)) ||
			d.HasChange(helper.GetFirstElementOf(specKey, tkgServiceVsphereKey, classKey)) ||
			d.HasChange(helper.GetFirstElementOf(specKey, tkgServiceVsphereKey, storageClassKey)) {
			updateRequired = true
		}

		if !updateRequired {
			return diags
		}

		incomingWorkerNodeCount := d.Get(helper.GetFirstElementOf(specKey, workerNodeCountKey))

		if incomingWorkerNodeCount.(string) != "" {
			getResp.Nodepool.Spec.WorkerNodeCount = incomingWorkerNodeCount.(string)
		}

		incomingTkgServiceVsphereClass := d.Get(helper.GetFirstElementOf(specKey, tkgServiceVsphereKey, classKey))

		if incomingTkgServiceVsphereClass.(string) != "" {
			getResp.Nodepool.Spec.TkgServiceVsphere.Class = incomingTkgServiceVsphereClass.(string)
		}

		incomingTkgServiceVsphereStorageClass := d.Get(helper.GetFirstElementOf(specKey, tkgServiceVsphereKey, storageClassKey))

		if incomingTkgServiceVsphereStorageClass.(string) != "" {
			getResp.Nodepool.Spec.TkgServiceVsphere.StorageClass = incomingTkgServiceVsphereStorageClass.(string)
		}

	case getResp.Nodepool.Spec.TkgVsphere != nil:
		if d.HasChange(helper.GetFirstElementOf(specKey, workerNodeCountKey)) {
			updateRequired = true
		}

		if !updateRequired {
			return diags
		}

		incomingWorkerNodeCount := d.Get(helper.GetFirstElementOf(specKey, workerNodeCountKey))

		if incomingWorkerNodeCount.(string) != "" {
			getResp.Nodepool.Spec.WorkerNodeCount = incomingWorkerNodeCount.(string)
		}
	}

	_, err = config.TMCConnection.NodePoolResourceService.ManageV1alpha1ClusterNodePoolResourceServiceUpdate(
		&nodepoolmodel.VmwareTanzuManageV1alpha1ClusterNodepoolCreateNodepoolRequest{
			Nodepool: getResp.Nodepool,
		},
	)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to update tanzu cluster node pool entry"))
	}

	return dataSourceClusterNodePoolRead(ctx, d, m)
}

func resourceClusterNodePoolDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	var diags diag.Diagnostics

	err := config.TMCConnection.NodePoolResourceService.ManageV1alpha1ClusterNodePoolResourceServiceDelete(constructFullName(d))
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete tanzu cluster node pool entry"))
	}

	d.SetId("")

	return diags
}
