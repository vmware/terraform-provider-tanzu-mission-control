// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package credential

import (
	"encoding/base64"

	"github.com/go-openapi/strfmt"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	credentialsmodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/credential"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey:        tfModelConverterHelper.BuildDefaultModelPath("fullname", "name"),
	common.MetaKey: common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	specKey: &tfModelConverterHelper.BlockToStruct{
		capabilityKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "capability"),
		providerKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "meta", "provider"),
		dataKey: &tfModelConverterHelper.BlockToStruct{
			genericCredentialKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "genericCredential"),
			awsCredentialKey: &tfModelConverterHelper.BlockToStruct{
				awsAccountIDKey:      tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "awsCredential", "accountId"),
				genericCredentialKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "awsCredential", "genericCredential"),
				awsIAMRoleKey: &tfModelConverterHelper.BlockToStruct{
					iamRoleARNKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "awsCredential", "iamRole", "arn"),
					iamRoleExtIDKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "awsCredential", "iamRole", "extId"),
				},
			},
			azureCredentialKey: &tfModelConverterHelper.BlockToStruct{
				servicePrincipalKey: &tfModelConverterHelper.BlockToStruct{
					subscriptionIDKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipal", "subscriptionId"),
					tenantIDKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipal", "tenantId"),
					resourceGroupKey:  tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipal", "resourceGroup"),
					clientIDKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipal", "clientId"),
					clientSecretKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipal", "clientSecret"),
					azureCloudNameKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipal", "azureCloudName"),
				},
				servicePrincipalWithCertKey: &tfModelConverterHelper.BlockToStruct{
					subscriptionIDKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipalWithCertificate", "subscriptionId"),
					tenantIDKey:             tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipalWithCertificate", "tenantId"),
					clientIDKey:             tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipalWithCertificate", "clientId"),
					clientCertificateKey:    tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipalWithCertificate", "clientCertificate"),
					azureCloudNameKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipalWithCertificate", "azureCloudName"),
					managedSubscriptionsKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "azureCredential", "servicePrincipalWithCertificate", "managedSubscriptions"),
				},
			},
			keyValueKey: &tfModelConverterHelper.BlockToStruct{
				dataKey: &tfModelConverterHelper.EvaluatedField{
					Field:    tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "keyValue", "data"),
					EvalFunc: tfModelConverterHelper.EvaluationFunc(keyValueEvalFunc),
				},
				typeKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "data", "keyValue", "type"),
			},
		},
	},
	statusKey: &tfModelConverterHelper.Map{
		tfModelConverterHelper.AllMapKeysFieldMarker: tfModelConverterHelper.BuildDefaultModelPath("status", tfModelConverterHelper.AllMapKeysFieldMarker),
	},
}

func keyValueEvalFunc(mode tfModelConverterHelper.EvaluationMode, valToEvaluate interface{}) interface{} {
	var returnValue interface{}

	if valToEvaluate != nil {
		if mode == tfModelConverterHelper.ConstructModel {
			returnValue = make(map[string]strfmt.Base64)

			for key, value := range valToEvaluate.(map[string]interface{}) {
				returnValue.(map[string]strfmt.Base64)[key] = []byte(value.(string))
			}
		} else {
			returnValue = make(map[string]interface{})

			for key, value := range valToEvaluate.(map[string]interface{}) {
				valueBytes, _ := base64.StdEncoding.DecodeString(value.(string))
				returnValue.(map[string]interface{})[key] = helper.ConvertToString(valueBytes, "")
			}
		}
	}

	return returnValue
}

var tfModelResourceConverter = tfModelConverterHelper.TFSchemaModelConverter[*credentialsmodels.VmwareTanzuManageV1alpha1AccountCredentialCredential]{
	TFModelMap: tfModelResourceMap,
}
