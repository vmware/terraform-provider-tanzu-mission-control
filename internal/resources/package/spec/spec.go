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
				ReleaseNotesKey: releaseNotesSchema,
				RepositoryNameKey: {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "Name of package repository to which this package belongs.",
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
				RawKey: rawSchema,
			},
		},
	}

	releaseNotesSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Release notes of Package.",
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				metadataNameKey: {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "metadata name.",
				},
				versionKey: {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "version selection.",
				},
				urlKey: {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "url for the package",
				},
			},
		},
	}

	rawSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Raw is the underlying serialization of object.",
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				examplesKey: {
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							namespaceKey: {
								Type:        schema.TypeString,
								Description: "Namespace name.",
								Computed:    true,
							},
						},
					},
				},
				propertiesKey: {
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							namespaceKey: {
								Type:     schema.TypeList,
								Computed: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										defaultKey: {
											Type:        schema.TypeString,
											Description: "Default value of object.",
											Computed:    true,
										},
										descriptionKey: {
											Type:        schema.TypeString,
											Description: "Description of the object of object.",
											Computed:    true,
										},
										typeKey: {
											Type:        schema.TypeString,
											Description: "Type of object.",
											Computed:    true,
										},
									},
								},
							},
						},
					},
				},
				titleKey: {
					Type:        schema.TypeString,
					Description: "Title of object.",
					Computed:    true,
				},
			},
		},
	}
)
