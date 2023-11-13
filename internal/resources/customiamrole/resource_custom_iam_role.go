/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrole

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	customiamrolemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/customiamrole"
)

func ResourceCustomIAMRole() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCustomIAMRoleCreate,
		UpdateContext: resourceCustomIAMRoleUpdate,
		ReadContext:   resourceCustomIAMRoleRead,
		DeleteContext: resourceCustomIAMRoleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCustomIAMRoleImporter,
		},
		CustomizeDiff: validateSchema,
		Schema:        customIAMRoleResourceSchema,
	}
}

func resourceCustomIAMRoleCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create Custom IAM Role."))
	}

	request := &customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData{
		Role: model,
	}

	_, err = config.TMCConnection.CustomIAMRoleResourceService.CustomIAMRoleResourceServiceCreate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't create Custom IAM Role.\nName: %s", model.FullName.Name))
	}

	return resourceCustomIAMRoleRead(helper.GetContextWithCaller(ctx, helper.CreateState), data, m)
}

func resourceCustomIAMRoleUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update Custom IAM Role."))
	}

	request := &customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData{
		Role: model,
	}

	_, err = config.TMCConnection.CustomIAMRoleResourceService.CustomIAMRoleResourceServiceUpdate(request)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't update Custom IAM Role.\nName: %s", model.FullName.Name))
	}

	return resourceCustomIAMRoleRead(helper.GetContextWithCaller(ctx, helper.UpdateState), data, m)
}

func resourceCustomIAMRoleRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	var resp *customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleData

	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't read Custom IAM Role."))
	}

	customIAMRoleFn := model.FullName
	resp, err = config.TMCConnection.CustomIAMRoleResourceService.CustomIAMRoleResourceServiceGet(customIAMRoleFn)

	if err != nil {
		if clienterrors.IsNotFoundError(err) {
			if !helper.IsContextCallerSet(ctx) {
				*data = schema.ResourceData{}

				return diags
			} else if helper.IsDeleteState(ctx) {
				// d.SetId("") is automatically called assuming delete returns no errors, but
				// it is added here for explicitness.
				_ = schema.RemoveFromState(data, m)

				return diags
			}
		}

		return diag.FromErr(errors.Wrapf(err, "Couldn't read Custom IAM Role.\nName: %s", customIAMRoleFn.Name))
	} else if resp != nil {
		oldSpecData := data.Get(SpecKey).([]interface{})[0].(map[string]interface{})
		aggregationRuleData, aggregationRuleExists := oldSpecData[AggregationRuleKey]
		allowedScopesData, allowedScopesExist := oldSpecData[AllowedScopesKey]
		err = tfModelConverter.FillTFSchema(resp.Role, data)

		if err != nil {
			return diag.FromErr(errors.Wrapf(err, "Couldn't read Custom IAM Role.\nName: %s", customIAMRoleFn.Name))
		}

		// API Inconsistency Fix
		if aggregationRuleExists || allowedScopesExist {
			newSpecData := data.Get(SpecKey).([]interface{})[0].(map[string]interface{})

			if aggregationRuleExists && len(aggregationRuleData.([]interface{})) > 0 {
				newSpecData[AggregationRuleKey] = formatAggregationRuleData(aggregationRuleData.([]interface{}), resp.Role.Spec.AggregationRule)
			}

			newSpecData[AllowedScopesKey] = formatResourcesData(allowedScopesData.([]interface{}), resp.Role.Spec.Resources)
			_ = data.Set(SpecKey, []interface{}{newSpecData})
		}

		data.SetId(customIAMRoleFn.Name)
	}

	return diags
}

func resourceCustomIAMRoleDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)
	model, err := tfModelConverter.ConvertTFSchemaToAPIModel(data, []string{NameKey})

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete Custom IAM Role."))
	}

	customIAMRoleFn := model.FullName
	err = config.TMCConnection.CustomIAMRoleResourceService.CustomIAMRoleResourceServiceDelete(customIAMRoleFn)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Couldn't delete Custom IAM Role.\nName: %s", customIAMRoleFn.Name))
	}

	return resourceCustomIAMRoleRead(helper.GetContextWithCaller(ctx, helper.DeleteState), data, m)
}

func resourceCustomIAMRoleImporter(_ context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	config := m.(authctx.TanzuContext)
	customIAMRoleName := data.Id()

	if customIAMRoleName == "" {
		return nil, errors.New("Cluster ID must be set to the custom IAM role name.")
	}

	customIAMRoleFn := &customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleFullName{
		Name: customIAMRoleName,
	}

	resp, err := config.TMCConnection.CustomIAMRoleResourceService.CustomIAMRoleResourceServiceGet(customIAMRoleFn)

	if err != nil || resp.Role == nil {
		return nil, errors.Wrapf(err, "Couldn't import Custom IAM Role.\nName: %s", customIAMRoleFn.Name)
	}

	err = tfModelConverter.FillTFSchema(resp.Role, data)

	if err != nil {
		return nil, errors.Wrapf(err, "Couldn't import Custom IAM Role.\nName: %s", customIAMRoleFn.Name)
	}

	return []*schema.ResourceData{data}, err
}

