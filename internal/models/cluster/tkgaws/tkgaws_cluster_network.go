/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgawsmodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork Network information for the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgaws.ClusterNetwork
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork struct {

	// APIServerPort specifies the port address for the cluster (optional).
	// The port value defaults to 6443.
	APIServerPort int32 `json:"apiServerPort,omitempty"`

	// Pod CIDR for Kubernetes pods. Defaults to 192.168.0.0/16.
	Pods []*VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange `json:"pods"`

	// Service CIDR for Kubernetes services. Defaults to 10.96.0.0/12.
	Services []*VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsNetworkRange `json:"services"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsClusterNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
