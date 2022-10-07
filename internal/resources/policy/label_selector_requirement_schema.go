/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policy

import (
	"context"
	"fmt"
	"strings"

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

func ValidateSpecLabelSelectorRequirement(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
	value, ok := diff.GetOk(SpecKey)
	if !ok {
		return fmt.Errorf("spec: %v is not valid: minimum one valid spec block is required", value)
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("spec data: %v is not valid: minimum one valid spec block is required", data)
	}

	specData := data[0].(map[string]interface{})

	namespaceSelector, ok := specData[NamespaceSelectorKey]
	if !ok {
		return fmt.Errorf("namespace_selector: %v is not valid: minimum one valid namespace_selector block is required", namespaceSelector)
	}

	namespaceData, ok := namespaceSelector.([]interface{})
	if !ok || len(namespaceData) == 0 || namespaceData[0] == nil {
		return fmt.Errorf("namespace_selector data: %v is not valid: minimum one valid namespace_selector block is required", namespaceData)
	}

	namespaceSelectorData, _ := namespaceData[0].(map[string]interface{})

	matchExpression, ok := namespaceSelectorData[MatchExpressionsKey]
	if !ok {
		return fmt.Errorf("match expressions: %v is not valid: minimum one valid match expressions block is required", matchExpression)
	}

	matchData, ok := matchExpression.([]interface{})
	if !ok {
		return fmt.Errorf("type of match expressions block data: %v is not valid", matchData)
	}

	var errStrings []string

	for _, raw := range matchData {
		if raw == nil {
			continue
		}

		labelSelectorRequirementData, _ := raw.(map[string]interface{})

		operatorData, ok := labelSelectorRequirementData[OperatorKey]
		if !ok {
			errStrings = append(errStrings, fmt.Errorf("- operator: %v is not valid: minimum one valid operator attribute is required", operatorData).Error())
		}

		operator := operatorData.(string)
		values := make([]string, 0)

		valueData, ok := labelSelectorRequirementData[ValuesKey]
		if !ok {
			errStrings = append(errStrings, fmt.Errorf("- values: %v is not valid: minimum one valid values attribute is required", valueData).Error())
		}

		vs1, ok := valueData.([]interface{})
		if !ok {
			errStrings = append(errStrings, fmt.Errorf("- type of values attribute data: %v is not valid", vs1).Error())
		}

		for _, raw1 := range vs1 {
			values = append(values, raw1.(string))
		}

		if (operator == "In" || operator == "NotIn") && len(values) == 0 {
			errStrings = append(errStrings, fmt.Errorf("- found label selector requirement with operator: \"%v\" and values: %v; If the operator is In or NotIn, the values array must be non-empty", operator, strings.Join(values, `, `)).Error())
		} else if (operator == "Exists" || operator == "DoesNotExist") && len(values) != 0 {
			errStrings = append(errStrings, fmt.Errorf("- found label selector requirement with operator: \"%v\" and values: %v; If the operator is Exists or DoesNotExist, the values array must be empty", operator, strings.Join(values, `, `)).Error())
		}
	}

	if len(errStrings) != 0 {
		return fmt.Errorf("error(s) in label selector requirement(s): \n%s", strings.Join(errStrings, "\n"))
	}

	return nil
}
