package models

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AksclusterChannel Channel options of auto upgrade config.
//
//   - CHANNEL_UNSPECIFIED: Unspecified channel.
//   - NONE: Disables auto-upgrades and keeps the cluster at its current version of Kubernetes.
//   - PATCH: Automatically upgrades the cluster to the latest supported patch version when it becomes available while keeping the minor version the same.
//   - STABLE: Automatically upgrades the cluster to the latest supported patch release on minor version N-1, where N is the latest supported minor version.
//   - RAPID: Automatically upgrades the cluster to the latest supported patch release on the latest supported minor version.
//   - NODE_IMAGE: Automatically upgrades the node image to the latest version available.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.Channel
type VmwareTanzuManageV1alpha1AksclusterChannel string

func NewVmwareTanzuManageV1alpha1AksclusterChannel(value VmwareTanzuManageV1alpha1AksclusterChannel) *VmwareTanzuManageV1alpha1AksclusterChannel {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AksclusterChannel.
func (m VmwareTanzuManageV1alpha1AksclusterChannel) Pointer() *VmwareTanzuManageV1alpha1AksclusterChannel {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AksclusterChannelCHANNELUNSPECIFIED captures enum value "CHANNEL_UNSPECIFIED".
	VmwareTanzuManageV1alpha1AksclusterChannelCHANNELUNSPECIFIED VmwareTanzuManageV1alpha1AksclusterChannel = "CHANNEL_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AksclusterChannelNONE captures enum value "NONE".
	VmwareTanzuManageV1alpha1AksclusterChannelNONE VmwareTanzuManageV1alpha1AksclusterChannel = "NONE"

	// VmwareTanzuManageV1alpha1AksclusterChannelPATCH captures enum value "PATCH".
	VmwareTanzuManageV1alpha1AksclusterChannelPATCH VmwareTanzuManageV1alpha1AksclusterChannel = "PATCH"

	// VmwareTanzuManageV1alpha1AksclusterChannelSTABLE captures enum value "STABLE".
	VmwareTanzuManageV1alpha1AksclusterChannelSTABLE VmwareTanzuManageV1alpha1AksclusterChannel = "STABLE"

	// VmwareTanzuManageV1alpha1AksclusterChannelRAPID captures enum value "RAPID".
	VmwareTanzuManageV1alpha1AksclusterChannelRAPID VmwareTanzuManageV1alpha1AksclusterChannel = "RAPID"

	// VmwareTanzuManageV1alpha1AksclusterChannelNODEIMAGE captures enum value "NODE_IMAGE".
	VmwareTanzuManageV1alpha1AksclusterChannelNODEIMAGE VmwareTanzuManageV1alpha1AksclusterChannel = "NODE_IMAGE"
)

// for schema.
var vmwareTanzuManageV1alpha1AksclusterChannelEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AksclusterChannel
	if err := json.Unmarshal([]byte(`["CHANNEL_UNSPECIFIED","NONE","PATCH","STABLE","RAPID","NODE_IMAGE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1AksclusterChannelEnum = append(vmwareTanzuManageV1alpha1AksclusterChannelEnum, v)
	}
}
