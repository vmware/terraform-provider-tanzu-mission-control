/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspections

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// Inspection Object Keys.
	ClusterNameKey           = "cluster_name"
	ManagementClusterNameKey = "management_cluster_name"
	ProvisionerNameKey       = "provisioner_name"
	NameKey                  = "name"
	StatusKey                = "status"

	// Inspection Object Status Keys.
	PhaseKey           = "phase"
	PhaseInfoKey       = "phase_info"
	ReportKey          = "report"
	TarballDownloadURL = "tarball_download_url"
)

var clusterNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Cluster name.",
	Required:    true,
	ForceNew:    true,
}

var managementClusterNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Management cluster name.",
	Required:    true,
	ForceNew:    true,
}

var provisionerNameSchema = &schema.Schema{
	Type:        schema.TypeString,
	Description: "Cluster provisioner name.",
	Required:    true,
	ForceNew:    true,
}

func getNameSchema(required bool) (nameSchema *schema.Schema) {
	nameSchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "Inspection name.",
		Optional:    !required,
		Required:    required,
		ForceNew:    true,
	}

	return nameSchema
}

var computedInspectionSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Inspection objects.",
	Computed:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			ClusterNameKey: {
				Type:        schema.TypeString,
				Description: "Cluster name.",
				Computed:    true,
			},
			ManagementClusterNameKey: {
				Type:        schema.TypeString,
				Description: "Management cluster name.",
				Computed:    true,
			},
			ProvisionerNameKey: {
				Type:        schema.TypeString,
				Description: "Provisioner name.",
				Computed:    true,
			},
			NameKey: {
				Type:        schema.TypeString,
				Description: "Inspection name.",
				Computed:    true,
			},
			StatusKey: {
				Type:        schema.TypeMap,
				Description: "Status of inspection resource",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	},
}
