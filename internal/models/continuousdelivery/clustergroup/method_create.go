/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package continuousdeliveryclustergroupmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest Request to create a ContinuousDelivery.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.continuousdelivery.CreateContinuousDeliveryRequest
type VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest struct {

	// ContinuousDelivery to create.
	ContinuousDelivery *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery `json:"continuousDelivery,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse Response from creating a ContinuousDelivery.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.continuousdelivery.CreateContinuousDeliveryResponse
type VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse struct {

	// ContinuousDelivery created.
	ContinuousDelivery *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery `json:"continuousDelivery,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDeliveryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
