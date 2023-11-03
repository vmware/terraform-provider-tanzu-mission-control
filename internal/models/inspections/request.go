/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionsmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInspectionScanData Request to create a Scan.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.CreateScanRequest
type VmwareTanzuManageV1alpha1ClusterInspectionScanData struct {

	// Scan to create.
	Scan *VmwareTanzuManageV1alpha1ClusterInspectionScan `json:"scan,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInspectionScanData

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterInspectionScanListData Response from listing Scans.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.ListScansResponse
type VmwareTanzuManageV1alpha1ClusterInspectionScanListData struct {

	// List of scans.
	Scans []*VmwareTanzuManageV1alpha1ClusterInspectionScan `json:"scans"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanListData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanListData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInspectionScanListData

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
