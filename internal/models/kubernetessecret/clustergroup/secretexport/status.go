// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustergroupsecretexport

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportStatus Status of Secret Export resource.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.namespace.secretexport.Status
type VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportStatus struct {

	// Details contains information about the Cluster Group secret export being applied on member Clusters.
	Details *statusmodel.VmwareTanzuManageV1alpha1CommonBatchDetails `json:"details,omitempty"`

	// Generation value at the time this status was updated.
	ObservedGeneration string `json:"observedGeneration,omitempty"`

	// Phase of the Cluster Group secret export application on member Clusters.
	Phase *statusmodel.VmwareTanzuManageV1alpha1CommonBatchPhase `json:"phase,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupNamespaceSecretexportStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
