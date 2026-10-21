// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzukubernetescluster

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	clusterclassmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clusterclass"
	tanzukubernetesclustermodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
	tkcnodepoolmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/nodepool"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clusterclass"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

// clusterResourceUpdateKeys keys in the spec of the resource that are related to the cluster object - needed for change checks on the resource for the cluster API.
var clusterResourceUpdateKeys = []string{
	strings.Join([]string{SpecKey, "0", ClusterGroupNameKey}, "."),
	strings.Join([]string{SpecKey, "0", ProxyNameKey}, "."),
	strings.Join([]string{SpecKey, "0", ImageRegistryKey}, "."),
	strings.Join([]string{SpecKey, "0", TopologyKey, "0", VersionKey}, "."),
	strings.Join([]string{SpecKey, "0", TopologyKey, "0", ClusterClassKey}, "."),
	strings.Join([]string{SpecKey, "0", TopologyKey, "0", ClusterVariablesKey}, "."),
	strings.Join([]string{SpecKey, "0", TopologyKey, "0", CoreAddonKey}, "."),
	strings.Join([]string{SpecKey, "0", TopologyKey, "0", NetworkKey}, "."),
	strings.Join([]string{SpecKey, "0", TopologyKey, "0", ControlPlaneKey}, "."),
	strings.Join([]string{common.MetaKey, "0", common.LabelsKey}, "."),
	strings.Join([]string{common.MetaKey, "0", common.AnnotationsKey}, "."),
	strings.Join([]string{common.MetaKey, "0", common.DescriptionKey}, "."),
}

// nodePoolResourceKey keys in the spec of the resource that are related to the node pool object - needed for change checks on the resource for the node pool API.
var nodePoolResourceKey = strings.Join([]string{SpecKey, "0", TopologyKey, "0", NodePoolKey}, ".")

func ResourceTanzuKubernetesCluster() *schema.Resource {
	return &schema.Resource{
		Schema:               tanzuKubernetesClusterSchema,
		CreateWithoutTimeout: resourceTanzuKubernetesClusterCreate,
		ReadWithoutTimeout:   resourceTanzuKubernetesClusterRead,
		UpdateWithoutTimeout: resourceTanzuKubernetesClusterUpdate,
		DeleteWithoutTimeout: resourceTanzuKubernetesClusterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceTanzuKubernetesClusterImporter,
		},
		CustomizeDiff: validateSchema,
	}
}

func resourceTanzuKubernetesClusterCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{})
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create TKG Cluster."))
	}

	// Detach node pools from the initial cluster spec because node pools must be created
	// separately via the dedicated node pool API once the cluster definition is accepted.
	modelNodePools := model.Spec.Topology.NodePools
	model.Spec.Topology.NodePools = nil

	// Retrieve the user-defined timeout policy or fallback to default timeout for bounded retry operations.
	timeoutPolicy := getTimeoutPolicy(data)
	timeoutMinutes := timeoutPolicy[TimeoutKey].(int)
	clusterTimeout := time.Duration(TimeoutDefaultValue) * time.Minute

	if timeoutMinutes > 0 {
		clusterTimeout = time.Duration(timeoutMinutes) * time.Minute
	}

	clusterRequest := &tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
		TanzuKubernetesCluster: model,
	}

	// Create cluster with adaptive retry to handle transient network/gateway drops, 429 rate limits,
	// or 409 conflict when a previously deleted cluster with the same name is still terminating in the background.
	err = createClusterWithRetry(ctx, &config, clusterRequest, clusterTimeout)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create TKG Cluster.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s",
			model.FullName.ManagementClusterName, model.FullName.ProvisionerName, model.FullName.Name))
	}

	// Set ID early so if node pool creation fails, times out, or is cancelled, the cluster is recorded in state and not orphaned.
	fullNameList := []string{model.FullName.ManagementClusterName, model.FullName.ProvisionerName, model.FullName.Name}
	data.SetId(strings.Join(fullNameList, "/"))

	// Create node pools sequentially with adaptive retry.
	// We use polling with exponential backoff and jitter instead of a static sleep to gracefully handle
	// admission webhook timing (e.g., SystemChecksSucceeded condition) and cluster initialization race conditions.
	for _, np := range modelNodePools {
		np.FullName.ManagementClusterName = model.FullName.ManagementClusterName
		np.FullName.ProvisionerName = model.FullName.ProvisionerName
		np.FullName.TanzuKubernetesClusterName = model.FullName.Name
		nodePoolRequest := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{
			Nodepool: np,
		}

		err = createNodePoolWithRetry(ctx, &config, nodePoolRequest, clusterTimeout)
		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "Couldn't create TKG Cluster Nodepool.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s, Node Pool Name: %s",
				np.FullName.ManagementClusterName, np.FullName.ProvisionerName, np.FullName.TanzuKubernetesClusterName, np.FullName.Name))
		}
	}

	return resourceTanzuKubernetesClusterRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func resourceTanzuKubernetesClusterRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var (
		resp *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData
	)

	config := m.(authctx.TanzuContext)

	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey, ProvisionerNameKey, ManagementClusterNameKey, TimeoutPolicyKey,
		strings.Join([]string{SpecKey, TopologyKey, NodePoolKey}, tfModelConverterHelper.DefaultModelPathSeparator)})
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read TKG Cluster."))
	}

	clusterFn := model.FullName

	if !helper.IsContextCallerSet(ctx) {
		resp, err = readFullClusterResource(&config, clusterFn)
	} else {
		timeoutPolicy := getTimeoutPolicy(data)
		resp, err = readResourceWait(ctx, &config, clusterFn, model.Spec.Topology.NodePools, timeoutPolicy)
	}

	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			if !helper.IsContextCallerSet(ctx) || helper.IsDeleteState(ctx) {
				// d.SetId("") is called to signal Terraform that the resource no longer exists
				// (e.g. drift or completed deletion) so it can be cleanly removed from state.
				_ = schema.RemoveFromState(data, m)

				return diags
			}
		}

		return diag.FromErr(errors.Wrapf(err, "Couldn't read TKG cluster.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name))
	} else if resp != nil {
		kubernetesClusterModel := resp.TanzuKubernetesCluster
		specData := data.Get(SpecKey).([]interface{})[0].(map[string]interface{})
		topologyData := specData[TopologyKey].([]interface{})[0].(map[string]interface{})
		clusterVariablesData := topologyData[ClusterVariablesKey].(string)
		controlPlaneData := topologyData[ControlPlaneKey].([]interface{})[0]
		nodePoolsData := topologyData[NodePoolKey].([]interface{})

		removeUnspecifiedClusterVariables(clusterVariablesData, kubernetesClusterModel)
		removeUnspecifiedControlPlaneOverrides(controlPlaneData, kubernetesClusterModel)
		removeUnspecifiedNodePoolsOverrides(nodePoolsData, kubernetesClusterModel)

		err = tfModelResourceConverter.FillTFSchema(kubernetesClusterModel, data)
		if err != nil {
			diags = diag.FromErr(err)
		}

		suppressNodePoolsOrderChanges(nodePoolsData, data)

		fullNameList := []string{kubernetesClusterModel.FullName.ManagementClusterName, kubernetesClusterModel.FullName.ProvisionerName, kubernetesClusterModel.FullName.Name}

		data.SetId(strings.Join(fullNameList, "/"))
	}

	return diags
}

func resourceTanzuKubernetesClusterDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey, ProvisionerNameKey, ManagementClusterNameKey,
		strings.Join([]string{SpecKey, TopologyKey, NodePoolKey}, tfModelConverterHelper.DefaultModelPathSeparator)})
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete TKG cluster"))
	}

	clusterFn := model.FullName
	err = config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterResourceServiceDelete(clusterFn, false)

	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete delete TKG cluster.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name))
	}

	return resourceTanzuKubernetesClusterRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func resourceTanzuKubernetesClusterUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	model, err := tfModelResourceConverter.ConvertTFSchemaToAPIModel(data, []string{})
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update TKG Cluster."))
	}

	if data.HasChangeExcept(TimeoutPolicyKey) {
		modelNodePools := model.Spec.Topology.NodePools

		if data.HasChanges(clusterResourceUpdateKeys...) {
			model.Spec.Topology.NodePools = nil

			clusterRequest := &tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData{
				TanzuKubernetesCluster: model,
			}

			_, err = config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterResourceServiceUpdate(clusterRequest)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Couldn't update TKG Cluster.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s",
					model.FullName.ManagementClusterName, model.FullName.ProvisionerName, model.FullName.Name))
			}
		}

		if data.HasChange(nodePoolResourceKey) {
			// Update node pools and propagate any diagnostics/errors to prevent silent failures.
			diags = resourceTanzuKubernetesClusterNodePoolsUpdate(ctx, config, data, modelNodePools, model.FullName)
			if diags.HasError() {
				return diags
			}
		}

		return resourceTanzuKubernetesClusterRead(helper.GetContextWithCaller(ctx, helper.UpdateState), data, m)
	}

	return diags
}

