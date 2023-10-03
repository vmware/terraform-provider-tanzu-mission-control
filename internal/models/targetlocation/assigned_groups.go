/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationmodels

import (
	"github.com/go-openapi/swag"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	clustergroupmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/clustergroup"
)

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAssignedGroup Group of resources the backup location will be assigned to.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.AssignedGroup.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAssignedGroup struct {

	// Full name of a cluster.
	Cluster *clustermodel.VmwareTanzuManageV1alpha1ClusterFullName `json:"cluster,omitempty"`

	// Full name of a cluster group.
	Clustergroup *clustergroupmodel.VmwareTanzuManageV1alpha1ClustergroupFullName `json:"clustergroup,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAssignedGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAssignedGroup) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAssignedGroup

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
