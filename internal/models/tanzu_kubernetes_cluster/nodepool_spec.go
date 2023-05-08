/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzukubernetescluster

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec Spec for the cluster nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.nodepool.Spec
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec struct {

	// The name of the machine deployment class used to create the nodepool.
	Class string `json:"class,omitempty"`

	// The failure domain the machines will be created in.
	FailureDomain string `json:"failureDomain,omitempty"`

	// The metadata of the nodepool.
	Metadata *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterMetadata `json:"metadata,omitempty"`

	// The OS image of the nodepool.
	OsImage *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterOSImage `json:"osImage,omitempty"`

	// Overrides can be used to override cluster level variables.
	Overrides []*VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterCommonClusterClusterVariable `json:"overrides"`

	// The replicas of the nodepool.
	Replicas int32 `json:"replicas,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerTanzukubernetesclusterNodepoolSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
