/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocation

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceTargetLocation() *schema.Resource {
	return &schema.Resource{
		Schema: backupTargetLocationResourceSchema,
	}
}
