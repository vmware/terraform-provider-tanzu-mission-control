// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	kustomizationclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/clustergroup"
)

func ConstructSpecForClusterGroupScope(d *schema.ResourceData) (spec *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec) {
	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	spec = &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec{}
	spec.AtomicSpec = ConstructSpecForClusterScope(d)

	return spec
}

func FlattenSpecForClusterGroupScope(spec *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationSpec) (data []interface{}) {
	if spec == nil || spec.AtomicSpec == nil {
		return data
	}

	return FlattenSpecForClusterScope(spec.AtomicSpec)
}
