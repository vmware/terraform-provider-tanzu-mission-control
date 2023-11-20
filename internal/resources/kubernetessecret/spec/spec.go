/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
)

const (
	DockerConfigjsonKey = "docker_config_json"
	DockerconfigKey     = ".dockerconfigjson"
	ImageRegistryURLKey = "image_registry_url"
	UsernameKey         = "username"
	PasswordKey         = "password"
	SpecKey             = "spec"
)

type dockerConfigJSON struct {
	Auths map[string]map[string]interface{} `json:"auths,omitempty"`
}

var SecretSpec = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the kubernetes secret",
	Required:    true,
	MaxItems:    1,
	MinItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			DockerConfigjsonKey: {
				Type:        schema.TypeList,
				Required:    true,
				Description: "SecretType definition - SECRET_TYPE_DOCKERCONFIGJSON, Kubernetes secrets type.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						UsernameKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Username of the registry.",
							Required:    true,
							ValidateFunc: validation.All(
								validation.StringLenBetween(1, 126),
								validation.StringIsNotEmpty,
								validation.StringIsNotWhiteSpace,
							),
						},
						PasswordKey: {
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
						ImageRegistryURLKey: {
							Type:        schema.TypeString,
							Description: "SecretType definition - Server URL of the registry.",
							Required:    true,
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

func HasSpecChanged(d *schema.ResourceData) bool {
	updateRequired := false

	if d.HasChange(helper.GetFirstElementOf(SpecKey, DockerConfigjsonKey, UsernameKey)) || d.HasChange(helper.GetFirstElementOf(SpecKey, DockerConfigjsonKey, PasswordKey)) {
		updateRequired = true
	}

	return updateRequired
}
