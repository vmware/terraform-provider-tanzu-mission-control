// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmchartorgmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartListResponse Response from listing Charts.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.fluxcd.helm.repository.chartmetadata.chart.ListChartsResponse
type VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartListResponse struct {

	// List of charts.
	Charts []*VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChart `json:"charts"`

	// Total count.
	TotalCount string `json:"totalCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartListResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
