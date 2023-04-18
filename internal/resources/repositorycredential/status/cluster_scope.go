/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import repositorycredentialclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/repositorycredential/cluster"

func FlattenStatusForClusterScope(status *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretStatus) (data interface{}) {
	if status == nil {
		return data
	}

	if status.Status == nil {
		return data
	}

	if status.Status.Phase == nil {
		return data
	}

	c := status.Status.Phase

	flattenStatusData := make(map[string]interface{})

	flattenStatusData[phaseKey] = string(*c)

	return flattenStatusData
}
