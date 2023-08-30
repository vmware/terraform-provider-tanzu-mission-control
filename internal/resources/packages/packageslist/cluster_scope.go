/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package listpackages

import (
	packageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/package/cluster"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/package/spec"
)

func FlattenSpecForClusterScope(resp *packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageListPackagesResponse) (packages []interface{}) {
	if resp == nil {
		return packages
	}

	for _, j := range resp.Packages {
		data := make(map[string]interface{})
		val := spec.FlattenSpecForClusterScope(j.Spec)

		data[nameKey] = j.FullName.Name
		data[SpecKey] = val

		packages = append(packages, data)
	}

	return packages
}
