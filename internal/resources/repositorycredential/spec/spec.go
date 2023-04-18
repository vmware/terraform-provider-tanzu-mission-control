/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/credential"
)

var SpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the Repository Credential.",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			secretTypeKey: {
				Type:        schema.TypeString,
				Description: "The type of credential that will be used for the repository. Options are SSH or USERNAME_PASSWORD",
				Required:    true,
			},
			dataKey: credential.KeyValueSpec,
		},
	},
}

var secretSpec = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the repository credential secret",
	Required:    true,
	MaxItems:    1,
	MinItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			usernamePasswordKey: {
				Type:        schema.TypeList,
				Required:    true,
				Description: "SecretType definition - SECRET_TYPE_DOCKERCONFIGJSON, Kubernetes secrets type.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						usernameKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Username of the registry.",
							Required:    true,
							ValidateFunc: validation.All(
								validation.StringLenBetween(1, 126),
								validation.StringIsNotEmpty,
								validation.StringIsNotWhiteSpace,
							),
						},
						passwordKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Password of the registry.",
							Required:    true,
							Sensitive:   true,
							ValidateFunc: validation.All(
								validation.StringLenBetween(1, 126),
								validation.StringIsNotEmpty,
								validation.StringIsNotWhiteSpace,
							),
						},
					},
				},
			},
			sshKey: {
				Type:        schema.TypeList,
				Required:    true,
				Description: "SecretType definition - SECRET_TYPE_DOCKERCONFIGJSON, Kubernetes secrets type.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						usernameKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Username of the registry.",
							Required:    true,
							ValidateFunc: validation.All(
								validation.StringLenBetween(1, 126),
								validation.StringIsNotEmpty,
								validation.StringIsNotWhiteSpace,
							),
						},
						passwordKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Password of the registry.",
							Required:    true,
							Sensitive:   true,
							ValidateFunc: validation.All(
								validation.StringLenBetween(1, 126),
								validation.StringIsNotEmpty,
								validation.StringIsNotWhiteSpace,
							),
						},
					},
				},
			},
		},
	},
}
