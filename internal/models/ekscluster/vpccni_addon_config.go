// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterVpcCniAddonConfig EKS vpc-cni addon configuration
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.VpcCniAddonConfig
type VmwareTanzuManageV1alpha1EksclusterVpcCniAddonConfig struct {

	// Additional configuration settings for vpc-cni addon. Eniconfig are repeateable and each eniConfig should have a mandatory subnet id and optional security group id(s).
	// Subnets need not be in the same VPC as the cluster.
	// The subnets provided across eniConfigs should be in different availability zones
	// If security group id(s) are not provided, the cluster security group will be used
	EniConfigs []*VmwareTanzuManageV1alpha1EksclusterEniConfig `json:"eniConfigs"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterVpcCniAddonConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterVpcCniAddonConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterVpcCniAddonConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
