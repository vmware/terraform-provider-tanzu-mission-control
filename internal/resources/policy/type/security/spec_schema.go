/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package security

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyrecipesecuritymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/security"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
)

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the security policy",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			inputKey:                    inputSchema,
			policy.NamespaceSelectorKey: policy.NamespaceSelector,
		},
	},
}

func constructSpec(d *schema.ResourceData) (spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec) {
	value, ok := d.GetOk(specKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	spec = &policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec{
		Type:          typeDefaultValue,
		RecipeVersion: recipeVersionDefaultValue,
	}

	v, ok := specData[inputKey]
	if !ok {
		return spec
	}

	v1, ok := v.([]interface{})
	if !ok {
		return spec
	}

	inputRecipeData := constructInput(v1)

	if inputRecipeData == nil || inputRecipeData.recipe == "" {
		return spec
	}

	spec.Recipe = string(inputRecipeData.recipe)

	switch inputRecipeData.recipe {
	case baselineRecipe:
		if inputRecipeData.inputBaseline != nil {
			spec.Input = *inputRecipeData.inputBaseline
		}
	case customRecipe:
		if inputRecipeData.inputCustom != nil {
			spec.Input = *inputRecipeData.inputCustom
		}
	case strictRecipe:
		if inputRecipeData.inputStrict != nil {
			spec.Input = *inputRecipeData.inputStrict
		}
	case unknownRecipe:
		fmt.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(recipesAllowed[:], `, `))
	}

	if v, ok := specData[policy.NamespaceSelectorKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.NamespaceSelector = policy.ConstructNamespaceSelector(v1)
		}
	}

	return spec
}

func flattenSpec(spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	if spec.Input == nil {
		return data
	}

	v1, ok := spec.Input.(map[string]interface{})
	if !ok {
		return data
	}

	var inputRecipeData *inputRecipe

	byteSlice, err := json.Marshal(v1)
	if err != nil {
		return data
	}

	switch spec.Recipe {
	case string(baselineRecipe):
		var baselineRecipeInput policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Baseline

		err = baselineRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:        baselineRecipe,
			inputBaseline: &baselineRecipeInput,
		}
	case string(customRecipe):
		var customRecipeInput policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom

		err = customRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:      customRecipe,
			inputCustom: &customRecipeInput,
		}
	case string(strictRecipe):
		var strictRecipeInput policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict

		err = strictRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:      strictRecipe,
			inputStrict: &strictRecipeInput,
		}
	case string(unknownRecipe):
		fmt.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(recipesAllowed[:], `, `))
	}

	flattenSpecData[inputKey] = flattenInput(inputRecipeData)

	if spec.NamespaceSelector != nil {
		flattenSpecData[policy.NamespaceSelectorKey] = policy.FlattenNamespaceSelector(spec.NamespaceSelector)
	}

	return []interface{}{flattenSpecData}
}

func validateSpecLabelSelectorRequirement(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
	value, ok := diff.GetOk(specKey)
	if !ok {
		return fmt.Errorf("spec: %v is not valid: minimum one valid spec block is required", value)
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("spec data: %v is not valid: minimum one valid spec block is required", data)
	}

	specData := data[0].(map[string]interface{})

	v, ok := specData[policy.NamespaceSelectorKey]
	if !ok {
		return fmt.Errorf("namespace_selector: %v is not valid: minimum one valid namespace_selector block is required", v)
	}

	v1, ok := v.([]interface{})
	if !ok || len(v1) == 0 || v1[0] == nil {
		return fmt.Errorf("namespace_selector data: %v is not valid: minimum one valid namespace_selector block is required", v1)
	}

	namespaceSelectorData, _ := v1[0].(map[string]interface{})

	v2, ok := namespaceSelectorData[policy.MatchExpressionsKey]
	if !ok {
		return fmt.Errorf("match expressions: %v is not valid: minimum one valid match expressions block is required", v2)
	}

	vs, ok := v2.([]interface{})
	if !ok {
		return fmt.Errorf("type of match expressions block data: %v is not valid", vs)
	}

	errStrings := make([]string, 0)

	for _, raw := range vs {
		if raw != nil {
			labelSelectorRequirementData, _ := raw.(map[string]interface{})

			v3, ok := labelSelectorRequirementData[policy.OperatorKey]
			if !ok {
				errStrings = append(errStrings, fmt.Errorf("- operator: %v is not valid: minimum one valid operator attribute is required", v3).Error())
			}

			operator := v3.(string)
			values := make([]string, 0)

			v4, ok := labelSelectorRequirementData[policy.ValuesKey]
			if !ok {
				errStrings = append(errStrings, fmt.Errorf("- values: %v is not valid: minimum one valid values attribute is required", v4).Error())
			}

			vs1, ok := v4.([]interface{})
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
	}

	if len(errStrings) != 0 {
		return fmt.Errorf("error(s) in label selector requirement(s): \n%s", strings.Join(errStrings, "\n"))
	}

	return nil
}

func hasSpecChanged(d *schema.ResourceData) bool {
	updateRequired := false

	switch {
	case d.HasChange(helper.GetFirstElementOf(specKey, inputKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(specKey, policy.NamespaceSelectorKey)):
		updateRequired = true
	}

	return updateRequired
}