func resourceTanzuKubernetesClusterNodePoolsUpdate(ctx context.Context, config authctx.TanzuContext, data *schema.ResourceData, modelNodePools []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool, clusterFn *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (diags diag.Diagnostics) {
	oldTFValue, newTFValue := data.GetChange(nodePoolResourceKey)
	existingNodePools := make(map[string]interface{})

	// Retrieve timeout for node pool lifecycle operations.
	timeoutPolicy := getTimeoutPolicy(data)
	timeoutMinutes := timeoutPolicy[TimeoutKey].(int)
	nodepoolTimeout := time.Duration(TimeoutDefaultValue) * time.Minute

	if timeoutMinutes > 0 {
		nodepoolTimeout = time.Duration(timeoutMinutes) * time.Minute
	}

	// Create New Node Pools & Update Existing Node Pools.
	for _, newNodePool := range newTFValue.([]interface{}) {
		newNodePoolMap := newNodePool.(map[string]interface{})
		newNodePoolName := newNodePoolMap[NameKey].(string)
		isNodePoolNew := true

		for _, oldNodePool := range oldTFValue.([]interface{}) {
			oldNodePoolMap := oldNodePool.(map[string]interface{})
			oldNodePoolName := oldNodePoolMap[NameKey].(string)

			if newNodePoolName == oldNodePoolName {
				isNodePoolNew = false
				existingNodePools[oldNodePoolName] = map[string]interface{}{
					"OldNodePoolMap": oldNodePoolMap,
					"NewNodePoolMap": newNodePoolMap,
				}

				break
			}
		}

		for _, np := range modelNodePools {
			if np.FullName.Name == newNodePoolName {
				np.FullName.ManagementClusterName = clusterFn.ManagementClusterName
				np.FullName.ProvisionerName = clusterFn.ProvisionerName
				np.FullName.TanzuKubernetesClusterName = clusterFn.Name
				nodePoolRequest := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData{
					Nodepool: np,
				}

				if isNodePoolNew {
					err := createNodePoolWithRetry(ctx, &config, nodePoolRequest, nodepoolTimeout)
					if err != nil {
						return diag.FromErr(errors.Wrapf(err, "Couldn't create TKG Cluster Nodepool.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s, Node Pool Name: %s",
							np.FullName.ManagementClusterName, np.FullName.ProvisionerName, np.FullName.TanzuKubernetesClusterName, np.FullName.Name))
					}
				} else {
					existingNodePoolMap := existingNodePools[newNodePoolName].(map[string]interface{})
					oldNodePoolMap := existingNodePoolMap["OldNodePoolMap"]
					newNodePoolMap := existingNodePoolMap["NewNodePoolMap"]

					if nodePoolHasChanged(oldNodePoolMap, newNodePoolMap) {
						err := updateNodePoolWithRetry(ctx, &config, nodePoolRequest, nodepoolTimeout)
						if err != nil {
							return diag.FromErr(errors.Wrapf(err, "Couldn't update TKG Cluster Nodepool.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s, Node Pool Name: %s",
								np.FullName.ManagementClusterName, np.FullName.ProvisionerName, np.FullName.TanzuKubernetesClusterName, np.FullName.Name))
						}
					}
				}

				break
			}
		}
	}

	// Delete Old Node Pools.
	for _, oldNodePool := range oldTFValue.([]interface{}) {
		oldNodePoolMap := oldNodePool.(map[string]interface{})
		oldNodePoolName := oldNodePoolMap[NameKey].(string)
		_, stillExists := existingNodePools[oldNodePoolName]

		if !stillExists {
			nodePoolFn := &tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName{
				ManagementClusterName:      clusterFn.ManagementClusterName,
				ProvisionerName:            clusterFn.ProvisionerName,
				TanzuKubernetesClusterName: clusterFn.Name,
				Name:                       oldNodePoolName,
			}

			err := deleteNodePoolWithRetry(ctx, &config, nodePoolFn, nodepoolTimeout)
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Couldn't delete TKG Cluster Nodepool.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s, Node Pool Name: %s",
					nodePoolFn.ManagementClusterName, nodePoolFn.ProvisionerName, nodePoolFn.TanzuKubernetesClusterName, nodePoolFn.Name))
			}
		}
	}

	return diags
}

