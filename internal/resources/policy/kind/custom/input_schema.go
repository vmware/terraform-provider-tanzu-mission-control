/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindcustom

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/custom/recipe"
)

var (
	inputSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Input for the custom policy, having one of the valid recipes: tmc_block_nodeport_service, tmc_block_resources, tmc_block_rolebinding_subjects, tmc_external_ips, tmc_https_ingress or tmc_require_labels.",
		Required:    true,
		MaxItems:    1,
		MinItems:    1,
		ForceNew:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				reciperesource.TMCBlockNodeportServiceKey:     reciperesource.TMCBlockNodeportService,
				reciperesource.TMCBlockResourcesKey:           reciperesource.TMCBlockResources,
				reciperesource.TMCBlockRolebindingSubjectsKey: reciperesource.TMCBlockRolebindingSubjects,
				reciperesource.TMCExternalIPSKey:              reciperesource.TMCExternalIps,
				reciperesource.TMCHTTPSIngressKey:             reciperesource.TMCHTTPSIngress,
				reciperesource.TMCRequireLabelsKey:            reciperesource.TMCRequireLabels,
			},
		},
	}
	RecipesAllowed = [...]string{reciperesource.TMCBlockNodeportServiceKey, reciperesource.TMCBlockResourcesKey, reciperesource.TMCBlockRolebindingSubjectsKey, reciperesource.TMCExternalIPSKey, reciperesource.TMCHTTPSIngressKey, reciperesource.TMCRequireLabelsKey}
)

type (
	Recipe string
	// InputRecipe is a struct for all types of custom policy inputs.
	inputRecipe struct {
		recipe                           Recipe
		inputTMCBlockNodeportService     *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe
		inputTMCBlockResources           *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe
		inputTMCBlockRolebindingSubjects *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjects
		inputTMCExternalIps              *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPS
		inputTMCHTTPSIngress             *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe
		inputTMCRequireLabels            *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabels
	}
)

