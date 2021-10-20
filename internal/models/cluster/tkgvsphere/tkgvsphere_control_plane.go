/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgvspheremodel

import (
	"github.com/go-openapi/swag"

	nodepoolmodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/cluster/nodepool"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereControlPlane VSphere specific control plane configuration for workload cluster object.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgvsphere.ControlPlane
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereControlPlane struct {

	// High Availability or Non High Availability Cluster. HA cluster
	// creates three controlplane machines, and non HA creates just one.
	HighAvailability bool `json:"highAvailability,omitempty"`

	// VM specific configuration.
	VMConfig *nodepoolmodel.VmwareTanzuManageV1alpha1CommonClusterTKGVsphereVMConfig `json:"vmConfig,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereControlPlane) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereControlPlane) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereControlPlane
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
