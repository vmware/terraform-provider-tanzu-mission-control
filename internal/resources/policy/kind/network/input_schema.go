// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policykindnetwork

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network/recipe"
)

var (
	inputSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Input for the network policy, having one of the valid recipes: allow-all, allow-all-to-pods, allow-all-egress, deny-all, deny-all-to-pods, deny-all-egress, custom-egress or custom-ingress.",
		Required:    true,
		MaxItems:    1,
		MinItems:    1,
		ForceNew:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				reciperesource.AllowAllKey:       reciperesource.AllowAll,
				reciperesource.AllowAllToPodsKey: reciperesource.AllowAllToPods,
				reciperesource.AllowAllEgressKey: reciperesource.AllowAllEgress,
				reciperesource.DenyAllKey:        reciperesource.DenyAll,
				reciperesource.DenyAllToPodsKey:  reciperesource.DenyAllToPods,
				reciperesource.DenyAllEgressKey:  reciperesource.DenyAllEgress,
				reciperesource.CustomEgressKey:   reciperesource.CustomEgress,
				reciperesource.CustomIngressKey:  reciperesource.CustomIngress,
			},
		},
	}
	RecipesAllowed = [...]string{reciperesource.AllowAllKey, reciperesource.AllowAllToPodsKey, reciperesource.AllowAllEgressKey, reciperesource.DenyAllKey, reciperesource.DenyAllToPodsKey, reciperesource.DenyAllEgressKey, reciperesource.CustomEgressKey, reciperesource.CustomIngressKey}
)

type (
	Recipe string
	// InputRecipe is a struct for all types of network policy inputs.
	inputRecipe struct {
		recipe              Recipe
		inputAllowAll       *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAll
		inputAllowAllToPods *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAllToPods
		inputDenyAllToPods  *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1DenyAllToPods
		inputCustomEgress   *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomEgress
		inputCustomIngress  *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomIngress
	}
)

func constructInput(data []interface{}) (inputRecipeData *inputRecipe) {
	if len(data) == 0 || data[0] == nil {
		return inputRecipeData
	}

	inputData, _ := data[0].(map[string]interface{})

	for recipeKey, input := range inputData {
		switch recipeKey {
		case reciperesource.AllowAllKey:
			if ir, ok := input.([]interface{}); ok && len(ir) != 0 {
				inputRecipeData = &inputRecipe{
					recipe:        AllowAllRecipe,
					inputAllowAll: reciperesource.ConstructAllowAll(ir),
				}
			}
		case reciperesource.AllowAllToPodsKey:
			if ir, ok := input.([]interface{}); ok && len(ir) != 0 {
				inputRecipeData = &inputRecipe{
					recipe:              AllowAllToPodsRecipe,
					inputAllowAllToPods: reciperesource.ConstructAllowAllToPods(ir),
				}
			}
		case reciperesource.AllowAllEgressKey:
			if ir, ok := input.([]interface{}); ok && len(ir) != 0 {
				inputRecipeData = &inputRecipe{
					recipe: AllowAllEgressRecipe,
				}
			}
		case reciperesource.DenyAllKey:
			if ir, ok := input.([]interface{}); ok && len(ir) != 0 {
				inputRecipeData = &inputRecipe{
					recipe: DenyAllRecipe,
				}
			}
		case reciperesource.DenyAllToPodsKey:
			if ir, ok := input.([]interface{}); ok && len(ir) != 0 {
				inputRecipeData = &inputRecipe{
					recipe:             DenyAllToPodsRecipe,
					inputDenyAllToPods: reciperesource.ConstructDenyAllToPods(ir),
				}
			}
		case reciperesource.DenyAllEgressKey:
			if ir, ok := input.([]interface{}); ok && len(ir) != 0 {
				inputRecipeData = &inputRecipe{
					recipe: DenyAllEgressRecipe,
				}
			}
		case reciperesource.CustomEgressKey:
			if ir, ok := input.([]interface{}); ok && len(ir) != 0 {
				inputRecipeData = &inputRecipe{
					recipe:            CustomEgressRecipe,
					inputCustomEgress: reciperesource.ConstructCustomEgress(ir),
				}
			}
		case reciperesource.CustomIngressKey:
			if ir, ok := input.([]interface{}); ok && len(ir) != 0 {
				inputRecipeData = &inputRecipe{
					recipe:             CustomIngressRecipe,
					inputCustomIngress: reciperesource.ConstructCustomIngress(ir),
				}
			}
		}
	}

	return inputRecipeData
}

