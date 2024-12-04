// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmfeatureclustermodel

import (
	"github.com/go-openapi/swag"

	optionsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/options"
)

// VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmRequestParameters Request parameters to list Helm.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.helm.ListHelmRequestParameters
type VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmRequestParameters struct {

	// Scope to search by, any fields left empty will be considered all (*).
	SearchScope *VmwareTanzuManageV1alpha1ClusterFluxcdHelmSearchScope `json:"search_scope,omitempty"`

	// Sort Order.
	SortBy string `json:"sort_by,omitempty"`

	// TQL query string.
	Query string `json:"query,omitempty"`

	// Pagination.
	Pagination *optionsmodel.VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions `json:"pagination,omitempty"`

	// Include total count.
	IncludeTotalCount bool `json:"include_total_count,omitempty"`
}

// VmwareTanzuManageV1alpha1ClusterFluxcdHelmSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.helm.SearchScope
type VmwareTanzuManageV1alpha1ClusterFluxcdHelmSearchScope struct {

	// Scope search to the specified cluster_name; supports globbing; default (*).
	ClusterName string `json:"clusterName,omitempty"`

	// Scope search to the specified management_cluster_name; supports globbing; default (*).
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Scope search to the specified provisioner_name; supports globbing; default (*).
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmSearchScope) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdHelmSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse Response from listing Helms.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.helm.ListHelmsResponse
type VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse struct {

	// List of helms.
	Helms []*VmwareTanzuManageV1alpha1ClusterFluxcdHelmHelm `json:"helms"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdHelmListHelmsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
