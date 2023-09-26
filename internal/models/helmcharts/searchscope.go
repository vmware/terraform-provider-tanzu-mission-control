/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmchartorgmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope Scope to search by, any fields left empty will be considered all (*).
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.fluxcd.helm.repository.chartmetadata.chart.SearchScope
type VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope struct {

	// Scope search to the specified chart_metadata_name; supports globbing; default (*).
	ChartMetadataName string `json:"chartMetadataName,omitempty"`

	// Scope search to the specified name; supports globbing; default (*).
	Name string `json:"name,omitempty"`

	// Scope search to the specified repository_name; supports globbing; default (*).
	RepositoryName string `json:"repositoryName,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSearchScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
