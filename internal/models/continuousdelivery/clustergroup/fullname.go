// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package continuousdeliveryclustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryFullName Full name of the Continuous Delivery.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.continuousdelivery.FullName
type VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryFullName struct {

	// Name of cluster group.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
