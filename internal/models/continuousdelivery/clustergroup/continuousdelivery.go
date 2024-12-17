// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package continuousdeliveryclustergroupmodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery Represents continuous delivery feature configuration for a cluster group.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.continuousdelivery.ContinuousDelivery
type VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery struct {

	// Full name for the Continuous Delivery.
	FullName *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryFullName `json:"fullName,omitempty"`

	// Metadata for the Continuous Delivery object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Status for the Continuous Delivery.
	Status *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
