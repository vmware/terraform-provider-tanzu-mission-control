// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

// Package scope contains schema and helper functions for different policy scopes.
// nolint: dupl
package scope

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/clustergroup"
)

var ClusterGroupPolicyFullname = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The schema for cluster group policy full name",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ClusterGroupNameKey: {
				Type:        schema.TypeString,
				Description: "Name of this cluster group",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

func ConstructClusterGroupPolicyFullname(data []interface{}, name string) (fullname *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName{}

	if v, ok := fullNameData[ClusterGroupNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ClusterGroupName, ClusterGroupNameKey)
	}

	fullname.Name = name

	return fullname
}

func FlattenClusterGroupPolicyFullname(fullname *policyclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupPolicyFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[ClusterGroupNameKey] = fullname.ClusterGroupName

	return []interface{}{flattenFullname}
}
