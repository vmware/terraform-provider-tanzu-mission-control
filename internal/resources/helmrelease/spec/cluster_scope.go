/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	releaseclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/cluster"
)

func ConstructSpecForClusterScope(d *schema.ResourceData) (spec *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec, err error) {
	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec, nil
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec, nil
	}

	specData := data[0].(map[string]interface{})

	spec = &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec{}

	if inlineConfigValueFile, ok := specData[InlineConfigKey]; ok {
		if !(fileExists(inlineConfigValueFile.(string))) {
			return spec, errors.Errorf("File %s does not exists.", inlineConfigValueFile)
		}

		spec.InlineConfiguration, err = readYamlFile(inlineConfigValueFile.(string))
		if err != nil {
			return spec, nil
		}
	}

	if targerNamespaceName, ok := specData[TargetNamespaceKey]; ok {
		helper.SetPrimitiveValue(targerNamespaceName, &spec.TargetNamespace, TargetNamespaceKey)
	}

	if intervalValue, ok := specData[IntervalKey]; ok {
		helper.SetPrimitiveValue(intervalValue, &spec.Interval, IntervalKey)
	}

	if ref, ok := specData[ChartRefKey]; ok {
		if refData, ok := ref.([]interface{}); ok {
			spec.ChartRef = expandRef(refData)
		}
	}

	return spec, nil
}

func FlattenSpecForClusterScope(spec *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	flattenSpecData[TargetNamespaceKey] = spec.TargetNamespace
	flattenSpecData[InlineConfigKey] = spec.InlineConfiguration
	flattenSpecData[IntervalKey] = spec.Interval

	var chartRef = make(map[string]interface{})

	if *spec.ChartRef.RepositoryType == *releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeGIT) {
		var gitRepoType = make(map[string]interface{})

		gitRepoType[RepositoryNameKey] = spec.ChartRef.RepositoryName
		gitRepoType[RepositoryNamespaceNameKey] = spec.ChartRef.RepositoryNamespace
		gitRepoType[ChartPathKey] = spec.ChartRef.Chart

		chartRef[GitRepositoryKey] = []interface{}{gitRepoType}
	}

	if *spec.ChartRef.RepositoryType == *releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeHELM) {
		var helmRepoType = make(map[string]interface{})

		helmRepoType[RepositoryNameKey] = spec.ChartRef.RepositoryName
		helmRepoType[RepositoryNamespaceNameKey] = spec.ChartRef.RepositoryNamespace
		helmRepoType[ChartNameKey] = spec.ChartRef.Chart
		helmRepoType[VersionKey] = spec.ChartRef.Version

		chartRef[HelmRepositorykey] = []interface{}{helmRepoType}
	}

	flattenSpecData[ChartRefKey] = []interface{}{chartRef}

	return []interface{}{flattenSpecData}
}
