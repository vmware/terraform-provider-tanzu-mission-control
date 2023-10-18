/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"encoding/base64"
	"strings"

	"github.com/go-openapi/strfmt"
	"gopkg.in/yaml.v2"

	packageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/package/cluster"
)

func FlattenSpecForClusterScope(spec *packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	var templatedata = make(map[string]interface{})

	var valuedata = make(map[string]interface{})

	var releaseNotes = make(map[string]interface{})

	flattenSpecData[ReleasedAtKey] = (spec.ReleasedAt).String()
	flattenSpecData[CapacityRequirementsDescriptionKey] = spec.CapacityRequirementsDescription
	releaseNotesValue := strings.Split(spec.ReleaseNotes, seperatorStr)

	releaseNotes[metadataNameKey] = releaseNotesValue[0]
	releaseNotes[versionKey] = releaseNotesValue[1]
	releaseNotes[urlKey] = releaseNotesValue[2]
	flattenSpecData[ReleaseNotesKey] = []interface{}{releaseNotes}
	flattenSpecData[RepositoryNameKey] = spec.RepositoryName

	if len(spec.Licenses) > 0 {
		flattenSpecData[LicensesKey] = spec.Licenses
	}

	rawdata, err := getDecodedSpecData(spec.ValuesSchema.Template.Raw)
	if err != nil {
		return data
	}

	var out RawData

	err1 := yaml.Unmarshal([]byte(rawdata), &out)
	if err1 != nil {
		return data
	}

	var rawSchemaData = make(map[string]interface{})

	rawSchemaData[examplesKey] = []interface{}{
		map[string]interface{}{
			namespaceKey: out.Examples[0].Namespace,
		},
	}

	rawSchemaData[propertiesKey] = []interface{}{
		map[string]interface{}{
			namespaceKey: []interface{}{
				map[string]interface{}{
					defaultKey:     out.Properties.Namespace.Default,
					descriptionKey: out.Properties.Namespace.Description,
					typeKey:        out.Properties.Namespace.Type,
				},
			},
		},
	}

	rawSchemaData[titleKey] = out.Title

	templatedata[RawKey] = []interface{}{rawSchemaData}

	valuedata[TemplateKey] = []interface{}{templatedata}

	flattenSpecData[ValuesSchemaKey] = []interface{}{valuedata}

	return []interface{}{flattenSpecData}
}

func getDecodedSpecData(data strfmt.Base64) (string, error) {
	rawData, err := base64.StdEncoding.DecodeString(data.String())
	if err != nil {
		return "", err
	}

	return string(rawData), nil
}
