/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	kustomizationclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/cluster"
)

func ConstructSpecForClusterScope(d *schema.ResourceData) (spec *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec) {
	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	spec = &kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec{}

	if source, ok := specData[sourceKey]; ok {
		if sourceData, ok := source.([]interface{}); ok {
			spec.Source = expandSource(sourceData)
		}
	}

	if pathValue, ok := specData[pathKey]; ok {
		helper.SetPrimitiveValue(pathValue, &spec.Path, pathKey)
	}

	if pruneValue, ok := specData[pruneKey]; ok {
		helper.SetPrimitiveValue(pruneValue, &spec.Prune, pruneKey)
	}

	if intervalValue, ok := specData[intervalKey]; ok {
		helper.SetPrimitiveValue(intervalValue, &spec.Interval, intervalKey)
	}

	if targetNamespaceValue, ok := specData[targetNamespaceKey]; ok {
		helper.SetPrimitiveValue(targetNamespaceValue, &spec.TargetNamespace, targetNamespaceKey)
	}

	return spec
}

func FlattenSpecForClusterScope(spec *kustomizationclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdKustomizationSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	if spec.Source != nil {
		flattenSpecData[sourceKey] = flattenSource(spec.Source)
	}

	flattenSpecData[pathKey] = spec.Path
	flattenSpecData[pruneKey] = spec.Prune
	flattenSpecData[intervalKey] = spec.Interval
	flattenSpecData[targetNamespaceKey] = spec.TargetNamespace

	return []interface{}{flattenSpecData}
}
