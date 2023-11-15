/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetescluster

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	openapiv3 "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/openapi_v3_schema_validator"
	legacyclustermodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	kubeconfigmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/kubeconfig"
	tanzukubernetesclustermodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster"
	tkccommonmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/common"
	tkcnodepoolmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/nodepool"
)

type ClusterClassModifierFunc func(tfVariable interface{}, modelVariable interface{}) interface{}

// readResourceWait helps read operations where wait is needed for the Tanzu Kubernetes Cluster and its assets to be in a stop status.
// This function determines whether a timeout is needed and whether to fail the request if a deadline has exceeded.
func readResourceWait(ctx context.Context, config *authctx.TanzuContext, clusterFn *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName, existingNodePools []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool, timeoutPolicy map[string]interface{}) (resp *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, err error) {
	timeout := timeoutPolicy[TimeoutKey].(int)
	waitForKubeConfig := timeoutPolicy[WaitForKubeConfigKey].(bool)
	failOnTimeOut := timeoutPolicy[FailOnTimeOutKey].(bool)

	if timeout > 0 {
		var cancel context.CancelFunc

		ctx, cancel = context.WithTimeout(ctx, time.Duration(timeout)*time.Minute)

		defer cancel()
	}

	resp, err = readFullClusterResourceWait(ctx, config, clusterFn, existingNodePools, waitForKubeConfig)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) && !failOnTimeOut {
			return resp, nil
		}

		resp = nil
	}

	return resp, err
}

// readFullClusterResourceWait orchestrates the entire operation of waiting for a full Tanzu Kubernetes Cluster model to be in a stop status.
func readFullClusterResourceWait(ctx context.Context, config *authctx.TanzuContext, clusterFn *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName, existingNodePools []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool, waitForKubeConfig bool) (resp *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, err error) {
	// We call this first in case we receive a timeout while waiting for the cluster to be ready, therefore in order
	// to return the last results in case of a timeout, the calling function will have to get some cluster response.
	resp, err = readFullClusterResource(config, clusterFn)

	if err != nil {
		return resp, err
	}

	err = waitClusterReady(ctx, config, clusterFn)

	if err != nil {
		return resp, err
	}

	// Calling again to get updated status
	resp, _ = readFullClusterResource(config, clusterFn)

	var nodePoolsToCheck []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool

	for _, existingNp := range existingNodePools {
		for _, clusterNp := range resp.TanzuKubernetesCluster.Spec.Topology.NodePools {
			if existingNp.FullName.Name == clusterNp.FullName.Name {
				nodePoolsToCheck = append(nodePoolsToCheck, clusterNp)
				break
			}
		}
	}

	resp.TanzuKubernetesCluster.Spec.Topology.NodePools = nodePoolsToCheck
	err = waitNodePoolsReady(ctx, config, clusterFn, nodePoolsToCheck)

	if err != nil {
		return resp, err
	}

	if waitForKubeConfig && resp.TanzuKubernetesCluster.Spec.KubeConfig == "" {
		resp.TanzuKubernetesCluster.Spec.KubeConfig, err = waitKubeConfigReady(ctx, config, clusterFn)

		if err != nil {
			return resp, err
		}
	}

	return resp, err
}

