// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	secretclustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/clustergroup"
	commonscope "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common/scope"
)

func ConstructClusterGroupSecretFullname(data []interface{}, name, namespace string) (fullname *secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName{}

	if nameValue, ok := fullNameData[commonscope.NameKey]; ok {
		helper.SetPrimitiveValue(nameValue, &fullname.ClusterGroupName, commonscope.NameKey)
	}

	fullname.Name = name
	fullname.NamespaceName = namespace

	return fullname
}

func FlattenClusterGroupSecretFullname(fullname *secretclustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[commonscope.NameKey] = fullname.ClusterGroupName

	return []interface{}{flattenFullname}
}
