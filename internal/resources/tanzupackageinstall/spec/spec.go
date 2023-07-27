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
				InlineValuesKey: {
					Type:        schema.TypeMap,
					Description: "Inline values to configure the Package Install.",
					Optional:    true,
					Sensitive:   true,
				},
			},
		},
	}

	PackageRefKeySpec = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Reference to the Package which will be installed.",
		Optional:    true,
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
		Optional:    true,
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
	case d.HasChange(helper.GetFirstElementOf(SpecKey, InlineValuesKey)):
		updateRequired = true
	}

	return updateRequired
}