// waitClusterReady waits for a cluster to be in a stop state using the legacy cluster API.
func waitClusterReady(ctx context.Context, config *authctx.TanzuContext, clusterFn *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (err error) {
	clusterStatus := legacyclustermodels.VmwareTanzuManageV1alpha1ClusterPhasePHASEUNSPECIFIED
	clusterHealth := legacyclustermodels.VmwareTanzuManageV1alpha1CommonClusterHealthHEALTHUNSPECIFIED
	isStopStatus := false
	legacyClusterFn := &legacyclustermodels.VmwareTanzuManageV1alpha1ClusterFullName{
		ManagementClusterName: clusterFn.ManagementClusterName,
		ProvisionerName:       clusterFn.ProvisionerName,
		Name:                  clusterFn.Name,
	}

	for !isStopStatus {
		time.Sleep(5 * time.Second)

		err := ctx.Err()

		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				err = errors.Wrapf(err, "Timeout exceeded while waiting for the cluster to be ready. Cluster Status: %s, Cluster Health: %s", clusterStatus, clusterHealth)
			}

			return err
		}

		legacyClusterResp, err := config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceGet(legacyClusterFn)

		if clienterrors.IsUnauthorizedError(err) && clusterStatus != legacyclustermodels.VmwareTanzuManageV1alpha1ClusterPhasePHASEUNSPECIFIED {
			authctx.RefreshUserAuthContext(config, clienterrors.IsUnauthorizedError, err)
		} else {
			if err != nil {
				return err
			}

			cluster := legacyClusterResp.Cluster

			if cluster.Status.Phase != nil {
				clusterStatus = *cluster.Status.Phase

				if cluster.Status.Health != nil {
					clusterHealth = *cluster.Status.Health
				}

				isStopStatus = isClusterStopStatus(&clusterStatus, &clusterHealth)
			}
		}
	}

	if clusterStatus == legacyclustermodels.VmwareTanzuManageV1alpha1ClusterPhaseERROR {
		err = errors.Errorf("TKG Cluster errored.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name)

		return err
	} else if clusterStatus == legacyclustermodels.VmwareTanzuManageV1alpha1ClusterPhaseUPGRADEFAILED {
		err = errors.Errorf("TKG Cluster upgrade failed,\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name)

		return err
	}

	return err
}

// waitNodePoolsReady waits for Node Pools to be in a stop state.
func waitNodePoolsReady(ctx context.Context, config *authctx.TanzuContext, clusterFn *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName, nodePoolsToCheck []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool) (err error) {
	nodePoolsReady := false

	for !nodePoolsReady {
		nodePoolsReady = true

		for _, np := range nodePoolsToCheck {
			nodePoolStatus := *np.Status.Phase

			switch nodePoolStatus {
			case tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseREADY:
			case tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseERROR:
				err = errors.Errorf("TKG Cluster node pool errored.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s, Node Pool Name: %s",
					clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name, np.FullName.Name)

				return err
			case tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseUPGRADEFAILED:
				err = errors.Errorf("TKG Cluster node pool upgrade failed.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s, Node Pool Name: %s",
					clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name, np.FullName.Name)

				return err
			default:
				nodePoolsReady = false
			}
		}

		if !nodePoolsReady {
			time.Sleep(5 * time.Second)

			err := ctx.Err()

			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					errMsg := "Timeout exceeded while waiting for the cluster node pools to be ready."

					for _, np := range nodePoolsToCheck {
						npStatusMsg := fmt.Sprintf("Node pool '%s' is in status %s", np.FullName.Name, *np.Status.Phase)
						errMsg = fmt.Sprintf("%s\n%s", errMsg, npStatusMsg)
					}

					err = errors.Wrapf(err, errMsg)
				}

				return err
			}

			nodePoolsResp, err := config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterNodePoolResourceServiceList(clusterFn)

			if clienterrors.IsUnauthorizedError(err) {
				authctx.RefreshUserAuthContext(config, clienterrors.IsUnauthorizedError, err)
			} else {
				if err != nil {
					return err
				}

				for i, np := range nodePoolsToCheck {
					for _, respNp := range nodePoolsResp.Nodepools {
						if np.FullName.Name == respNp.FullName.Name {
							nodePoolsToCheck[i] = respNp

							break
						}
					}
				}
			}
		}
	}

	return err
}

