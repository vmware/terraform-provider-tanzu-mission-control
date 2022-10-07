/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package security

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
			policy.InputKey:             inputSchema,
			policy.NamespaceSelectorKey: policy.NamespaceSelector,
		},
	},
}

func constructSpec(d *schema.ResourceData) (spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec) {
	value, ok := d.GetOk(policy.SpecKey)
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
		RecipeVersion: policy.RecipeVersionDefaultValue,
	}

	v, ok := specData[policy.InputKey]
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

	flattenSpecData[policy.InputKey] = flattenInput(inputRecipeData)

	if spec.NamespaceSelector != nil {
		flattenSpecData[policy.NamespaceSelectorKey] = policy.FlattenNamespaceSelector(spec.NamespaceSelector)
	}

	return []interface{}{flattenSpecData}
}
