/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policyclustergroupmodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
)

// VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy A Policy to apply on a group of Kubernetes clusters.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.policy.Policy
type VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy struct {

	// Full name for the ClusterGroup policy.
	FullName *VmwareTanzuManageV1alpha1ClustergroupPolicyFullName `json:"fullName,omitempty"`

	// Metadata for the ClusterGroup policy.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the ClusterGroup policy.
	Spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec `json:"spec,omitempty"`

	// Metadata describing the type of the resource.
	Type *policymodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupPolicyPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
