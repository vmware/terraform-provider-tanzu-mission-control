// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmchartorgmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartFullName Full name of the helm chart.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.fluxcd.helm.repository.chartmetadata.chart.FullName
type VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartFullName struct {

	// Name of the helm chart.
	ChartMetadataName string `json:"chartMetadataName,omitempty"`

	// Version of helm chart such as 0.5.1
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of helm repository.
	RepositoryName string `json:"repositoryName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
