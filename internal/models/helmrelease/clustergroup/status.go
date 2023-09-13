/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmreleaseclustergroupmodel

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus Status of the Release.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.fluxcd.helm.release.Status
type VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus struct {

	// Details contains information about the Cluster Group helm release being applied on member Clusters.
	Details *statusmodel.VmwareTanzuManageV1alpha1CommonBatchDetails `json:"details,omitempty"`

	// Generation value at the time this status was updated.
	ObservedGeneration string `json:"observedGeneration,omitempty"`

	// Phase of the Cluster Group helm release application on member Clusters.
	Phase *statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhase `json:"phase,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceFluxcdHelmReleaseStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
