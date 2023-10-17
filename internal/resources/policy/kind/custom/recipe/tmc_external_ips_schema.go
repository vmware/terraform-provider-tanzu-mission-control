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

var TMCExternalIps = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for custom policy tmc_external_ips recipe version v1",
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
						allowedIPsKey: {
							Type:        schema.TypeList,
							Description: "Allowed IPs.",
							Required:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			TargetKubernetesResourcesKey: common.TargetKubernetesResourcesSchema,
		},
	},
}

func ConstructTMCExternalIPS(data []interface{}) (externalIPs *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPS) {
	if len(data) == 0 || data[0] == nil {
		return externalIPs
	}

	externalIPsData, _ := data[0].(map[string]interface{})

	externalIPs = &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPS{}

	if v, ok := externalIPsData[AuditKey]; ok {
		helper.SetPrimitiveValue(v, &externalIPs.Audit, AuditKey)
	}

	if v, ok := externalIPsData[ParametersKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			externalIPs.Parameters = expandExternalIPSParameters(v1)
		}
	}

	if v, ok := externalIPsData[TargetKubernetesResourcesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				externalIPs.TargetKubernetesResources = make([]*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources, 0)

				for _, raw := range vs {
					externalIPs.TargetKubernetesResources = append(externalIPs.TargetKubernetesResources, common.ExpandTargetKubernetesResources(raw))
				}
			}
		}
	}

	return externalIPs
}

func expandExternalIPSParameters(data []interface{}) (parameters *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPSParameters) {
	if len(data) == 0 || data[0] == nil {
		return parameters
	}

	parametersData, ok := data[0].(map[string]interface{})
	if !ok {
		return parameters
	}

	parameters = &policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPSParameters{}

	if v, ok := parametersData[allowedIPsKey]; ok {
		vs, _ := v.([]interface{})
		for _, raw := range vs {
			parameters.AllowedIPs = append(parameters.AllowedIPs, raw.(string))
		}
	}

	return parameters
}

func FlattenTMCExternalIPS(externalIPs *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPS) (data []interface{}) {
	if externalIPs == nil {
		return data
	}

	flattenExternalIPs := make(map[string]interface{})

	flattenExternalIPs[AuditKey] = externalIPs.Audit

	if externalIPs.Parameters != nil {
		flattenExternalIPs[ParametersKey] = flattenExternalIPSParameters(externalIPs.Parameters)
	}

	if externalIPs.TargetKubernetesResources != nil {
		var tkrs []interface{}

		for _, tkr := range externalIPs.TargetKubernetesResources {
			tkrs = append(tkrs, common.FlattenTargetKubernetesResources(tkr))
		}

		flattenExternalIPs[TargetKubernetesResourcesKey] = tkrs
	}

	return []interface{}{flattenExternalIPs}
}

func flattenExternalIPSParameters(parameters *policyrecipecustommodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCExternalIPSParameters) (data []interface{}) {
	if parameters == nil {
		return data
	}

	flattenParameters := make(map[string]interface{})

	flattenParameters[allowedIPsKey] = parameters.AllowedIPs

	return []interface{}{flattenParameters}
}
