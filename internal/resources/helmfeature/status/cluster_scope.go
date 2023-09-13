/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import (
	helmclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/cluster"
)

func FlattenStatusForClusterScope(status *helmclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus) (data interface{}) {
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
