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

var Baseline = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for security policy baseline recipe version v1",
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

func ConstructBaseline(data []interface{}) (baseline *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Baseline) {
	if len(data) == 0 || data[0] == nil {
		return baseline
	}

	baselineData, _ := data[0].(map[string]interface{})

	baseline = &policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Baseline{}

	if v, ok := baselineData[AuditKey]; ok {
		helper.SetPrimitiveValue(v, &baseline.Audit, AuditKey)
	}

	if v, ok := baselineData[DisableNativePspKey]; ok {
		helper.SetPrimitiveValue(v, &baseline.DisableNativePsp, DisableNativePspKey)
	}

	return baseline
}

func FlattenBaseline(baseline *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Baseline) (data []interface{}) {
	if baseline == nil {
		return data
	}

	flattenBaseline := make(map[string]interface{})

	flattenBaseline[AuditKey] = baseline.Audit
	flattenBaseline[DisableNativePspKey] = baseline.DisableNativePsp

	return []interface{}{flattenBaseline}
}
