/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindcustom

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
)

var SpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the custom policy",
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
	case TMCBlockNodeportServiceRecipe:
		if inputRecipeData.inputTMCBlockNodeportService != nil {
			spec.Input = *inputRecipeData.inputTMCBlockNodeportService
		}
	case TMCBlockResourcesRecipe:
		if inputRecipeData.inputTMCBlockResources != nil {
			spec.Input = *inputRecipeData.inputTMCBlockResources
		}
	case TMCBlockRolebindingSubjectsRecipe:
		if inputRecipeData.inputTMCBlockRolebindingSubjects != nil {
			spec.Input = *inputRecipeData.inputTMCBlockRolebindingSubjects
		}
	case TMCExternalIPSRecipe:
		if inputRecipeData.inputTMCExternalIps != nil {
			spec.Input = *inputRecipeData.inputTMCExternalIps
		}
	case TMCHTTPSIngressRecipe:
		if inputRecipeData.inputTMCHTTPSIngress != nil {
			spec.Input = *inputRecipeData.inputTMCHTTPSIngress
		}
	case TMCRequireLabelsRecipe:
		if inputRecipeData.inputTMCRequireLabels != nil {
			spec.Input = *inputRecipeData.inputTMCRequireLabels
		}
	case TMCCustomRecipe:
		spec.Recipe = inputRecipeData.recipeTMCCustom

		if inputRecipeData.inputTMCCustom != nil {
			spec.Input = *inputRecipeData.inputTMCCustom
		}
	case UnknownRecipe:
		fmt.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(RecipesAllowed[:], `, `))
	}

	if namespace, ok := specData[policy.NamespaceSelectorKey]; ok {
		if namespaceData, ok := namespace.([]interface{}); ok {
			spec.NamespaceSelector = policy.ConstructNamespaceSelector(namespaceData)
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
	case string(TMCBlockNodeportServiceRecipe):
		var tmcBlockNodeportServiceRecipeInput policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe

		err = tmcBlockNodeportServiceRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:                       TMCBlockNodeportServiceRecipe,
			inputTMCBlockNodeportService: &tmcBlockNodeportServiceRecipeInput,
		}
	case string(TMCBlockResourcesRecipe):
		var tmcBlockResourcesRecipeInput policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe

		err = tmcBlockResourcesRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:                 TMCBlockResourcesRecipe,
			inputTMCBlockResources: &tmcBlockResourcesRecipeInput,
		}
	case string(TMCBlockRolebindingSubjectsRecipe):
		var tmcBlockRolebindingSubjectsRecipeInput policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjects

		err = tmcBlockRolebindingSubjectsRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:                           TMCBlockRolebindingSubjectsRecipe,
			inputTMCBlockRolebindingSubjects: &tmcBlockRolebindingSubjectsRecipeInput,
		}
	case string(TMCExternalIPSRecipe):
		var tmcExternalIPSRecipeInput policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPS

		err = tmcExternalIPSRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:              TMCExternalIPSRecipe,
			inputTMCExternalIps: &tmcExternalIPSRecipeInput,
		}
	case string(TMCHTTPSIngressRecipe):
		var tmcHTTPSIngressRecipeInput policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe

		err = tmcHTTPSIngressRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:               TMCHTTPSIngressRecipe,
			inputTMCHTTPSIngress: &tmcHTTPSIngressRecipeInput,
		}
	case string(TMCRequireLabelsRecipe):
		var tmcRequireLabelsRecipeInput policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabels

		err = tmcRequireLabelsRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:                TMCRequireLabelsRecipe,
			inputTMCRequireLabels: &tmcRequireLabelsRecipeInput,
		}
	case string(UnknownRecipe):
		fmt.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(RecipesAllowed[:], `, `))
	default:
		var tmcCustomRecipeInput policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustom

		err = tmcCustomRecipeInput.UnmarshalBinary(byteSlice)

		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:          TMCCustomRecipe,
			recipeTMCCustom: spec.Recipe,
			inputTMCCustom:  &tmcCustomRecipeInput,
		}
	}

	flattenSpecData[policy.InputKey] = flattenInput(inputRecipeData)

	if spec.NamespaceSelector != nil {
		flattenSpecData[policy.NamespaceSelectorKey] = policy.FlattenNamespaceSelector(spec.NamespaceSelector)
	}

	return []interface{}{flattenSpecData}
}
