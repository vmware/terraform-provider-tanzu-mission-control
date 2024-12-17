// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1EksclusterNodepoolSpec Spec for the cluster nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.ekscluster.nodepool.Spec
type VmwareTanzuManageV1alpha1EksclusterNodepoolSpec struct {

	// AMI info for the nodepool.
	AmiInfo *VmwareTanzuManageV1alpha1EksclusterNodepoolAmiInfo `json:"amiInfo,omitempty"`

	// AMI type.
	AmiType string `json:"amiType,omitempty"`

	// Capacity type.
	CapacityType string `json:"capacityType,omitempty"`

	// Nodepool instance types.
	// The potential values could be found using cluster:options api.
	InstanceTypes []string `json:"instanceTypes"`

	// Launch template for the nodepool.
	LaunchTemplate *VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate `json:"launchTemplate,omitempty"`

	// Kubernetes node labels.
	NodeLabels map[string]string `json:"nodeLabels,omitempty"`

	// Remote access to worker nodes.
	RemoteAccess *VmwareTanzuManageV1alpha1EksclusterNodepoolRemoteAccess `json:"remoteAccess,omitempty"`

	// ARN of the IAM role that provides permissions for the Kubernetes nodepool to make calls to AWS API operations.
	RoleArn string `json:"roleArn,omitempty"`

	// Root disk size in GiB. Defaults to 20 GiB.
	RootDiskSize int32 `json:"rootDiskSize,omitempty"`

	// Nodepool scaling config.
	ScalingConfig *VmwareTanzuManageV1alpha1EksclusterNodepoolScalingConfig `json:"scalingConfig,omitempty"`

	// Subnets required for the nodepool.
	SubnetIds []string `json:"subnetIds"`

	// EKS specific tags.
	Tags map[string]string `json:"tags,omitempty"`

	// If specified, the node's taints.
	Taints []*VmwareTanzuManageV1alpha1EksclusterNodepoolTaint `json:"taints"`

	// Update config for the nodepool.
	UpdateConfig *VmwareTanzuManageV1alpha1EksclusterNodepoolUpdateConfig `json:"updateConfig,omitempty"`

	// AMI release version. This field is used to update the AMI release version for that k8s version.
	// This field should not be specfied for cluster CREATE as TMC uses the latest ami release version for it.
	ReleaseVersion string `json:"releaseVersion,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1EksclusterNodepoolSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
