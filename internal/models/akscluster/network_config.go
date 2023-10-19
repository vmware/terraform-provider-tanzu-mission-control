/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package models

import (
	"github.com/go-openapi/swag"
)

const (
	// VmwareTanzuManageV1alpha1AksClusterNetworkPluginKubenet captures value "kubenet".
	VmwareTanzuManageV1alpha1AksClusterNetworkPluginKubenet = "kubenet"

	// VmwareTanzuManageV1alpha1AksClusterNetworkPluginAzure captures value "azure".
	VmwareTanzuManageV1alpha1AksClusterNetworkPluginAzure = "azure"

	// VmwareTanzuManageV1alpha1AksClusterNetworkPluginModeOverlay captures value "overlay".
	VmwareTanzuManageV1alpha1AksClusterNetworkPluginModeOverlay = "overlay"
)

// VmwareTanzuManageV1alpha1AksclusterNetworkConfig The network configuration for the AKS cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.NetworkConfig
type VmwareTanzuManageV1alpha1AksclusterNetworkConfig struct {

	// DNS prefix of the cluster.
	DNSPrefix string `json:"dnsPrefix,omitempty"`

	// An IP address assigned to the Kubernetes DNS service.
	// It must be within the Kubernetes service address range specified in serviceCidr.
	DNSServiceIP string `json:"dnsServiceIp,omitempty"`

	// A CIDR notation IP range assigned to the Docker bridge network.
	// It must not overlap with any Subnet IP ranges or the Kubernetes service address range.
	DockerBridgeCidr string `json:"dockerBridgeCidr,omitempty"`

	// The load balancer SKU for the cluster. The valid value is basic and standard.
	LoadBalancerSku string `json:"loadBalancerSku,omitempty"`

	// Network plugin of the cluster. The valid value is azure, kubenet and none.
	NetworkPlugin string `json:"networkPlugin,omitempty"`

	// Network plugin mode of the cluster. The valid values are overlay and ''.
	NetworkPluginMode string `json:"networkPluginMode,omitempty"`

	// Network policy of the cluster. The valid value is azure and calico.
	NetworkPolicy string `json:"networkPolicy,omitempty"`

	// The CIDR notation IP ranges from which to assign pod IPs.
	PodCidrs []string `json:"podCidrs"`

	// The CIDR notation IP ranges from which to assign service cluster IPs.
	ServiceCidrs []string `json:"serviceCidrs"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterNetworkConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterNetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
