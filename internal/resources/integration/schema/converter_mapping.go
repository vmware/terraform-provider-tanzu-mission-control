/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integrationschema

import (
	"encoding/json"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var TFModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	ScopeKey: &tfModelConverterHelper.BlockToStruct{
		ClusterScopeKey: &tfModelConverterHelper.BlockToStruct{
			NameKey:                  tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterName"),
			ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
			ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("fullName", "provisionerName"),
		},
		ClusterGroupScopeKey: &tfModelConverterHelper.BlockToStruct{
			NameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterGroupName"),
		},
	},
	common.MetaKey: common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	SpecKey: &tfModelConverterHelper.BlockToStruct{
		CredentialsKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "credentialName"),
		ConfigurationsKey: &tfModelConverterHelper.EvaluatedField{
			Field:    tfModelConverterHelper.BuildDefaultModelPath("spec", "configurations"),
			EvalFunc: evaluateConfigurations,
		},
		SecretsKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "secrets"),
	},
}

func evaluateConfigurations(mode tfModelConverterHelper.EvaluationMode, value interface{}) (reportData interface{}) {
	if mode == tfModelConverterHelper.ConstructModel {
		reportData = make(map[string]interface{})
		_ = json.Unmarshal([]byte(value.(string)), &reportData)
	} else {
		reportJSONBytes, _ := json.Marshal(value)
		reportData = helper.ConvertToString(reportJSONBytes, "")
	}

	return reportData
}
