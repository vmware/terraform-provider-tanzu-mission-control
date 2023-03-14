/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	gitrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/gitrepository/cluster"
)

func ConstructSpecForClusterScope(d *schema.ResourceData) (spec *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec) {
	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	spec = &gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec{}

	if v, ok := specData[URLKey]; ok {
		helper.SetPrimitiveValue(v, &spec.Interval, URLKey)
	}

	if v, ok := specData[secretRefKey]; ok {
		helper.SetPrimitiveValue(v, &spec.SecretRef, secretRefKey)
	}

	if v, ok := specData[intervalKey]; ok {
		helper.SetPrimitiveValue(v, &spec.Interval, intervalKey)
	}

	if v, ok := specData[gitImplementationKey]; ok {
		spec.GitImplementation = gitrepositoryclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(v.(string)))
	}

	if ref, ok := specData[refKey]; ok {
		if refData, ok := ref.([]interface{}); ok {
			spec.Ref = expandRef(refData)
		}
	}

	return spec
}

func FlattenSpecForClusterScope(spec *gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositorySpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	flattenSpecData[URLKey] = spec.URL
	flattenSpecData[secretRefKey] = spec.SecretRef
	flattenSpecData[intervalKey] = spec.Interval

	if spec.GitImplementation != nil {
		flattenSpecData[gitImplementationKey] = string(*spec.GitImplementation)
	}

	if spec.Ref != nil {
		flattenSpecData[refKey] = flattenRef(spec.Ref)
	}

	return []interface{}{flattenSpecData}
}
