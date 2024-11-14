// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package status

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var StatusSchema = &schema.Schema{
	Type:        schema.TypeMap,
	Description: "Status for the Repository.",
	Computed:    true,
	Elem:        &schema.Schema{Type: schema.TypeString},
}
