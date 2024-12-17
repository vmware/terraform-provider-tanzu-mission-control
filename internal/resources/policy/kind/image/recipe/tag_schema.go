// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipeimagecommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image/common"
)

var tag = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Allowed image tag, wildcards are supported (for example: v1.*). No validation is performed on tag if the field is empty.",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			NegateKey: {
				Type:        schema.TypeBool,
				Description: "The negate flag used to exclude certain tag patterns.",
				Optional:    true,
				Default:     false,
			},
			ValueKey: {
				Type:        schema.TypeString,
				Description: "The value (support wildcard) is used to validate against the tag of the image.",
				Optional:    true,
				Default:     "",
			},
		},
	},
}

func expandTag(data []interface{}) (tag *policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag) {
	if len(data) == 0 || data[0] == nil {
		return tag
	}

	tagsData, ok := data[0].(map[string]interface{})
	if !ok {
		return tag
	}

	tag = &policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag{}

	if v, ok := tagsData[NegateKey]; ok {
		tag.Negate = helper.BoolPointer(v.(bool))
	}

	if v, ok := tagsData[ValueKey]; ok {
		helper.SetPrimitiveValue(v, &tag.Value, ValueKey)
	}

	return tag
}

func flattenTag(tag *policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag) (data []interface{}) {
	if tag == nil {
		return data
	}

	flattenTag := make(map[string]interface{})

	if tag.Negate != nil {
		flattenTag[NegateKey] = *tag.Negate
	}

	flattenTag[ValueKey] = tag.Value

	return []interface{}{flattenTag}
}
