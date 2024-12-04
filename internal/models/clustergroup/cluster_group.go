// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustergroupmodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClustergroupClusterGroup A group of Kubernetes clusters.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.ClusterGroup
type VmwareTanzuManageV1alpha1ClustergroupClusterGroup struct {

	// Full name for the ClusterGroup.
	FullName *VmwareTanzuManageV1alpha1ClustergroupFullName `json:"fullName,omitempty"`

	// Metadata for the ClusterGroup object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupClusterGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupClusterGroup) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupClusterGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
