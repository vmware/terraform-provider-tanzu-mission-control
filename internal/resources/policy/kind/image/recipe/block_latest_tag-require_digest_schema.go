/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipeimagemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image"
)

var BlockLatestTag = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for image policy block-latest-tag recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			AuditKey: {
				Type:        schema.TypeBool,
				Description: "Audit (dry-run). Violations will be logged but not denied.",
				Optional:    true,
				Default:     false,
			},
		},
	},
}

var RequireDigest = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for image policy require-digest recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			AuditKey: {
				Type:        schema.TypeBool,
				Description: "Audit (dry-run). Violations will be logged but not denied.",
				Optional:    true,
				Default:     false,
			},
		},
	},
}

func ConstructCommonRecipe(data []interface{}) (common *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CommonRecipe) {
	if len(data) == 0 || data[0] == nil {
		return common
	}

	commonRecipeData, _ := data[0].(map[string]interface{})

	common = &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CommonRecipe{}

	if v, ok := commonRecipeData[AuditKey]; ok {
		helper.SetPrimitiveValue(v, &common.Audit, AuditKey)
	}

	return common
}

func FlattenCommonRecipe(common *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CommonRecipe) (data []interface{}) {
	if common == nil {
		return data
	}

	flattenCommonRecipe := make(map[string]interface{})

	flattenCommonRecipe[AuditKey] = common.Audit

	return []interface{}{flattenCommonRecipe}
}
