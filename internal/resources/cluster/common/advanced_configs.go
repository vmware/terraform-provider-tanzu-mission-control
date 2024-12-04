// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	clustercommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/common"
)

const (
	advancedConfigurationKey      = "key"
	advancedConfigurationValueKey = "value"
)

var AdvancedConfigs = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Advanced configuration for TKGm cluster",
	Optional:    true,
	MinItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			advancedConfigurationKey: {
				Type:        schema.TypeString,
				Description: "The key of the advanced configuration parameters",
				Required:    true,
			},
			advancedConfigurationValueKey: {
				Type:        schema.TypeString,
				Description: "The value of the advanced configuration parameters",
				Required:    true,
			},
		},
	},
}

func ExpandAdvancedConfig(data interface{}) (advancedConfig *clustercommon.VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig) {
	lookUpAdvancedConfig, ok := data.(map[string]interface{})
	if !ok {
		return advancedConfig
	}

	advancedConfig = &clustercommon.VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig{}

	if v, ok := lookUpAdvancedConfig[advancedConfigurationKey]; ok {
		advancedConfig.Key, _ = v.(string)
	}

	if v, ok := lookUpAdvancedConfig[advancedConfigurationValueKey]; ok {
		advancedConfig.Value, _ = v.(string)
	}

	return advancedConfig
}

func FlattenAdvancedConfig(advancedConfig *clustercommon.VmwareTanzuManageV1alpha1CommonClusterAdvancedConfig) (data interface{}) {
	flattenAdvancedConfig := make(map[string]interface{})

	if advancedConfig == nil {
		return nil
	}

	flattenAdvancedConfig[advancedConfigurationKey] = advancedConfig.Key
	flattenAdvancedConfig[advancedConfigurationValueKey] = advancedConfig.Value

	return flattenAdvancedConfig
}
