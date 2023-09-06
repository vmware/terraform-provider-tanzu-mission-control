/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package listpackages

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	packagespec "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/package/spec"
)

var (
	PackageSchema = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				nameKey: {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "Name of the package.This denotes semantic version of the package.",
				},
				SpecKey: packagespec.SpecSchema,
			},
		},
	}
)
