/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionsmodels

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult Result describes the status of the inspection.
//
//   - RESULT_UNSPECIFIED: Unspecified result indicates the scan result is unknown.
//   - SUCCESS: Success - to be used if all the tests are part of the scan are successful.
//   - FAILURE: Failure - to indicate that one or more test as part of the scan failed.
//   - WARNING: Warning - to indicate that one or more test as part of the scan has a warning error.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.ReportInfo.Result
type VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult string

func NewVmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult(value VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult) *VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult.
func (m VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult) Pointer() *VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResultRESULTUNSPECIFIED captures enum value "RESULT_UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResultRESULTUNSPECIFIED VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult = "RESULT_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResultSUCCESS captures enum value "SUCCESS".
	VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResultSUCCESS VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult = "SUCCESS"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResultFAILURE captures enum value "FAILURE".
	VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResultFAILURE VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult = "FAILURE"

	// VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResultWARNING captures enum value "WARNING".
	VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResultWARNING VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult = "WARNING"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResultEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult

	if err := json.Unmarshal([]byte(`["RESULT_UNSPECIFIED","SUCCESS","FAILURE","WARNING"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResultEnum = append(vmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResultEnum, v)
	}
}
