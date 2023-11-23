/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package clustergroupintegrationmodels

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
)

// VmwareTanzuManageV1alpha1ClusterGroupIntegrationIntegration An integration configuration for a cluster group.
// This integration configuration will be applied to child clusters of the cluster group.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.integration.Integration
type VmwareTanzuManageV1alpha1ClusterGroupIntegrationIntegration struct {

	// Full name for the Cluster Group integration.
	FullName *VmwareTanzuManageV1alpha1ClusterGroupIntegrationFullName `json:"fullName,omitempty"`

	// Metadata for the Cluster Group integration.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Cluster Group integration configuration.
	Spec *VmwareTanzuManageV1alpha1ClusterGroupIntegrationSpec `json:"spec,omitempty"`

	// Status for the Integration.
	Status *VmwareTanzuManageV1alpha1ClusterGroupIntegrationStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *policymodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupIntegrationIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterGroupIntegrationIntegration) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterGroupIntegrationIntegration

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
