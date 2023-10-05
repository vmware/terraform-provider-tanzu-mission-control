/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"encoding/base64"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	sourcesecretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/cluster"
)

// nolint: nestif
func ConstructSpecForClusterScope(d *schema.ResourceData) (spec *sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec) {
	value, ok := d.GetOk(SpecKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	spec = &sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec{}

	if data, ok := specData[DataKey]; ok {
		if v1, ok := data.([]interface{}); ok && len(v1) != 0 {
			specType := v1[0].(map[string]interface{})

			if usernamePassword, ok := specType[UsernamePasswordKey]; ok {
				if v1, ok := usernamePassword.([]interface{}); ok && len(v1) != 0 {
					data := v1[0].(map[string]interface{})

					var username, password string

					if v, ok := data[usernameKey]; ok {
						username = v.(string)
					}

					if v, ok := data[PasswordKey]; ok {
						password = v.(string)
					}

					spec.SourceSecretType = sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(
						sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeUSERNAMEPASSWORD,
					)

					spec.Data = &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{}

					spec.Data.Type = sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType(
						sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretTypeOPAQUESECRETTYPE,
					)

					unamekeyData, _ := GetEncodedSpecData(username)
					passwordkeyData, _ := GetEncodedSpecData(password)
					specData := map[string]strfmt.Base64{
						usernameKey: unamekeyData,
						PasswordKey: passwordkeyData,
					}

					spec.Data.Data = specData

					return spec
				}
			}

			if ssh, ok := specType[SSHKey]; ok {
				if v1, ok := ssh.([]interface{}); ok && len(v1) != 0 {
					data := v1[0].(map[string]interface{})

					var identity, knownhosts string

					if v, ok := data[IdentityKey]; ok {
						identity = v.(string)
					}

					if v, ok := data[KnownhostsKey]; ok {
						knownhosts = v.(string)
					}

					spec.SourceSecretType = sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(
						sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeSSH,
					)

					spec.Data = &sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{}

					spec.Data.Type = sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType(
						sourcesecretclustermodel.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretTypeOPAQUESECRETTYPE,
					)

					identitydata, _ := GetEncodedSpecData(identity)
					knownhostsdata, _ := GetEncodedSpecData(knownhosts)
					specData := map[string]strfmt.Base64{
						IdentityKey:   identitydata,
						KnownhostsKey: knownhostsdata,
					}

					spec.Data.Data = specData

					return spec
				}
			}
		}
	}

	return spec
}

func FlattenSpecForClusterScope(spec *sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec, specTypeData string) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	var Data = make(map[string]interface{})

	if *spec.SourceSecretType == *sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeUSERNAMEPASSWORD) {
		username, ok := spec.Data.Data[usernameKey]
		if !ok {
			return data
		}

		var usernamePassword = make(map[string]interface{})

		unamedata, err := getDecodedSpecData(username)
		if err != nil {
			return data
		}

		usernamePassword[usernameKey] = unamedata
		usernamePassword[PasswordKey] = specTypeData

		Data[UsernamePasswordKey] = []interface{}{usernamePassword}
	}

	if *spec.SourceSecretType == *sourcesecretclustermodel.NewVmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretType(sourcesecretclustermodel.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretTypeSSH) {
		knownhosts, ok := spec.Data.Data[KnownhostsKey]
		if !ok {
			return data
		}

		var sshspecData = make(map[string]interface{})

		knownhostsdata, err := getDecodedSpecData(knownhosts)
		if err != nil {
			return data
		}

		sshspecData[IdentityKey] = specTypeData
		sshspecData[KnownhostsKey] = knownhostsdata

		Data[SSHKey] = []interface{}{sshspecData}
	}

	flattenSpecData[DataKey] = []interface{}{Data}

	return []interface{}{flattenSpecData}
}

func GetEncodedSpecData(data string) (strfmt.Base64, error) {
	var secretspecdata strfmt.Base64

	err := secretspecdata.Scan(base64.StdEncoding.EncodeToString([]byte(data)))
	if err != nil {
		return nil, err
	}

	return secretspecdata, nil
}

func getDecodedSpecData(data strfmt.Base64) (string, error) {
	rawData, err := base64.StdEncoding.DecodeString(data.String())
	if err != nil {
		return "", err
	}

	return string(rawData), nil
}
