// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package spec

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
)

var (
	SpecSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Spec for source secret.",
		Required:    true,
		MinItems:    1,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				DataKey: dataSchema,
			},
		},
	}

	specsAllowed = [...]string{UsernamePasswordKey, SSHKey}

	dataSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The schema for spec credential type.",
		Required:    true,
		MinItems:    1,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				UsernamePasswordKey: usernamePasswordDataSpec,
				SSHKey:              sshDataSpec,
			},
		},
	}

	usernamePasswordDataSpec = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The schema for Username/Password credential type spec.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				usernameKey: {
					Type:        schema.TypeString,
					Description: "Username for the basic authorization.",
					Required:    true,
				},
				PasswordKey: {
					Type:        schema.TypeString,
					Description: "Password for the basic authorization.",
					Required:    true,
					Sensitive:   true,
				},
			},
		},
	}

	sshDataSpec = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The schema for SSH credential type spec.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				IdentityKey: {
					Type:        schema.TypeString,
					Description: "SSH Identity file.",
					Required:    true,
					Sensitive:   true,
				},
				KnownhostsKey: {
					Type:        schema.TypeString,
					Description: "Known Hosts file path.",
					Required:    true,
				},
			},
		},
	}
)

func HasSpecChanged(d *schema.ResourceData) bool {
	updateRequired := false

	switch {
	case d.Get(helper.GetFirstElementOf(SpecKey, DataKey, UsernamePasswordKey)) != nil:
		if d.HasChange(helper.GetFirstElementOf(SpecKey, DataKey, UsernamePasswordKey, usernameKey)) || d.HasChange(helper.GetFirstElementOf(SpecKey, DataKey, UsernamePasswordKey, PasswordKey)) {
			updateRequired = true
		}

		fallthrough
	case d.Get(helper.GetFirstElementOf(SpecKey, DataKey, SSHKey)) != nil:
		if d.HasChange(helper.GetFirstElementOf(SpecKey, DataKey, SSHKey, IdentityKey)) || d.HasChange(helper.GetFirstElementOf(SpecKey, DataKey, SSHKey, KnownhostsKey)) {
			updateRequired = true
		}
	}

	return updateRequired
}

func ValidateSpec(_ context.Context, diff *schema.ResourceDiff, i interface{}) error {
	value, ok := diff.GetOk(SpecKey)
	if !ok {
		return fmt.Errorf("scope: %v is not valid: minimum one valid scope block is required", value)
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("spec data: %v is not valid: minimum one valid spec type block is required among: %v", data, strings.Join(specsAllowed[:], `, `))
	}

	specData := data[0].(map[string]interface{})

	specesFound := make([]string, 0)

	if data, ok := specData[DataKey]; ok {
		if v1, ok := data.([]interface{}); ok && len(v1) != 0 {
			if len(v1) == 0 || v1[0] == nil {
				return fmt.Errorf("spec data: %v is not valid: minimum one valid spec type block is required among: %v", data, strings.Join(specsAllowed[:], `, `))
			}

			specType := v1[0].(map[string]interface{})

			if usernamePassword, ok := specType[UsernamePasswordKey]; ok {
				if v1, ok := usernamePassword.([]interface{}); ok && len(v1) != 0 {
					specesFound = append(specesFound, UsernamePasswordKey)
				}
			}

			if ssh, ok := specType[SSHKey]; ok {
				if v1, ok := ssh.([]interface{}); ok && len(v1) != 0 {
					specesFound = append(specesFound, SSHKey)
				}
			}
		}
	}

	if len(specesFound) == 0 {
		return fmt.Errorf("no valid spec type block found: minimum one valid spec type block is required among: %v", strings.Join(specsAllowed[:], `, `))
	} else if len(specesFound) > 1 {
		return fmt.Errorf("found spec types: %v are not valid: maximum one valid spec type block is allowed", strings.Join(specesFound, `, `))
	}

	return nil
}
