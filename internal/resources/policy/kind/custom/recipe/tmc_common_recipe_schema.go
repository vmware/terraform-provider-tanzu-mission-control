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
			TargetKubernetesResourcesKey: targetKubernetesResources,
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
			TargetKubernetesResourcesKey: targetKubernetesResources,
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
			TargetKubernetesResourcesKey: targetKubernetesResources,
		},
	},
}

func ConstructTMCCommonRecipe(data []interface{}) (common *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe) {
	if len(data) == 0 || data[0] == nil {
		return common
	}

	commonRecipeData, _ := data[0].(map[string]interface{})

	common = &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe{}

	if v, ok := commonRecipeData[AuditKey]; ok {
		helper.SetPrimitiveValue(v, &common.Audit, AuditKey)
	}

	if v, ok := commonRecipeData[TargetKubernetesResourcesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				common.TargetKubernetesResources = make([]*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources, 0)

				for _, raw := range vs {
					common.TargetKubernetesResources = append(common.TargetKubernetesResources, expandTargetKubernetesResources(raw))
				}
			}
		}
	}

	return common
}

func FlattenTMCCommonRecipe(common *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe) (data []interface{}) {
	if common == nil {
		return data
	}

	flattenCommonRecipe := make(map[string]interface{})

	flattenCommonRecipe[AuditKey] = common.Audit

	if common.TargetKubernetesResources != nil {
		var tkrs []interface{}

		for _, tkr := range common.TargetKubernetesResources {
			tkrs = append(tkrs, flattenTargetKubernetesResources(tkr))
		}

		flattenCommonRecipe[TargetKubernetesResourcesKey] = tkrs
	}

	return []interface{}{flattenCommonRecipe}
}