// waitKubeConfigReady waits for the KubeConfig to be in a stop state and returns it.
func waitKubeConfigReady(ctx context.Context, config *authctx.TanzuContext, clusterFn *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (kubeConfig string, err error) {
	kubeConfigStatus := kubeconfigmodels.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusSTATUSUNSPECIFIED
	stopStatuses := map[kubeconfigmodels.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus]bool{
		kubeconfigmodels.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusUNAVAILABLE: true,
		kubeconfigmodels.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusREADY:       true,
	}
	isStopStatus := false

	for !isStopStatus {
		time.Sleep(5 * time.Second)

		err := ctx.Err()

		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				err = errors.Wrapf(err, "Timeout exceeded while waiting for the cluster kubeconfig to be ready. KubeConfig Status: %s", kubeConfigStatus)
			}

			return kubeConfig, err
		}

		kubeConfigResp, err := config.TMCConnection.TanzuKubernetesClusterResourceService.KubeConfigResourceServiceGet(clusterFn)

		if clienterrors.IsUnauthorizedError(err) && kubeConfigStatus != kubeconfigmodels.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusSTATUSUNSPECIFIED {
			authctx.RefreshUserAuthContext(config, clienterrors.IsUnauthorizedError, err)
		} else {
			if err != nil {
				return "", err
			}

			kubeConfig = kubeConfigResp.Kubeconfig
			kubeConfigStatus = *kubeConfigResp.Status
			_, isStopStatus = stopStatuses[kubeConfigStatus]
		}
	}

	if kubeConfigStatus == kubeconfigmodels.VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatusUNAVAILABLE {
		err = errors.Errorf("TKG Cluster's kubeconfig is unavailable.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name)

		return "", err
	}

	return kubeConfig, err
}

// readFullClusterResource returns a full Tanzu Kubernetes Cluster model populated with Node Pools and KubeConfig.
func readFullClusterResource(config *authctx.TanzuContext, clusterFn *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName) (resp *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData, err error) {
	resp, err = config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterResourceServiceGet(clusterFn)

	if err != nil || resp.TanzuKubernetesCluster == nil {
		return nil, err
	}

	nodePoolsResp, err := config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterNodePoolResourceServiceList(clusterFn)

	if err != nil {
		return nil, err
	}

	resp.TanzuKubernetesCluster.Spec.Topology.NodePools = nodePoolsResp.Nodepools

	kubeConfigResp, err := config.TMCConnection.TanzuKubernetesClusterResourceService.KubeConfigResourceServiceGet(clusterFn)

	if err != nil {
		return nil, err
	}

	resp.TanzuKubernetesCluster.Spec.KubeConfig = kubeConfigResp.Kubeconfig

	return resp, err
}

// isClusterStopStatus checks whether the cluster has reached a status where waiting for it should be stopped.
func isClusterStopStatus(clusterPhase *legacyclustermodels.VmwareTanzuManageV1alpha1ClusterPhase, clusterHealth *legacyclustermodels.VmwareTanzuManageV1alpha1CommonClusterHealth) bool {
	return (*clusterPhase == legacyclustermodels.VmwareTanzuManageV1alpha1ClusterPhaseERROR) ||
		(*clusterPhase == legacyclustermodels.VmwareTanzuManageV1alpha1ClusterPhaseUPGRADEFAILED) ||
		(*clusterPhase == legacyclustermodels.VmwareTanzuManageV1alpha1ClusterPhaseREADY && *clusterHealth == legacyclustermodels.VmwareTanzuManageV1alpha1CommonClusterHealthHEALTHY)
}

