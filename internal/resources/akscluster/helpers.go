/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster

import (
	"context"
	"encoding/json"
	"sort"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	models "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type retryInterval string

func setResourceState(data *schema.ResourceData, cluster *models.VmwareTanzuManageV1alpha1AksclusterAksCluster, nodepools []*models.VmwareTanzuManageV1alpha1AksclusterNodepoolNodepool) error {
	data.SetId(cluster.Meta.UID)

	if err := data.Set(common.MetaKey, common.FlattenMeta(cluster.Meta)); err != nil {
		return err
	}

	sort.Slice(nodepools, func(i, j int) bool { return nodepools[i].FullName.Name < nodepools[j].FullName.Name })

	specMap := toClusterSpecMap(cluster.Spec, nodepools)
	if err := data.Set(clusterSpecKey, specMap); err != nil {
		return err
	}

	return nil
}

func clusterIsReady(resp *models.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse) bool {
	if resp == nil || resp.AksCluster == nil || resp.AksCluster.Status == nil || *resp.AksCluster.Status.Phase != models.VmwareTanzuManageV1alpha1AksclusterPhaseREADY {
		return false
	}

	return true
}

func clusterHasFatalError(resp *models.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse) bool {
	if resp == nil || resp.AksCluster == nil || resp.AksCluster.Status == nil || *resp.AksCluster.Status.Phase != models.VmwareTanzuManageV1alpha1AksclusterPhaseERROR {
		return false
	}

	return true
}

func nodepoolIsReady(resp *models.VmwareTanzuManageV1alpha1AksclusterNodepoolGetNodepoolResponse) bool {
	if resp == nil || resp.Nodepool == nil || resp.Nodepool.Status == nil || *resp.Nodepool.Status.Phase != models.VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseREADY {
		return false
	}

	return true
}

func nodepoolHasFatalError(resp *models.VmwareTanzuManageV1alpha1AksclusterNodepoolGetNodepoolResponse) bool {
	if resp == nil || resp.Nodepool == nil || resp.Nodepool.Status == nil || *resp.Nodepool.Status.Phase != models.VmwareTanzuManageV1alpha1AksclusterNodepoolPhaseERROR {
		return false
	}

	return true
}

func getErrorReason(conditions map[string]models.VmwareTanzuCoreV1alpha1StatusCondition) string {
	msg, err := json.Marshal(conditions)
	if err != nil {
		return "unknown error"
	}

	return string(msg)
}

func getTimeOut(data *schema.ResourceData) time.Duration {
	timeout, ok := data.GetOk(waitKey)
	if !ok {
		return defaultTimeout
	}

	duration, err := time.ParseDuration(timeout.(string))
	if err != nil {
		return defaultTimeout
	}

	return duration
}

func getPollInterval(ctx context.Context) time.Duration {
	value := ctx.Value(RetryInterval)
	if value != nil {
		return value.(time.Duration)
	}

	return defaultInterval
}

func extractClusterFullName(d *schema.ResourceData) *models.VmwareTanzuManageV1alpha1AksclusterFullName {
	fn := &models.VmwareTanzuManageV1alpha1AksclusterFullName{}
	fn.CredentialName, _ = d.Get(CredentialNameKey).(string)
	fn.SubscriptionID, _ = d.Get(SubscriptionIDKey).(string)
	fn.ResourceGroupName, _ = d.Get(ResourceGroupNameKey).(string)
	fn.Name, _ = d.Get(NameKey).(string)

	return fn
}
