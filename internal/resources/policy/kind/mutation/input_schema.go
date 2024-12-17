// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package policykindmutation

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipemutationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/mutation/recipe"
)

var (
	inputSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Input for the mutation policy.",
		Required:    true,
		MaxItems:    1,
		MinItems:    1,
		ForceNew:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				reciperesource.PodSecurityKey: reciperesource.PodSecuritySchema,
				reciperesource.LabelKey:       reciperesource.LabelSchema,
				reciperesource.AnnotationKey:  reciperesource.AnnotationSchema,
			},
		},
	}
	RecipesAllowed = [...]string{reciperesource.PodSecurityKey, reciperesource.LabelKey, reciperesource.AnnotationKey}
)

type (
	Recipe      string
	inputRecipe struct {
		recipe      Recipe
		podSecurity *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1PodSecurity
		label       *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label
		annotation  *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation
	}
)

func constructInput(data []interface{}) (inputRecipeData *inputRecipe) {
	if len(data) == 0 || data[0] == nil {
		return inputRecipeData
	}

	inputData, _ := data[0].(map[string]interface{})

	for recipeKey, input := range inputData {
		switch recipeKey {
		case reciperesource.PodSecurityKey:
			if ir, ok := input.([]interface{}); ok && len(ir) != 0 {
				inputRecipeData = &inputRecipe{
					recipe:      PodSecurityRecipe,
					podSecurity: reciperesource.ConstructPodSecurity(ir),
				}
			}
		case reciperesource.LabelKey:
			if ir, ok := input.([]interface{}); ok && len(ir) != 0 {
				inputRecipeData = &inputRecipe{
					recipe: LabelRecipe,
					label:  reciperesource.ConstructLabel(ir),
				}
			}
		case reciperesource.AnnotationKey:
			if ir, ok := input.([]interface{}); ok && len(ir) != 0 {
				inputRecipeData = &inputRecipe{
					recipe:     AnnotationRecipe,
					annotation: reciperesource.ConstructAnnotation(ir),
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
	case PodSecurityRecipe:
		flattenInputData[reciperesource.PodSecurityKey] = reciperesource.FlattenPodSecurity(inputRecipeData.podSecurity)
	case LabelRecipe:
		flattenInputData[reciperesource.LabelKey] = reciperesource.FlattenLabel(inputRecipeData.label)
	case AnnotationRecipe:
		flattenInputData[reciperesource.AnnotationKey] = reciperesource.FlattenAnnotation(inputRecipeData.annotation)
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
	recipeKeys := []string{reciperesource.PodSecurityKey, reciperesource.LabelKey, reciperesource.AnnotationKey}

	for _, recipeKey := range recipeKeys {
		if recipeData, ok := inputData[recipeKey]; ok {
			if recipeType, ok := recipeData.([]interface{}); ok && len(recipeType) != 0 {
				recipesFound = append(recipesFound, recipeKey)
			}
		}
	}

	return recipesFound
}
