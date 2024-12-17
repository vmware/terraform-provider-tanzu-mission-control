// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmfeatureclustermodel

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus Status of the Helm.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.helm.Status
type VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus struct {

	// Conditions of the Helm feature resource.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdHelmStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
