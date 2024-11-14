// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

var (
	port = &schema.Resource{
		Schema: map[string]*schema.Schema{
			portKey: {
				Type:        schema.TypeString,
				Description: "The port on the given protocol. This can either be a numerical or named port on a pod.",
				Optional:    true,
			},
			protocolKey: {
				Type:        schema.TypeString,
				Description: "The protocol (TCP or UDP) which traffic must match.",
				Optional:    true,
			},
		},
	}

	ip = &schema.Resource{
		Schema: map[string]*schema.Schema{
			ipBlockKey: {
				Type:        schema.TypeList,
				Description: "IPBlock defines policy on a particular IPBlock. If this field is set then neither of the namespaceSelector and PodSelector can be set.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cidrKey: {
							Type:        schema.TypeString,
							Description: "CIDR is a string representing the IP Block Valid examples are \"192.168.1.1/24\" or \"2001:db9::/64\"",
							Required:    true,
						},
						exceptKey: {
							Type:        schema.TypeList,
							Description: "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are \"192.168.1.1/24\" or \"2001:db9::/64\" Except values will be rejected if they are outside the CIDR range",
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}

	selector = &schema.Resource{
		Schema: map[string]*schema.Schema{
			namespaceSelectorKey: {
				Type:        schema.TypeMap,
				Description: "Use a label selector to identify the namespaces to allow as egress destinations.",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			podSelectorKey: {
				Type:        schema.TypeMap,
				Description: "Use a label selector to identify the pods to allow as egress destinations.",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
)

// Custom is a struct for all types of rule specs.
type custom struct {
	customIPBlock  *policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec0
	customSelector *policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec1
}

func expandCustomRule(data interface{}) (rule policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRules) {
	if data == nil {
		return rule
	}

	ruleData, _ := data.(map[string]interface{})

	rule = policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRules{}

	if portsData, ok := ruleData[portsKey]; ok {
		if pt, ok := portsData.([]interface{}); ok {
			if len(pt) != 0 && pt[0] != nil {
				ports := make([]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesPorts, 0)
				rule.Ports = &ports

				for _, p := range pt {
					ports = append(*rule.Ports, expandPort(p))
					rule.Ports = &ports
				}
			}
		}
	}

	ruleSpec, ok := ruleData[ruleSpecKey]
	if !ok {
		return rule
	}

	ruleSpecData, ok := ruleSpec.([]interface{})
	if !ok {
		return rule
	}

	if len(ruleSpecData) != 0 && ruleSpecData[0] != nil {
		rs := make([]interface{}, 0)

		for _, raw := range ruleSpecData {
			ruleSpecTypeData := constructRuleSpec(raw)

			if ruleSpecTypeData == nil {
				return rule
			}

			if ruleSpecTypeData.customIPBlock != nil {
				rs = append(rs, ruleSpecTypeData.customIPBlock)
			}

			if ruleSpecTypeData.customSelector != nil {
				rs = append(rs, ruleSpecTypeData.customSelector)
			}
		}

		rule.RuleSpec = rs
	}

	return rule
}

func expandPort(data interface{}) (port policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesPorts) {
	if data == nil {
		return port
	}

	portData, _ := data.(map[string]interface{})

	port = policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesPorts{}

	if pt, ok := portData[portKey]; ok {
		port.Port = helper.StringPointer(pt.(string))
	}

	if proto, ok := portData[protocolKey]; ok {
		protocol := policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesPortsProtocol(proto.(string))
		port.Protocol = &protocol
	}

	return port
}

func constructRuleSpec(data interface{}) (customData *custom) {
	if data == nil {
		return customData
	}

	inputData, _ := data.(map[string]interface{})

	if ruleIP, ok := inputData[ruleSpecIPKey]; ok {
		if specIP, ok := ruleIP.([]interface{}); ok && len(specIP) != 0 {
			customData = &custom{
				customIPBlock: expandRuleSpec0(specIP),
			}
		}
	}

	if ruleSelector, ok := inputData[ruleSpecSelectorKey]; ok {
		if specSelector, ok := ruleSelector.([]interface{}); ok && len(specSelector) != 0 {
			customData = &custom{
				customSelector: expandRuleSpec1(specSelector),
			}
		}
	}

	return customData
}

func expandRuleSpec0(data []interface{}) (spec0 *policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec0) {
	if len(data) != 0 && data[0] == nil {
		return spec0
	}

	spec0Data, _ := data[0].(map[string]interface{})

	spec0 = &policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec0{}

	if ipData, ok := spec0Data[ipBlockKey]; ok {
		if ips, ok := ipData.([]interface{}); ok {
			spec0.IpBlock = expandIP(ips)
		}
	}

	return spec0
}

func expandIP(data []interface{}) (ip *policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec0IpBlock) {
	if len(data) != 0 && data[0] == nil {
		return ip
	}

	ipData, _ := data[0].(map[string]interface{})

	ip = &policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec0IpBlock{}

	if cidrData, ok := ipData[cidrKey]; ok {
		helper.SetPrimitiveValue(cidrData, &ip.Cidr, cidrKey)
	}

	if exceptData, ok := ipData[exceptKey]; ok {
		except, _ := exceptData.([]interface{})

		s := make([]string, 0)

		for _, et := range except {
			s = append(s, et.(string))
		}

		ip.Except = &s
	}

	return ip
}

func expandRuleSpec1(data []interface{}) (spec1 *policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec1) {
	if len(data) != 0 && data[0] == nil {
		return spec1
	}

	spec1Data, _ := data[0].(map[string]interface{})

	spec1 = &policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec1{}

	if namespaceSelectorData, ok := spec1Data[namespaceSelectorKey]; ok {
		if namespaceSelectors, ok := namespaceSelectorData.(map[string]interface{}); ok {
			if len(namespaceSelectors) != 0 {
				selectors := make([]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels, 0)
				spec1.NamespaceSelector = &selectors

				for key, value := range namespaceSelectors {
					value := fmt.Sprintf("%v", value)

					selectors = append(selectors, policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
						Key:   key,
						Value: value,
					})
				}

				spec1.NamespaceSelector = &selectors
			}
		}
	}

	if podSelectorData, ok := spec1Data[podSelectorKey]; ok {
		if podSelectors, ok := podSelectorData.(map[string]interface{}); ok {
			if len(podSelectors) != 0 {
				selectors := make([]policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels, 0)
				spec1.PodSelector = &selectors

				for key, value := range podSelectors {
					value := fmt.Sprintf("%v", value)

					selectors = append(selectors, policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
						Key:   key,
						Value: value,
					})
				}

				spec1.PodSelector = &selectors
			}
		}
	}

	return spec1
}

func flattenCustomRule(rule *policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRules) (data interface{}) {
	if rule == nil {
		return data
	}

	flattenRule := make(map[string]interface{})

	if rule.Ports != nil {
		var ports []interface{}

		for _, port := range *rule.Ports {
			port := port
			ports = append(ports, flattenPort(&port))
		}

		flattenRule[portsKey] = ports
	}

	if rule.RuleSpec != nil && rule.RuleSpec[0] != nil {
		var rss []interface{}

		for _, rs := range rule.RuleSpec {
			ruleSpec, ok := rs.(map[string]interface{})
			if !ok {
				return data
			}

			var (
				ruleSpecData        *custom
				customIPBlock       policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec0
				customSelectorBlock policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec1
			)

			byteSlice, err := json.Marshal(ruleSpec)
			if err != nil {
				return data
			}

			err = customIPBlock.UnmarshalBinary(byteSlice)
			if err != nil {
				return data
			}

			if _, ok := ruleSpec[IPBlock]; ok {
				ruleSpecData = &custom{
					customIPBlock: &customIPBlock,
				}
			}

			err = customSelectorBlock.UnmarshalBinary(byteSlice)
			if err != nil {
				return data
			}

			_, isPodSelectorPresent := ruleSpec[PodSelector]
			_, isNamespaceSelectorPresent := ruleSpec[NamespaceSelector]

			if isPodSelectorPresent || isNamespaceSelectorPresent {
				ruleSpecData = &custom{
					customSelector: &customSelectorBlock,
				}
			}

			if ruleSpecData == nil {
				return data
			}

			rss = append(rss, flattenRuleSpec(ruleSpecData))
		}

		flattenRule[ruleSpecKey] = rss
	}

	return flattenRule
}

func flattenPort(port *policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesPorts) (data interface{}) {
	if port == nil {
		return data
	}

	flattenPortData := make(map[string]interface{})

	flattenPortData[portKey] = *port.Port

	flattenPortData[protocolKey] = string(*port.Protocol)

	return flattenPortData
}

func flattenRuleSpec(customData *custom) (data interface{}) {
	if customData == nil {
		return data
	}

	flattenRuleSpecData := make(map[string]interface{})

	if customData.customIPBlock != nil {
		flattenRuleSpecData[ruleSpecIPKey] = flattenRuleSpec0(customData.customIPBlock)
	}

	if customData.customSelector != nil {
		flattenRuleSpecData[ruleSpecSelectorKey] = flattenRuleSpec1(customData.customSelector)
	}

	return flattenRuleSpecData
}

func flattenRuleSpec0(spec0 *policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec0) (data []interface{}) {
	if spec0 == nil {
		return data
	}

	flattenSpec0 := make(map[string]interface{})

	if spec0.IpBlock != nil {
		flattenSpec0[ipBlockKey] = flattenIP(spec0.IpBlock)
	}

	return []interface{}{flattenSpec0}
}

func flattenIP(ip *policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec0IpBlock) (data []interface{}) {
	if ip == nil {
		return data
	}

	flattenIPData := make(map[string]interface{})

	flattenIPData[cidrKey] = ip.Cidr
	flattenIPData[exceptKey] = *ip.Except

	if ip.Except != nil {
		flattenIPData[exceptKey] = *ip.Except
	}

	return []interface{}{flattenIPData}
}

func flattenRuleSpec1(spec1 *policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1CustomRulesRuleSpec1) (data []interface{}) {
	if spec1 == nil {
		return data
	}

	flattenSpec1 := make(map[string]interface{})

	if spec1.NamespaceSelector != nil {
		namespaceSelectors := make(map[string]interface{})

		for _, namespaceSelector := range *spec1.NamespaceSelector {
			namespaceSelectors[namespaceSelector.Key] = namespaceSelector.Value
		}

		flattenSpec1[namespaceSelectorKey] = namespaceSelectors
	}

	if spec1.PodSelector != nil {
		podSelectors := make(map[string]interface{})

		for _, podSelector := range *spec1.PodSelector {
			podSelectors[podSelector.Key] = podSelector.Value
		}

		flattenSpec1[podSelectorKey] = podSelectors
	}

	return []interface{}{flattenSpec1}
}
