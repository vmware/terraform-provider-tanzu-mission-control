/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgvspheremodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkSettings Network related settings for VSphere cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgvsphere.NetworkSettings
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkSettings struct {

	// ControlPlaneEndpoint specifies the control plane virtual IP address.
	ControlPlaneEndpoint string `json:"controlPlaneEndpoint,omitempty"`

	// Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16.
	Pods *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkRanges `json:"pods,omitempty"`

	// Service CIDR for kubernetes services defaults to 10.96.0.0/12.
	Services *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkRanges `json:"services,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkSettings) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereNetworkSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
