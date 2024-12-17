// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	pkgrepositoryclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzupackagerepository"
)

func ConstructSpecForClusterScope(d *schema.ResourceData) (spec *pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec) {
	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	spec = &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec{
		ImgpkgBundle: &pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositoryImgPkgBundleSpec{},
	}

	if data, ok := specData[ImgpkgBundleKey]; ok {
		if v1, ok := data.([]interface{}); ok && len(v1) != 0 {
			specData := v1[0].(map[string]interface{})

			var image string

			if v, ok := specData[ImageKey]; ok {
				image = v.(string)
			}

			spec.ImgpkgBundle.Image = image
		}
	}

	return spec
}

func FlattenSpecForClusterScope(spec *pkgrepositoryclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageRepositorySpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	image := spec.ImgpkgBundle.Image

	var specImgpkgBundle = make(map[string]interface{})

	specImgpkgBundle[ImageKey] = image

	flattenSpecData[ImgpkgBundleKey] = []interface{}{specImgpkgBundle}

	return []interface{}{flattenSpecData}
}
