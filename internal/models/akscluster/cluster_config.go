/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterClusterConfig The config of an AKS cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.ClusterConfig
type VmwareTanzuManageV1alpha1AksclusterClusterConfig struct {

	// The access config.
	AccessConfig *VmwareTanzuManageV1alpha1AksclusterAccessConfig `json:"accessConfig,omitempty"`

	// The Addons config.
	AddonsConfig *VmwareTanzuManageV1alpha1AksclusterAddonsConfig `json:"addonsConfig,omitempty"`

	// The access config for the cluster API server.
	APIServerAccessConfig *VmwareTanzuManageV1alpha1AksclusterAPIServerAccessConfig `json:"apiServerAccessConfig,omitempty"`

	// The auto upgrade config.
	AutoUpgradeConfig *VmwareTanzuManageV1alpha1AksclusterAutoUpgradeConfig `json:"autoUpgradeConfig,omitempty"`

	// The resource ID of the disk encryption set to use for enabling
	DiskEncryptionSetID string `json:"diskEncryptionSetId,omitempty"`

	// The linux VMs config.
	LinuxConfig *VmwareTanzuManageV1alpha1AksclusterLinuxConfig `json:"linuxConfig,omitempty"`

	// The geo-location where the resource lives for the cluster.
	Location string `json:"location,omitempty"`

	// Kubernetes network config.
	NetworkConfig *VmwareTanzuManageV1alpha1AksclusterNetworkConfig `json:"networkConfig,omitempty"`

	// The name of the resource group containing node pool nodes.
	NodeResourceGroupName string `json:"nodeResourceGroupName,omitempty"`

	// SKU of the cluster.
	Sku *VmwareTanzuManageV1alpha1AksclusterClusterSKU `json:"sku,omitempty"`

	// The storage config.
	StorageConfig *VmwareTanzuManageV1alpha1AksclusterStorageConfig `json:"storageConfig,omitempty"`

	// The metadata to apply to the cluster to assist with categorization and organization.
	Tags map[string]string `json:"tags,omitempty"`

	// The managed identity to apply to the cluster.
	IdentityConfig *VmwareTanzuManageV1alpha1AksclusterManagedIdentityConfig `json:"identityConfig,omitempty"`

	// Kubernetes version of the cluster.
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterClusterConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
