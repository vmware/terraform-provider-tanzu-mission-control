/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	dataprotectioncgmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup/dataprotection"
)

func ConstructClusterGroupScopeFullname(data []interface{}) (fullname *dataprotectioncgmodels.VmwareTanzuManageV1alpha1ClustergroupDataprotectionFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &dataprotectioncgmodels.VmwareTanzuManageV1alpha1ClustergroupDataprotectionFullName{}

	if v, ok := fullNameData[ClusterGroupNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ClusterGroupName, ClusterGroupNameKey)
	}

	return fullname
}
