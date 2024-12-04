// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool Nodepool associated with a EKS cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.Nodepool
type VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool struct {

	// Full name for the Nodepool.
	FullName *VmwareTanzuManageV1alpha1EksclusterNodepoolFullName `json:"fullName,omitempty"`

	// Metadata for the Nodepool object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Nodepool.
	Spec *VmwareTanzuManageV1alpha1EksclusterNodepoolSpec `json:"spec,omitempty"`

	// Status of the Nodepool.
	Status *VmwareTanzuManageV1alpha1EksclusterNodepoolStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
