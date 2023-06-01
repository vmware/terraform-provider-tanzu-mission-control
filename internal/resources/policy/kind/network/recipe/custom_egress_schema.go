/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

//nolint:dupl //ignore 7-151 lines are duplicate of `internal/resources/policy/kind/network/recipe/custom_ingress_schema.go:7-151`
package recipe

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

var (
	CustomEgress = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The input schema for network policy custom egress recipe version v1",
		Optional:    true,
		ForceNew:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				rulesKey: {
					Type:        schema.TypeList,
					Description: "This specifies list of egress rules to be applied to the selected pods.",
					Required:    true,
					MinItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							portsKey: {
								Type:        schema.TypeList,
								Description: "List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. Default is this rule matches all ports (traffic not restricted by port).",
								Required:    true,
								MinItems:    1,
								Elem:        port,
							},
							ruleSpecKey: {
								Type:        schema.TypeList,
								Description: "List of destinations for outgoing traffic of pods selected for this rule. Default is the rule matches all destinations (traffic not restricted by destinations).",
								Required:    true,
								MinItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										ruleSpecIPKey:       ruleSpecIPDestination,
										ruleSpecSelectorKey: ruleSpecSelectorDestination,
									},
								},
							},
						},
					},
				},
				ToPodLabelsKey: toPodLabel,
			},
		},
	}

	ruleSpecIPDestination = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The rule Spec (destination) for IP Block.",
		Optional:    true,
		MinItems:    1,
		Elem:        ip,
	}

	ruleSpecSelectorDestination = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The rule Spec (destination) for Selectors.",
		Optional:    true,
		MinItems:    1,
		Elem:        selector,
	}
)

func ConstructCustomEgress(data []interface{}) (customEgress *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomEgress) {
	if len(data) == 0 || data[0] == nil {
		return customEgress
	}

	customEgressData, _ := data[0].(map[string]interface{})

	customEgress = &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomEgress{}

	if v, ok := customEgressData[rulesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				customEgress.Rules = make([]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRules, 0)

				for _, raw := range vs {
					customEgress.Rules = append(customEgress.Rules, expandCustomRule(raw))
				}
			}
		}
	}

	if podLabelsData, ok := customEgressData[ToPodLabelsKey]; ok {
		if podLabels, ok := podLabelsData.(map[string]interface{}); ok {
			if len(podLabels) != 0 {
				labels := make([]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels, 0)
				customEgress.ToPodLabels = &labels

				for key, value := range podLabels {
					value := fmt.Sprintf("%v", value)

					labels = append(labels, policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
						Key:   key,
						Value: value,
					})
				}

				customEgress.ToPodLabels = &labels
			}
		}
	}

	return customEgress
}

func FlattenCustomEgress(customEgress *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomEgress) (data []interface{}) {
	if customEgress == nil {
		return data
	}

	flattenCustomEgress := make(map[string]interface{})

	if customEgress.Rules != nil {
		var rules []interface{}

		for _, rule := range customEgress.Rules {
			rule := rule
			rules = append(rules, flattenCustomRule(&rule))
		}

		flattenCustomEgress[rulesKey] = rules
	}

	if customEgress.ToPodLabels != nil {
		labels := make(map[string]interface{})

		for _, label := range *customEgress.ToPodLabels {
			labels[label.Key] = label.Value
		}

		flattenCustomEgress[ToPodLabelsKey] = labels
	}

	return []interface{}{flattenCustomEgress}
}
