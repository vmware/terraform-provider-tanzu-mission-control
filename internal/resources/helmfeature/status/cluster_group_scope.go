/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import helmclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmfeature/clustergroup"

func FlattenStatusForClusterGroupScope(status *helmclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmStatus) (data interface{}) {
	if status == nil {
		return data
	}

	if status.Phase == nil {
		return data
	}

	flattenStatusData := make(map[string]interface{})

	flattenStatusData[phaseKey] = string(*status.Phase)

	return flattenStatusData
}
