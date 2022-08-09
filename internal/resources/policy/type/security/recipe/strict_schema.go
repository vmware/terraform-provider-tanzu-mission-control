/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

// Package recipe contains schema and helper functions for different input recipes.
// nolint: dupl
package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipesecuritymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/security"
)

var Strict = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for security policy strict recipe version v1",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			AuditKey: {
				Type:        schema.TypeBool,
				Description: "Audit (dry-run)",
				Optional:    true,
				Default:     false,
			},
			DisableNativePspKey: {
				Type:        schema.TypeBool,
				Description: "Disable native pod security policy",
				Optional:    true,
				Default:     false,
			},
		},
	},
}

func ConstructStrict(data []interface{}) (strict *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict) {
	if len(data) == 0 || data[0] == nil {
		return strict
	}

	strictData, _ := data[0].(map[string]interface{})

	strict = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict{}

	if v, ok := strictData[AuditKey]; ok {
		helper.SetPrimitiveValue(v, &strict.Audit, AuditKey)
	}

	if v, ok := strictData[DisableNativePspKey]; ok {
		helper.SetPrimitiveValue(v, &strict.DisableNativePsp, DisableNativePspKey)
	}

	return strict
}

func FlattenStrict(strict *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict) (data []interface{}) {
	if strict == nil {
		return data
	}

	flattenStrict := make(map[string]interface{})

	flattenStrict[AuditKey] = strict.Audit
	flattenStrict[DisableNativePspKey] = strict.DisableNativePsp

	return []interface{}{flattenStrict}
}
