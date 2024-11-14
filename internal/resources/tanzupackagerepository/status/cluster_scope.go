// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package status

import (
	pkgrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackagerepository"
)

func FlattenStatusForClusterScope(status *pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryStatus) (data interface{}) {
	if status == nil {
		return data
	}

	if status.Conditions == nil {
		return data
	}

	condition, ok := status.Conditions[conditionReady]
	if !ok {
		return data
	}

	flattenStatusData := make(map[string]interface{})

	flattenStatusData[packageRepositoryPhaseKey] = condition.Message
	flattenStatusData[subscribedKey] = status.Subscribed
	flattenStatusData[disabledKey] = status.Disabled
	flattenStatusData[managedKey] = status.Managed

	return []interface{}{flattenStatusData}
}
