/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionsmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpec CIS security inspection scan specification.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.CISSpec
type VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpec struct {

	// List of Targets that the CIS plugin will run against.
	CisTargets []*VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpecTargets `json:"cisTargets"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
