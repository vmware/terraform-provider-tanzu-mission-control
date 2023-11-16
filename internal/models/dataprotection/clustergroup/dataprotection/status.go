/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package dataprotectionclustergroupmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatus Status of the DataProtection configure resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.dataprotection.Status.
type VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatus struct {
	// Generation value at the time this status was updated.
	ObservedGeneration string `json:"observedGeneration,omitempty"`

	// Phase of the Cluster Group Data Protection on member Clusters.
	Phase *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusPhase `json:"phase,omitempty"`

	// Details contains information about the Cluster Group Data Protection being applied on member Clusters.
	Details *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatusDetails `json:"details,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupDataprotectionStatus

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
