/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func ConstructClusterGitRepositoryFullname(data []interface{}, name, namespace, orgID string) (fullname *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName{}

	if v, ok := fullNameData[commonscope.ManagementClusterNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ManagementClusterName, commonscope.ManagementClusterNameKey)
	}

	if v, ok := fullNameData[commonscope.ProvisionerNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ProvisionerName, commonscope.ProvisionerNameKey)
	}

	if v, ok := fullNameData[commonscope.ClusterNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ClusterName, commonscope.ClusterNameKey)
	}

	fullname.Name = name
	fullname.NamespaceName = namespace

	if orgID != "" {
		fullname.OrgID = orgID
	}

	return fullname
}

func FlattenClusterGitRepositoryFullname(fullname *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[commonscope.ManagementClusterNameKey] = fullname.ManagementClusterName
	flattenFullname[commonscope.ProvisionerNameKey] = fullname.ProvisionerName
	flattenFullname[commonscope.ClusterNameKey] = fullname.ClusterName

	return []interface{}{flattenFullname}
}
