// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package namespacemodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceNamespace A managed Kubernetes namespace.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.Namespace
type VmwareTanzuManageV1alpha1ClusterNamespaceNamespace struct {

	// Full name for the Namespace.
	FullName *VmwareTanzuManageV1alpha1ClusterNamespaceFullName `json:"fullName,omitempty"`

	// Metadata for the Namespace object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Namespace.
	Spec *VmwareTanzuManageV1alpha1ClusterNamespaceSpec `json:"spec,omitempty"`

	// Status for the Namespace.
	Status *VmwareTanzuManageV1alpha1ClusterNamespaceStatus `json:"status,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceNamespace) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
