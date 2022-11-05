/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package credential

import (
	"encoding/json"
	"log"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
)

type CredMapping struct {
	Capability string `json:"capability,omitempty"`
	Data       []struct {
		AwsCredential []struct {
			AccountID         string `json:"account_id,omitempty"`
			GenericCredential string `json:"generic_credential,omitempty"`
			IamRole           []struct {
				Arn   string `json:"arn,omitempty"`
				ExtID string `json:"ext_id,omitempty"`
			} `json:"iam_role,omitempty"`
		} `json:"aws_credential,omitempty"`
		GenericCredential string `json:"generic_credential,omitempty"`
		KeyValue          []struct {
			Data map[string]string `json:"data,omitempty"`
			Type string            `json:"type,omitempty"`
		} `json:"key_value,omitempty"`
	} `json:"data,omitempty"`
	Provider string `json:"provider,omitempty"`
}

func constructSpec(d *schema.ResourceData) (spec *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialSpec) {
	spec = &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialSpec{
		Meta: &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialMeta{},
		Data: &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialData{},
	}

	value, ok := d.GetOk(specKey)
	if !ok {
		return spec
	}

	data := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	jsonData, err := json.Marshal(specData)
	if err != nil {
		log.Println("[ERROR] marshall, ", err.Error())
		return
	}

	credMapping := &CredMapping{}

	err = json.Unmarshal(jsonData, credMapping)
	if err != nil {
		log.Println("[ERROR] mapping, ", err.Error())
		return
	}

	spec.Capability = credMapping.Capability
	spec.Meta.Provider = credentialsmodels.NewVmwareTanzuManageV1alpha1AccountCredentialProvider(credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProvider(credMapping.Provider))

	switch len(credMapping.Data) != 0 {
	case false:
		return
	case len(credMapping.Data[0].AwsCredential) != 0:
		spec.Data.AwsCredential = &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialTypeAwsSpec{
			AccountID:         credMapping.Data[0].AwsCredential[0].AccountID,
			GenericCredential: credMapping.Data[0].AwsCredential[0].GenericCredential,
		}

		if len(credMapping.Data[0].AwsCredential[0].IamRole) != 0 {
			spec.Data.AwsCredential.IamRole = &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialTypeAwsIAMRole{
				Arn:   credMapping.Data[0].AwsCredential[0].IamRole[0].Arn,
				ExtID: credMapping.Data[0].AwsCredential[0].IamRole[0].ExtID,
			}
		}
	case len(credMapping.Data[0].KeyValue) != 0:
		data := make(map[string]strfmt.Base64)
		for key, value := range credMapping.Data[0].KeyValue[0].Data {
			data[key] = []byte(value)
		}

		spec.Data.KeyValue = &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec{
			Type: credentialsmodels.NewVmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType(credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType(credMapping.Data[0].KeyValue[0].Type)),
			Data: data,
		}
	case credMapping.Data[0].GenericCredential != "":
		spec.Data.GenericCredential = credMapping.Data[0].GenericCredential
	}

	return spec
}

func constructFullname(d *schema.ResourceData) (fullname *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialFullName) {
	fullname = &credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialFullName{}
	fullname.Name, _ = d.Get(NameKey).(string)

	return fullname
}

// nolint:deadcode,unused
func flattenSpec(spec *credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})
	flattenSpecData[capabilityKey] = spec.Capability
	flattenSpecData[providerKey] = credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialProviderPROVIDERUNSPECIFIED

	if spec.Meta != nil {
		flattenSpecData[providerKey] = spec.Meta.Provider
	}

	flattenSpecCredData := make(map[string]interface{})

	switch spec.Data != nil {
	case false:
		return
	case spec.Data.AwsCredential != nil:
		iamRole := map[string]interface{}{}
		if spec.Data.AwsCredential.IamRole != nil {
			iamRole[iamRoleARNKey] = spec.Data.AwsCredential.IamRole.Arn
			iamRole[iamRoleExtIDKey] = spec.Data.AwsCredential.IamRole.ExtID
		}

		flattenSpecCredData[awsCredentialKey] = []interface{}{map[string]interface{}{
			awsAccountIDKey:      spec.Data.AwsCredential.AccountID,
			genericCredentialKey: spec.Data.AwsCredential.GenericCredential,
			awsIAMRoleKey:        []interface{}{iamRole},
		}}
	case spec.Data.KeyValue != nil:
		data := make(map[string]string)
		for key, value := range spec.Data.KeyValue.Data {
			data[key] = string(value)
		}

		flattenSpecCredData[keyValueKey] = []interface{}{
			map[string]interface{}{
				typeKey: spec.Data.KeyValue.Type,
				dataKey: data,
			},
		}
	case spec.Data.GenericCredential != "":
		flattenSpecCredData[genericCredentialKey] = spec.Data.GenericCredential
	}

	flattenSpecData[dataKey] = []interface{}{flattenSpecCredData}

	return []interface{}{flattenSpecData}
}
