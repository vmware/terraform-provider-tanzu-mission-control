/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package continuousdeliveryclustermodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest Request to create a ContinuousDelivery.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.continuousdelivery.CreateContinuousDeliveryRequest
type VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest struct {

	// ContinuousDelivery to create.
	ContinuousDelivery *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDelivery `json:"continuousDelivery,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse Response from creating a ContinuousDelivery.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.continuousdelivery.CreateContinuousDeliveryResponse
type VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse struct {

	// ContinuousDelivery created.
	ContinuousDelivery *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDelivery `json:"continuousDelivery,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDeliveryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
