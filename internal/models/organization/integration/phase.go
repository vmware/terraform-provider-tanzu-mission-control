/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package integration

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1OrganizationIntegrationPhase Integration Lifecycle Phase.
//
//   - UNSPECIFIED: Unspecified  phase.
//   - ENABLING: Enabling phase when process for enabling integration for an organization is started.
//   - REGISTERED: Registered phase when communication between inter product services is completed successfully for enabling integration.
//   - ENABLED: Enabled phase when integration is enabled for an organization.
//   - INITIATION_ERROR: Initiation Error phase when there is any issue during enablement of the integration.
//   - DISABLING: Disabling phase when process for removing integration from organization is started.
//   - UNREGISTERED: Unregistered phase when communication between inter product services is completed successfully
//
// for disabling integration.
//   - TERMINATION_ERROR: Termination Error phase when there is any issue during disablement of the integration.
//   - DISABLED: Disabled phase when integration is disabled from org / yet to be enabled.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.integration.Phase
type VmwareTanzuManageV1alpha1OrganizationIntegrationPhase string

func NewVmwareTanzuManageV1alpha1OrganizationIntegrationPhase(value VmwareTanzuManageV1alpha1OrganizationIntegrationPhase) *VmwareTanzuManageV1alpha1OrganizationIntegrationPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1OrganizationIntegrationPhase.
func (m VmwareTanzuManageV1alpha1OrganizationIntegrationPhase) Pointer() *VmwareTanzuManageV1alpha1OrganizationIntegrationPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseUNSPECIFIED captures enum value "UNSPECIFIED".
	VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseUNSPECIFIED VmwareTanzuManageV1alpha1OrganizationIntegrationPhase = "UNSPECIFIED"

	// VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseENABLING captures enum value "ENABLING".
	VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseENABLING VmwareTanzuManageV1alpha1OrganizationIntegrationPhase = "ENABLING"

	// VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseREGISTERED captures enum value "REGISTERED".
	VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseREGISTERED VmwareTanzuManageV1alpha1OrganizationIntegrationPhase = "REGISTERED"

	// VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseENABLED captures enum value "ENABLED".
	VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseENABLED VmwareTanzuManageV1alpha1OrganizationIntegrationPhase = "ENABLED"

	// VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseINITIATIONERROR captures enum value "INITIATION_ERROR".
	VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseINITIATIONERROR VmwareTanzuManageV1alpha1OrganizationIntegrationPhase = "INITIATION_ERROR"

	// VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseDISABLING captures enum value "DISABLING".
	VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseDISABLING VmwareTanzuManageV1alpha1OrganizationIntegrationPhase = "DISABLING"

	// VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseUNREGISTERED captures enum value "UNREGISTERED".
	VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseUNREGISTERED VmwareTanzuManageV1alpha1OrganizationIntegrationPhase = "UNREGISTERED"

	// VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseTERMINATIONERROR captures enum value "TERMINATION_ERROR".
	VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseTERMINATIONERROR VmwareTanzuManageV1alpha1OrganizationIntegrationPhase = "TERMINATION_ERROR"

	// VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseDISABLED captures enum value "DISABLED".
	VmwareTanzuManageV1alpha1OrganizationIntegrationPhaseDISABLED VmwareTanzuManageV1alpha1OrganizationIntegrationPhase = "DISABLED"
)

// for schema.
var vmwareTanzuManageV1alpha1OrganizationIntegrationPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1OrganizationIntegrationPhase
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","ENABLING","REGISTERED","ENABLED","INITIATION_ERROR","DISABLING","UNREGISTERED","TERMINATION_ERROR","DISABLED"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1OrganizationIntegrationPhaseEnum = append(vmwareTanzuManageV1alpha1OrganizationIntegrationPhaseEnum, v)
	}
}
