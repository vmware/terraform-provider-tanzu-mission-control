// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package credential

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var keyValueSpec = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	Description: "Key Value credential",
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			typeKey: {
				Type:        schema.TypeString,
				Description: "Type of Secret data, usually mapped to k8s secret type. Supported types: [SECRET_TYPE_UNSPECIFIED,OPAQUE_SECRET_TYPE,DOCKERCONFIGJSON_SECRET_TYPE]",
				Default:     "SECRET_TYPE_UNSPECIFIED",
				Optional:    true,
			},
			dataKey: {
				Type:        schema.TypeMap,
				Description: "Data secret data in the format of key-value pair",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	},
}
