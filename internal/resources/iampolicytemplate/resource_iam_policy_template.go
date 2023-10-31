/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package iampolicytemplate

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIAMPolicyTemplate() *schema.Resource {
	return &schema.Resource{
		Schema: iamPolicyTemplateResourceSchema,
	}
}
