// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

// Package scope contains schema and helper functions for different policy scopes.
// nolint: dupl
package scope

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyworkspacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/workspace"
)

var WorkspacePolicyFullname = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The schema for workspace policy full name",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			WorkspaceNameKey: {
				Type:        schema.TypeString,
				Description: "Name of this workspace",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

func ConstructWorkspacePolicyFullname(data []interface{}, name string) (fullname *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName{}

	if v, ok := fullNameData[WorkspaceNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.WorkspaceName, WorkspaceNameKey)
	}

	fullname.Name = name

	return fullname
}

func FlattenWorkspacePolicyFullname(fullname *policyworkspacemodel.VmwareTanzuManageV1alpha1WorkspacePolicyFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[WorkspaceNameKey] = fullname.WorkspaceName

	return []interface{}{flattenFullname}
}
