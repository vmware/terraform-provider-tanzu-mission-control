/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyclustermodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	policymodel "github.com/vmware-tanzu/terraform-provider-tanzu-mission-control/internal/models/policy"
)

// VmwareTanzuManageV1alpha1ClusterPolicyPolicy A Policy to apply on a Kubernetes cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.policy.Policy
type VmwareTanzuManageV1alpha1ClusterPolicyPolicy struct {

	// Full name for the Cluster policy.
	FullName *VmwareTanzuManageV1alpha1ClusterPolicyFullName `json:"fullName,omitempty"`

	// Metadata for the Cluster policy.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Cluster policy.
	Spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec `json:"spec,omitempty"`

	// Metadata describing the type of the resource.
	Type *policymodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPolicyPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterPolicyPolicy) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterPolicyPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
