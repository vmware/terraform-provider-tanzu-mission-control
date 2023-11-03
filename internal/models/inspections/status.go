/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionsmodels

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClusterInspectionScanStatus Status of the scan inspection.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.inspection.scan.Status
type VmwareTanzuManageV1alpha1ClusterInspectionScanStatus struct {

	// Available phases of the inspection scan resource.
	AvailablePhases []*VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase `json:"availablePhases"`

	// Condition 'Scheduled' with Status 'Unknown' indicates that the inspection is pending
	// Condition 'Scheduled' with 'True' and Condition 'Ready' with 'Unknown' indicates that the inspection is running
	// Condition 'Ready' with 'True' indicates that the inspection is complete
	// Condition 'Ready' with 'False' indicates that the inspection is in error state.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Phase of the inspection scan based on conditions. If state is 'PHASE_UNSPECIFIED', use conditions to
	// interpret the state of the inspection.
	Phase *VmwareTanzuManageV1alpha1ClusterInspectionScanStatusPhase `json:"phase,omitempty"`

	// Additional information e.g., reason for ERROR state.
	PhaseInfo string `json:"phaseInfo,omitempty"`

	// Report details.
	Report *VmwareTanzuManageV1alpha1ClusterInspectionScanReport `json:"report,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInspectionScanStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInspectionScanStatus

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
