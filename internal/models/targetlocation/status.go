/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationmodels

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatus Status of the backup location resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.Status.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatus struct {

	// A list of available phases for backup location object.
	AvailablePhases []*VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase `json:"availablePhases"`

	// The Conditions attached to a backup location.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// The resource generation the current status applies to.
	ObservedGeneration string `json:"observedGeneration,omitempty"`

	// The overall phase of the backup location.
	Phase *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusPhase `json:"phase,omitempty"`

	// Additional info about the phase.
	PhaseInfo string `json:"phaseInfo,omitempty"`

	// Type of the backup location.
	Type *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatusType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationStatus

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
