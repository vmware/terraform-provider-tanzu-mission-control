/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ChartsKey       = "charts"
	SpecKey         = "spec"
	APIVersionKey   = "api_version"
	AppVersionKey   = "app_version"
	DeprecatedKey   = "deprecated"
	KubeVersionKey  = "kube_version"
	ReleasedAtKey   = "released_at"
	SourcesKey      = "sources"
	UrlsKey         = "urls"
	ValuesConfigKey = "values_config"
	DependenciesKey = "dependencies"

	AliasKey        = "alias"
	ChartNameKey    = "chart_name"
	ChartVersionKey = "chart_version"
	ConditionKey    = "condition"
	ImportValuesKey = "import_values"
	RepositoryKey   = "repository"
	TagsKey         = "tags"

	nameKey              = "name"
	chartMetadataNameKey = "chart_metadata_name"
)

var Charts = &schema.Schema{
	Type:        schema.TypeList,
	Description: "List Helm charts.",
	Computed:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			nameKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of helm chart such as 0.5.1",
			},
			chartMetadataNameKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the helm chart.",
			},
			SpecKey: specSchema,
		},
	},
}

var dependencySchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "List of the chart requirements.",
	Computed:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			AliasKey: {
				Type:        schema.TypeString,
				Description: "Alias to be used for the chart.",
				Computed:    true,
			},
			ChartNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the chart.",
				Computed:    true,
			},
			ChartVersionKey: {
				Type:        schema.TypeBool,
				Description: "Version of the chart.",
				Computed:    true,
			},
			ConditionKey: {
				Type:        schema.TypeString,
				Description: "Yaml path that resolves to a boolean, used for enabling/disabling charts.",
				Computed:    true,
			},
			ImportValuesKey: {
				Type:        schema.TypeList,
				Description: "Holds the mapping of source values to parent key to be imported.",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			RepositoryKey: {
				Type:        schema.TypeList,
				Description: "Repository URL.",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			TagsKey: {
				Type:        schema.TypeList,
				Description: "Tags can be used to group charts for enabling/disabling together.",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	},
}

var specSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the Helm chart.",
	Computed:    true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			APIVersionKey: {
				Type:        schema.TypeString,
				Description: "The chart API version.",
				Computed:    true,
			},
			AppVersionKey: {
				Type:        schema.TypeString,
				Description: "Application version of the chart.",
				Computed:    true,
			},
			DeprecatedKey: {
				Type:        schema.TypeBool,
				Description: "Whether this chart is deprecated.",
				Computed:    true,
			},
			KubeVersionKey: {
				Type:        schema.TypeString,
				Description: "A SemVer range of compatible Kubernetes versions.",
				Computed:    true,
			},
			ReleasedAtKey: {
				Type:        schema.TypeString,
				Description: "Date on which helm chart is released.",
				Computed:    true,
			},
			SourcesKey: {
				Type:        schema.TypeList,
				Description: "List of URLs to source code for this project.",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			UrlsKey: {
				Type:        schema.TypeList,
				Description: "List of URLs to download helm chart bundle.",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			ValuesConfigKey: {
				Type:        schema.TypeString,
				Description: "Default configuration values for this chart.",
				Computed:    true,
			},
			DependenciesKey: dependencySchema,
		},
	},
}
