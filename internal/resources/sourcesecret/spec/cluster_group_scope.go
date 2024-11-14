// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	sourcesecretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/clustergroup"
)

func ConstructSpecForClusterGroupScope(d *schema.ResourceData) (spec *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec) {
	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	spec = &sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec{}
	spec.AtomicSpec = ConstructSpecForClusterScope(d)

	return spec
}

func FlattenSpecForClusterGroupScope(spec *sourcesecretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec, specTypeData string) (data []interface{}) {
	if spec == nil || spec.AtomicSpec == nil {
		return data
	}

	return FlattenSpecForClusterScope(spec.AtomicSpec, specTypeData)
}