func resourceTanzuKubernetesClusterImporter(ctx context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	config, ok := m.(authctx.TanzuContext)

	if !ok {
		return nil, errors.New("error while retrieving Tanzu auth config")
	}

	clusterFullName := data.Id()
	clusterFullNameParts := strings.Split(clusterFullName, "/")

	if len(clusterFullNameParts) != 3 {
		return nil, errors.New("Cluster ID must be comprised of management_cluster_name, provisioner_name and cluster_name - separated by /")
	}

	clusterFn := &tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName{
		ManagementClusterName: clusterFullNameParts[0],
		ProvisionerName:       clusterFullNameParts[1],
		Name:                  clusterFullNameParts[2],
	}

	clusterResp, err := readFullClusterResource(&config, clusterFn)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't import TKG cluster.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name)
	} else if clusterResp == nil || clusterResp.TanzuKubernetesCluster == nil {
		return nil, errors.Errorf("Couldn't import TKG cluster.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name)
	}

	clusterClassFn := &clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassFullName{
		ManagementClusterName: clusterResp.TanzuKubernetesCluster.FullName.ManagementClusterName,
		ProvisionerName:       clusterResp.TanzuKubernetesCluster.FullName.ProvisionerName,
		Name:                  clusterResp.TanzuKubernetesCluster.Spec.Topology.ClusterClass,
	}

	clusterClassResp, err := config.TMCConnection.ClusterClassResourceService.ClusterClassResourceServiceGet(clusterClassFn)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't find cluster class.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Class Name: %s.",
			clusterClassFn.ManagementClusterName, clusterClassFn.ProvisionerName, clusterClassFn.Name)
	} else if len(clusterClassResp.ClusterClasses) == 0 {
		return nil, errors.Errorf("Couldn't find cluster class.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Class Name: %s.",
			clusterClassFn.ManagementClusterName, clusterClassFn.ProvisionerName, clusterClassFn.Name)
	}

	clusterClassVariablesMap := clusterclass.BuildClusterClassMap(clusterClassResp.ClusterClasses[0].Spec)

	modifyClusterClassVariables(clusterClassVariablesMap, clusterResp.TanzuKubernetesCluster, removeModelVariable)

	err = tfModelResourceConverter.FillTFSchema(clusterResp.TanzuKubernetesCluster, data)
	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't import TKG cluster.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name)
	}

	return []*schema.ResourceData{data}, nil
}

func validateSchema(_ context.Context, data *schema.ResourceDiff, value interface{}) error {
	config, ok := value.(authctx.TanzuContext)
	if !ok || config.TMCConnection == nil || config.TMCConnection.ClusterClassResourceService == nil {
		return nil
	}

	specRaw := data.Get(SpecKey)
	if specRaw == nil {
		return nil
	}

	specList, ok := specRaw.([]interface{})
	if !ok || len(specList) == 0 || specList[0] == nil {
		return nil
	}

	specData, ok := specList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	topologyRaw, ok := specData[TopologyKey]
	if !ok || topologyRaw == nil {
		return nil
	}

	topologyList, ok := topologyRaw.([]interface{})
	if !ok || len(topologyList) == 0 || topologyList[0] == nil {
		return nil
	}

	topologyData, ok := topologyList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	clusterClass, clusterClassExist := topologyData[ClusterClassKey].(string)
	if !clusterClassExist || clusterClass == "" {
		return errors.New("Cluster Class must be set during creation.")
	}

	mgmtClusterName, _ := data.Get(ManagementClusterNameKey).(string)
	provisionerName, _ := data.Get(ProvisionerNameKey).(string)

	clusterClassFn := &clusterclassmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerClusterClassFullName{
		ManagementClusterName: mgmtClusterName,
		ProvisionerName:       provisionerName,
		Name:                  clusterClass,
	}

	resp, err := config.TMCConnection.ClusterClassResourceService.ClusterClassResourceServiceGet(clusterClassFn)
	if err != nil {
		return errors.Wrapf(err, "Couldn't find cluster class.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Class Name: %s.",
			clusterClassFn.ManagementClusterName, clusterClassFn.ProvisionerName, clusterClassFn.Name)
	} else if len(resp.ClusterClasses) == 0 || resp.ClusterClasses[0] == nil {
		return errors.Errorf("Couldn't find cluster class.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Class Name: %s.",
			clusterClassFn.ManagementClusterName, clusterClassFn.ProvisionerName, clusterClassFn.Name)
	}

	clusterClassValidator := NewClusterClassValidator(resp.ClusterClasses[0].Spec)

	if clusterVarsStr, ok := topologyData[ClusterVariablesKey].(string); ok && clusterVarsStr != "" {
		clusterVariablesErrs := clusterClassValidator.ValidateClusterVariables(clusterVarsStr, true)
		if len(clusterVariablesErrs) > 0 {
			errStr := "Cluster Variables validation failed:\n"

			for _, e := range clusterVariablesErrs {
				errStr = fmt.Sprintf("%s%s\n", errStr, e.Error())
			}

			err = errors.New(errStr)
		}
	}

	if nodePoolsList, ok := topologyData[NodePoolKey].([]interface{}); ok && len(nodePoolsList) > 0 {
		nodePoolsErrs := clusterClassValidator.ValidateNodePools(nodePoolsList)
		if len(nodePoolsErrs) > 0 {
			errStr := "Node pools validation failed:\n"

			if err != nil {
				errStr = fmt.Sprintf("%s%s", err.Error(), errStr)
			}

			for _, e := range nodePoolsErrs {
				errStr = fmt.Sprintf("%s%s\n", errStr, e.Error())
			}

			err = errors.New(errStr)
		}
	}

	return err
}
