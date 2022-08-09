/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
)

var NamespaceSelector = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Label based Namespace Selector for the policy",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			MatchExpressionsKey: {
				Type:        schema.TypeList,
				Description: "Match expressions is a list of label selector requirements, the requirements are ANDed",
				Required:    true,
				Elem:        labelSelectorRequirement,
			},
		},
	},
}

func ConstructNamespaceSelector(data []interface{}) (namespaceSelector *policymodel.VmwareTanzuManageV1alpha1CommonPolicyLabelSelector) {
	if len(data) == 0 || data[0] == nil {
		return namespaceSelector
	}

	namespaceSelectorData, _ := data[0].(map[string]interface{})

	namespaceSelector = &policymodel.VmwareTanzuManageV1alpha1CommonPolicyLabelSelector{
		MatchExpressions: make([]*policymodel.K8sIoApimachineryPkgApisMetaV1LabelSelectorRequirement, 0),
	}

	if v, ok := namespaceSelectorData[MatchExpressionsKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			for _, raw := range vs {
				namespaceSelector.MatchExpressions = append(namespaceSelector.MatchExpressions, constructLabelSelectorRequirement(raw))
			}
		}
	}

	return namespaceSelector
}

func FlattenNamespaceSelector(namespaceSelector *policymodel.VmwareTanzuManageV1alpha1CommonPolicyLabelSelector) (data []interface{}) {
	if namespaceSelector == nil {
		return data
	}

	flattenNamespaceSelector := make(map[string]interface{})

	mes := make([]interface{}, 0)

	for _, me := range namespaceSelector.MatchExpressions {
		mes = append(mes, flattenLabelSelectorRequirement(me))
	}

	flattenNamespaceSelector[MatchExpressionsKey] = mes

	return []interface{}{flattenNamespaceSelector}
}
