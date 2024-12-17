// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package nodepool

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool TKGAWSNodepool is the nodepool spec for TKG aws cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.TKGAWSNodepool
type VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool struct {

	// Availability zone for the nodepool. Should be one of the availability zones chosen for the cluster.
	// Use this field only if you are creating a nodepool for cluster in TMC hosted AWS solution. To create a nodepool for TKG
	// workload cluster please use TKGAWSNodePlacement
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// Nodepool instance type.
	// The potential values could be found using cluster:options api.
	InstanceType string `json:"instanceType,omitempty"`

	// List of AZs to place the AWS nodes on. Please use this field to provision a nodepool for workload cluster on an attached TKG AWS management cluster.
	// Please specify 1 AZ for a dev cluster and up to 3 AZs for production cluster.
	NodePlacement []*VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodePlacement `json:"nodePlacement"`

	// Subnet ID of the private subnet in which you want the nodes to be created in. If specified, availability zone is
	// ignored.
	SubnetID string `json:"subnetId,omitempty"`

	// Kubernetes version of the node pool.
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolTKGAWSNodepool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
