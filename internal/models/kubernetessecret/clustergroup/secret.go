/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupsecret

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret Represents Tanzu Secret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.secret.Secret
type VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret struct {

	// Full name for the Secret.
	FullName *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretFullName `json:"fullName,omitempty"`

	// Metadata for the Secret  object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Secret.
	Spec *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec `json:"spec,omitempty"`

	// Status for the Secret.
	Status *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
