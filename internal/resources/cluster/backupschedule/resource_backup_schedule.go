/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package backupschedule

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceBackupSchedule() *schema.Resource {
	return &schema.Resource{
		Schema: backupScheduleResourceSchema,
	}
}
