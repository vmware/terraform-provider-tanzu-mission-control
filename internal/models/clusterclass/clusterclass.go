/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterclass

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassClusterClass A Kubernetes Cluster Class.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.provisioner.clusterclass.ClusterClass
type VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassClusterClass struct {

	// Full name for the cluster class.
	FullName *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassFullName `json:"fullName,omitempty"`

	// Metadata for the cluster class object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the cluster class.
	Spec *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassSpec `json:"spec,omitempty"`

	// Status of the cluster class.
	Status *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassClusterClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassClusterClass) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterProvisionerClusterclassClusterClass

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
