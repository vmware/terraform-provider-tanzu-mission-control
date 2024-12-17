// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgawsmodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet Subnet configuration for the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgaws.Subnet
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet struct {

	// AWS availability zone e.g. us-west-2a
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// CIDR for AWS subnet.
	// This CIDR block must be in the range of AWS VPC CIDR block.
	CidrBlock string `json:"cidrBlock,omitempty"`

	// AWS subnet ID. The rest of the fields are ignored if this field is specified.
	ID string `json:"id,omitempty"`

	// Public subnet or private subnet.
	IsPublic bool `json:"isPublic,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
