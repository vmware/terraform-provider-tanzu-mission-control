/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var (
	SpecSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Spec for the Repository.",
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				LicensesKey: {
					Type:        schema.TypeSet,
					Computed:    true,
					Description: "Licenses under which Package is released.",
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				ReleasedAtKey: {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "Date on which Package is released.",
				},
				CapacityRequirementsDescriptionKey: {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "Minimum capacity requirements to install Package on a cluster.",
				},
				ReleaseNotesKey: {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "Release notes of Package.",
				},
				ValuesSchemaKey: valueSchema,
			},
		},
	}

	valueSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Values schema is used to show template values that can be configured by users.",
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				TemplateKey: templateSchema,
			},
		},
	}

	templateSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Template values in OpenAPI V3 schema format.",
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				RawKey: {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "Raw is the underlying serialization of object.",
				},
			},
		},
	}
)
