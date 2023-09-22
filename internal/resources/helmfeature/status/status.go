/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package status

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var StatusSchema = &schema.Schema{
	Type:        schema.TypeMap,
	Description: "Status for the helm feature.",
	Computed:    true,
	Elem:        &schema.Schema{Type: schema.TypeString},
}
