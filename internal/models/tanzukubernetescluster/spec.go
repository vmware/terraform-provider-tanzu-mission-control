/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcmodels

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterSpec Spec of the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.Spec
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterSpec struct {

	// Name of the cluster group to which this cluster belongs.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Name of the image registry configuration to use.
	ImageRegistry string `json:"imageRegistry,omitempty"`

	// Name of the proxy configuration to use.
	ProxyName string `json:"proxyName,omitempty"`

	// TMC-managed flag indicates if the cluster is managed by tmc.
	TmcManaged bool `json:"tmcManaged"`

	// The cluster topology.
	Topology *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterTopology `json:"topology,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
