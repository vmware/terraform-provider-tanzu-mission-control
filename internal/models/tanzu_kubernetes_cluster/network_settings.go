/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetescluster

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings Network related settings for the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.NetworkSettings
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings struct {

	// Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16.
	Pods *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkRanges `json:"pods,omitempty"`

	// Domain name for services.
	ServiceDomain string `json:"serviceDomain,omitempty"`

	// Service CIDR for kubernetes services defaults to 10.96.0.0/12.
	Services *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkRanges `json:"services,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNetworkSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
