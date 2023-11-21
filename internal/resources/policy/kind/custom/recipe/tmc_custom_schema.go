/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	openapiv3 "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/openapi_v3_schema_validator"
	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
	recipemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/recipe"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/common"
)

var TMCCustomSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for custom policy tmc_external_ips recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			TemplateNameKey: {
				Type:        schema.TypeString,
				Description: "Name of custom template.",
				Required:    true,
			},
			AuditKey: {
				Type:        schema.TypeBool,
				Description: "Audit (dry-run).",
				Optional:    true,
				Default:     false,
			},
			ParametersKey: {
				Type:        schema.TypeString,
				Description: "JSON encoded template parameters.",
				Optional:    true,
			},
			TargetKubernetesResourcesKey: common.TargetKubernetesResourcesSchema,
		},
	},
}

func FlattenTMCCustom(recipeName string, customRecipe *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustom) []interface{} {
	if customRecipe == nil {
		return nil
	}

	customInputMap := make(map[string]interface{})
	customInputMap[AuditKey] = customRecipe.Audit
	customInputMap[TemplateNameKey] = recipeName

	if customRecipe.Parameters != nil {
		parametersJSONBytes, _ := json.Marshal(customRecipe.Parameters)
		customInputMap[ParametersKey] = helper.ConvertToString(parametersJSONBytes, "")
	}

	targetKubernetesResources := make([]interface{}, 0)

	for _, tkr := range customRecipe.TargetKubernetesResources {
		targetKubernetesResources = append(targetKubernetesResources, common.FlattenTargetKubernetesResources(tkr))
	}

	customInputMap[TargetKubernetesResourcesKey] = targetKubernetesResources

	return []interface{}{customInputMap}
}

func ConstructTMCCustom(customRecipe []interface{}) (customInputModel *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustom) {
	if len(customRecipe) != 0 && customRecipe[0] != nil {
		customInputMap := customRecipe[0].(map[string]interface{})

		customInputModel = &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustom{
			Audit: customInputMap[AuditKey].(bool),
		}

		parametersData := customInputMap[ParametersKey].(string)

		if parametersData != "" {
			parametersJSON := make(map[string]interface{})

			_ = json.Unmarshal([]byte(parametersData), &parametersJSON)

			customInputModel.Parameters = parametersJSON
		}

		targetKubernetesResourcesData := customInputMap[TargetKubernetesResourcesKey].([]interface{})
		customInputModel.TargetKubernetesResources = make([]*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources, 0)

		for _, targetKubernetesResource := range targetKubernetesResourcesData {
			customInputModel.TargetKubernetesResources = append(customInputModel.TargetKubernetesResources, common.ExpandTargetKubernetesResources(targetKubernetesResource))
		}
	}

	return customInputModel
}

func ValidateCustomRecipe(config authctx.TanzuContext, customRecipe map[string]interface{}) error {
	errMessages := make([]string, 0)
	customTemplateName := customRecipe[TemplateNameKey].(string)

	recipeModel := &recipemodels.VmwareTanzuManageV1alpha1PolicyTypeRecipeFullName{
		TypeName: "custom-policy",
		Name:     customTemplateName,
	}

	recipeData, err := config.TMCConnection.RecipeResourceService.RecipeResourceServiceGet(recipeModel)

	if err != nil {
		errMessages = append(errMessages, err.Error())
	} else {
		errs := ValidateRecipeParameters(recipeData.Recipe.Spec.InputSchema, customRecipe[ParametersKey].(string))

		if len(errs) > 0 {
			errMsg := ""

			for _, e := range errs {
				if errMsg == "" {
					errMsg = e.Error()
				} else {
					errMsg = fmt.Sprintf("%s\n%s", errMsg, e.Error())
				}
			}

			errMessages = append(errMessages, errMsg)
		}
	}

	if len(errMessages) > 0 {
		errMsg := strings.Join(errMessages, "\n")

		return errors.New(errMsg)
	}

	return nil
}

func ValidateRecipeParameters(recipeSchema string, recipeParameters string) (errs []error) {
	recipeSchemaJSON := make(map[string]interface{})
	_ = json.Unmarshal([]byte(recipeSchema), &recipeSchemaJSON)

	recipeParametersSchema, parametersSchemaExist := recipeSchemaJSON["properties"].(map[string]interface{})["parameters"]

	if parametersSchemaExist {
		openAPIV3Validator := &openapiv3.OpenAPIV3SchemaValidator{
			Schema: recipeParametersSchema.(map[string]interface{})["properties"].(map[string]interface{}),
		}

		recipeParametersJSON := make(map[string]interface{})
		_ = json.Unmarshal([]byte(recipeParameters), &recipeParametersJSON)
		errs = make([]error, 0)

		errs = append(errs, openAPIV3Validator.ValidateRequiredFields(recipeParametersJSON)...)
		errs = append(errs, openAPIV3Validator.ValidateFormat(recipeParametersJSON)...)
	}

	return errs
}
