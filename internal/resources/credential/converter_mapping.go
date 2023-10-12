/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

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
	NameKey:        "fullName.name",
	common.MetaKey: common.MetaConverterMap,
	specKey: &tfModelConverterHelper.BlockToStruct{
		capabilityKey: "spec.capability",
		providerKey:   "spec.meta.provider",
		dataKey: &tfModelConverterHelper.BlockToStruct{
			genericCredentialKey: "spec.data.genericCredential",
			awsCredentialKey: &tfModelConverterHelper.BlockToStruct{
				awsAccountIDKey:      "spec.data.awsCredential.accountId",
				genericCredentialKey: "spec.data.awsCredential.genericCredential",
				awsIAMRoleKey: &tfModelConverterHelper.BlockToStruct{
					iamRoleARNKey:   "spec.data.awsCredential.iamRole.arn",
					iamRoleExtIDKey: "spec.data.awsCredential.iamRole.extId",
				},
			},
			keyValueKey: &tfModelConverterHelper.BlockToStruct{
				dataKey: &tfModelConverterHelper.EvaluatedField{
					Field:    "spec.data.keyValue.data",
					EvalFunc: tfModelConverterHelper.EvaluationFunc(keyValueEvalFunc),
				},
				typeKey: "spec.data.keyValue.type",
			},
		},
	},
	statusKey: &tfModelConverterHelper.Map{
		"*": "status.*",
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
