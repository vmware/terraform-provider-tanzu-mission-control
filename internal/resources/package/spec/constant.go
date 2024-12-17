// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package spec

const (
	SpecKey                            = "spec"
	LicensesKey                        = "licenses"
	ReleasedAtKey                      = "released_at"
	CapacityRequirementsDescriptionKey = "capacity_requirements_description"
	ReleaseNotesKey                    = "release_notes"
	ValuesSchemaKey                    = "values_schema"
	TemplateKey                        = "template"
	RepositoryNameKey                  = "repository_name"
	RawKey                             = "raw"
	metadataNameKey                    = "metadata_name"
	versionKey                         = "version"
	urlKey                             = "url"
	namespaceKey                       = "namespace"
	examplesKey                        = "examples"
	propertiesKey                      = "properties"
	defaultKey                         = "default"
	descriptionKey                     = "description"
	typeKey                            = "type"
	titleKey                           = "title"

	seperatorStr = " "
)

type RawData struct {
	Examples   []ExampleData `json:"examples"`
	Properties NamespaceData `json:"properties"`
	Title      string        `json:"title"`
}

type ExampleData struct {
	Namespace string `json:"namespace"`
}

type NamespaceData struct {
	Namespace Data `json:"namespace"`
}

type Data struct {
	Default     string `json:"default"`
	Description string `json:"description"`
	Type        string `json:"type"`
}
