// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package status

import releaseclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/cluster"

func FlattenStatusForClusterScope(status *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseStatus) (data interface{}) {
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

	flattenStatusData[phaseKey] = condition.Reason

	flattenStatusData[generatedResourcesKey] = []interface{}{
		map[string]interface{}{
			clusterRoleNameKey:    status.GeneratedResources.ClusterRoleName,
			serviceAccountNameKey: status.GeneratedResources.ServiceAccountName,
			roleBindingNameKey:    status.GeneratedResources.RoleBindingName,
		},
	}

	return []interface{}{flattenStatusData}
}
