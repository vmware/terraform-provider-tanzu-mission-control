/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionsmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInspectionScanSpec Spec of the inspection scan.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.Spec
type VmwareTanzuManageV1alpha1ClusterInspectionScanSpec struct {

	// CIS security inspection scan specification.
	CisSpec *VmwareTanzuManageV1alpha1ClusterInspectionScanCISSpec `json:"cisSpec,omitempty"`

	// Conformance inspection scan specification.
	ConformanceSpec VmwareTanzuManageV1alpha1ClusterInspectionScanConformanceSpec `json:"conformanceSpec,omitempty"`

	// E2E inspection scan specification.
	E2eSpec VmwareTanzuManageV1alpha1ClusterInspectionScanE2ESpec `json:"e2eSpec,omitempty"`

	// Lite inspection scan specification.
	LiteSpec VmwareTanzuManageV1alpha1ClusterInspectionScanLiteSpec `json:"liteSpec,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInspectionScanSpec

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
