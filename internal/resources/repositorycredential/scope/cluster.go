/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	repositorycredentialclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/repositorycredential/cluster"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func ConstructClusterRepositoryCredentialFullname(data []interface{}, name, orgID string) (fullname *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName{}

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

	if orgID != "" {
		fullname.OrgId = orgID
	}

	return fullname
}

func FlattenClusterKustomizationFullname(fullname *repositorycredentialclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[commonscope.ManagementClusterNameKey] = fullname.ManagementClusterName
	flattenFullname[commonscope.ProvisionerNameKey] = fullname.ProvisionerName
	flattenFullname[commonscope.ClusterNameKey] = fullname.ClusterName

	return []interface{}{flattenFullname}
}