func constructInput(data []interface{}) (inputRecipeData *inputRecipe) {
	if len(data) == 0 || data[0] == nil {
		return inputRecipeData
	}

	inputData, _ := data[0].(map[string]interface{})

	if input, ok := inputData[reciperesource.TMCBlockNodeportServiceKey]; ok {
		if recipeType, ok := input.([]interface{}); ok && len(recipeType) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:                       TMCBlockNodeportServiceRecipe,
				inputTMCBlockNodeportService: reciperesource.ConstructTMCCommonRecipe(recipeType),
			}
		}
	}

	if input, ok := inputData[reciperesource.TMCBlockResourcesKey]; ok {
		if recipeType, ok := input.([]interface{}); ok && len(recipeType) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:                 TMCBlockResourcesRecipe,
				inputTMCBlockResources: reciperesource.ConstructTMCCommonRecipe(recipeType),
			}
		}
	}

	if input, ok := inputData[reciperesource.TMCBlockRolebindingSubjectsKey]; ok {
		if recipeType, ok := input.([]interface{}); ok && len(recipeType) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:                           TMCBlockRolebindingSubjectsRecipe,
				inputTMCBlockRolebindingSubjects: reciperesource.ConstructTMCBlockRolebindingSubjects(recipeType),
			}
		}
	}

	if input, ok := inputData[reciperesource.TMCExternalIPSKey]; ok {
		if recipeType, ok := input.([]interface{}); ok && len(recipeType) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:              TMCExternalIPSRecipe,
				inputTMCExternalIps: reciperesource.ConstructTMCExternalIPS(recipeType),
			}
		}
	}

	if input, ok := inputData[reciperesource.TMCHTTPSIngressKey]; ok {
		if recipeType, ok := input.([]interface{}); ok && len(recipeType) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:               TMCHTTPSIngressRecipe,
				inputTMCHTTPSIngress: reciperesource.ConstructTMCCommonRecipe(recipeType),
			}
		}
	}

	if input, ok := inputData[reciperesource.TMCRequireLabelsKey]; ok {
		if recipeType, ok := input.([]interface{}); ok && len(recipeType) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:                TMCRequireLabelsRecipe,
				inputTMCRequireLabels: reciperesource.ConstructTMCRequireLabels(recipeType),
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
	case TMCBlockNodeportServiceRecipe:
		flattenInputData[reciperesource.TMCBlockNodeportServiceKey] = reciperesource.FlattenTMCCommonRecipe(inputRecipeData.inputTMCBlockNodeportService)
	case TMCBlockResourcesRecipe:
		flattenInputData[reciperesource.TMCBlockResourcesKey] = reciperesource.FlattenTMCCommonRecipe(inputRecipeData.inputTMCBlockResources)
	case TMCBlockRolebindingSubjectsRecipe:
		flattenInputData[reciperesource.TMCBlockRolebindingSubjectsKey] = reciperesource.FlattenTMCBlockRolebindingSubjects(inputRecipeData.inputTMCBlockRolebindingSubjects)
	case TMCExternalIPSRecipe:
		flattenInputData[reciperesource.TMCExternalIPSKey] = reciperesource.FlattenTMCExternalIPS(inputRecipeData.inputTMCExternalIps)
	case TMCHTTPSIngressRecipe:
		flattenInputData[reciperesource.TMCHTTPSIngressKey] = reciperesource.FlattenTMCCommonRecipe(inputRecipeData.inputTMCHTTPSIngress)
	case TMCRequireLabelsRecipe:
		flattenInputData[reciperesource.TMCRequireLabelsKey] = reciperesource.FlattenTMCRequireLabels(inputRecipeData.inputTMCRequireLabels)

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
	recipesFound := make([]string, 0)

	if recipeData, ok := inputData[reciperesource.TMCBlockNodeportServiceKey]; ok {
		if recipeType, ok := recipeData.([]interface{}); ok && len(recipeType) != 0 {
			recipesFound = append(recipesFound, reciperesource.TMCBlockNodeportServiceKey)
		}
	}

	if recipeData, ok := inputData[reciperesource.TMCBlockResourcesKey]; ok {
		if recipeType, ok := recipeData.([]interface{}); ok && len(recipeType) != 0 {
			recipesFound = append(recipesFound, reciperesource.TMCBlockResourcesKey)
		}
	}

	if recipeData, ok := inputData[reciperesource.TMCBlockRolebindingSubjectsKey]; ok {
		if recipeType, ok := recipeData.([]interface{}); ok && len(recipeType) != 0 {
			recipesFound = append(recipesFound, reciperesource.TMCBlockRolebindingSubjectsKey)
		}
	}

	if recipeData, ok := inputData[reciperesource.TMCExternalIPSKey]; ok {
		if recipeType, ok := recipeData.([]interface{}); ok && len(recipeType) != 0 {
			recipesFound = append(recipesFound, reciperesource.TMCExternalIPSKey)
		}
	}

	if recipeData, ok := inputData[reciperesource.TMCHTTPSIngressKey]; ok {
		if recipeType, ok := recipeData.([]interface{}); ok && len(recipeType) != 0 {
			recipesFound = append(recipesFound, reciperesource.TMCHTTPSIngressKey)
		}
	}

	if recipeData, ok := inputData[reciperesource.TMCRequireLabelsKey]; ok {
		if recipeType, ok := recipeData.([]interface{}); ok && len(recipeType) != 0 {
			recipesFound = append(recipesFound, reciperesource.TMCRequireLabelsKey)
		}
	}

	if len(recipesFound) == 0 {
		return fmt.Errorf("no valid input recipe block found: minimum one valid input recipe block is required among: %v", strings.Join(RecipesAllowed[:], `, `))
	} else if len(recipesFound) > 1 {
		return fmt.Errorf("found input recipes: %v are not valid: maximum one valid input recipe block is allowed", strings.Join(recipesFound, `, `))
	}

	return nil
}
