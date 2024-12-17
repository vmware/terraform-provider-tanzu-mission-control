// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package continuousdeliveryclustergroupmodel

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryStatus Status of the Continuous Delivery.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.continuousdelivery.Status
type VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryStatus struct {

	// Details contains information about the Cluster Group continuous delivery being applied on member Clusters.
	Details *statusmodel.VmwareTanzuManageV1alpha1CommonBatchDetails `json:"details,omitempty"`

	// Phase of the Cluster Group continuous delivery feature on member Clusters.
	Phase *statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhase `json:"phase,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
