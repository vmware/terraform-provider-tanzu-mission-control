/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	dataprotectionmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/dataprotection"
)

func ConstructClusterScopeFullname(data []interface{}) (fullname *dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &dataprotectionmodels.VmwareTanzuManageV1alpha1ClusterDataprotectionFullName{}

	if v, ok := fullNameData[ManagementClusterNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ManagementClusterName, ManagementClusterNameKey)
	}

	if v, ok := fullNameData[ProvisionerNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ProvisionerName, ProvisionerNameKey)
	}

	if v, ok := fullNameData[ClusterNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ClusterName, ClusterNameKey)
	}

	return fullname
}
