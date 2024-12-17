// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var toPodLabel = &schema.Schema{
	Type:        schema.TypeMap,
	Description: "Pod Labels on which traffic should be allowed/denied. Use a label selector to identify the pods to which the policy applies.",
	Optional:    true,
	Elem: &schema.Schema{
		Type: schema.TypeString,
		ValidateFunc: validation.All(
			validation.StringLenBetween(1, 63),
			validation.StringIsValidRegExp,
		),
	},
}
