// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmfeatureclustergroupmodel

import (
	"github.com/go-openapi/swag"

	optionsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/options"
)

// VmwareTanzuManageV1alpha1ClustergroupHelmListHelmRequestParameters Request parameters to list Helm.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.helm.ListHelmsRequestParameters
type VmwareTanzuManageV1alpha1ClustergroupHelmListHelmRequestParameters struct {

	// Scope to search by, any fields left empty will be considered all (*).
	SearchScope *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmSearchScope `json:"search_scope,omitempty"`

	// Sort Order.
	SortBy string `json:"sort_by,omitempty"`

	// TQL query string.
	Query string `json:"query,omitempty"`

	// Pagination.
	Pagination *optionsmodel.VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions `json:"pagination,omitempty"`

	// Include total count.
	IncludeTotalCount bool `json:"include_total_count,omitempty"`
}

// VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.helm.SearchScope
type VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmSearchScope struct {

	// Scope search to the specified cluster_group_name; supports globbing; default (*).
	ClusterGroupName string `json:"clusterGroupName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmSearchScope) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmListHelmsResponse Response from listing Helms.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.helm.ListHelmsResponse
type VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmListHelmsResponse struct {

	// List of helms.
	Helms []*VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmHelm `json:"helms"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmListHelmsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmListHelmsResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdHelmListHelmsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
