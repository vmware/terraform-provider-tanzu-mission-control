// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

const (
	AllowAllKey          = "allow_all"
	AllowAllToPodsKey    = "allow_all_to_pods"
	AllowAllEgressKey    = "allow_all_egress"
	DenyAllKey           = "deny_all"
	DenyAllToPodsKey     = "deny_all_to_pods"
	DenyAllEgressKey     = "deny_all_egress"
	CustomEgressKey      = "custom_egress"
	CustomIngressKey     = "custom_ingress"
	FromOwnNamespaceKey  = "from_own_namespace"
	ToPodLabelsKey       = "to_pod_labels"
	LabelKey             = "key"
	LabelValueKey        = "value"
	rulesKey             = "rules"
	portsKey             = "ports"
	portKey              = "port"
	protocolKey          = "protocol"
	ruleSpecKey          = "rule_spec"
	ruleSpecIPKey        = "custom_ip"
	ruleSpecSelectorKey  = "custom_selector"
	ipBlockKey           = "ip_block"
	cidrKey              = "cidr"
	exceptKey            = "except"
	namespaceSelectorKey = "namespace_selector"
	podSelectorKey       = "pod_selector"

	IPBlock           = "ipBlock"
	PodSelector       = "podSelector"
	NamespaceSelector = "namespaceSelector"
)
