/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var Small = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for namespace quota policy small recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{},
	},
}

var Medium = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for namespace quota policy medium recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{},
	},
}

var Large = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for namespace quota policy large recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{},
	},
}
