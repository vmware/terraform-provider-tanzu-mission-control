// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
)

var TargetKubernetesResourcesSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "A list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds.",
	Required:    true,
	MinItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			APIGroupsKey: {
				Type:        schema.TypeList,
				Description: "APIGroup is a group containing the resource type.",
				Required:    true,
				MinItems:    1,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			KindsKey: {
				Type:        schema.TypeList,
				Description: "Kind is the name of the object schema (resource type).",
				Required:    true,
				MinItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.All(
						validation.StringIsNotEmpty,
						validation.StringIsNotWhiteSpace,
					),
				},
			},
		},
	},
}

func ExpandTargetKubernetesResources(data interface{}) (kubernetesResources *policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources) {
	if data == nil {
		return kubernetesResources
	}

	kubernetesResourcesData, ok := data.(map[string]interface{})
	if !ok {
		return kubernetesResources
	}

	kubernetesResources = &policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources{}

	if v, ok := kubernetesResourcesData[APIGroupsKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			var value string
			if raw != nil {
				value = raw.(string)
			}

			kubernetesResources.APIGroups = append(kubernetesResources.APIGroups, value)
		}
	}

	if v, ok := kubernetesResourcesData[KindsKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			kubernetesResources.Kinds = append(kubernetesResources.Kinds, raw.(string))
		}
	}

	return kubernetesResources
}

func FlattenTargetKubernetesResources(kubernetesResources *policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources) (data interface{}) {
	if kubernetesResources == nil {
		return data
	}

	flattenTargetKubernetesResources := make(map[string]interface{})

	flattenTargetKubernetesResources[APIGroupsKey] = kubernetesResources.APIGroups
	flattenTargetKubernetesResources[KindsKey] = kubernetesResources.Kinds

	return flattenTargetKubernetesResources
}
