/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	gitrepositoryclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func ConstructClusterGroupGitRepositoryFullname(data []interface{}, name, namespace, orgID string) (fullname *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName{}

	if v, ok := fullNameData[commonscope.ClusterGroupNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ClusterGroupName, commonscope.ClusterGroupNameKey)
	}

	fullname.Name = name
	fullname.NamespaceName = namespace

	if orgID != "" {
		fullname.OrgID = orgID
	}

	return fullname
}

func FlattenClusterGroupGitRepositoryFullname(fullname *gitrepositoryclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdGitrepositoryFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[commonscope.ClusterGroupNameKey] = fullname.ClusterGroupName

	return []interface{}{flattenFullname}
}
