/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
)

var AllowAll = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for network policy allow-all recipe version v1",
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
		},
	},
}

func ConstructAllowAll(data []interface{}) (allowAll *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAll) {
	if len(data) == 0 || data[0] == nil {
		return allowAll
	}

	allowAllData, _ := data[0].(map[string]interface{})

	allowAll = &policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAll{}

	if v, ok := allowAllData[FromOwnNamespaceKey]; ok {
		allowAll.FromOwnNamespace = helper.BoolPointer(v.(bool))
	}

	return allowAll
}

func FlattenAllowAll(allowAll *policyrecipenetworkmodel.V1alpha1CommonPolicySpecNetworkV1AllowAll) (data []interface{}) {
	if allowAll == nil {
		return data
	}

	flattenAllowAll := make(map[string]interface{})

	if allowAll.FromOwnNamespace == nil {
		return data
	}

	flattenAllowAll[FromOwnNamespaceKey] = *allowAll.FromOwnNamespace

	return []interface{}{flattenAllowAll}
}
