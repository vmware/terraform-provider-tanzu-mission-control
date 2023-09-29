/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import (
	helmrepoclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrepository"
)

func FlattenStatusForClusterScope(status *helmrepoclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositoryStatus) (data interface{}) {
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

	return flattenStatusData
}
