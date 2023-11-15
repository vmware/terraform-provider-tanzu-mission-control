/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcnodepool

import (
	"github.com/go-openapi/swag"

	tkccommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/common"
)

// VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolSpec Spec for the cluster nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.nodepool.Spec
type VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolSpec struct {

	// The name of the machine deployment class used to create the nodepool.
	Class string `json:"class,omitempty"`

	// The failure domain the machines will be created in.
	FailureDomain string `json:"failureDomain,omitempty"`

	// The metadata of the nodepool.
	Metadata *tkccommon.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterMetadata `json:"metadata,omitempty"`

	// The OS image of the nodepool.
	OsImage *tkccommon.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterOSImage `json:"osImage,omitempty"`

	// Overrides can be used to override cluster level variables.
	Overrides []*tkccommon.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterClusterVariable `json:"overrides"`

	// The replicas of the nodepool.
	Replicas int32 `json:"replicas,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
