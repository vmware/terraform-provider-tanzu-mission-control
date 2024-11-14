// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package spec

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var (
	SpecSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "spec for package repository.",
		Required:    true,
		MinItems:    1,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				ImgpkgBundleKey: imgpkgBundleKeySpec,
			},
		},
	}

	imgpkgBundleKeySpec = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Docker image url; unqualified, tagged, or digest references supported.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				ImageKey: {
					Type:        schema.TypeString,
					Description: "image url string.",
					Required:    true,
				},
			},
		},
	}
)
