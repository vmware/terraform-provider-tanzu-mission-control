// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterAddonsConfig EKS addons configuration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.AddonsConfig
type VmwareTanzuManageV1alpha1EksclusterAddonsConfig struct {

	// Enable the Kubernetes vpc-cni addon.
	VpcCniAddonConfig *VmwareTanzuManageV1alpha1EksclusterVpcCniAddonConfig `json:"vpcCniAddonConfig,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterAddonsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterAddonsConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterAddonsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
