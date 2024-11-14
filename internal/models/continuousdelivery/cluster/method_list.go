// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package continuousdeliveryclustermodel

import (
	"github.com/go-openapi/swag"

	optionsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/options"
)

// VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesRequestParameters Request parameters to list ContinuousDeliveries.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.continuousdelivery.ListContinuousDeliveriesRequestParameters
type VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesRequestParameters struct {

	// Scope to search by, any fields left empty will be considered all (*).
	SearchScope *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliverySearchScope `json:"search_scope,omitempty"`

	// Sort Order.
	SortBy string `json:"sort_by,omitempty"`

	// TQL query string.
	Query string `json:"query,omitempty"`

	// Pagination.
	Pagination *optionsmodel.VmwareTanzuCoreV1alpha1OptionsOffsetPaginationOptions `json:"pagination,omitempty"`

	// Include total count.
	IncludeTotalCount bool `json:"include_total_count,omitempty"`
}

// VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliverySearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.continuousdelivery.SearchScope
type VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliverySearchScope struct {

	// Scope search to the specified cluster_name; supports globbing; default (*).
	ClusterName string `json:"clusterName,omitempty"`

	// Scope search to the specified management_cluster_name; supports globbing; default (*).
	ManagementClusterName string `json:"managementClusterName,omitempty"`

	// Scope search to the specified provisioner_name; supports globbing; default (*).
	ProvisionerName string `json:"provisionerName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliverySearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliverySearchScope) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliverySearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse Response from listing ContinuousDeliveries.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.continuousdelivery.ListContinuousDeliveriesResponse
type VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse struct {

	// List of continuousdeliveries.
	ContinuousDeliveries []*VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryContinuousDelivery `json:"continuousDeliveries"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdContinuousdeliveryListContinuousDeliveriesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
