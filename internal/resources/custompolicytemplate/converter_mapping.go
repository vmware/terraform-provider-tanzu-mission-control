/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package custompolicytemplate

import (
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	custompolicytemplatemodels "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/custompolicytemplate"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

var (
	dataInventoryArrayField = tfModelConverterHelper.BuildArrayField("dataInventory")
)

var tfModelResourceMap = &tfModelConverterHelper.BlockToStruct{
	NameKey:        tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	common.MetaKey: common.GetMetaConverterMap(tfModelConverterHelper.DefaultModelPathSeparator),
	SpecKey: &tfModelConverterHelper.BlockToStruct{
		IsDeprecatedKey:     tfModelConverterHelper.BuildDefaultModelPath("spec", "deprecated"),
		TemplateManifestKey: tfModelConverterHelper.BuildDefaultModelPath("spec", "object"),
		ObjectTypeKey:       tfModelConverterHelper.BuildDefaultModelPath("spec", "objectType"),
		TemplateTypeKey:     tfModelConverterHelper.BuildDefaultModelPath("spec", "templateType"),
		DataInventoryKey: &tfModelConverterHelper.BlockSliceToStructSlice{
			{
				GroupKey:   tfModelConverterHelper.BuildDefaultModelPath("spec", dataInventoryArrayField, "group"),
				VersionKey: tfModelConverterHelper.BuildDefaultModelPath("spec", dataInventoryArrayField, "kind"),
				KindKey:    tfModelConverterHelper.BuildDefaultModelPath("spec", dataInventoryArrayField, "version"),
			},
		},
	},
}

var tfModelConverter = tfModelConverterHelper.TFSchemaModelConverter[*custompolicytemplatemodels.VmwareTanzuManageV1alpha1PolicyTemplate]{
	TFModelMap: tfModelResourceMap,
}
