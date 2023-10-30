/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package customiamrole

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCustomIAMRole() *schema.Resource {
	return &schema.Resource{
		Schema: customIAMRoleResourceSchema,
	}
}
