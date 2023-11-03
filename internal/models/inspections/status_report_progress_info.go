/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionsmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInspectionScanProgressInfo Progress information about the inspection scan.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.ProgressInfo
type VmwareTanzuManageV1alpha1ClusterInspectionScanProgressInfo struct {

	// Number of tests completed.
	NumTestsCompleted string `json:"numTestsCompleted,omitempty"`

	// Number of tests run as part of the inspection scan.
	TotalNumTests string `json:"totalNumTests,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanProgressInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanProgressInfo) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInspectionScanProgressInfo

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
