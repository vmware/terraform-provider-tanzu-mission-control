/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterAPIServerAccessConfig The access config for the cluster API server.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.ApiServerAccessConfig
type VmwareTanzuManageV1alpha1AksclusterAPIServerAccessConfig struct {

	// The IP ranges authorized to access the Kubernetes API server.
	AuthorizedIPRanges []string `json:"authorizedIpRanges"`

	// Whether to create the cluster as a private cluster or not.
	EnablePrivateCluster bool `json:"enablePrivateCluster,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAPIServerAccessConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAPIServerAccessConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterAPIServerAccessConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
