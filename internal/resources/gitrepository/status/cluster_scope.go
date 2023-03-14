/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"

func FlattenStatusForClusterScope(status *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryStatus) (data interface{}) {
	if status == nil {
		return data
	}

	if status.Conditions == nil {
		return data
	}

	c, ok := status.Conditions[conditionReady]
	if !ok {
		return data
	}

	if c.Status == nil {
		return data
	}

	flattenStatusData := make(map[string]interface{})

	flattenStatusData[stateKey] = string(*c.Status)

	return flattenStatusData
}
