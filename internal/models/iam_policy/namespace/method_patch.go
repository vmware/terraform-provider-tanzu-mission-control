/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespaceiammodel

import (
	"github.com/go-openapi/swag"

	iammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
	namespacemodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/namespace"
)

// VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyRequest PatchNamespaceIAMPolicy request message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.PatchNamespaceIAMPolicyRequest
type VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyRequest struct {

	// Binding delta to be applied.
	BindingDeltaList []*iammodel.VmwareTanzuCoreV1alpha1PolicyBindingDelta `json:"bindingDeltaList"`

	// Namespace full_name.
	FullName *namespacemodel.VmwareTanzuManageV1alpha1ClusterNamespaceFullName `json:"fullName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyResponse PatchNamespaceIAMPolicy response message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.PatchNamespaceIAMPolicyResponse
type VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyResponse struct {

	// New policy object.
	Policy *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespacePatchNamespaceIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
