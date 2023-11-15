/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkcnodepool

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool Nodepool associated with a tanzu Kubernetes cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.tanzukubernetescluster.nodepool.Nodepool
type VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool struct {
	// Full name of the nodepool. This includes the object name along with any parents or further identifiers.
	FullName *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolFullName `json:"fullName,omitempty"`

	// Holds general shared object metadatas.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the cluster nodepool.
	Spec *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolSpec `json:"spec,omitempty"`

	// Status of node pool resource.
	Status *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepoolStatus `json:"status,omitempty"`

	// Holds general type metadatas.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementClusterProvisionerTanzukubernetesClusterNodepool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
