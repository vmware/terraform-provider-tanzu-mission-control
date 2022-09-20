/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package workspace

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	workspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/workspace"
)

var WorkspaceFullname = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The schema for workspace iam policy full name",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			NameKey: {
				Type:        schema.TypeString,
				Description: "Name of the workspace",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

func ConstructWorkspaceFullname(data []interface{}) (fullname *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName{}

	if v, ok := fullNameData[NameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.Name, NameKey)
	}

	return fullname
}

func FlattenWorkspaceFullname(fullname *workspacemodel.VmwareTanzuManageV1alpha1WorkspaceFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[NameKey] = fullname.Name

	return []interface{}{flattenFullname}
}
