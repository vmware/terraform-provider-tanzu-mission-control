/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgawsmodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsDistribution Distribution of the AWS cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.infrastructure.tkgaws.Distribution
type VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsDistribution struct {

	// Arch of the OS used for the cluster.
	OsArch string `json:"osArch,omitempty"`

	// Name of the OS used for the cluster.
	OsName string `json:"osName,omitempty"`

	// Version of the OS used for the cluster.
	OsVersion string `json:"osVersion,omitempty"`

	// Name of the account (provisioner credential) in which to create the cluster.
	ProvisionerCredentialName string `json:"provisionerCredentialName,omitempty"`

	// Region of the cluster.
	Region string `json:"region,omitempty"`

	// Version of the cluster.
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsDistribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsDistribution) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsDistribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
