/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionsmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInspectionScanReport Encapsulates the data for a Inspection scan run.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.Report
type VmwareTanzuManageV1alpha1ClusterInspectionScanReport struct {

	// Map of all the diagnostic files in the report tarball.
	Diagnostic map[string]string `json:"diagnostic,omitempty"`

	// Map of files containing host information (config, healthz).
	Hosts map[string]string `json:"hosts,omitempty"`

	// Meta-info of this report.
	Info *VmwareTanzuManageV1alpha1ClusterInspectionScanReportInfo `json:"info,omitempty"`

	// Map of files with metadata information about the scan (Config, query time, run).
	Meta map[string]string `json:"meta,omitempty"`

	// Map of all the files ending in .xml.
	Results map[string]string `json:"results,omitempty"`

	// Download URL for the .tar.gz file with this full report.
	TarballDownloadURL string `json:"tarballDownloadUrl,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanReport) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInspectionScanReport

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
