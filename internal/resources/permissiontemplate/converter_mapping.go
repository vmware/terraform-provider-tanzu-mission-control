// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package permissiontemplate

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	permissiontemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/permissiontemplate"
)

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	CredentialsNameKey:         tfModelConverterHelper.BuildDefaultModelPath("fullname", "name"),
	CapabilityKey:              "capability",
	ProviderKey:                "provider",
	TemplateKey:                "permissionTemplate",
	TemplateURLKey:             "templateUrl",
	TemplateValuesKey:          "templateValues",
	UndefinedTemplateValuesKey: "undefinedTemplateValues",
}

var tfModelRequestConverter = &tfModelConverterHelper.TFSchemaModelConverter[*permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateRequest]{
	TFModelMap: tfModelResourceMap,
}

var tfModelResponseConverter = &tfModelConverterHelper.TFSchemaModelConverter[*permissiontemplatemodels.VmwareTanzuManageV1alpha1AccountCredentialPermissionTemplateResponse]{
	TFModelMap: tfModelResourceMap,
}
