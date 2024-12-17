// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkcmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkSettings Network related settings for the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.NetworkSettings
type VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkSettings struct {

	// Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16.
	Pods *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkRanges `json:"pods,omitempty"`

	// Domain name for services.
	ServiceDomain string `json:"serviceDomain,omitempty"`

	// Service CIDR for kubernetes services defaults to 10.96.0.0/12.
	Services *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkRanges `json:"services,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkSettings) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNetworkSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
