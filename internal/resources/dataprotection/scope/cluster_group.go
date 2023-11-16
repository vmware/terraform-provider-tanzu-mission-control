/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionscope

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	dataprotectionclustergroupmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/dataprotection/clustergroup/dataprotection"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func ConstructDataProtectionClusterGroupFullname(data []interface{}) (fullname *dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName{}

	if nameValue, ok := fullNameData[commonscope.NameKey]; ok {
		helper.SetPrimitiveValue(nameValue, &fullname.ClusterGroupName, commonscope.NameKey)
	}

	return fullname
}

func FlattenDataProtectionClusterGroupFullname(fullname *dataprotectionclustergroupmodels.VmwareTanzuManageV1alpha1ClusterGroupDataprotectionFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[commonscope.NameKey] = fullname.ClusterGroupName

	return []interface{}{flattenFullname}
}
