/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
)

var labelSelectorRequirement = &schema.Resource{
	Description: "Metadata describing the type of the resource",
	Schema: map[string]*schema.Schema{
		KeyKey: {
			Type:        schema.TypeString,
			Description: "Key is the label key that the selector applies to",
			Optional:    true,
		},
		OperatorKey: {
			Type:         schema.TypeString,
			Description:  "Operator represents a key's relationship to a set of values",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"In", "NotIn", "Exists", "DoesNotExist"}, false),
		},
		ValuesKey: {
			Type:        schema.TypeList,
			Description: "Values is an array of string values",
			Required:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	},
}

func constructLabelSelectorRequirement(data interface{}) (labelSelectorRequirement *policymodel.K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement) {
	if data == nil {
		return labelSelectorRequirement
	}

	labelSelectorRequirementData, _ := data.(map[string]interface{})

	labelSelectorRequirement = &policymodel.K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement{
		Values: make([]string, 0),
	}

	if v, ok := labelSelectorRequirementData[KeyKey]; ok {
		helper.SetPrimitiveValue(v, &labelSelectorRequirement.Key, KeyKey)
	}

	if v, ok := labelSelectorRequirementData[OperatorKey]; ok {
		helper.SetPrimitiveValue(v, &labelSelectorRequirement.Operator, OperatorKey)
	}

	if v, ok := labelSelectorRequirementData[ValuesKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			labelSelectorRequirement.Values = append(labelSelectorRequirement.Values, raw.(string))
		}
	}

	return labelSelectorRequirement
}

func flattenLabelSelectorRequirement(labelSelectorRequirement *policymodel.K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement) (data interface{}) {
	if labelSelectorRequirement == nil {
		return data
	}

	flattenLabelSelectorRequirement := make(map[string]interface{})

	flattenLabelSelectorRequirement[KeyKey] = labelSelectorRequirement.Key
	flattenLabelSelectorRequirement[OperatorKey] = labelSelectorRequirement.Operator
	flattenLabelSelectorRequirement[ValuesKey] = labelSelectorRequirement.Values

	return flattenLabelSelectorRequirement
}
