/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionsmodels

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfo Contains the metadata for a single report
// (e.g. report id, etc).
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.ReportInfo
type VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfo struct {

	// Kubernetes Server Version.
	KubeServerVersion string `json:"kubeServerVersion,omitempty"`

	// Number of inspections failed.
	NumFailed string `json:"numFailed,omitempty"`

	// Total number of inspections as part of the scan.
	NumInspections string `json:"numInspections,omitempty"`

	// Number of inspection tests in warning state.
	NumWarning string `json:"numWarning,omitempty"`

	// Progress information about the inspection scan.
	ProgressInfo *VmwareTanzuManageV1alpha1ClusterInspectionScanProgressInfo `json:"progressInfo,omitempty"`

	// Internal ID of the run.
	ReportID string `json:"reportId,omitempty"`

	// Result is a success / failure condition based on the result of the scan.
	Result *VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfoResult `json:"result,omitempty"`

	// Date and time of the run.
	// Format: date-time
	RunDatetime strfmt.DateTime `json:"runDatetime,omitempty"`

	// The scan type.
	ScanType string `json:"scanType,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfo) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfo

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
