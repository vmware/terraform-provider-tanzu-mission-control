/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var StatusSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "status for package repository.",
	Computed:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			packageRepositoryPhaseKey: {
				Type:        schema.TypeString,
				Description: "One-word reason for the condition's last transition.",
			},
			subscribedKey: {
				Type:        schema.TypeBool,
				Description: "If true, the Package Repository has been subscribed by user organization.",
			},
			disabledKey: {
				Type:        schema.TypeBool,
				Description: "If true, the Package Repository is disabled.",
			},
			managedKey: {
				Type:        schema.TypeBool,
				Description: "If true, the Package Repository is managed by TMC.",
			},
		},
	},
}
