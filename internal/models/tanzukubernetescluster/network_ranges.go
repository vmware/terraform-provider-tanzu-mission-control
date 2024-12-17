// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkcmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkRanges NetworkRanges describes a collection of IP addresses as a list of ranges.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.NetworkRanges
type VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkRanges struct {

	// CIDRBlocks specifies one or more ranges of IP addresses.
	CidrBlocks []string `json:"cidrBlocks"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkRanges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkRanges) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkRanges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
