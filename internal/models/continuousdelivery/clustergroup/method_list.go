/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package continuousdeliveryclustergroupmodel

import (
	"github.com/go-openapi/swag"

	optionsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/options"
)

// VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesRequestParameters Request parameters to list ContinuousDeliveries.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.continuousdelivery.ListContinuousDeliveriesRequestParameters
type VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesRequestParameters struct {

	// Scope to search by, any fields left empty will be considered all (*).
	SearchScope *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliverySearchScope `json:"search_scope,omitempty"`

	// Sort Order.
	SortBy string `json:"sort_by,omitempty"`

	// TQL query string.
	Query string `json:"query,omitempty"`

	// Pagination.
	Pagination *optionsmodel.VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions `json:"pagination,omitempty"`

	// Include total count.
	IncludeTotalCount bool `json:"include_total_count,omitempty"`
}

// VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliverySearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.continuousdelivery.SearchScope
type VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliverySearchScope struct {

	// Scope search to the specified cluster_group_name; supports globbing; default (*).
	ClusterGroupName string `json:"clusterGroupName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliverySearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliverySearchScope) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliverySearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse Response from listing ContinuousDeliveries.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.continuousdelivery.ListContinuousDeliveriesResponse
type VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse struct {

	// List of continuousdeliveries.
	ContinuousDeliveries []*VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryContinuousDelivery `json:"continuousDeliveries"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdContinuousdeliveryListContinuousDeliveriesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