func flattenInput(inputRecipeData *inputRecipe) (data []interface{}) {
	if inputRecipeData == nil {
		return data
	}

	flattenInputData := make(map[string]interface{})

	switch inputRecipeData.recipe {
	case AllowAllRecipe:
		flattenInputData[reciperesource.AllowAllKey] = reciperesource.FlattenAllowAll(inputRecipeData.inputAllowAll)
	case AllowAllToPodsRecipe:
		flattenInputData[reciperesource.AllowAllToPodsKey] = reciperesource.FlattenAllowAllToPods(inputRecipeData.inputAllowAllToPods)
	case AllowAllEgressRecipe:
		flattenInputData[reciperesource.AllowAllEgressKey] = []interface{}{make(map[string]interface{})}
	case DenyAllRecipe:
		flattenInputData[reciperesource.DenyAllKey] = []interface{}{make(map[string]interface{})}
	case DenyAllToPodsRecipe:
		flattenInputData[reciperesource.DenyAllToPodsKey] = reciperesource.FlattenDenyAllToPods(inputRecipeData.inputDenyAllToPods)
	case DenyAllEgressRecipe:
		flattenInputData[reciperesource.DenyAllEgressKey] = []interface{}{make(map[string]interface{})}
	case CustomEgressRecipe:
		flattenInputData[reciperesource.CustomEgressKey] = reciperesource.FlattenCustomEgress(inputRecipeData.inputCustomEgress)
	case CustomIngressRecipe:
		flattenInputData[reciperesource.CustomIngressKey] = reciperesource.FlattenCustomIngress(inputRecipeData.inputCustomIngress)
	case UnknownRecipe:
		fmt.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(RecipesAllowed[:], `, `))
	}

	return []interface{}{flattenInputData}
}

func ValidateInput(_ context.Context, diff *schema.ResourceDiff, _ interface{}) error {
	value, ok := diff.GetOk(policy.SpecKey)
	if !ok {
		return fmt.Errorf("spec: %v is not valid: minimum one valid spec block is required", value)
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("spec data: %v is not valid: minimum one valid spec block is required among: %v", data, strings.Join(RecipesAllowed[:], `, `))
	}

	specData := data[0].(map[string]interface{})

	input, ok := specData[policy.InputKey]
	if !ok {
		return fmt.Errorf("input: %v is not valid: minimum one valid input block is required", input)
	}

	inputType, ok := input.([]interface{})
	if !ok {
		return fmt.Errorf("type of input block data: %v is not valid", inputType)
	}

	if len(inputType) == 0 || inputType[0] == nil {
		return fmt.Errorf("input data: %v is not valid: minimum one valid input block is required", inputType)
	}

	inputData, _ := inputType[0].(map[string]interface{})

	recipesFound := appendRecipeFromInput(inputData)
	numberOfRecipes := len(recipesFound)

	if numberOfRecipes == 0 {
		return fmt.Errorf("no valid input recipe block found: minimum one valid input recipe block is required among: %v", strings.Join(RecipesAllowed[:], `, `))
	} else if numberOfRecipes > 1 {
		return fmt.Errorf("found input recipes: %v are not valid: maximum one valid input recipe block is allowed", strings.Join(recipesFound, `, `))
	}

	return nil
}

func appendRecipeFromInput(inputData map[string]interface{}) (recipesFound []string) {
	recipeKeys := []string{
		reciperesource.AllowAllKey,
		reciperesource.AllowAllToPodsKey,
		reciperesource.AllowAllEgressKey,
		reciperesource.DenyAllKey,
		reciperesource.DenyAllToPodsKey,
		reciperesource.DenyAllEgressKey,
		reciperesource.CustomEgressKey,
		reciperesource.CustomIngressKey,
	}

	for _, recipeKey := range recipeKeys {
		if recipeData, ok := inputData[recipeKey]; ok {
			if recipeType, ok := recipeData.([]interface{}); ok && len(recipeType) != 0 {
				recipesFound = append(recipesFound, recipeKey)
			}
		}
	}

	return recipesFound
}
