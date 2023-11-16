/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustermodels

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClusterDataprotectionStatus Status of the DataProtection configure resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.dataprotection.Status.
type VmwareTanzuManageV1alpha1ClusterDataprotectionStatus struct {

	// A list of available phases for data protection object.
	AvailablePhases []*VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhase `json:"availablePhases"`

	// The conditions attached to this data protection object.
	// The description of the conditions is as follows:
	// - "Scheduled" with status 'Unknown' indicates the request has not been applied to the cluster yet.
	// - "Scheduled" with status 'True' and "Ready" with status 'Unknown' indicates the data protection create / delete intent has been applied / deleted but not yet acted upon.
	// - "Ready" with status 'True' indicates the the creation of data protection is complete.
	// - "Ready" with status 'False' indicates the the creation of data protection is in error state.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// The namespace used to install backup solution.
	Namespace string `json:"namespace,omitempty"`

	// The resource generation the current status applies to.
	ObservedGeneration string `json:"observedGeneration,omitempty"`

	// The overall phase of the data protection.
	Phase *VmwareTanzuManageV1alpha1ClusterDataprotectionStatusPhase `json:"phase,omitempty"`

	// Additional info about the phase.
	PhaseInfo string `json:"phaseInfo,omitempty"`

	// The version information of backup solution.
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterDataprotectionStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterDataprotectionStatus

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
