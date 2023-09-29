/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	URLKey  = "url"
	SpecKey = "spec"
)

var SpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the Helm Repository.",
	Computed:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			URLKey: {
				Type:        schema.TypeString,
				Description: "URL of helm repository.",
				Computed:    true,
			},
		},
	},
}
