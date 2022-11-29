/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig The EKS cluster config.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.ControlPlaneConfig
type VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig struct {

	// Kubernetes Network Config.
	KubernetesNetworkConfig *VmwareTanzuManageV1alpha1EksclusterKubernetesNetworkConfig `json:"kubernetesNetworkConfig,omitempty"`

	// EKS logging configuration.
	Logging *VmwareTanzuManageV1alpha1EksclusterLogging `json:"logging,omitempty"`

	// ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations.
	RoleArn string `json:"roleArn,omitempty"`

	// The metadata to apply to the cluster to assist with categorization and organization.
	Tags map[string]string `json:"tags,omitempty"`

	// Kubernetes version of the cluster.
	Version string `json:"version,omitempty"`

	// VPC config.
	Vpc *VmwareTanzuManageV1alpha1EksclusterVPCConfig `json:"vpc,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
