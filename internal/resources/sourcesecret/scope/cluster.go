// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	sourcesecretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func ConstructClusterSourcesecretFullname(data []interface{}, name string) (fullname *sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName{}

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

	return fullname
}

func FlattenClusterSourcesecretFullname(fullname *sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[commonscope.ManagementClusterNameKey] = fullname.ManagementClusterName
	flattenFullname[commonscope.ProvisionerNameKey] = fullname.ProvisionerName
	flattenFullname[commonscope.NameKey] = fullname.ClusterName

	return []interface{}{flattenFullname}
}
