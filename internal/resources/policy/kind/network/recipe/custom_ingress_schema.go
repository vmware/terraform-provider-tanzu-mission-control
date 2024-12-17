// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

//nolint:dupl //ignore 7-151 lines are duplicate of `internal/resources/policy/kind/network/recipe/custom_egress_schema.go:7-151`
package recipe

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

var (
	CustomIngress = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The input schema for network policy custom ingress recipe version v1",
		Optional:    true,
		ForceNew:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				rulesKey: {
					Type:        schema.TypeList,
					Description: "This specifies list of ingress rules to be applied to the selected pods.",
					Required:    true,
					MinItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							portsKey: {
								Type:        schema.TypeList,
								Description: "List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. Default is this rule matches all ports (traffic not restricted by port).",
								Required:    true,
								MinItems:    1,
								Elem:        port,
							},
							ruleSpecKey: {
								Type:        schema.TypeList,
								Description: "List of sources which should be able to access the pods selected for this rule. Default is the rule matches all sources (traffic not restricted by source). List of items of type V1alpha1CommonPolicySpecNetworkV1CustomIngressRulesRuleSpec0 OR V1alpha1CommonPolicySpecNetworkV1CustomIngressRulesRuleSpec1.",
								Required:    true,
								MinItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										ruleSpecIPKey:       ruleSpecIPSource,
										ruleSpecSelectorKey: ruleSpecSelectorSource,
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

	ruleSpecIPSource = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The rule Spec (source) for IP Block.",
		Optional:    true,
		MinItems:    1,
		Elem:        ip,
	}

	ruleSpecSelectorSource = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The rule Spec (source) for Selectors.",
		Optional:    true,
		MinItems:    1,
		Elem:        selector,
	}
)

func ConstructCustomIngress(data []interface{}) (customIngress *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomIngress) {
	if len(data) == 0 || data[0] == nil {
		return customIngress
	}

	customIngressData, _ := data[0].(map[string]interface{})

	customIngress = &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomIngress{}

	if v, ok := customIngressData[rulesKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				customIngress.Rules = make([]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRules, 0)

				for _, raw := range vs {
					customIngress.Rules = append(customIngress.Rules, expandCustomRule(raw))
				}
			}
		}
	}

	if podLabelsData, ok := customIngressData[ToPodLabelsKey]; ok {
		if podLabels, ok := podLabelsData.(map[string]interface{}); ok {
			if len(podLabels) != 0 {
				labels := make([]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels, 0)
				customIngress.ToPodLabels = &labels

				for key, value := range podLabels {
					value := fmt.Sprintf("%v", value)

					labels = append(labels, policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
						Key:   key,
						Value: value,
					})
				}

				customIngress.ToPodLabels = &labels
			}
		}
	}

	return customIngress
}

func FlattenCustomIngress(customIngress *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1CustomIngress) (data []interface{}) {
	if customIngress == nil {
		return data
	}

	flattenCustomIngress := make(map[string]interface{})

	if customIngress.Rules != nil {
		var rules []interface{}

		for _, rule := range customIngress.Rules {
			rules = append(rules, flattenCustomRule(&rule))
		}

		flattenCustomIngress[rulesKey] = rules
	}

	if customIngress.ToPodLabels != nil {
		labels := make(map[string]interface{})

		for _, label := range *customIngress.ToPodLabels {
			labels[label.Key] = label.Value
		}

		flattenCustomIngress[ToPodLabelsKey] = labels
	}

	return []interface{}{flattenCustomIngress}
}
