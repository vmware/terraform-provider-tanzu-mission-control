/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package credential

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var keyValueSpec = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			typeKey: {
				Type:     schema.TypeString,
				Default:  "SECRET_TYPE_UNSPECIFIED",
				Optional: true,
			},
			dataKey: {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	},
}
