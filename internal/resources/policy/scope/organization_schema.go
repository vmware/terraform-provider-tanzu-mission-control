/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

// Package scope contains schema and helper functions for different policy scopes.
// nolint: dupl
package scope

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyorganizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/organization"
)

var OrganizationPolicyFullname = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The schema for organization policy full name",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			OrganizationIDKey: {
				Type:        schema.TypeString,
				Description: "ID of this organization",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

func ConstructOrganizationPolicyFullname(data []interface{}, name string) (fullname *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName{}

	if v, ok := fullNameData[OrganizationIDKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.OrgID, OrganizationIDKey)
	}

	fullname.Name = name

	return fullname
}

func FlattenOrganizationPolicyFullname(fullname *policyorganizationmodel.VmwareTanzuManageV1alpha1OrganizationPolicyFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[OrganizationIDKey] = fullname.OrgID

	return []interface{}{flattenFullname}
}
