/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterEniConfig EKS vpc-cni addon ENI configuration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.EniConfig
type VmwareTanzuManageV1alpha1EksclusterEniConfig struct {

	// Subnet Id  for the pods.
	SubnetID string `json:"subnetId,omitempty"`

	// Security group for the pods.
	SecurityGroupIds []string `json:"securityGroupIds"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterEniConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterEniConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterEniConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
