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

var Custom = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for image policy custom recipe version v1",
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
						HostNameKey: {
							Type:        schema.TypeString,
							Description: "Allowed image hostnames, wildcards are supported(for example: *.mycompany.com). Empty field is equivalent to *.",
							Optional:    true,
							Default:     "",
						},
						ImageNameKey: {
							Type:        schema.TypeString,
							Description: "Allowed image names, wildcards are supported(for example: fooservice/*). Empty field is equivalent to *.",
							Optional:    true,
							Default:     "",
						},
						PortKey: {
							Type:        schema.TypeString,
							Description: "Allowed port(if presented) of the image hostname, must associate with valid hostname. Wildcards are supported.",
							Optional:    true,
							Default:     "",
						},
						RequireKey: {
							Type:        schema.TypeBool,
							Description: "The flag used to enforce digest to appear in container images.",
							Optional:    true,
							Default:     false,
						},
						TagKey: tag,
					},
				},
			},
		},
	},
}

func ConstructCustom(data []interface{}) (custom *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1Custom) {
	if len(data) == 0 || data[0] == nil {
		return custom
	}

	customData, _ := data[0].(map[string]interface{})

	custom = &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1Custom{}

	if v, ok := customData[AuditKey]; ok {
		custom.Audit = helper.BoolPointer(v.(bool))
	}

	if v, ok := customData[RulesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				custom.Rules = make([]*policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CustomRules, 0)

				for _, raw := range vs {
					custom.Rules = append(custom.Rules, expandCustomRules(raw))
				}
			}
		}
	}

	return custom
}

func expandCustomRules(data interface{}) (rules *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CustomRules) {
	if data == nil {
		return rules
	}

	rulesData, ok := data.(map[string]interface{})
	if !ok {
		return rules
	}

	rules = &policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CustomRules{}

	if v, ok := rulesData[HostNameKey]; ok {
		helper.SetPrimitiveValue(v, &rules.Hostname, HostNameKey)
	}

	if v, ok := rulesData[ImageNameKey]; ok {
		helper.SetPrimitiveValue(v, &rules.ImageName, ImageNameKey)
	}

	if v, ok := rulesData[PortKey]; ok {
		helper.SetPrimitiveValue(v, &rules.Port, PortKey)
	}

	if v, ok := rulesData[RequireKey]; ok {
		rules.RequireDigest = helper.BoolPointer(v.(bool))
	}

	if v, ok := rulesData[TagKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			rules.Tag = expandTag(v1)
		}
	}

	return rules
}

func FlattenCustom(custom *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1Custom) (data []interface{}) {
	if custom == nil {
		return data
	}

	flattenCustom := make(map[string]interface{})

	if custom.Audit != nil {
		flattenCustom[AuditKey] = *custom.Audit
	}

	if custom.Rules != nil {
		var rules []interface{}

		for _, rule := range custom.Rules {
			rules = append(rules, flattenCustomRules(rule))
		}

		flattenCustom[RulesKey] = rules
	}

	return []interface{}{flattenCustom}
}

func flattenCustomRules(rules *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CustomRules) (data interface{}) {
	if rules == nil {
		return data
	}

	flattenRules := make(map[string]interface{})

	flattenRules[HostNameKey] = rules.Hostname
	flattenRules[ImageNameKey] = rules.ImageName
	flattenRules[PortKey] = rules.Port

	if rules.RequireDigest != nil {
		flattenRules[RequireKey] = *rules.RequireDigest
	}

	if rules.Tag != nil {
		flattenRules[TagKey] = flattenTag(rules.Tag)
	}

	return flattenRules
}
