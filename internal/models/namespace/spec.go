/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespacemodel

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceSpec The Namespace spec.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.Spec
type VmwareTanzuManageV1alpha1ClusterNamespaceSpec struct {

	// Attach specifies whether the namespace is being created or attached.
	Attach bool `json:"attach,omitempty"`

	// Name of Workspace which this Namespace belongs to.
	WorkspaceName string `json:"workspaceName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
