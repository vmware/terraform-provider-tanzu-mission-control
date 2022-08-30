/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicy

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	organizationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/organization"
)

var organizationFullname = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The schema for organization iam policy full name",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			organizationIDKey: {
				Type:        schema.TypeString,
				Description: "ID of the Organization",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

func constructOrganizationFullname(data []interface{}) (fullname *organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName{}

	if v, ok := fullNameData[organizationIDKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.OrgID, organizationIDKey)
	}

	return fullname
}

func flattenOrganizationFullname(fullname *organizationmodel.VmwareTanzuManageV1alpha1OrganizationFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[organizationIDKey] = fullname.OrgID

	return []interface{}{flattenFullname}
}
