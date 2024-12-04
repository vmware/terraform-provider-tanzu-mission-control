// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policykindnetwork

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
)

var SpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the network policy",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			policy.InputKey:             inputSchema,
			policy.NamespaceSelectorKey: policy.NamespaceSelector,
		},
	},
}

func ConstructSpec(d *schema.ResourceData) (spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec) {
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
		Type:          typePolicy,
		RecipeVersion: policy.RecipeVersionDefaultValue,
	}

	input, ok := specData[policy.InputKey]
	if !ok {
		return spec
	}

	inputData, ok := input.([]interface{})
	if !ok {
		return spec
	}

	inputRecipeData := constructInput(inputData)

	if inputRecipeData == nil || inputRecipeData.recipe == "" {
		return spec
	}

	spec.Recipe = strings.ReplaceAll(string(inputRecipeData.recipe), "_", "-")

	switch inputRecipeData.recipe {
	case AllowAllRecipe:
		if inputRecipeData.inputAllowAll != nil {
			spec.Input = *inputRecipeData.inputAllowAll
		}
	case AllowAllToPodsRecipe:
		if inputRecipeData.inputAllowAllToPods != nil {
			spec.Input = *inputRecipeData.inputAllowAllToPods
		}
	case DenyAllToPodsRecipe:
		if inputRecipeData.inputDenyAllToPods != nil {
			spec.Input = *inputRecipeData.inputDenyAllToPods
		}
	case CustomEgressRecipe:
		if inputRecipeData.inputCustomEgress != nil {
			spec.Input = *inputRecipeData.inputCustomEgress
		}
	case CustomIngressRecipe:
		if inputRecipeData.inputCustomIngress != nil {
			spec.Input = *inputRecipeData.inputCustomIngress
		}
	case UnknownRecipe:
		fmt.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(RecipesAllowed[:], `, `))
	}

	if v, ok := specData[policy.NamespaceSelectorKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.NamespaceSelector = policy.ConstructNamespaceSelector(v1)
		}
	}

	return spec
}

func FlattenSpec(spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	if spec.Input == nil {
		return data
	}

	input, ok := spec.Input.(map[string]interface{})
	if !ok {
		return data
	}

	var inputRecipeData *inputRecipe

	byteSlice, err := json.Marshal(input)
	if err != nil {
		return data
	}

	switch strings.ReplaceAll(spec.Recipe, "-", "_") {
	case string(AllowAllRecipe):
		var allowAllRecipeInput policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAll

		err = allowAllRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:        AllowAllRecipe,
			inputAllowAll: &allowAllRecipeInput,
		}
	case string(AllowAllToPodsRecipe):
		var allowAllToPodsRecipeInput policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAllToPods

		err = allowAllToPodsRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:              AllowAllToPodsRecipe,
			inputAllowAllToPods: &allowAllToPodsRecipeInput,
		}
	case string(AllowAllEgressRecipe):
		inputRecipeData = &inputRecipe{
			recipe: AllowAllEgressRecipe,
		}
	case string(DenyAllRecipe):
		inputRecipeData = &inputRecipe{
			recipe: DenyAllRecipe,
		}
	case string(DenyAllToPodsRecipe):
		var denyAllToPodsRecipeInput policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1DenyAllToPods

		err = denyAllToPodsRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:             DenyAllToPodsRecipe,
			inputDenyAllToPods: &denyAllToPodsRecipeInput,
		}
	case string(DenyAllEgressRecipe):
		inputRecipeData = &inputRecipe{
			recipe: DenyAllEgressRecipe,
		}
	case string(CustomEgressRecipe):
		var customEgressRecipeInput policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomEgress

		err = customEgressRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:            CustomEgressRecipe,
			inputCustomEgress: &customEgressRecipeInput,
		}
	case string(CustomIngressRecipe):
		var customIngressRecipeInput policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomIngress

		err = customIngressRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:             CustomIngressRecipe,
			inputCustomIngress: &customIngressRecipeInput,
		}
	case string(UnknownRecipe):
		fmt.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(RecipesAllowed[:], `, `))
	}

	flattenSpecData[policy.InputKey] = flattenInput(inputRecipeData)

	if spec.NamespaceSelector != nil {
		flattenSpecData[policy.NamespaceSelectorKey] = policy.FlattenNamespaceSelector(spec.NamespaceSelector)
	}

	return []interface{}{flattenSpecData}
}
