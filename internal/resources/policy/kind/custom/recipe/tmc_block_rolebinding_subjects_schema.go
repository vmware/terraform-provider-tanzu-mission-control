/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

// Package recipe contains schema and helper functions for different input recipes.
// nolint: dupl
package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/common"
)

var TMCBlockRolebindingSubjects = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for custom policy tmc_block_rolebinding_subjects recipe version v1",
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
						disallowedSubjectsKey: {
							Type:        schema.TypeList,
							Description: "Disallowed Subjects.",
							Required:    true,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									kindKey: {
										Type:         schema.TypeString,
										Description:  "The kind of subject to disallow, can be User/Group/ServiceAccount.",
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"User", "Group", "ServiceAccount"}, false),
									},
									nameKey: {
										Type:        schema.TypeString,
										Description: "The name of the subject to disallow.",
										Required:    true,
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

func ConstructTMCBlockRolebindingSubjects(data []interface{}) (roleBindingSubjects *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjects) {
	if len(data) == 0 || data[0] == nil {
		return roleBindingSubjects
	}

	roleBindingSubjectsData, _ := data[0].(map[string]interface{})

	roleBindingSubjects = &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjects{}

	if v, ok := roleBindingSubjectsData[AuditKey]; ok {
		helper.SetPrimitiveValue(v, &roleBindingSubjects.Audit, AuditKey)
	}

	if v, ok := roleBindingSubjectsData[ParametersKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			roleBindingSubjects.Parameters = expandBlockRoleBindingParameters(v1)
		}
	}

	if v, ok := roleBindingSubjectsData[TargetKubernetesResourcesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				roleBindingSubjects.TargetKubernetesResources = make([]*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources, 0)

				for _, raw := range vs {
					roleBindingSubjects.TargetKubernetesResources = append(roleBindingSubjects.TargetKubernetesResources, common.ExpandTargetKubernetesResources(raw))
				}
			}
		}
	}

	return roleBindingSubjects
}

func expandBlockRoleBindingParameters(data []interface{}) (parameters *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParameters) {
	if len(data) == 0 || data[0] == nil {
		return parameters
	}

	parametersData, ok := data[0].(map[string]interface{})
	if !ok {
		return parameters
	}

	parameters = &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParameters{}

	if v, ok := parametersData[disallowedSubjectsKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				parameters.DisallowedSubjects = make([]*policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParametersDisallowedSubjects, 0)

				for _, raw := range vs {
					parameters.DisallowedSubjects = append(parameters.DisallowedSubjects, expandDisallowedSubjects(raw))
				}
			}
		}
	}

	return parameters
}

func expandDisallowedSubjects(data interface{}) (disallowedSubjects *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParametersDisallowedSubjects) {
	if data == nil {
		return disallowedSubjects
	}

	disallowedSubjectsData, ok := data.(map[string]interface{})
	if !ok {
		return disallowedSubjects
	}

	disallowedSubjects = &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParametersDisallowedSubjects{}

	if v, ok := disallowedSubjectsData[kindKey]; ok {
		helper.SetPrimitiveValue(v, &disallowedSubjects.Kind, kindKey)
	}

	if v, ok := disallowedSubjectsData[nameKey]; ok {
		helper.SetPrimitiveValue(v, &disallowedSubjects.Name, nameKey)
	}

	return disallowedSubjects
}

func FlattenTMCBlockRolebindingSubjects(roleBindingSubjects *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjects) (data []interface{}) {
	if roleBindingSubjects == nil {
		return data
	}

	flattenRoleBindingSubjects := make(map[string]interface{})

	flattenRoleBindingSubjects[AuditKey] = roleBindingSubjects.Audit

	if roleBindingSubjects.Parameters != nil {
		flattenRoleBindingSubjects[ParametersKey] = flattenBlockRoleBindingParameters(roleBindingSubjects.Parameters)
	}

	if roleBindingSubjects.TargetKubernetesResources != nil {
		var tkrs []interface{}

		for _, tkr := range roleBindingSubjects.TargetKubernetesResources {
			tkrs = append(tkrs, common.FlattenTargetKubernetesResources(tkr))
		}

		flattenRoleBindingSubjects[TargetKubernetesResourcesKey] = tkrs
	}

	return []interface{}{flattenRoleBindingSubjects}
}

func flattenBlockRoleBindingParameters(parameters *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParameters) (data []interface{}) {
	if parameters == nil {
		return data
	}

	flattenParameters := make(map[string]interface{})

	if parameters.DisallowedSubjects != nil {
		var labels []interface{}

		for _, label := range parameters.DisallowedSubjects {
			labels = append(labels, flattenDisallowedSubjects(label))
		}

		flattenParameters[disallowedSubjectsKey] = labels
	}

	return []interface{}{flattenParameters}
}

func flattenDisallowedSubjects(disallowedSubjects *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCBlockRoleBindingSubjectsParametersDisallowedSubjects) (data interface{}) {
	if disallowedSubjects == nil {
		return data
	}

	flattenDisallowedSubjects := make(map[string]interface{})

	flattenDisallowedSubjects[kindKey] = disallowedSubjects.Kind
	flattenDisallowedSubjects[nameKey] = disallowedSubjects.Name

	return flattenDisallowedSubjects
}
