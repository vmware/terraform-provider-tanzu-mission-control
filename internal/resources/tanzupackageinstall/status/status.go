// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package status

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var StatusSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "status for package install.",
	Computed:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			packageInstallPhaseKey: {
				Type:        schema.TypeString,
				Description: "One-word reason for the condition's last transition.",
				Computed:    true,
			},
			resolvedVersionKey: {
				Type:        schema.TypeString,
				Description: "Resolved version of the Package Install.",
				Computed:    true,
			},
			managedKey: {
				Type:        schema.TypeBool,
				Description: "If true, the Package Install is managed by TMC.",
				Computed:    true,
			},
			referredByKey: {
				Type:        schema.TypeList,
				Description: "TMC services/features referencing the package install.",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
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
