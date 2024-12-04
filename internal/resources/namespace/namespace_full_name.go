// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package namespace

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	namespacemodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/namespace"
)

var NamespaceFullname = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The schema for namespace iam policy full name",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			NameKey: {
				Type:        schema.TypeString,
				Description: "Name of the Namespace",
				Required:    true,
				ForceNew:    true,
			},
			ManagementClusterNameKey: {
				Type:        schema.TypeString,
				Description: "Name of ManagementCluster",
				Default:     attachedValue,
				Optional:    true,
				ForceNew:    true,
			},
			ProvisionerNameKey: {
				Type:        schema.TypeString,
				Description: "Name of Provisioner",
				Default:     attachedValue,
				Optional:    true,
				ForceNew:    true,
			},
			ClusterNameKey: {
				Type:        schema.TypeString,
				Description: "Name of Cluster",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

func ConstructNamespaceFullname(data []interface{}) (fullname *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName{}

	if v, ok := fullNameData[ManagementClusterNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ManagementClusterName, ManagementClusterNameKey)
	}

	if v, ok := fullNameData[ProvisionerNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ProvisionerName, ProvisionerNameKey)
	}

	if v, ok := fullNameData[ClusterNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ClusterName, ClusterNameKey)
	}

	if v, ok := fullNameData[NameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.Name, NameKey)
	}

	return fullname
}

func FlattenNamespaceFullname(fullname *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[ManagementClusterNameKey] = fullname.ManagementClusterName
	flattenFullname[ProvisionerNameKey] = fullname.ProvisionerName
	flattenFullname[ClusterNameKey] = fullname.ClusterName
	flattenFullname[NameKey] = fullname.Name

	return []interface{}{flattenFullname}
}