// nodePoolHasChanged recursively comparing old value & new value of a node pool to determine whether it was changed.
func nodePoolHasChanged(oldValue interface{}, newValue interface{}) bool {
	switch oldValue := oldValue.(type) {
	case map[string]interface{}:
		newValueMap := newValue.(map[string]interface{})
		allKeys := helper.GetAllMapsKeys(oldValue, newValueMap)

		for k := range allKeys {
			subOldValue, subOldValueExists := oldValue[k]
			subNewValue, subNewValueExists := newValueMap[k]

			if subOldValueExists != subNewValueExists {
				return true
			} else if subOldValueExists && subNewValueExists {
				if (k == OverridesKey && !isVariablesValuesEqual(k, subOldValue.(string), subNewValue.(string), nil)) ||
					(k != OverridesKey && nodePoolHasChanged(subOldValue, subNewValue)) {
					return true
				}
			}
		}
	case []interface{}:
		newValueArr := newValue.([]interface{})

		if len(oldValue) != len(newValueArr) {
			return true
		} else {
			for i := 0; i < len(oldValue); i++ {
				if nodePoolHasChanged(oldValue[i], newValueArr[i]) {
					return true
				}
			}
		}
	default:
		if oldValue != newValue {
			return true
		}
	}

	return false
}

// getTimeoutPolicy returns a timeout policy based on the default values and Terraform config values provided.
func getTimeoutPolicy(data *schema.ResourceData) map[string]interface{} {
	timeoutPolicy := map[string]interface{}{
		TimeoutKey:           TimeoutDefaultValue,
		WaitForKubeConfigKey: WaitForKubeConfigDefaultValue,
		FailOnTimeOutKey:     FailOnTimeOutDefaultValue,
	}

	tfTimeoutPolicy := data.Get(TimeoutPolicyKey).([]interface{})

	if len(tfTimeoutPolicy) > 0 {
		tfTimeoutPolicyMap := tfTimeoutPolicy[0].(map[string]interface{})

		for k, v := range tfTimeoutPolicyMap {
			timeoutPolicy[k] = v
		}
	}

	return timeoutPolicy
}

// removeUnspecifiedClusterVariables removed cluster variables returning in the API which do not exist in the Cluster Class schema.
func removeUnspecifiedClusterVariables(tfClusterVariables string, kubernetesClusterModel *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster) {
	tfClusterVariablesJSON := make(map[string]interface{})
	_ = json.Unmarshal([]byte(tfClusterVariables), &tfClusterVariablesJSON)
	modifyClusterClassVariables(tfClusterVariablesJSON, kubernetesClusterModel, modifyModelVariable)
}

// removeUnspecifiedNodePoolsOverrides removed node pools overrides returning in the API which do not exist in the Cluster Class schema.
func removeUnspecifiedNodePoolsOverrides(nodePools []interface{}, kubernetesClusterModel *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster) {
	for i, np := range nodePools {
		tfNodePoolSpec := np.(map[string]interface{})[SpecKey].([]interface{})[0].(map[string]interface{})
		tfNodePoolOverrides := tfNodePoolSpec[OverridesKey]

		if tfNodePoolOverrides != nil && tfNodePoolOverrides.(string) != "" {
			tfOverridesVariablesJSON := make(map[string]interface{})
			_ = json.Unmarshal([]byte(tfNodePoolOverrides.(string)), &tfOverridesVariablesJSON)

			overridesToKeep := make([]*tkccommonmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterClusterVariable, 0)

			for _, modelOverride := range kubernetesClusterModel.Spec.Topology.NodePools[i].Spec.Overrides {
				varKey := modelOverride.Name

				if tfOverridesVariableValue, exist := tfOverridesVariablesJSON[varKey]; exist {
					// This is necessary because some inner values have defaults and are being returned even when not filled
					modifiedModelVariableValue := modifyModelVariable(tfOverridesVariableValue, modelOverride.Value)
					modelOverride.Value = modifiedModelVariableValue
					overridesToKeep = append(overridesToKeep, modelOverride)
				}
			}

			kubernetesClusterModel.Spec.Topology.NodePools[i].Spec.Overrides = overridesToKeep
		}
	}
}

