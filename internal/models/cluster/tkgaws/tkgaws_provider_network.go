/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgawsmodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork Provider related network settings for the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgaws.ProviderNetwork
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork struct {

	// Optional list of subnets used to place the nodes in the cluster.
	Subnets []*VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSubnet `json:"subnets"`

	// AWS VPC configuration for the cluster.
	Vpc *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsVPC `json:"vpc,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsProviderNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
