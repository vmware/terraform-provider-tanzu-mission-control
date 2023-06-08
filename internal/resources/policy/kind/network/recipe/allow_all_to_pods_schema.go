/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

var AllowAllToPods = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for network policy allow-all-to-pods recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			FromOwnNamespaceKey: {
				Type:        schema.TypeBool,
				Description: "Allow traffic only from own namespace. Allow traffic only from pods in the same namespace as the destination pod.",
				Optional:    true,
				Default:     false,
			},
			ToPodLabelsKey: toPodLabel,
		},
	},
}

func ConstructAllowAllToPods(data []interface{}) (allowAllToPods *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAllToPods) {
	if len(data) == 0 || data[0] == nil {
		return allowAllToPods
	}

	allowAllToPodsData, _ := data[0].(map[string]interface{})

	allowAllToPods = &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAllToPods{}

	if v, ok := allowAllToPodsData[FromOwnNamespaceKey]; ok {
		allowAllToPods.FromOwnNamespace = helper.BoolPointer(v.(bool))
	}

	if podLabelsData, ok := allowAllToPodsData[ToPodLabelsKey]; ok {
		if podLabels, ok := podLabelsData.(map[string]interface{}); ok {
			if len(podLabels) != 0 {
				labels := make([]*policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels, 0)
				allowAllToPods.ToPodLabels = labels

				for key, value := range podLabels {
					value := fmt.Sprintf("%v", value)

					labels = append(labels, &policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
						Key:   key,
						Value: value,
					})
				}

				allowAllToPods.ToPodLabels = labels
			}
		}
	}

	return allowAllToPods
}

func FlattenAllowAllToPods(allowAllToPods *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAllToPods) (data []interface{}) {
	if allowAllToPods == nil {
		return data
	}

	flattenAllToPods := make(map[string]interface{})

	flattenAllToPods[FromOwnNamespaceKey] = *allowAllToPods.FromOwnNamespace

	if allowAllToPods.ToPodLabels != nil {
		labels := make(map[string]interface{})

		for _, label := range allowAllToPods.ToPodLabels {
			labels[label.Key] = label.Value
		}

		flattenAllToPods[ToPodLabelsKey] = labels
	}

	return []interface{}{flattenAllToPods}
}
