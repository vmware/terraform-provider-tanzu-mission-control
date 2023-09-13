/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var StatusSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "status for package install.",
	Computed:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			phaseKey: {
				Type:        schema.TypeString,
				Description: "One-word reason for the condition's last transition.",
				Computed:    true,
			},
			generatedResourcesKey: generatedResourcesStatus,
		},
	},
}

var generatedResourcesStatus = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Kuberenetes RBAC resources and service account created on the cluster by TMC for Package Install.",
	Computed:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			clusterRoleNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the cluster role used for Package Install.",
				Computed:    true,
			},
			serviceAccountNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the role binding used for Package Install.",
				Computed:    true,
			},
			roleBindingNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the service account used for Package Install.",
				Computed:    true,
			},
		},
	},
}
