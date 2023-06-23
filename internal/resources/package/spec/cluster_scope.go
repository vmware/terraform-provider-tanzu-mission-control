/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"encoding/base64"

	"github.com/go-openapi/strfmt"

	packageclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/package/cluster"
)

func FlattenSpecForClusterScope(spec *packageclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageMetadataPackageSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	var templatedata = make(map[string]interface{})

	var valuedata = make(map[string]interface{})

	flattenSpecData[ReleasedAtKey] = (spec.ReleasedAt).String()
	flattenSpecData[CapacityRequirementsDescriptionKey] = spec.CapacityRequirementsDescription
	flattenSpecData[ReleaseNotesKey] = spec.ReleaseNotes
	flattenSpecData[RepositoryNameKey] = spec.RepositoryName

	if len(spec.Licenses) > 0 {
		flattenSpecData[LicensesKey] = spec.Licenses
	}

	rawdata, err := getDecodedSpecData(spec.ValuesSchema.Template.Raw)
	if err != nil {
		return data
	}

	templatedata[RawKey] = rawdata

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
