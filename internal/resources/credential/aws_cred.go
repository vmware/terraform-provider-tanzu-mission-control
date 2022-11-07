/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package credential

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var awsCredSpec = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	Description: "AWS credential data type",
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			awsAccountIDKey: {
				Description: "Account ID of the AWS credential",
				Type:        schema.TypeString,
				Optional:    true,
			},
			genericCredentialKey: {
				Description: "Generic credential",
				Type:        schema.TypeString,
				Optional:    true,
			},
			awsIAMRoleKey: iamRoleSpec,
		},
	},
}

var iamRoleSpec = &schema.Schema{
	Type:        schema.TypeList,
	Description: "AWS IAM role ARN and external ID",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			iamRoleARNKey: {
				Description: "AWS IAM role ARN",
				Type:        schema.TypeString,
				Optional:    true,
			},
			iamRoleExtIDKey: {
				Description: "An external ID used to assume an AWS IAM role",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	},
}
