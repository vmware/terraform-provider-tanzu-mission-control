/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindimage

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipeimagemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/image/recipe"
)

var (
	inputSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Input for the image registry policy, having one of the valid recipes: allowed-name-tag, custom, block-latest-tag or require-digest.",
		Required:    true,
		MaxItems:    1,
		MinItems:    1,
		ForceNew:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				reciperesource.AllowedNameTagKey: reciperesource.AllowedNameTag,
				reciperesource.CustomKey:         reciperesource.Custom,
				reciperesource.BlockLatestTagKey: reciperesource.BlockLatestTag,
				reciperesource.RequireDigestKey:  reciperesource.RequireDigest,
			},
		},
	}
	RecipesAllowed = [...]string{reciperesource.AllowedNameTagKey, reciperesource.CustomKey, reciperesource.BlockLatestTagKey, reciperesource.RequireDigestKey}
)

type (
	Recipe string
	// InputRecipe is a struct for all types of image registry policy inputs.
	inputRecipe struct {
		recipe              Recipe
		inputAllowedNameTag *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag
		inputCustom         *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1Custom
		inputBlockLatestTag *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CommonRecipe
		inputRequireDigest  *policyrecipeimagemodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1CommonRecipe
	}
)

func constructInput(data []interface{}) (inputRecipeData *inputRecipe) {
	if len(data) == 0 || data[0] == nil {
		return inputRecipeData
	}

	inputData, _ := data[0].(map[string]interface{})

	if v, ok := inputData[reciperesource.AllowedNameTagKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:              AllowedNameTagRecipe,
				inputAllowedNameTag: reciperesource.ConstructAllowedNameTag(v1),
			}
		}
	}

	if v, ok := inputData[reciperesource.CustomKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:      CustomRecipe,
				inputCustom: reciperesource.ConstructCustom(v1),
			}
		}
	}

	if v, ok := inputData[reciperesource.BlockLatestTagKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:              BlockLatestTagRecipe,
				inputBlockLatestTag: reciperesource.ConstructCommonRecipe(v1),
			}
		}
	}

	if v, ok := inputData[reciperesource.RequireDigestKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:             RequireDigestRecipe,
				inputRequireDigest: reciperesource.ConstructCommonRecipe(v1),
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
	case AllowedNameTagRecipe:
		flattenInputData[reciperesource.AllowedNameTagKey] = reciperesource.FlattenAllowedNameTag(inputRecipeData.inputAllowedNameTag)
	case CustomRecipe:
		flattenInputData[reciperesource.CustomKey] = reciperesource.FlattenCustom(inputRecipeData.inputCustom)
	case BlockLatestTagRecipe:
		flattenInputData[reciperesource.BlockLatestTagKey] = reciperesource.FlattenCommonRecipe(inputRecipeData.inputBlockLatestTag)
	case RequireDigestRecipe:
		flattenInputData[reciperesource.RequireDigestKey] = reciperesource.FlattenCommonRecipe(inputRecipeData.inputRequireDigest)
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

	if v, ok := inputData[reciperesource.AllowedNameTagKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.AllowedNameTagKey)
		}
	}

	if v, ok := inputData[reciperesource.CustomKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.CustomKey)
		}
	}

	if v, ok := inputData[reciperesource.BlockLatestTagKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.BlockLatestTagKey)
		}
	}

	if v, ok := inputData[reciperesource.RequireDigestKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.RequireDigestKey)
		}
	}

	if len(recipesFound) == 0 {
		return fmt.Errorf("no valid input recipe block found: minimum one valid input recipe block is required among: %v", strings.Join(RecipesAllowed[:], `, `))
	} else if len(recipesFound) > 1 {
		return fmt.Errorf("found input recipes: %v are not valid: maximum one valid input recipe block is allowed", strings.Join(recipesFound, `, `))
	}

	return nil
}
