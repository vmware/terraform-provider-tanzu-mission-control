/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

// Package recipe contains schema and helper functions for different input recipes.
// Contains recipe schema for tmc-https-ingress, tmc-block-nodeport-service and tmc-block-resources recipe.
package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipecustommodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom"
	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/common"
)

var TMCBlockNodeportService = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for custom policy tmc_block_nodeport_service recipe version v1",
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
			TargetKubernetesResourcesKey: common.TargetKubernetesResourcesSchema,
		},
	},
}

var TMCBlockResources = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for custom policy tmc_block_resources recipe version v1",
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
			TargetKubernetesResourcesKey: common.TargetKubernetesResourcesSchema,
		},
	},
}

var TMCHTTPSIngress = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for custom policy tmc_https_ingress recipe version v1",
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
			TargetKubernetesResourcesKey: common.TargetKubernetesResourcesSchema,
		},
	},
}

func ConstructTMCCommonRecipe(data []interface{}) (commonRecipe *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe) {
	if len(data) == 0 || data[0] == nil {
		return commonRecipe
	}

	commonRecipeData, _ := data[0].(map[string]interface{})

	commonRecipe = &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe{}

	if v, ok := commonRecipeData[AuditKey]; ok {
		helper.SetPrimitiveValue(v, &commonRecipe.Audit, AuditKey)
	}

	if v, ok := commonRecipeData[TargetKubernetesResourcesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				commonRecipe.TargetKubernetesResources = make([]*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources, 0)

				for _, raw := range vs {
					commonRecipe.TargetKubernetesResources = append(commonRecipe.TargetKubernetesResources, common.ExpandTargetKubernetesResources(raw))
				}
			}
		}
	}

	return commonRecipe
}

func FlattenTMCCommonRecipe(commonRecipe *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe) (data []interface{}) {
	if commonRecipe == nil {
		return data
	}

	flattenCommonRecipe := make(map[string]interface{})

	flattenCommonRecipe[AuditKey] = commonRecipe.Audit

	if commonRecipe.TargetKubernetesResources != nil {
		var tkrs []interface{}

		for _, tkr := range commonRecipe.TargetKubernetesResources {
			tkrs = append(tkrs, common.FlattenTargetKubernetesResources(tkr))
		}

		flattenCommonRecipe[TargetKubernetesResourcesKey] = tkrs
	}

	return []interface{}{flattenCommonRecipe}
}
