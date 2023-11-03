/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspections

import (
	"encoding/json"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	tfModelConverterHelper "github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper/converter"
	inspectionsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/inspections"
)

var tfInspectionModelMap = &tfModelConverterHelper.BlockToStruct{
	NameKey:                  tfModelConverterHelper.BuildDefaultModelPath("fullName", "name"),
	ClusterNameKey:           tfModelConverterHelper.BuildDefaultModelPath("fullName", "clusterName"),
	ManagementClusterNameKey: tfModelConverterHelper.BuildDefaultModelPath("fullName", "managementClusterName"),
	ProvisionerNameKey:       tfModelConverterHelper.BuildDefaultModelPath("fullName", "provisionerName"),
	StatusKey: &tfModelConverterHelper.Map{
		PhaseKey:     tfModelConverterHelper.BuildDefaultModelPath("status", "phase"),
		PhaseInfoKey: tfModelConverterHelper.BuildDefaultModelPath("status", "phaseInfo"),
		ReportKey: &tfModelConverterHelper.EvaluatedField{
			Field:    tfModelConverterHelper.BuildDefaultModelPath("status", "report"),
			EvalFunc: evaluateReport,
		},
		TarballDownloadURL: tfModelConverterHelper.BuildDefaultModelPath("status", "tarballDownloadUrl"),
	},
}

func evaluateReport(mode tfModelConverterHelper.EvaluationMode, value interface{}) (reportData interface{}) {
	if mode == tfModelConverterHelper.ConstructTFSchema {
		reportJSONBytes, _ := json.Marshal(value)
		reportData = helper.ConvertToString(reportJSONBytes, "")
	}

	return reportData
}

var tfInspectionModelConverter = tfModelConverterHelper.TFSchemaModelConverter[*inspectionsmodel.VmwareTanzuManageV1alpha1ClusterInspectionScan]{
	TFModelMap: tfInspectionModelMap,
}

var tfInspectionListModelMap = &tfModelConverterHelper.BlockToStruct{
	InspectionListKey: &tfModelConverterHelper.BlockSliceToStructSlice{
		// UNPACK tfModelResourceMap HERE.
	},
	TotalCountKey: "totalCount",
}

var tfListModelConverter = tfModelConverterHelper.TFSchemaModelConverter[*inspectionsmodel.VmwareTanzuManageV1alpha1ClusterInspectionScanListData]{
	TFModelMap: tfInspectionListModelMap,
}

func constructTFListModelDataMap() {
	tfListModelSchema := tfInspectionModelConverter.UnpackSchema(tfModelConverterHelper.BuildArrayField("scans"))

	statusKey := (*tfListModelSchema)[StatusKey]
	(*tfListModelSchema)[StatusKey] = statusKey.(*tfModelConverterHelper.Map).Copy([]string{TarballDownloadURL})

	*(*tfInspectionListModelMap)[InspectionListKey].(*tfModelConverterHelper.BlockSliceToStructSlice) = append(
		*(*tfInspectionListModelMap)[InspectionListKey].(*tfModelConverterHelper.BlockSliceToStructSlice),
		tfListModelSchema,
	)
}
