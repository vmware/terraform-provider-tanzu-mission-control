/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import releaseclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/clustergroup"

func FlattenStatusForClusterGroupScope(status *releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus) (data interface{}) {
	if status == nil {
		return data
	}

	if status.Phase == nil {
		return data
	}

	flattenStatusData := make(map[string]interface{})

	flattenStatusData[phaseKey] = string(*status.Phase)

	return []interface{}{flattenStatusData}
}
