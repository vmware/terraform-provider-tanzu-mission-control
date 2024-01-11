/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	backupscheduleclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/backupschedule/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func ConstructClusterGroupBackupScheduleFullname(data []interface{}, name string) (fullname *backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &backupscheduleclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupDataprotectionScheduleFullName{}

	if nameValue, ok := fullNameData[commonscope.ClusterGroupNameKey]; ok {
		helper.SetPrimitiveValue(nameValue, &fullname.ClusterGroupName, commonscope.ClusterGroupNameKey)
	}

	fullname.Name = name

	return fullname
}
