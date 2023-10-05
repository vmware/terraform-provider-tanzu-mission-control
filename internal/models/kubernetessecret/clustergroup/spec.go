/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupsecret

import (
	"github.com/go-openapi/swag"

	secretclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/kubernetessecret/cluster"
)

// VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec Spec of the Secret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.secret.Spec
type VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec struct {

	// Spec of secret as defined at atomic level.
	AtomicSpec *secretclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec `json:"atomicSpec,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
