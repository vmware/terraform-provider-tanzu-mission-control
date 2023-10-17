/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponse Response with cluster kubeconfig.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.kubeconfig.GetKubeconfigResponse
type VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponse struct {

	// Provides the server endpoint used in the kubeconfig.
	Endpoint string `json:"endpoint,omitempty"`

	// Cluster Kubeconfig.
	Kubeconfig string `json:"kubeconfig,omitempty"`

	// Provides the detail message for PENDING/UNAVAILABLE status.
	Msg string `json:"msg,omitempty"`

	// Kubeconfig status.
	Status *VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponseStatus `json:"status,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterKubeconfigGetKubeconfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
