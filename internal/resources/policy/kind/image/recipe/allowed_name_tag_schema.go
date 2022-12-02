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

var AllowedNameTag = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for image policy allowed-name-tag recipe version v1",
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
			RulesKey: {
				Type:        schema.TypeList,
				Description: "It specifies a list of rules that defines allowed image patterns.",
				Required:    true,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						ImageNameKey: {
							Type:        schema.TypeString,
							Description: "Allowed image names, wildcards are supported(for example: fooservice/*). Empty field is equivalent to *.",
							Optional:    true,
							Default:     "",
						},
						TagKey: tag,
					},
				},
			},
		},
	},
}

func ConstructAllowedNameTag(data []interface{}) (nameTag *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag) {
	if len(data) == 0 || data[0] == nil {
		return nameTag
	}

	allowedNameTagData, _ := data[0].(map[string]interface{})

	nameTag = &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag{}

	if v, ok := allowedNameTagData[AuditKey]; ok {
		helper.SetPrimitiveValue(v, &nameTag.Audit, AuditKey)
	}

	if v, ok := allowedNameTagData[RulesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				nameTag.Rules = make([]*policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules, 0)

				for _, raw := range vs {
					nameTag.Rules = append(nameTag.Rules, expandNameTagRules(raw))
				}
			}
		}
	}

	return nameTag
}

func expandNameTagRules(data interface{}) (rules *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules) {
	if data == nil {
		return rules
	}

	rulesData, ok := data.(map[string]interface{})
	if !ok {
		return rules
	}

	rules = &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules{}

	if v, ok := rulesData[ImageNameKey]; ok {
		helper.SetPrimitiveValue(v, &rules.ImageName, ImageNameKey)
	}

	if v, ok := rulesData[TagKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			rules.Tag = expandTag(v1)
		}
	}

	return rules
}

func FlattenAllowedNameTag(nameTag *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag) (data []interface{}) {
	if nameTag == nil {
		return data
	}

	flattenAllowedNameTag := make(map[string]interface{})

	flattenAllowedNameTag[AuditKey] = nameTag.Audit

	if nameTag.Rules != nil {
		var rules []interface{}

		for _, rule := range nameTag.Rules {
			rules = append(rules, flattenNameTagRules(rule))
		}

		flattenAllowedNameTag[RulesKey] = rules
	}

	return []interface{}{flattenAllowedNameTag}
}

func flattenNameTagRules(rules *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules) (data interface{}) {
	if rules == nil {
		return data
	}

	flattenRules := make(map[string]interface{})

	flattenRules[ImageNameKey] = rules.ImageName

	if rules.Tag != nil {
		flattenRules[TagKey] = flattenTag(rules.Tag)
	}

	return flattenRules
}
