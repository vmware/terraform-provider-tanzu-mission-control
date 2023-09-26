/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	helmchart "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmcharts"
)

func flattenSpec(spec *helmchart.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	dependencies := make([]interface{}, len(spec.Dependencies))

	flattenSpecData[ReleasedAtKey] = (spec.ReleasedAt).String()
	flattenSpecData[APIVersionKey] = spec.APIVersion
	flattenSpecData[AppVersionKey] = spec.AppVersion
	flattenSpecData[DeprecatedKey] = spec.Deprecated
	flattenSpecData[KubeVersionKey] = spec.KubeVersion
	flattenSpecData[KubeVersionKey] = spec.KubeVersion
	flattenSpecData[SourcesKey] = spec.Sources
	flattenSpecData[UrlsKey] = spec.Urls
	flattenSpecData[ValuesConfigKey] = spec.ValuesConfig

	for _, val := range spec.Dependencies {
		var dependency = make(map[string]interface{})

		dependency[AliasKey] = val.Alias
		dependency[ChartNameKey] = val.ChartName
		dependency[ChartVersionKey] = val.ChartVersion
		dependency[ConditionKey] = val.Condition
		dependency[ImportValuesKey] = val.ImportValues
		dependency[RepositoryKey] = val.Repository
		dependency[TagsKey] = val.Tags

		dependencies = append(dependencies, dependency)
	}

	flattenSpecData[DependenciesKey] = dependencies

	return []interface{}{flattenSpecData}
}

func FlattenCharts(resp *helmchart.VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartListResponse) (charts []interface{}) {
	if resp == nil {
		return charts
	}

	for _, chart := range resp.Charts {
		data := make(map[string]interface{})
		val := flattenSpec(chart.Spec)

		data[nameKey] = chart.FullName.Name
		data[chartMetadataNameKey] = chart.FullName.ChartMetadataName
		data[SpecKey] = val

		charts = append(charts, data)
	}

	return charts
}
