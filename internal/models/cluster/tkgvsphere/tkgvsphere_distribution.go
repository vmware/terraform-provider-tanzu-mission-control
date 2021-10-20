/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgvspheremodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereDistribution VSphere specific distribution.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgvsphere.Distribution
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereDistribution struct {

	// Version specifies the version of the Kubernetes cluster.
	Version string `json:"version,omitempty"`

	// Workspace defines a workspace configuration for the vSphere cloud provider.
	Workspace *VmwareTanzuManageV1alpha1CommonClusterTKGVsphereWorkspace `json:"workspace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereDistribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereDistribution) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereDistribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
