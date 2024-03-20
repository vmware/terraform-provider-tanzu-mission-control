/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
)

var (
	SpecSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "spec for package install.",
		Required:    true,
		MinItems:    1,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				PackageRefKey: PackageRefKeySpec,
				RoleBindingScopeKey: {
					Type:        schema.TypeString,
					Description: "Role binding scope for service account which will be used by Package Install.",
					Computed:    true,
				},
				PathToInlineValuesKey: {
					Type:        schema.TypeString,
					Description: "File to read inline values from (in yaml format). User needs to specify the file path for inline values.",
					Optional:    true,
					DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
						newInlineValues, err := helper.ReadYamlFile(new)
						if err != nil {
							return false
						}

						return old == newInlineValues
					},
				},
				InlineValuesKey: {
					Type:        schema.TypeMap,
					Description: "Inline values to configure the Package Install.",
					Optional:    true,
					Sensitive:   true,
					Deprecated:  "This field is deprecated. For providing the inline values, use the new field: path_to_inline_values",
				},
			},
		},
	}

	PackageRefKeySpec = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Reference to the Package which will be installed.",
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				PackageMetadataNameKey: {
					Type:        schema.TypeString,
					Description: "Name of the Package Metadata.",
					Required:    true,
					ForceNew:    true,
				},
				VersionSelectionKey: versionSelectionSpec,
			},
		},
	}

	versionSelectionSpec = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Version Selection of the Package.",
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				ConstraintsKey: {
					Type:        schema.TypeString,
					Description: "Constraints to select Package. Example: constraints: 'v1.2.3', constraints: '<v1.4.0' etc.",
					Required:    true,
				},
			},
		},
	}
)

func HasSpecChanged(d *schema.ResourceData) bool {
	updateRequired := false

	switch {
	case d.HasChange(helper.GetFirstElementOf(SpecKey, PackageRefKey, PackageMetadataNameKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, PackageRefKey, VersionSelectionKey, ConstraintsKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, RoleBindingScopeKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, PathToInlineValuesKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, InlineValuesKey)):
		updateRequired = true
	}

	return updateRequired
}
