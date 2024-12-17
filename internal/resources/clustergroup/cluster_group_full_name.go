// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustergroup

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
)

var ClusterGroupFullname = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The schema for cluster group full name",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			NameKey: {
				Type:        schema.TypeString,
				Description: "Name of the cluster group",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

func ConstructClusterGroupFullname(data []interface{}) (fullname *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName{}

	if v, ok := fullNameData[NameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.Name, NameKey)
	}

	return fullname
}

func FlattenClusterGroupFullname(fullname *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[NameKey] = fullname.Name

	return []interface{}{flattenFullname}
}
