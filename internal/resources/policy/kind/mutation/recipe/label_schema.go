/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

// nolint: dupl
package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
	policyrecipemutationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation"
	policyrecipemutationcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/mutation/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/common"
)

var LabelSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for custom policy tmc_block_nodeport_service recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			targetKubernetesResourcesKey: common.TargetKubernetesResourcesSchema,
			scopeKey: {
				Type:         schema.TypeString,
				Description:  "Scope",
				Optional:     true,
				Default:      "All",
				ValidateFunc: validation.StringInSlice([]string{"All", "Cluster", "Namespaced"}, false),
			},
			LabelKey: {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						keyKey: {
							Type:     schema.TypeString,
							Required: true,
						},
						valueKey: {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	},
}

func ConstructLabel(data []interface{}) (labelModel *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label) {
	if len(data) == 0 || data[0] == nil {
		return labelModel
	}

	requireLabelData, _ := data[0].(map[string]interface{})

	labelModel = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label{}

	if v, ok := requireLabelData[LabelKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				keyValuePair, _ := vs[0].(map[string]interface{})

				keyValue := &policyrecipemutationcommonmodel.KeyValue{}
				if v, ok := keyValuePair[keyKey]; ok {
					helper.SetPrimitiveValue(v, &keyValue.Key, keyKey)
				}

				if v, ok := keyValuePair[valueKey]; ok {
					helper.SetPrimitiveValue(v, &keyValue.Value, valueKey)
				}

				labelModel.Label = keyValue
			}
		}
	}

	if scope, ok := requireLabelData[scopeKey]; ok {
		mutationScope := policyrecipemutationcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope(scope.(string))
		labelModel.Scope = &mutationScope
	}

	if v, ok := requireLabelData[targetKubernetesResourcesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				labelModel.TargetKubernetesResources = make([]*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources, 0)

				for _, raw := range vs {
					labelModel.TargetKubernetesResources = append(labelModel.TargetKubernetesResources, common.ExpandTargetKubernetesResources(raw))
				}
			}
		}
	}

	return labelModel
}

func FlattenLabel(mutationLabel *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Label) (data []interface{}) {
	if mutationLabel == nil {
		return data
	}

	flattenLabel := make(map[string]interface{})

	if mutationLabel.Label != nil {
		flattenLabel[LabelKey] = flattenKeyValuesFromLabel(mutationLabel.Label)
	}

	if mutationLabel.Scope != nil {
		flattenLabel[scopeKey] = string(*mutationLabel.Scope)
	}

	if mutationLabel.TargetKubernetesResources != nil {
		var targetKubernetesResources []interface{}

		for _, tkr := range mutationLabel.TargetKubernetesResources {
			targetKubernetesResources = append(targetKubernetesResources, common.FlattenTargetKubernetesResources(tkr))
		}

		flattenLabel[targetKubernetesResourcesKey] = targetKubernetesResources
	}

	return []interface{}{flattenLabel}
}

func flattenKeyValuesFromLabel(keyValue *policyrecipemutationcommonmodel.KeyValue) []interface{} {
	var labelKeyValue []interface{}

	flattenLabelValue := make(map[string]interface{})
	flattenLabelValue[keyKey] = keyValue.Key
	flattenLabelValue[valueKey] = keyValue.Value

	labelKeyValue = append(labelKeyValue, flattenLabelValue)

	return labelKeyValue
}
