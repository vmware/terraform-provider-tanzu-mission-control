/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterNodepoolSpec Spec for the cluster nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.nodepool.Spec
type VmwareTanzuManageV1alpha1AksclusterNodepoolSpec struct {

	// Auto scaling config.
	AutoScaling *VmwareTanzuManageV1alpha1AksclusterNodepoolAutoScalingConfig `json:"autoScaling,omitempty"`

	// The list of Availability zones to use for nodepool.
	// This can only be specified if the type of the nodepool is AvailabilitySet.
	AvailabilityZones []string `json:"availabilityZones"`

	// Count is the number of nodes.
	Count int32 `json:"count,omitempty"`

	// Whether each node is allocated its own public IP.
	EnableNodePublicIP bool `json:"enableNodePublicIp,omitempty"`

	// The maximum number of pods that can run on a node.
	MaxPods int32 `json:"maxPods,omitempty"`

	// The mode of the nodepool
	// A cluster must have at least one 'System' nodepool at all times.
	Mode *VmwareTanzuManageV1alpha1AksclusterNodepoolMode `json:"mode,omitempty"`

	// The node image version of the nodepool.
	NodeImageVersion string `json:"nodeImageVersion,omitempty"`

	// The node labels to be persisted across all nodes in nodepool.
	NodeLabels map[string]string `json:"nodeLabels,omitempty"`

	// The taints added to new nodes during nodepool create and scale.
	NodeTaints []*VmwareTanzuManageV1alpha1AksclusterNodepoolTaint `json:"nodeTaints"`

	// OS Disk Size in GB to be used to specify the disk size for every machine in the nodepool.
	// If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGb int32 `json:"osDiskSizeGb,omitempty"`

	// The OS disk type of the nodepool.
	OsDiskType *VmwareTanzuManageV1alpha1AksclusterNodepoolOsDiskType `json:"osDiskType,omitempty"`

	// The operation system type of the nodepool.
	OsType *VmwareTanzuManageV1alpha1AksclusterNodepoolOsType `json:"osType,omitempty"`

	// The Virtual Machine Scale Set eviction policy to use.
	// This cannot be specified unless the scaleSetPriority is 'Spot'.
	ScaleSetEvictionPolicy *VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetEvictionPolicy `json:"scaleSetEvictionPolicy,omitempty"`

	// The Virtual Machine Scale Set priority.
	ScaleSetPriority *VmwareTanzuManageV1alpha1AksclusterNodepoolScaleSetPriority `json:"scaleSetPriority,omitempty"`

	// The max price (in US Dollars) you are willing to pay for spot instances.
	// Possible values are any decimal value greater than zero or -1 which indicates default price to be up-to on-demand.
	SpotMaxPrice float32 `json:"spotMaxPrice,omitempty"`

	// The metadata to apply to the nodepool.
	Tags map[string]string `json:"tags,omitempty"`

	// The type of the nodepool.
	Type *VmwareTanzuManageV1alpha1AksclusterNodepoolType `json:"type,omitempty"`

	// The upgrade config.
	UpgradeConfig *VmwareTanzuManageV1alpha1AksclusterNodepoolUpgradeConfig `json:"upgradeConfig,omitempty"`

	// The size of the nodepool VMs.
	VMSize string `json:"vmSize,omitempty"`

	// If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to
	// nodes and pods, otherwise it applies to just nodes. This is of the form:
	VnetSubnetID string `json:"vnetSubnetId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNodepoolSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNodepoolSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
