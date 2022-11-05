/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package credential

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var awsCredSpec = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			awsAccountIDKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			genericCredentialKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			awsIAMRoleKey: iamRoleSpec,
		},
	},
}

var iamRoleSpec = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			iamRoleARNKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
			iamRoleExtIDKey: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	},
}
