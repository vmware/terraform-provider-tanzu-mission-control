/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package nodepool

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterNodepoolDefinition Definition is the definition of nodepool for cluster
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.Definition
type VmwareTanzuManageV1alpha1ClusterNodepoolDefinition struct {

	// Info for the nodepool.
	Info *VmwareTanzuManageV1alpha1ClusterNodepoolInfo `json:"info,omitempty"`

	// Spec for the nodepool.
	Spec *VmwareTanzuManageV1alpha1ClusterNodepoolSpec `json:"spec,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolDefinition) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNodepoolNodepool A group of Kubernetes clusters.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.Nodepool
type VmwareTanzuManageV1alpha1ClusterNodepoolNodepool struct {

	// Full name for the Nodepool.
	FullName *VmwareTanzuManageV1alpha1ClusterNodepoolFullName `json:"fullName,omitempty"`

	// Metadata for the Nodepool object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Nodepool.
	Spec *VmwareTanzuManageV1alpha1ClusterNodepoolSpec `json:"spec,omitempty"`

	// Status of the Nodepool.
	Status *VmwareTanzuManageV1alpha1ClusterNodepoolStatus `json:"status,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolNodepool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolNodepool) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolNodepool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNodepoolFullName Full name of the nodepool. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.FullName
type VmwareTanzuManageV1alpha1ClusterNodepoolFullName struct {

	// Name of the cluster.
	ClusterName string `json:"clusterName,omitempty"`

	// Name of the management cluster.
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Name of this nodepool.
	Name string `json:"name,omitempty"`

	// Provisioner of the cluster.
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
