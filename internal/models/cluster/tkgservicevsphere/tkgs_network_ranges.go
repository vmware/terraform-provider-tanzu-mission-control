// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkgservicevspheremodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkRanges Network ranges for the workload cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgservicevsphere.NetworkRanges
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkRanges struct {

	// CIDRBlocks specifies one or more ranges of IP addresses.
	CidrBlocks []string `json:"cidrBlocks"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkRanges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkRanges) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkRanges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
