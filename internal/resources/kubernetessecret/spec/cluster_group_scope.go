/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	secertclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup"
)

func ConstructSpecForClusterGroupScope(d *schema.ResourceData) (spec *secertclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec) {
	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	spec = &secertclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec{}
	spec.AtomicSpec = ConstructSpecForClusterScope(d)

	return spec
}

func FlattenSpecForClusterGroupScope(spec *secertclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec, password string, opaqueData map[string]interface{}) (data []interface{}) {
	if spec == nil || spec.AtomicSpec == nil {
		return data
	}

	return FlattenSpecForClusterScope(spec.AtomicSpec, password, opaqueData)
}
