/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	secretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
)

var ClusterFullname = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The schema for cluster secret full name",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ManagementClusterNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the management cluster",
				Default:     AttachedValue,
				Optional:    true,
				ForceNew:    true,
			},
			ProvisionerNameKey: {
				Type:        schema.TypeString,
				Description: "Provisioner of the cluster",
				Default:     AttachedValue,
				Optional:    true,
				ForceNew:    true,
			},
			ClusterNameKey: {
				Type:        schema.TypeString,
				Description: "Name of this cluster",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

func ConstructClusterFullname(data []interface{}, name string) (fullname *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName{}

	if v, ok := fullNameData[ManagementClusterNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ManagementClusterName, ManagementClusterNameKey)
	}

	if v, ok := fullNameData[ProvisionerNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ProvisionerName, ProvisionerNameKey)
	}

	if v, ok := fullNameData[ClusterNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ClusterName, ClusterNameKey)
	}

	fullname.Name = name

	return fullname
}

func FlattenClusterFullname(fullname *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[ManagementClusterNameKey] = fullname.ManagementClusterName
	flattenFullname[ProvisionerNameKey] = fullname.ProvisionerName
	flattenFullname[ClusterNameKey] = fullname.ClusterName

	return []interface{}{flattenFullname}
}
