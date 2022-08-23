/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package scope

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/cluster"
)

var ClusterPolicyFullname = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The schema for cluster policy full name",
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

func ConstructClusterPolicyFullname(data []interface{}, name string) (fullname *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName) {
	if len(data) == 0 || data[0] == nil {
		return fullname
	}

	fullNameData, _ := data[0].(map[string]interface{})

	fullname = &policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName{}

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

func FlattenClusterPolicyFullname(fullname *policyclustermodel.VmwareTanzuManageV1alpha1ClusterPolicyFullName) (data []interface{}) {
	if fullname == nil {
		return data
	}

	flattenFullname := make(map[string]interface{})

	flattenFullname[ManagementClusterNameKey] = fullname.ManagementClusterName
	flattenFullname[ProvisionerNameKey] = fullname.ProvisionerName
	flattenFullname[ClusterNameKey] = fullname.ClusterName

	return []interface{}{flattenFullname}
}
