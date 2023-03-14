/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	kustomizationclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kustomization/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func ConstructClusterGroupKustomizationFullname(data []interface{}, name, namespace, orgID string) (fullname *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName{}

	if v, ok := fullNameData[commonscope.NameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ClusterGroupName, commonscope.NameKey)
	}

	fullname.Name = name
	fullname.NamespaceName = namespace

	if orgID != "" {
		fullname.OrgID = orgID
	}

	return fullname
}

func FlattenClusterGroupKustomizationFullname(fullname *kustomizationclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdKustomizationFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[commonscope.NameKey] = fullname.ClusterGroupName

	return []interface{}{flattenFullname}
}
