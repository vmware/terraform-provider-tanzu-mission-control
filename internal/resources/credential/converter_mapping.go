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
	NameKey:        tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "fullname", "name"),
	common.MetaKey: common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	specKey: &tfModelConverterHelper.BlockToStruct{
		capabilityKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "capability"),
		providerKey:   tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "meta", "provider"),
		dataKey: &tfModelConverterHelper.BlockToStruct{
			genericCredentialKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "data", "genericCredential"),
			awsCredentialKey: &tfModelConverterHelper.BlockToStruct{
				awsAccountIDKey:      tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "data", "awsCredential", "accountId"),
				genericCredentialKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "data", "awsCredential", "genericCredential"),
				awsIAMRoleKey: &tfModelConverterHelper.BlockToStruct{
					iamRoleARNKey:   tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "data", "awsCredential", "iamRole", "arn"),
					iamRoleExtIDKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "data", "awsCredential", "iamRole", "extId"),
				},
			},
			keyValueKey: &tfModelConverterHelper.BlockToStruct{
				dataKey: &tfModelConverterHelper.EvaluatedField{
					Field:    tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "data", "keyValue", "data"),
					EvalFunc: tfModelConverterHelper.EvaluationFunc(keyValueEvalFunc),
				},
				typeKey: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "spec", "data", "keyValue", "type"),
			},
		},
	},
	statusKey: &tfModelConverterHelper.Map{
		tfModelConverterHelper.AllMapKeysFieldMarker: tfModelConverterHelper.BuildModelPath(tfModelConverterHelper.DefaultModelPathSeparator, "status", tfModelConverterHelper.AllMapKeysFieldMarker),
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