func validateSchema(ctx context.Context, data *schema.ResourceDiff, m interface{}) (err error) {
	specData := data.Get(SpecKey).([]interface{})[0].(map[string]interface{})
	ruleData := specData[KubernetesPermissionsKey].([]interface{})[0].(map[string]interface{})[RuleKey].([]interface{})
	errMsg := ""

	for i, r := range ruleData {
		resourcesLen := len(r.(map[string]interface{})[ResourcesKey].([]interface{}))
		urlPathsLen := len(r.(map[string]interface{})[URLPathsKey].([]interface{}))

		if (resourcesLen > 0 && urlPathsLen > 0) || (resourcesLen == 0 && urlPathsLen == 0) {
			if errMsg == "" {
				errMsg = "Custom IAM Role Rules Validation Failed:"
			}

			errMsg = fmt.Sprintf("%s\n%s", errMsg, fmt.Sprintf("Rule #%d - Must include %s or %s but not both.", i+1, ResourcesKey, URLPathsKey))
		}
	}

	if errMsg != "" {
		err = errors.New(errMsg)
	}

	return err
}

func formatAggregationRuleData(tfAggregationRule []interface{}, modelAggregationRule *customiamrolemodels.VmwareTanzuManageV1alpha1IamRoleAggregationRule) (aggregationRule []interface{}) {
	clusterRoleSelector := make([]interface{}, 0)
	tfClusterRoleSelectors := tfAggregationRule[0].(map[string]interface{})[ClusterRoleSelectorKey].([]interface{})
	modelClusterRoleSelectors := modelAggregationRule.ClusterRoleSelectors

	for _, tfSelector := range tfClusterRoleSelectors {
		tfSelectorMap := tfSelector.(map[string]interface{})

		if tfSelectorMap[MatchExpressionKey] != nil {
			clusterRoleSelector = append(clusterRoleSelector, tfSelector)
		} else if tfSelectorMap[MatchLabelsKey] != nil {
			for _, modelSelector := range modelClusterRoleSelectors {
				if compareClusterRoleSelectors(tfSelectorMap, modelSelector) {
					clusterRoleSelector = append(clusterRoleSelector, tfSelector)
					break
				}
			}
		}
	}

	if len(clusterRoleSelector) != len(modelClusterRoleSelectors) {
		for _, modelSelector := range modelClusterRoleSelectors {
			modelSelectorFound := false

			if modelSelector.MatchLabels != nil {
				for _, tfSelector := range tfClusterRoleSelectors {
					if compareClusterRoleSelectors(tfSelector.(map[string]interface{}), modelSelector) {
						modelSelectorFound = true
						break
					}
				}

				if !modelSelectorFound {
					clusterRoleSelector = append(clusterRoleSelector, modelSelector)
				}
			}
		}
	}

	aggregationRule = make([]interface{}, 0)
	aggregationRuleMap := make(map[string]interface{})
	aggregationRuleMap[ClusterRoleSelectorKey] = clusterRoleSelector
	aggregationRule = append(aggregationRule, aggregationRuleMap)

	return aggregationRule
}

func compareClusterRoleSelectors(tfSelector map[string]interface{}, modelSelector *customiamrolemodels.K8sIoApimachineryPkgApisMetaV1LabelSelector) bool {
	isEqual := false
	tfMatchLabels, _ := tfSelector[MatchLabelsKey].(map[string]interface{})

	if modelSelector.MatchLabels != nil && tfMatchLabels != nil {
		isEqual = true

		for k, v := range tfMatchLabels {
			modelValue, keyExist := modelSelector.MatchLabels[k]

			if !keyExist || modelValue != v {
				isEqual = false
				break
			}
		}
	}

	return isEqual
}

func formatResourcesData(tfResources []interface{}, modelResources []*customiamrolemodels.VmwareTanzuManageV1alpha1IamPermissionResource) (resources []interface{}) {
	resources = make([]interface{}, 0)

	for _, tfRes := range tfResources {
		for _, modelRes := range modelResources {
			if tfRes.(string) == string(*modelRes) {
				resources = append(resources, tfRes)
				break
			}
		}
	}

	if len(resources) != len(modelResources) {
		for _, modelRes := range modelResources {
			modelResFound := false

			for _, tfRes := range tfResources {
				if tfRes.(string) == string(*modelRes) {
					modelResFound = true
					break
				}
			}

			if !modelResFound {
				resources = append(resources, modelRes)
			}
		}
	}

	return resources
}
