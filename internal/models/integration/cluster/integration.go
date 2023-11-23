/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clusterintegrationmodels

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
)

// VmwareTanzuManageV1alpha1ClusterIntegration An integration configuration for a cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.integration.Integration
type VmwareTanzuManageV1alpha1ClusterIntegration struct {
	// Full name for the Cluster integration.
	FullName *VmwareTanzuManageV1alpha1ClusterIntegrationFullName `json:"fullName,omitempty"`

	// Metadata for the Cluster policy.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Cluster integration configuration.
	Spec *VmwareTanzuManageV1alpha1ClusterIntegrationSpec `json:"spec,omitempty"`

	// Status for the Integration.
	Status *VmwareTanzuManageV1alpha1ClusterIntegrationStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *policymodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterIntegration) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterIntegration

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
