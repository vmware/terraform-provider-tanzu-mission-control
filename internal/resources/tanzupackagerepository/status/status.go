/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var StatusSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "",
	Computed:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			packageRepositoryPhaseKey: {
				Type:        schema.TypeString,
				Description: "",
			},
			subscribedKey: {
				Type:        schema.TypeBool,
				Description: "",
			},
			disabledKey: {
				Type:        schema.TypeBool,
				Description: "",
			},
			managedKey: {
				Type:        schema.TypeBool,
				Description: "",
			},
		},
	},
}
