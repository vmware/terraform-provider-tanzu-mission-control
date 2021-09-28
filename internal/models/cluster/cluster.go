/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustermodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterCluster A Kubernetes Cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.Cluster
type VmwareTanzuManageV1alpha1ClusterCluster struct {

	// Full name for the cluster.
	FullName *VmwareTanzuManageV1alpha1ClusterFullName `json:"fullName,omitempty"`

	// Metadata for the cluster object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the cluster.
	Spec *VmwareTanzuManageV1alpha1ClusterSpec `json:"spec,omitempty"`

	// Status for the cluster.
	Status *VmwareTanzuManageV1alpha1ClusterStatus `json:"status,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterCluster) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
