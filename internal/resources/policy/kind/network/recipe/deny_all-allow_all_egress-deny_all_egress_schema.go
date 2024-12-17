// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var DenyAll = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for network policy deny-all recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{},
	},
}

var AllowAllEgress = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for network policy allow-all-egress recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{},
	},
}

var DenyAllEgress = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for network policy deny-all-egress recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{},
	},
}
