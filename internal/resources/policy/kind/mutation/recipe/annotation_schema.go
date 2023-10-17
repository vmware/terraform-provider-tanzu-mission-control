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

var AnnotationSchema = &schema.Schema{
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
				Default:      "*",
				ValidateFunc: validation.StringInSlice([]string{"*", "Cluster", "Namespaced"}, false),
			},
			AnnotationKey: {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
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

func ConstructAnnotation(data []interface{}) (annotationModel *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation) {
	if len(data) == 0 || data[0] == nil {
		return annotationModel
	}

	requireAnnotationData, _ := data[0].(map[string]interface{})

	annotationModel = &policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation{}

	if v, ok := requireAnnotationData[AnnotationKey]; ok {
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

				annotationModel.Annotation = keyValue
			}
		}
	}

	if scope, ok := requireAnnotationData[scopeKey]; ok {
		mutationScope := policyrecipemutationcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Scope(scope.(string))
		annotationModel.Scope = &mutationScope
	}

	if v, ok := requireAnnotationData[targetKubernetesResourcesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				annotationModel.TargetKubernetesResources = make([]*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources, 0)

				for _, raw := range vs {
					annotationModel.TargetKubernetesResources = append(annotationModel.TargetKubernetesResources, common.ExpandTargetKubernetesResources(raw))
				}
			}
		}
	}

	return annotationModel
}

func FlattenAnnotation(mutationAnnotation *policyrecipemutationmodel.VmwareTanzuManageV1alpha1CommonPolicySpecMutationV1Annotation) (data []interface{}) {
	if mutationAnnotation == nil {
		return data
	}

	flattenAnnotation := make(map[string]interface{})

	if mutationAnnotation.Annotation != nil {
		flattenAnnotation[AnnotationKey] = flattenKeyValuesFromAnnotation(mutationAnnotation.Annotation)
	}

	if mutationAnnotation.Scope != nil {
		flattenAnnotation[scopeKey] = string(*mutationAnnotation.Scope)
	}

	if mutationAnnotation.TargetKubernetesResources != nil {
		var targetKubernetesResources []interface{}

		for _, tkr := range mutationAnnotation.TargetKubernetesResources {
			targetKubernetesResources = append(targetKubernetesResources, common.FlattenTargetKubernetesResources(tkr))
		}

		flattenAnnotation[targetKubernetesResourcesKey] = targetKubernetesResources
	}

	return []interface{}{flattenAnnotation}
}

func flattenKeyValuesFromAnnotation(keyValue *policyrecipemutationcommonmodel.KeyValue) []interface{} {
	var annotationKeyValue []interface{}

	flattenAnnotationValue := make(map[string]interface{})
	flattenAnnotationValue[keyKey] = keyValue.Key
	flattenAnnotationValue[valueKey] = keyValue.Value

	annotationKeyValue = append(annotationKeyValue, flattenAnnotationValue)

	return annotationKeyValue
}
