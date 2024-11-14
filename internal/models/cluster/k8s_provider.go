// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustermodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1CommonClusterKubernetesProvider KubernetesProvider definition - indicates the k8s provider type and version.
//
// swagger:model vmware.tanzu.manage.v1alpha1.common.cluster.KubernetesProvider
type VmwareTanzuManageV1alpha1CommonClusterKubernetesProvider struct {

	// Indicates the k8s provider type.
	Type *VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType `json:"type,omitempty"`

	// Indicates the k8s provider version.
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterKubernetesProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonClusterKubernetesProvider) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonClusterKubernetesProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
