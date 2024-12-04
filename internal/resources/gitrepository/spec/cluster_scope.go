// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

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

	if URLValue, ok := specData[URLKey]; ok {
		helper.SetPrimitiveValue(URLValue, &spec.URL, URLKey)
	}

	if secretRefValue, ok := specData[secretRefKey]; ok {
		helper.SetPrimitiveValue(secretRefValue, &spec.SecretRef, secretRefKey)
	}

	if intervalValue, ok := specData[intervalKey]; ok {
		helper.SetPrimitiveValue(intervalValue, &spec.Interval, intervalKey)
	}

	if gitImplementationValue, ok := specData[gitImplementationKey]; ok {
		spec.GitImplementation = gitrepositoryclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementation(gitImplementationValue.(string)))
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

	// The response from the GitRepository Get API does not contain the git implementation field, but only if the value is GO_GIT. If the value is LIB_GIT2, the response contains the git implementation field.
	// Since the value of the GO_GIT type is set to 0 and omitempty flag is set, it gets dropped when the proto is converted to JSON.
	flattenSpecData[gitImplementationKey] = string(gitrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdGitrepositoryGitImplementationGOGIT)

	if spec.GitImplementation != nil {
		flattenSpecData[gitImplementationKey] = string(*spec.GitImplementation)
	}

	if spec.Ref != nil {
		flattenSpecData[refKey] = flattenRef(spec.Ref)
	}

	return []interface{}{flattenSpecData}
}
