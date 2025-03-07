// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tkcmodels

import (
	"github.com/go-openapi/swag"

	tkccommon "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/tanzukubernetescluster/common"
)

// VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterControlPlane The cluster specific control plane configuration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.ControlPlane
type VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterControlPlane struct {

	// The metadata of the control plane.
	Metadata *tkccommon.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterMetadata `json:"metadata,omitempty"`

	// The OS image of the control plane.
	OsImage *tkccommon.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterOSImage `json:"osImage,omitempty"`

	// Overrides can be used to override cluster level variables.
	Overrides []*tkccommon.VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterCommonClusterClusterVariable `json:"overrides"`

	// The replicas of the control plane.
	Replicas int32 `json:"replicas,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterControlPlane) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterControlPlane) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterControlPlane
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
