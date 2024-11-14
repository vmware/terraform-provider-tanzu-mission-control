// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package nodepool

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement TKGAWSNodePlacement is the structure to indicate the AZ to place the nodes on.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.TKGAWSNodePlacement
type VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement struct {

	// The AZ where the AWS nodes are placed.
	AvailabilityZone string `json:"availabilityZone,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
