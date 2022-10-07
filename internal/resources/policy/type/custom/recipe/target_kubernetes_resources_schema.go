/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
)

var targetKubernetesResources = &schema.Schema{
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
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	},
}

func expandTargetKubernetesResources(data interface{}) (kubernetesResources *policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources) {
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
			kubernetesResources.APIGroups = append(kubernetesResources.APIGroups, raw.(string))
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

func flattenTargetKubernetesResources(kubernetesResources *policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources) (data interface{}) {
	if kubernetesResources == nil {
		return data
	}

	flattenTargetKubernetesResources := make(map[string]interface{})

	flattenTargetKubernetesResources[APIGroupsKey] = kubernetesResources.APIGroups
	flattenTargetKubernetesResources[KindsKey] = kubernetesResources.Kinds

	return flattenTargetKubernetesResources
}
