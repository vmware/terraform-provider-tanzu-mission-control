/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzupackageinstall

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall Represents an instance of Package in the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.tanzupackage.install.Install.
type VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall struct {

	// Full name for the Package Install.
	FullName *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallFullName `json:"fullName,omitempty"`

	// Metadata for the Package Install object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Package Install.
	Spec *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallSpec `json:"spec,omitempty"`

	// Status for the Package Install.
	Status *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceTanzupackageInstallInstall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
