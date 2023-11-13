/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrole

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	customiamrolemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/customiamrole"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var (
	clusterRoleSelectorsArrayField = tfModelConverterHelper.BuildArrayField("clusterRoleSelectors")
	rulesArrayField                = tfModelConverterHelper.BuildArrayField("rules")
	matchExpressionsArrayField     = tfModelConverterHelper.BuildArrayField("matchExpressions")
)

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey:        tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	common.MetaKey: common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	SpecKey: &tfModelConverterHelper.BlockToStruct{
		IsDeprecatedKey:     tfModelConverterHelper.BuildDefaultModelPath("spec", "isDeprecated"),
		AllowedScopesKey:    tfModelConverterHelper.BuildDefaultModelPath("spec", "resources"),
		TanzuPermissionsKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "tanzuPermissions"),
		KubernetesPermissionsKey: &tfModelConverterHelper.BlockToStruct{
			RuleKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					APIGroupsKey:     tfModelConverterHelper.BuildDefaultModelPath("spec", rulesArrayField, "apiGroups"),
					URLPathsKey:      tfModelConverterHelper.BuildDefaultModelPath("spec", rulesArrayField, "nonResourceUrls"),
					ResourceNamesKey: tfModelConverterHelper.BuildDefaultModelPath("spec", rulesArrayField, "resourceNames"),
					ResourcesKey:     tfModelConverterHelper.BuildDefaultModelPath("spec", rulesArrayField, "resources"),
					VerbsKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", rulesArrayField, "verbs"),
				},
			},
		},
		AggregationRuleKey: &tfModelConverterHelper.BlockToStruct{
			ClusterRoleSelectorKey: &tfModelConverterHelper.BlockSliceToStructSlice{
				{
					MatchExpressionKey: &tfModelConverterHelper.BlockSliceToStructSlice{
						{
							MeKey:         tfModelConverterHelper.BuildDefaultModelPath("spec", "aggregationRule", clusterRoleSelectorsArrayField, matchExpressionsArrayField, "key"),
							MeOperatorKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "aggregationRule", clusterRoleSelectorsArrayField, matchExpressionsArrayField, "operator"),
							MeValuesKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "aggregationRule", clusterRoleSelectorsArrayField, matchExpressionsArrayField, "values"),
						},
					},
					MatchLabelsKey: &tfModelConverterHelper.Map{
						tfModelConverterHelper.AllMapKeysFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("spec", "aggregationRule", clusterRoleSelectorsArrayField, "matchLabels", tfModelConverterHelper.AllMapKeysFieldMarker),
					},
				},
			},
		},
	},
}

var tfModelConverter = tfModelConverterHelper.TFSchemaModelConverter[*customiamrolemodels.VmwareTanzuManageV1alpha1IamRole]{
	TFModelMap: tfModelResourceMap,
}
