/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

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
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			NegateKey: {
				Type:        schema.TypeBool,
				Description: "The negate flag used to exclude certain tag patterns.",
				Required:    true,
			},
			ValueKey: {
				Type:        schema.TypeString,
				Description: "The value (support wildcard) is used to validate against the tag of the image.",
				Optional:    true,
			},
		},
	},
}

func expandTag(data interface{}) (tag *policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag) {
	if data == nil {
		return tag
	}

	tagsData, ok := data.(map[string]interface{})
	if !ok {
		return tag
	}

	tag = &policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag{}

	if v, ok := tagsData[NegateKey]; ok {
		helper.SetPrimitiveValue(v, &tag.Negate, NegateKey)
	}

	if v, ok := tagsData[ValueKey]; ok {
		helper.SetPrimitiveValue(v, &tag.Value, ValueKey)
	}

	return tag
}

func flattenTag(tag *policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag) (data interface{}) {
	if tag == nil {
		return data
	}

	flattenTag := make(map[string]interface{})

	flattenTag[NegateKey] = tag.Negate
	flattenTag[ValueKey] = tag.Value

	return flattenTag
}
