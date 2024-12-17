// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterSpec Spec of the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.Spec
type VmwareTanzuManageV1alpha1EksclusterSpec struct {

	// Name of the cluster group to which this cluster belongs.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// EKS config for the cluster control plane.
	Config *VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig `json:"config,omitempty"`

	// Optional proxy name is the name of the Proxy Config
	// to be used for the cluster.
	ProxyName string `json:"proxyName,omitempty"`

	// Agent name of the cluster.
	AgentName string `json:"agentName,omitempty"`

	// Arn of the cluster.
	Arn string `json:"arn,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1EksclusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1EksclusterSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
