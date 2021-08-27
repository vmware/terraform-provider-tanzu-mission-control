/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespacemodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceStatus Status of the namespace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.Status
type VmwareTanzuManageV1alpha1ClusterNamespaceStatus struct {

	// available_phases is a list of available phases for namespace
	AvailablePhases []*VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase `json:"availablePhases"`

	// phase of the namespace.
	Phase *VmwareTanzuManageV1alpha1ClusterNamespaceStatusPhase `json:"phase,omitempty"`

	// phase_info contains additional info about the phase
	PhaseInfo string `json:"phaseInfo,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
