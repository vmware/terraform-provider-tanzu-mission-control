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

func ConstructClusterGitRepositoryFullname(data []interface{}, name, namespace string) (fullname *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName{}

	if managementClusterNameValue, ok := fullNameData[commonscope.ManagementClusterNameKey]; ok {
		helper.SetPrimitiveValue(managementClusterNameValue, &fullname.ManagementClusterName, commonscope.ManagementClusterNameKey)
	}

	if provisionerNameValue, ok := fullNameData[commonscope.ProvisionerNameKey]; ok {
		helper.SetPrimitiveValue(provisionerNameValue, &fullname.ProvisionerName, commonscope.ProvisionerNameKey)
	}

	if nameValue, ok := fullNameData[commonscope.NameKey]; ok {
		helper.SetPrimitiveValue(nameValue, &fullname.ClusterName, commonscope.NameKey)
	}

	fullname.Name = name
	fullname.NamespaceName = namespace

	return fullname
}

func FlattenClusterGitRepositoryFullname(fullname *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[commonscope.ManagementClusterNameKey] = fullname.ManagementClusterName
	flattenFullname[commonscope.ProvisionerNameKey] = fullname.ProvisionerName
	flattenFullname[commonscope.NameKey] = fullname.ClusterName

	return []interface{}{flattenFullname}
}
