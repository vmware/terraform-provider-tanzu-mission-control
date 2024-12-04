// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package status

import secretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"

func FlattenStatusForClusterScope(status *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretStatus) (data interface{}) {
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

	if condition.Status == nil {
		return data
	}

	flattenStatusData := make(map[string]interface{})

	flattenStatusData[stateKey] = string(*condition.Status)

	return flattenStatusData
}
