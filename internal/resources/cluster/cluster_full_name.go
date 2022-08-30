/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
)

var ClusterFullname = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The schema for cluster iam policy full name",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ManagementClusterNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the management cluster",
				Default:     attachedValue,
				Optional:    true,
				ForceNew:    true,
			},
			ProvisionerNameKey: {
				Type:        schema.TypeString,
				Description: "Provisioner of the cluster",
				Default:     attachedValue,
				Optional:    true,
				ForceNew:    true,
			},
			NameKey: {
				Type:        schema.TypeString,
				Description: "Name of this cluster",
				Required:    true,
				ForceNew:    true,
			},
		},
	},
}

func ConstructClusterFullname(data []interface{}) (fullname *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &clustermodel.VmwareTanzuManageV1alpha1ClusterFullName{}

	if v, ok := fullNameData[ManagementClusterNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ManagementClusterName, ManagementClusterNameKey)
	}

	if v, ok := fullNameData[ProvisionerNameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.ProvisionerName, ProvisionerNameKey)
	}

	if v, ok := fullNameData[NameKey]; ok {
		helper.SetPrimitiveValue(v, &fullname.Name, NameKey)
	}

	return fullname
}

func FlattenClusterFullname(fullname *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[ManagementClusterNameKey] = fullname.ManagementClusterName
	flattenFullname[ProvisionerNameKey] = fullname.ProvisionerName
	flattenFullname[NameKey] = fullname.Name

	return []interface{}{flattenFullname}
}
