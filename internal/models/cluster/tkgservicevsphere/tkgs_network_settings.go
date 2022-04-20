/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgservicevspheremodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkSettings Kubernetes specific network information for workload cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgservicevsphere.NetworkSettings
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkSettings struct {

	// Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16.
	Pods *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkRanges `json:"pods,omitempty"`

	// Service CIDR for kubernetes services defaults to 10.96.0.0/12.
	Services *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkRanges `json:"services,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkSettings) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereNetworkSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
