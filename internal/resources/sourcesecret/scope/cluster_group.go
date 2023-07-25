/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	sourcesecretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func ConstructClusterGroupSourcesecretFullname(data []interface{}, name string) (fullname *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName{}

	if nameValue, ok := fullNameData[commonscope.NameKey]; ok {
		helper.SetPrimitiveValue(nameValue, &fullname.ClusterGroupName, commonscope.NameKey)
	}

	fullname.Name = name

	return fullname
}

func FlattenClusterGroupSourcesecretFullname(fullname *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[commonscope.NameKey] = fullname.ClusterGroupName

	return []interface{}{flattenFullname}
}
