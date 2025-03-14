// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policykindquota

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipequotamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/quota"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/quota/recipe"
)

var (
	inputSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Input for the namespace quota policy, having one of the valid recipes: small, medium, large or custom.",
		Required:    true,
		MaxItems:    1,
		MinItems:    1,
		ForceNew:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				reciperesource.CustomKey: reciperesource.Custom,
				reciperesource.SmallKey:  reciperesource.Small,
				reciperesource.MediumKey: reciperesource.Medium,
				reciperesource.LargeKey:  reciperesource.Large,
			},
		},
	}
	RecipesAllowed = [...]string{reciperesource.CustomKey, reciperesource.SmallKey, reciperesource.MediumKey, reciperesource.LargeKey}
)

type (
	Recipe string
	// InputRecipe is a struct for all types of quota policy inputs.
	inputRecipe struct {
		recipe Recipe
		input  *policyrecipequotamodel.VmwareTanzuManageV1alpha1CommonPolicySpecQuotaV1Custom
	}
)

func constructInput(data []interface{}) (inputRecipeData *inputRecipe) {
	if len(data) == 0 || data[0] == nil {
		return inputRecipeData
	}

	inputData, _ := data[0].(map[string]interface{})

	if v, ok := inputData[reciperesource.CustomKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe: CustomRecipe,
				input:  reciperesource.ConstructCustom(v1),
			}
		}
	}

	if v, ok := inputData[reciperesource.SmallKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe: SmallRecipe,
			}
		}
	}

	if v, ok := inputData[reciperesource.MediumKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe: MediumRecipe,
			}
		}
	}

	if v, ok := inputData[reciperesource.LargeKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe: LargeRecipe,
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
	case CustomRecipe:
		flattenInputData[reciperesource.CustomKey] = reciperesource.FlattenCustom(inputRecipeData.input)
	case SmallRecipe:
		flattenInputData[reciperesource.SmallKey] = []interface{}{make(map[string]interface{})}
	case MediumRecipe:
		flattenInputData[reciperesource.MediumKey] = []interface{}{make(map[string]interface{})}
	case LargeRecipe:
		flattenInputData[reciperesource.LargeKey] = []interface{}{make(map[string]interface{})}
	case UnknownRecipe:
		fmt.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(RecipesAllowed[:], `, `))
	}

	return []interface{}{flattenInputData}
}

func ValidateInput(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
	value, ok := diff.GetOk(policy.SpecKey)
	if !ok {
		return fmt.Errorf("spec: %v is not valid: minimum one valid spec block is required", value)
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("spec data: %v is not valid: minimum one valid spec block is required among: %v", data, strings.Join(RecipesAllowed[:], `, `))
	}

	specData := data[0].(map[string]interface{})

	v, ok := specData[policy.InputKey]
	if !ok {
		return fmt.Errorf("input: %v is not valid: minimum one valid input block is required", v)
	}

	v1, ok := v.([]interface{})
	if !ok {
		return fmt.Errorf("type of input block data: %v is not valid", v1)
	}

	if len(v1) == 0 || v1[0] == nil {
		return fmt.Errorf("input data: %v is not valid: minimum one valid input block is required", v1)
	}

	inputData, _ := v1[0].(map[string]interface{})
	recipesFound := make([]string, 0)

	if v, ok := inputData[reciperesource.CustomKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.CustomKey)
		}
	}

	if v, ok := inputData[reciperesource.SmallKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.SmallKey)
		}
	}

	if v, ok := inputData[reciperesource.MediumKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.MediumKey)
		}
	}

	if v, ok := inputData[reciperesource.LargeKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.LargeKey)
		}
	}

	if len(recipesFound) == 0 {
		return fmt.Errorf("no valid input recipe block found: minimum one valid input recipe block is required among: %v", strings.Join(RecipesAllowed[:], `, `))
	} else if len(recipesFound) > 1 {
		return fmt.Errorf("found input recipes: %v are not valid: maximum one valid input recipe block is allowed", strings.Join(recipesFound, `, `))
	}

	return nil
}
