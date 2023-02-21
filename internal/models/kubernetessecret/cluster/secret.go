/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package secret

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceSecret Represents Tanzu Secret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.secret.Secret
type VmwareTanzuManageV1alpha1ClusterNamespaceSecret struct {

	// Full name for the Secret.
	FullName *VmwareTanzuManageV1alpha1ClusterNamespaceSecretFullName `json:"fullName,omitempty"`

	// Metadata for the Secret  object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Secret.
	Spec *VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec `json:"spec,omitempty"`

	// Status for the Secret.
	Status *statusmodel.VmwareTanzuManageV1alpha1ClusterNamespaceStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecret) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
