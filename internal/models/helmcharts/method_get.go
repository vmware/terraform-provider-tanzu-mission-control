/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmchartorgmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartGetResponse Response from getting a Chart.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.fluxcd.helm.repository.chartmetadata.chart.GetChartResponse
type VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartGetResponse struct {

	// Chart returned.
	Chart *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChart `json:"chart,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartGetResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
