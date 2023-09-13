/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	releaseclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/clustergroup"
)

func ConstructSpecForClusterGroupScope(d *schema.ResourceData) (spec *releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseSpec, err error) {
	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec, nil
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec, nil
	}

	spec = &releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseSpec{}
	spec.AtomicSpec, err = ConstructSpecForClusterScope(d)

	return spec, err
}

func FlattenSpecForClusterGroupScope(spec *releaseclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseSpec) (data []interface{}) {
	if spec == nil || spec.AtomicSpec == nil {
		return data
	}

	return FlattenSpecForClusterScope(spec.AtomicSpec)
}
