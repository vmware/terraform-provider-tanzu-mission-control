// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

var DenyAllToPods = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for network policy deny-all-to-pods recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ToPodLabelsKey: toPodLabel,
		},
	},
}

func ConstructDenyAllToPods(data []interface{}) (denyAllToPods *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1DenyAllToPods) {
	if len(data) == 0 || data[0] == nil {
		return denyAllToPods
	}

	denyAllToPodsData, _ := data[0].(map[string]interface{})

	denyAllToPods = &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1DenyAllToPods{}

	if podLabelsData, ok := denyAllToPodsData[ToPodLabelsKey]; ok {
		if podLabels, ok := podLabelsData.(map[string]interface{}); ok {
			if len(podLabels) != 0 {
				labels := make([]*policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels, 0)
				denyAllToPods.ToPodLabels = labels

				for key, value := range podLabels {
					value := fmt.Sprintf("%v", value)

					labels = append(labels, &policyrecipenetworkcommonmodel.V1alpha1CommonPolicySpecNetworkV1Labels{
						Key:   key,
						Value: value,
					})
				}

				denyAllToPods.ToPodLabels = labels
			}
		}
	}

	return denyAllToPods
}

func FlattenDenyAllToPods(denyAllToPods *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1DenyAllToPods) (data []interface{}) {
	if denyAllToPods == nil {
		return data
	}

	flattenDenyAllToPods := make(map[string]interface{})

	if denyAllToPods.ToPodLabels != nil {
		labels := make(map[string]interface{})

		for _, label := range denyAllToPods.ToPodLabels {
			labels[label.Key] = label.Value
		}

		flattenDenyAllToPods[ToPodLabelsKey] = labels
	}

	return []interface{}{flattenDenyAllToPods}
}