// modifyModelVariable helps when certain variables do no return from the API or in a case where some values are returning
// from the API with defaults but were not set in the Terraform config.
func modifyModelVariable(tfVariable interface{}, modelVariable interface{}) interface{} {
	switch modelVariable := modelVariable.(type) {
	case map[string]interface{}:
		modifiedMap := make(map[string]interface{})

		for k, v := range modelVariable {
			tfVariableValue, tfKeyExists := tfVariable.(map[string]interface{})[k]

			if tfKeyExists {
				modifiedMap[k] = modifyModelVariable(tfVariableValue, v)
			}
		}

		return modifiedMap
	default:
		if !helper.IsEmptyInterface(tfVariable) && helper.IsEmptyInterface(modelVariable) {
			return tfVariable
		}
	}

	return modelVariable
}

// removeModelVariable helps when importing a Tanzu Kubernetes Cluster and the API returns cluster variables which are
// not present in the Cluster Class schema in a recursive manner.
// This is useful when a certain root value exists in the Cluster Class schema but having fields that do not exist.
func removeModelVariable(clusterClassSchema interface{}, modelVariable interface{}) interface{} {
	modelVariableMap, isModelVarMap := modelVariable.(map[string]interface{})

	if isModelVarMap {
		modelVarProperties, propertiesExist := clusterClassSchema.(map[string]interface{})[string(openapiv3.PropertiesKey)]

		if propertiesExist {
			modelVariable = make(map[string]interface{})

			for k, v := range modelVarProperties.(map[string]interface{}) {
				property, propertyExists := modelVariableMap[k]

				if propertyExists {
					modelVariable.(map[string]interface{})[k] = removeModelVariable(v.(map[string]interface{}), property)
				}
			}
		} else {
			modelVarAdditionalProperties := clusterClassSchema.(map[string]interface{})[string(openapiv3.AdditionalPropertiesKey)]
			_, propertiesExist := modelVarAdditionalProperties.(map[string]interface{})[string(openapiv3.PropertiesKey)]

			if propertiesExist {
				modelVariable = make(map[string]interface{})

				for k, v := range modelVariableMap {
					modelVariable.(map[string]interface{})[k] = removeModelVariable(modelVarAdditionalProperties, v)
				}
			}
		}
	}

	return modelVariable
}

// modifyClusterClassVariables helps delete or modify the values of a variable like array using a modifier function for the
// values that exist.
func modifyClusterClassVariables(clusterVariablesMap map[string]interface{}, kubernetesClusterModel *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster, modifierFunc ClusterClassModifierFunc) {
	variablesToKeep := make([]*tkccommonmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterClusterVariable, 0)

	for _, modelVariable := range kubernetesClusterModel.Spec.Topology.Variables {
		varKey := modelVariable.Name

		if tfClusterVariableValue, exist := clusterVariablesMap[varKey]; exist {
			// This is necessary because some inner values have defaults and are being returned even when not filled
			modifiedModelVariableValue := modifierFunc(tfClusterVariableValue, modelVariable.Value)
			modelVariable.Value = modifiedModelVariableValue
			variablesToKeep = append(variablesToKeep, modelVariable)
		}
	}

	kubernetesClusterModel.Spec.Topology.Variables = variablesToKeep
}

// suppressNodePoolsOrderChanges helps suppress node pools order difference when receiving the node pools from the API.
func suppressNodePoolsOrderChanges(oldNodePoolsArray []interface{}, data *schema.ResourceData) {
	specData := data.Get(SpecKey).([]interface{})[0].(map[string]interface{})
	topologyData := specData[TopologyKey].([]interface{})[0].(map[string]interface{})
	newNodePoolsArray := topologyData[NodePoolKey].([]interface{})
	orderedNodePoolsArray := make([]interface{}, 0, len(newNodePoolsArray))

	for _, oldNp := range oldNodePoolsArray {
		for _, newNp := range newNodePoolsArray {
			if oldNp.(map[string]interface{})[NameKey] == newNp.(map[string]interface{})[NameKey] {
				orderedNodePoolsArray = append(orderedNodePoolsArray, newNp)

				break
			}
		}
	}

	topologyData[NodePoolKey] = orderedNodePoolsArray

	_ = data.Set(SpecKey, []interface{}{specData})
}
