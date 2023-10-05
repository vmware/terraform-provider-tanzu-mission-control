/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

// Package recipe contains schema and helper functions for different input recipes.
// nolint: dupl
package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/common"
)

var TMCRequireLabels = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for custom policy tmc_require_labels recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			AuditKey: {
				Type:        schema.TypeBool,
				Description: "Audit (dry-run).",
				Optional:    true,
				Default:     false,
			},
			ParametersKey: {
				Type:        schema.TypeList,
				Description: "Parameters.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						parametersLabelKey: {
							Type:        schema.TypeList,
							Description: "Labels.",
							Required:    true,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									labelKey: {
										Type:        schema.TypeString,
										Description: "The label key to enforce.",
										Required:    true,
									},
									labelValueKey: {
										Type:        schema.TypeString,
										Description: "Optional label value to enforce (if left empty, only key will be enforced).",
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			TargetKubernetesResourcesKey: common.TargetKubernetesResourcesSchema,
		},
	},
}

func ConstructTMCRequireLabels(data []interface{}) (requireLabels *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabels) {
	if len(data) == 0 || data[0] == nil {
		return requireLabels
	}

	requireLabelsData, _ := data[0].(map[string]interface{})

	requireLabels = &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabels{}

	if v, ok := requireLabelsData[AuditKey]; ok {
		helper.SetPrimitiveValue(v, &requireLabels.Audit, AuditKey)
	}

	if v, ok := requireLabelsData[ParametersKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			requireLabels.Parameters = expandRequiredLabelsParameters(v1)
		}
	}

	if v, ok := requireLabelsData[TargetKubernetesResourcesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				requireLabels.TargetKubernetesResources = make([]*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources, 0)

				for _, raw := range vs {
					requireLabels.TargetKubernetesResources = append(requireLabels.TargetKubernetesResources, common.ExpandTargetKubernetesResources(raw))
				}
			}
		}
	}

	return requireLabels
}

func expandRequiredLabelsParameters(data []interface{}) (parameters *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParameters) {
	if len(data) == 0 || data[0] == nil {
		return parameters
	}

	parametersData, ok := data[0].(map[string]interface{})
	if !ok {
		return parameters
	}

	parameters = &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParameters{}

	if v, ok := parametersData[parametersLabelKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				parameters.Labels = make([]*policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParametersLabels, 0)

				for _, raw := range vs {
					parameters.Labels = append(parameters.Labels, expandLabels(raw))
				}
			}
		}
	}

	return parameters
}

func expandLabels(data interface{}) (labels *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParametersLabels) {
	if data == nil {
		return labels
	}

	labelsData, ok := data.(map[string]interface{})
	if !ok {
		return labels
	}

	labels = &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParametersLabels{}

	if v, ok := labelsData[labelKey]; ok {
		helper.SetPrimitiveValue(v, &labels.Key, labelKey)
	}

	if v, ok := labelsData[labelValueKey]; ok {
		helper.SetPrimitiveValue(v, &labels.Value, labelValueKey)
	}

	return labels
}

func FlattenTMCRequireLabels(requireLabels *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabels) (data []interface{}) {
	if requireLabels == nil {
		return data
	}

	flattenRequiredLabels := make(map[string]interface{})

	flattenRequiredLabels[AuditKey] = requireLabels.Audit

	if requireLabels.Parameters != nil {
		flattenRequiredLabels[ParametersKey] = flattenRequiredLabelsParameters(requireLabels.Parameters)
	}

	if requireLabels.TargetKubernetesResources != nil {
		var tkrs []interface{}

		for _, tkr := range requireLabels.TargetKubernetesResources {
			tkrs = append(tkrs, common.FlattenTargetKubernetesResources(tkr))
		}

		flattenRequiredLabels[TargetKubernetesResourcesKey] = tkrs
	}

	return []interface{}{flattenRequiredLabels}
}

func flattenRequiredLabelsParameters(parameters *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParameters) (data []interface{}) {
	if parameters == nil {
		return data
	}

	flattenParameters := make(map[string]interface{})

	if parameters.Labels != nil {
		var labels []interface{}

		for _, label := range parameters.Labels {
			labels = append(labels, flattenLabels(label))
		}

		flattenParameters[parametersLabelKey] = labels
	}

	return []interface{}{flattenParameters}
}

func flattenLabels(labels *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCRequireLabelsParametersLabels) (data interface{}) {
	if labels == nil {
		return data
	}

	flattenLabels := make(map[string]interface{})

	flattenLabels[labelKey] = labels.Key
	flattenLabels[labelValueKey] = labels.Value

	return flattenLabels
}
