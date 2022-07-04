/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package namespaceiammodel

import (
	"github.com/go-openapi/swag"

	iammodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/iam_policy"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceUpdateNamespaceIAMPolicyResponse UpdateNamespaceIAMPolicy response message.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.UpdateNamespaceIAMPolicyResponse
type VmwareTanzuManageV1alpha1ClusterNamespaceUpdateNamespaceIAMPolicyResponse struct {

	// Namespace policy set.
	Policy *iammodel.VmwareTanzuCoreV1alpha1PolicyIAMPolicy `json:"policy,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceUpdateNamespaceIAMPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceUpdateNamespaceIAMPolicyResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceUpdateNamespaceIAMPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
