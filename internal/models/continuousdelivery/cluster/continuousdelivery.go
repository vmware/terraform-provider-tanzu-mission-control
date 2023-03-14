/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package continuousdeliveryclustermodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDelivery Represents continuous delivery feature configuration for a cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.continuousdelivery.ContinuousDelivery
type VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDelivery struct {

	// Full name for the Continuous Delivery.
	FullName *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryFullName `json:"fullName,omitempty"`

	// Metadata for the Continuous Delivery object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Status for the Continuous Delivery.
	Status *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDelivery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDelivery) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDelivery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
