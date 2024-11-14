// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/gitrepository/spec"
)

const (
	OpaqueKey           = "opaque"
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
			OpaqueKey: {
				Type:        schema.TypeMap,
				Description: "SecretType definition - SECRET_TYPE_OPAQUE, Kubernetes secrets type.",
				Optional:    true,
				Sensitive:   true,

				Elem: &schema.Schema{Type: schema.TypeString},
			},
			DockerConfigjsonKey: {
				Type:     schema.TypeList,
				Optional: true,

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

	if d.HasChange(helper.GetFirstElementOf(SpecKey, OpaqueKey)) {
		updateRequired = true
	}

	return updateRequired
}

func ValidateInput(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
	value, ok := diff.GetOk(spec.SpecKey)
	if !ok {
		return fmt.Errorf("spec: %v is not valid: minimum one valid spec block is required", value)
	}

	data, _ := value.([]interface{})

	specData := data[0].(map[string]interface{})
	secretTypes := []string{
		OpaqueKey,
		DockerConfigjsonKey,
	}
	secretTypesFound := make([]string, 0)

	for _, secret := range secretTypes {
		if secretData, ok := specData[secret]; ok {
			if secret == OpaqueKey {
				converted := secretData.(map[string]interface{})
				if len(converted) != 0 {
					secretTypesFound = append(secretTypesFound, secret)
				}
			}

			if secret == DockerConfigjsonKey {
				converted := secretData.([]interface{})
				if len(converted) != 0 {
					secretTypesFound = append(secretTypesFound, secret)
				}
			}
		}
	}

	if len(secretTypesFound) == 0 {
		return fmt.Errorf("no valid spec block found: minimum one valid secret block is required among: %v", strings.Join(secretTypes, `, `))
	} else if len(secretTypesFound) > 1 {
		return fmt.Errorf("found secret blocks: %v are not valid: maximum one valid secret block is allowed", strings.Join(secretTypesFound, `, `))
	}

	return nil
}
