// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzukubernetescluster

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
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

	// waitNodePoolsReady actively polls and verifies that all expected node pools exist and reach READY phase.
	readyNodePools, err := waitNodePoolsReady(ctx, config, clusterFn, existingNodePools)
	if err != nil {
		return resp, err
	}

	if resp != nil && resp.TanzuKubernetesCluster != nil && resp.TanzuKubernetesCluster.Spec != nil && resp.TanzuKubernetesCluster.Spec.Topology != nil {
		resp.TanzuKubernetesCluster.Spec.Topology.NodePools = readyNodePools
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

var nodePoolPollInterval = 5 * time.Second

func findNodePoolByName(nodepools []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool, name string) *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool {
	for _, np := range nodepools {
		if np != nil && np.FullName != nil && np.FullName.Name == name {
			return np
		}
	}

	return nil
}

func checkExpectedNodePoolStatus(
	clusterFn *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName,
	expectedName string,
	foundNp *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool,
	npStatusMap map[string]string,
) (*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool, error) {
	if foundNp == nil {
		npStatusMap[expectedName] = "not found in cluster"

		return nil, nil
	}

	if foundNp.Status == nil || foundNp.Status.Phase == nil {
		npStatusMap[expectedName] = "initializing"

		return nil, nil
	}

	nodePoolStatus := *foundNp.Status.Phase
	npStatusMap[expectedName] = string(nodePoolStatus)

	switch nodePoolStatus {
	case tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseREADY:
		return foundNp, nil
	case tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseERROR:
		return nil, errors.Errorf("TKG Cluster node pool errored.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s, Node Pool Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name, expectedName)
	case tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatusPhaseUPGRADEFAILED:
		return nil, errors.Errorf("TKG Cluster node pool upgrade failed.\nManagement Cluster Name: %s, Provisioner: %s, Cluster Name: %s, Node Pool Name: %s",
			clusterFn.ManagementClusterName, clusterFn.ProvisionerName, clusterFn.Name, expectedName)
	default:
		return nil, nil
	}
}

func evaluateExpectedNodePools(
	clusterFn *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName,
	expectedNodePools []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool,
	remotePools []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool,
	npStatusMap map[string]string,
) ([]*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool, bool, error) {
	var matchedPools []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool

	allReady := true

	for _, expectedNp := range expectedNodePools {
		expectedName := ""
		if expectedNp.FullName != nil {
			expectedName = expectedNp.FullName.Name
		}

		if expectedName == "" {
			continue
		}

		foundNp := findNodePoolByName(remotePools, expectedName)

		readyNp, err := checkExpectedNodePoolStatus(clusterFn, expectedName, foundNp, npStatusMap)
		if err != nil {
			return nil, false, err
		}

		if readyNp == nil {
			allReady = false

			continue
		}

		matchedPools = append(matchedPools, readyNp)
	}

	return matchedPools, allReady && len(matchedPools) == len(expectedNodePools), nil
}

func buildNodePoolTimeoutError(
	ctxErr error,
	expectedNodePools []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool,
	npStatusMap map[string]string,
) error {
	if !errors.Is(ctxErr, context.DeadlineExceeded) {
		return ctxErr
	}

	errMsg := "Timeout exceeded while waiting for the cluster node pools to be ready."

	for _, expectedNp := range expectedNodePools {
		expectedName := ""
		if expectedNp.FullName != nil {
			expectedName = expectedNp.FullName.Name
		}

		status, exists := npStatusMap[expectedName]
		if !exists {
			status = "not found in cluster"
		}

		errMsg = fmt.Sprintf("%s\nNode pool %q is in status: %s", errMsg, expectedName, status)
	}

	return errors.Wrap(ctxErr, errMsg)
}

// waitNodePoolsReady actively polls the TMC node pool API until all expected node pools exist in the cluster
// and reach the READY phase. If any node pool enters an ERROR or UPGRADEFAILED phase, the function immediately
// returns a descriptive error. If the context deadline is exceeded before all node pools become ready, it constructs
// a detailed error report displaying the last observed status of every expected node pool.
func waitNodePoolsReady(
	ctx context.Context,
	config *authctx.TanzuContext,
	clusterFn *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterFullName,
	expectedNodePools []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool,
) ([]*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool, error) {
	if len(expectedNodePools) == 0 {
		return nil, nil
	}

	npStatusMap := make(map[string]string)

	for {
		nodePoolsResp, err := config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterNodePoolResourceServiceList(clusterFn)

		switch {
		case clienterrors.IsUnauthorizedError(err):
			authctx.RefreshUserAuthContext(config, clienterrors.IsUnauthorizedError, err)
		case err != nil:
			log.Printf("[DEBUG] Failed to list node pools for cluster %q: %v", clusterFn.Name, err)
		default:
			var remotePools []*tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool
			if nodePoolsResp != nil {
				remotePools = nodePoolsResp.Nodepools
			}

			matchedPools, allReady, evalErr := evaluateExpectedNodePools(clusterFn, expectedNodePools, remotePools, npStatusMap)
			if evalErr != nil {
				return nil, evalErr
			}

			if allReady {
				return matchedPools, nil
			}
		}

		select {
		case <-ctx.Done():
			return nil, buildNodePoolTimeoutError(ctx.Err(), expectedNodePools, npStatusMap)
		case <-time.After(nodePoolPollInterval):
		}
	}
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
	if kubernetesClusterModel == nil || kubernetesClusterModel.Spec == nil || kubernetesClusterModel.Spec.Topology == nil {
		return
	}

	for _, np := range nodePools {
		npMap, ok := np.(map[string]interface{})
		if !ok {
			continue
		}

		name, _ := npMap[NameKey].(string)
		if name == "" {
			continue
		}

		specList, ok := npMap[SpecKey].([]interface{})
		if !ok || len(specList) == 0 {
			continue
		}

		tfNodePoolSpec, ok := specList[0].(map[string]interface{})
		if !ok {
			continue
		}

		tfNodePoolOverrides, ok := tfNodePoolSpec[OverridesKey].(string)
		if !ok || tfNodePoolOverrides == "" {
			continue
		}

		targetNodePool := findNodePoolByName(kubernetesClusterModel.Spec.Topology.NodePools, name)
		if targetNodePool != nil && targetNodePool.Spec != nil {
			overridesToKeep := excludeUnspecifiedOverrides(tfNodePoolOverrides, targetNodePool.Spec.Overrides)
			targetNodePool.Spec.Overrides = overridesToKeep
		}
	}
}

// removeUnspecifiedControlPlaneOverrides removed node pools overrides returning in the API which do not exist in the Cluster Class schema.
func removeUnspecifiedControlPlaneOverrides(controlPlane interface{}, kubernetesClusterModel *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterTanzuKubernetesCluster) {
	if kubernetesClusterModel == nil || kubernetesClusterModel.Spec == nil || kubernetesClusterModel.Spec.Topology == nil || kubernetesClusterModel.Spec.Topology.ControlPlane == nil {
		return
	}

	controlPlaneMap, ok := controlPlane.(map[string]interface{})
	if !ok {
		return
	}

	tfControlPlaneOverrides, ok := controlPlaneMap[OverridesKey].(string)
	if ok && tfControlPlaneOverrides != "" {
		overridesToKeep := excludeUnspecifiedOverrides(tfControlPlaneOverrides, kubernetesClusterModel.Spec.Topology.ControlPlane.Overrides)
		kubernetesClusterModel.Spec.Topology.ControlPlane.Overrides = overridesToKeep
	}
}

// excludeUnspecifiedOverrides removed overrides returning in the API which do not exist in the Cluster Class schema.
func excludeUnspecifiedOverrides(overridesResource string, overrides []*tkccommonmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterClusterVariable) []*tkccommonmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterClusterVariable {
	tfOverridesVariablesJSON := make(map[string]interface{})
	_ = json.Unmarshal([]byte(overridesResource), &tfOverridesVariablesJSON)

	overridesToKeep := make([]*tkccommonmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterClusterVariable, 0)

	for _, modelOverride := range overrides {
		varKey := modelOverride.Name

		if tfOverridesVariableValue, exist := tfOverridesVariablesJSON[varKey]; exist {
			// This is necessary because some inner values have defaults and are being returned even when not filled
			modifiedModelVariableValue := modifyModelVariable(tfOverridesVariableValue, modelOverride.Value)
			modelOverride.Value = modifiedModelVariableValue
			overridesToKeep = append(overridesToKeep, modelOverride)
		}
	}

	return overridesToKeep
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
			modelVarAdditionalProperties, additionalPropertiesExist := clusterClassSchema.(map[string]interface{})[string(openapiv3.AdditionalPropertiesKey)]
			preserve, preserveExists := clusterClassSchema.(map[string]interface{})[string(openapiv3.PreserveUnknownFieldKey)]

			if additionalPropertiesExist {
				_, propertiesExist := modelVarAdditionalProperties.(map[string]interface{})[string(openapiv3.PropertiesKey)]

				if propertiesExist {
					modelVariable = make(map[string]interface{})

					for k, v := range modelVariableMap {
						modelVariable.(map[string]interface{})[k] = removeModelVariable(modelVarAdditionalProperties, v)
					}
				}
			} else if !preserveExists || !preserve.(bool) {
				modelVariable = make(map[string]interface{})
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
// It preserves the ordering of existing node pools matching oldNodePoolsArray, while appending any new
// node pools returned from the API that were not present in oldNodePoolsArray (preventing silent drops).
func suppressNodePoolsOrderChanges(oldNodePoolsArray []interface{}, data *schema.ResourceData) {
	specList, ok := data.Get(SpecKey).([]interface{})
	if !ok || len(specList) == 0 {
		return
	}

	specData, ok := specList[0].(map[string]interface{})
	if !ok {
		return
	}

	topologyList, ok := specData[TopologyKey].([]interface{})
	if !ok || len(topologyList) == 0 {
		return
	}

	topologyData, ok := topologyList[0].(map[string]interface{})
	if !ok {
		return
	}

	newNodePoolsArray, ok := topologyData[NodePoolKey].([]interface{})
	if !ok {
		return
	}

	newPoolsByName := make(map[string]interface{}, len(newNodePoolsArray))
	for _, np := range newNodePoolsArray {
		if npMap, ok := np.(map[string]interface{}); ok {
			if name, ok := npMap[NameKey].(string); ok && name != "" {
				newPoolsByName[name] = np
			}
		}
	}

	orderedNodePoolsArray := make([]interface{}, 0, len(newNodePoolsArray))
	matched := make(map[string]bool)

	// 1. Maintain old ordering for existing pools
	for _, oldNp := range oldNodePoolsArray {
		oldMap, ok := oldNp.(map[string]interface{})
		if !ok {
			continue
		}

		oldName, _ := oldMap[NameKey].(string)
		if newNp, exists := newPoolsByName[oldName]; exists {
			orderedNodePoolsArray = append(orderedNodePoolsArray, newNp)
			matched[oldName] = true
		}
	}

	// 2. Append newly added pools that did not exist in oldNodePoolsArray
	for _, newNp := range newNodePoolsArray {
		newMap, ok := newNp.(map[string]interface{})
		if !ok {
			orderedNodePoolsArray = append(orderedNodePoolsArray, newNp)
			continue
		}

		newName, _ := newMap[NameKey].(string)
		if !matched[newName] {
			orderedNodePoolsArray = append(orderedNodePoolsArray, newNp)
		}
	}

	topologyData[NodePoolKey] = orderedNodePoolsArray

	_ = data.Set(SpecKey, []interface{}{specData})
}

// isRetryableNodePoolError classifies whether an error returned from a node pool lifecycle operation
// (create, update, or delete) represents a transient condition that can be resolved via retries.
//
// Recognized retryable conditions include:
//  1. Webhook admission rejections: CAPI / Supervisor validating webhooks frequently reject early node pool
//     creation requests while the control plane is still initializing (e.g., "SystemChecksSucceeded condition is not True").
//  2. Initialization race conditions: Concurrent label synchronization on newly created clusters (e.g., duplicate "nodePoolLabels").
//  3. Optimistic locking / ResourceVersion conflicts: Concurrent updates in Kubernetes etcd ("the object has been modified").
//  4. Transient network / transport drops: Connection resets, broken pipes, unexpected EOFs, or gRPC Unavailable errors.
//  5. Rate limiting: HTTP 429 Too Many Requests or gRPC Code 8 (ResourceExhausted).
//  6. Gateway / Server-side 5xx errors: HTTP 500, 502, 503, or 504 transient outages.
func isRetryableNodePoolError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := strings.ToLower(err.Error())

	// 1. Webhook admission rejection due to pending SystemChecks or transient validation
	if strings.Contains(errMsg, "systemcheckssucceeded") ||
		strings.Contains(errMsg, "update cannot be initiated") ||
		strings.Contains(errMsg, "admission webhook") {
		return true
	}

	// 2. Race condition on cluster initialization or nodepool labels
	if strings.Contains(errMsg, "duplicate") ||
		strings.Contains(errMsg, "nodepoollabels") ||
		strings.Contains(errMsg, "being initialized") {
		return true
	}

	// 3. Optimistic locking / Kubernetes resource version conflicts
	if strings.Contains(errMsg, "the object has been modified") ||
		strings.Contains(errMsg, "conflict") ||
		strings.Contains(errMsg, "resourceversion") {
		return true
	}

	// 4. Transient network / RPC / gateway issues
	if strings.Contains(errMsg, "eof") ||
		strings.Contains(errMsg, "connection reset") ||
		strings.Contains(errMsg, "connection refused") ||
		strings.Contains(errMsg, "transport is closing") ||
		strings.Contains(errMsg, "unavailable") ||
		strings.Contains(errMsg, "broken pipe") ||
		strings.Contains(errMsg, "deadline exceeded") {
		return true
	}

	// 5. Rate limiting (HTTP 429 / gRPC Code 8 ResourceExhausted)
	if strings.Contains(errMsg, "status : 429") ||
		strings.Contains(errMsg, "too many requests") ||
		strings.Contains(errMsg, "rate limit") ||
		strings.Contains(errMsg, "rate limiting") ||
		strings.Contains(errMsg, "resourceexhausted") ||
		strings.Contains(errMsg, "\"code\":8") ||
		strings.Contains(errMsg, "code = resourceexhausted") {
		return true
	}

	// 6. Transient 5xx server errors
	if strings.Contains(errMsg, "status : 500") ||
		strings.Contains(errMsg, "status : 502") ||
		strings.Contains(errMsg, "status : 503") ||
		strings.Contains(errMsg, "status : 504") {
		return true
	}

	return false
}

var (
	nodePoolInitialBackoff = 5 * time.Second
	nodePoolMaxBackoff     = 60 * time.Second
	nodePoolBackoffFactor  = 1.5
	nodePoolJitterFactor   = 0.2 // +/-20% random jitter
)

// cryptoRandFloat64 returns a cryptographically secure random float64 in the range [0.0, 1.0).
func cryptoRandFloat64() float64 {
	n, err := rand.Int(rand.Reader, big.NewInt(1<<53))
	if err != nil {
		return 0.5
	}

	return float64(n.Int64()) / float64(1<<53)
}

// calculateBackoffWithJitter calculates the actual sleep duration for the current retry attempt (with randomized jitter)
// and computes the next base backoff interval using exponential progression.
//
// Randomized jitter (default +/-20%) prevents synchronized retry storms (the "thundering herd" problem)
// when multiple node pools or clusters are created concurrently across parallel Terraform runs.
func calculateBackoffWithJitter(currentBackoff, maxBackoff time.Duration, factor, jitterRatio float64) (time.Duration, time.Duration) {
	if currentBackoff <= 0 {
		currentBackoff = nodePoolInitialBackoff
	}

	if maxBackoff <= 0 {
		maxBackoff = nodePoolMaxBackoff
	}

	if factor <= 1.0 {
		factor = nodePoolBackoffFactor
	}

	if jitterRatio < 0 || jitterRatio > 1.0 {
		jitterRatio = nodePoolJitterFactor
	}

	// Apply jitter in range [1 - jitterRatio, 1 + jitterRatio]
	jitter := 1.0 + (cryptoRandFloat64()*2.0-1.0)*jitterRatio
	sleepDuration := time.Duration(float64(currentBackoff) * jitter)

	if sleepDuration > maxBackoff {
		sleepDuration = maxBackoff
	}

	if sleepDuration < 0 {
		sleepDuration = currentBackoff
	}

	nextBackoff := time.Duration(float64(currentBackoff) * factor)
	if nextBackoff > maxBackoff {
		nextBackoff = maxBackoff
	}

	return sleepDuration, nextBackoff
}

// createNodePoolWithRetry creates a node pool with polling retries and exponential backoff with jitter.
//
// Key behaviors:
//   - Respects the provided timeout context.
//   - Idempotency: If the API reports that the node pool "already exists" (e.g., due to a previous unacknowledged request),
//     it verifies existence via a GET call and treats it as a success.
//   - Retries transient errors (admission webhooks, rate limits, transport drops) until accepted or timeout.
func createNodePoolWithRetry(
	ctx context.Context,
	config *authctx.TanzuContext,
	nodePoolRequest *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData,
	timeout time.Duration,
) error {
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	currentBackoff := nodePoolInitialBackoff
	attempt := 0

	for {
		attempt++

		_, err := config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterNodePoolResourceServiceCreate(nodePoolRequest)
		if err == nil {
			return nil
		}

		// Idempotency: if already exists, verify whether nodepool is really present
		if clienterrors.IsAlreadyExistsError(err) || strings.Contains(strings.ToLower(err.Error()), "already exists") {
			_, getErr := config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterNodePoolResourceServiceGet(nodePoolRequest.Nodepool.FullName)
			if getErr == nil {
				return nil
			}
		}

		if !isRetryableNodePoolError(err) {
			return err
		}

		sleepDuration, next := calculateBackoffWithJitter(currentBackoff, nodePoolMaxBackoff, nodePoolBackoffFactor, nodePoolJitterFactor)
		currentBackoff = next

		log.Printf("[DEBUG] Retrying creation for node pool %q (attempt %d) in %v due to transient error: %s",
			nodePoolRequest.Nodepool.FullName.Name, attempt, sleepDuration, err.Error())

		select {
		case <-ctx.Done():
			return errors.Wrapf(ctx.Err(), "timed out waiting for nodepool %q creation to be accepted after %d attempts. Last error: %s",
				nodePoolRequest.Nodepool.FullName.Name, attempt, err.Error())
		case <-time.After(sleepDuration):
		}
	}
}

// updateNodePoolWithRetry updates an existing node pool with polling retries and exponential backoff with jitter,
// handling transient errors such as admission webhook validations or optimistic concurrency conflicts.
func updateNodePoolWithRetry(
	ctx context.Context,
	config *authctx.TanzuContext,
	nodePoolRequest *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolData,
	timeout time.Duration,
) error {
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	currentBackoff := nodePoolInitialBackoff
	attempt := 0

	for {
		attempt++

		_, err := config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterNodePoolResourceServiceUpdate(nodePoolRequest)
		if err == nil {
			return nil
		}

		if !isRetryableNodePoolError(err) {
			return err
		}

		sleepDuration, next := calculateBackoffWithJitter(currentBackoff, nodePoolMaxBackoff, nodePoolBackoffFactor, nodePoolJitterFactor)
		currentBackoff = next

		log.Printf("[DEBUG] Retrying update for node pool %q (attempt %d) in %v due to transient error: %s",
			nodePoolRequest.Nodepool.FullName.Name, attempt, sleepDuration, err.Error())

		select {
		case <-ctx.Done():
			return errors.Wrapf(ctx.Err(), "timed out waiting for nodepool %q update to be accepted after %d attempts. Last error: %s",
				nodePoolRequest.Nodepool.FullName.Name, attempt, err.Error())
		case <-time.After(sleepDuration):
		}
	}
}

// deleteNodePoolWithRetry deletes a node pool with polling retries and exponential backoff with jitter.
// Treats 404 Not Found as an immediate success (idempotent deletion).
func deleteNodePoolWithRetry(
	ctx context.Context,
	config *authctx.TanzuContext,
	nodePoolFn *tkcnodepoolmodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName,
	timeout time.Duration,
) error {
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	currentBackoff := nodePoolInitialBackoff
	attempt := 0

	for {
		attempt++

		err := config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterNodePoolResourceServiceDelete(nodePoolFn)
		if err == nil || clienterrors.IsNotFoundError(err) {
			return nil
		}

		if !isRetryableNodePoolError(err) {
			return err
		}

		sleepDuration, next := calculateBackoffWithJitter(currentBackoff, nodePoolMaxBackoff, nodePoolBackoffFactor, nodePoolJitterFactor)
		currentBackoff = next

		log.Printf("[DEBUG] Retrying deletion for node pool %q (attempt %d) in %v due to transient error: %s",
			nodePoolFn.Name, attempt, sleepDuration, err.Error())

		select {
		case <-ctx.Done():
			return errors.Wrapf(ctx.Err(), "timed out waiting for nodepool %q deletion to be accepted after %d attempts. Last error: %s",
				nodePoolFn.Name, attempt, err.Error())
		case <-time.After(sleepDuration):
		}
	}
}

// isRetryableClusterError determines if a cluster creation failure is transient and eligible for adaptive retry.
//
// Recognized retryable conditions include:
//  1. Rate limiting: HTTP 429 Too Many Requests or gRPC Code 8 (ResourceExhausted).
//  2. Lingering background deletions / 409 Conflicts: Occurs when recreating a cluster with the same name while
//     the prior cluster is still in the process of background teardown on TMC or Supervisor.
//  3. Transient network / transport drops: Connection resets, broken pipes, EOFs, or transport closing.
//  4. Server-side 5xx errors: HTTP 500, 502, 503, 504 gateway or internal errors.
func isRetryableClusterError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := strings.ToLower(err.Error())

	// 1. Rate limiting (HTTP 429 / gRPC Code 8 ResourceExhausted)
	if strings.Contains(errMsg, "status : 429") ||
		strings.Contains(errMsg, "too many requests") ||
		strings.Contains(errMsg, "rate limit") ||
		strings.Contains(errMsg, "rate limiting") ||
		strings.Contains(errMsg, "resourceexhausted") ||
		strings.Contains(errMsg, "\"code\":8") ||
		strings.Contains(errMsg, "code = resourceexhausted") {
		return true
	}

	// 2. Conflict when an old cluster is still deleting or terminating in background
	if clienterrors.IsAlreadyExistsError(err) ||
		strings.Contains(errMsg, "already exists") ||
		strings.Contains(errMsg, "status : 409") ||
		strings.Contains(errMsg, "conflict") {
		return true
	}

	// 3. Transient network / RPC / gateway issues
	if strings.Contains(errMsg, "eof") ||
		strings.Contains(errMsg, "connection reset") ||
		strings.Contains(errMsg, "connection refused") ||
		strings.Contains(errMsg, "transport is closing") ||
		strings.Contains(errMsg, "unavailable") ||
		strings.Contains(errMsg, "broken pipe") ||
		strings.Contains(errMsg, "deadline exceeded") {
		return true
	}

	// 4. Transient 5xx server errors
	if strings.Contains(errMsg, "status : 500") ||
		strings.Contains(errMsg, "status : 502") ||
		strings.Contains(errMsg, "status : 503") ||
		strings.Contains(errMsg, "status : 504") {
		return true
	}

	return false
}

// createClusterWithRetry creates a Tanzu Kubernetes Cluster using exponential backoff with randomized jitter.
//
// Key behaviors:
//   - Handles 409 Conflict / "already exists": If a cluster with the same name is found in DELETING phase,
//     the retry loop waits for the prior cluster deletion to complete before retrying creation.
//   - Idempotency: If an existing cluster is already in an active non-deleting state, it returns success.
//   - Retries transient rate limits (429), gateway drops, and 5xx errors until the cluster is accepted or timeout expires.
func createClusterWithRetry(
	ctx context.Context,
	config *authctx.TanzuContext,
	clusterRequest *tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterData,
	timeout time.Duration,
) error {
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	currentBackoff := nodePoolInitialBackoff
	attempt := 0
	clusterFn := clusterRequest.TanzuKubernetesCluster.FullName

	for {
		attempt++

		_, err := config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterResourceServiceCreate(clusterRequest)
		if err == nil {
			return nil
		}

		// Idempotency / lingering deletion handling on 409 Conflict
		if clienterrors.IsAlreadyExistsError(err) || strings.Contains(strings.ToLower(err.Error()), "already exists") {
			existingCluster, getErr := config.TMCConnection.TanzuKubernetesClusterResourceService.TanzuKubernetesClusterResourceServiceGet(clusterFn)
			if getErr == nil && existingCluster != nil && existingCluster.TanzuKubernetesCluster != nil {
				tkc := existingCluster.TanzuKubernetesCluster
				if tkc.Status != nil && tkc.Status.Phase != nil {
					phase := *tkc.Status.Phase
					// If the existing cluster is actively DELETING, wait for it to be fully removed
					if phase == tanzukubernetesclustermodels.VmwareTanzuManageV1alpha1CommonClusterStatusPhaseDELETING {
						log.Printf("[DEBUG] Cluster %q currently in DELETING phase, waiting for removal before recreation...", clusterFn.Name)
					} else {
						// Existing cluster is present and not deleting, idempotent success
						log.Printf("[INFO] Cluster %q already exists in phase %q, proceeding idempotently", clusterFn.Name, phase)
						return nil
					}
				} else {
					return nil
				}
			}
		}

		if !isRetryableClusterError(err) {
			return err
		}

		sleepDuration, next := calculateBackoffWithJitter(currentBackoff, nodePoolMaxBackoff, nodePoolBackoffFactor, nodePoolJitterFactor)
		currentBackoff = next

		log.Printf("[DEBUG] Retrying creation for cluster %q (attempt %d) in %v due to transient/conflict error: %s",
			clusterFn.Name, attempt, sleepDuration, err.Error())

		select {
		case <-ctx.Done():
			return errors.Wrapf(ctx.Err(), "timed out waiting for cluster %q creation to be accepted after %d attempts. Last error: %s",
				clusterFn.Name, attempt, err.Error())
		case <-time.After(sleepDuration):
		}
	}
}
