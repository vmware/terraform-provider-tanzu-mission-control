/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import (
	pkginstallclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackageinstall"
)

func FlattenStatusForClusterScope(status *pkginstallclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus) (data interface{}) {
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

	flattenStatusData[packageInstallPhaseKey] = condition.Message
	flattenStatusData[resolvedVersionKey] = status.ResolvedVersion
	flattenStatusData[managedKey] = status.Managed
	flattenStatusData[generatedResourcesKey] = []interface{}{
		map[string]interface{}{
			clusterRoleNameKey:    status.GeneratedResources.ClusterRoleName,
			serviceAccountNameKey: status.GeneratedResources.ServiceAccountName,
			roleBindingNameKey:    status.GeneratedResources.RoleBindingName,
		},
	}
	flattenStatusData[referredByKey] = status.ReferredBy

	return []interface{}{flattenStatusData}
}
