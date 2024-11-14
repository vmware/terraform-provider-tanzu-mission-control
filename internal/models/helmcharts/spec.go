// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmchartorgmodel

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSpec Spec of the helm chart.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.fluxcd.helm.repository.chartmetadata.chart.Spec
type VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSpec struct {

	// The chart API version.
	APIVersion string `json:"apiVersion,omitempty"`

	// Application version of the chart.
	AppVersion string `json:"appVersion,omitempty"`

	// List of the chart requirements.
	Dependencies []*VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartDependency `json:"dependencies"`

	// Whether this chart is deprecated.
	Deprecated bool `json:"deprecated,omitempty"`

	// A SemVer range of compatible Kubernetes versions.
	KubeVersion string `json:"kubeVersion,omitempty"`

	// Date on which helm chart is released.
	// Format: date-time
	ReleasedAt strfmt.DateTime `json:"releasedAt,omitempty"`

	// List of URLs to source code for this project.
	Sources []string `json:"sources"`

	// List of URLs to download helm chart bundle.
	Urls []string `json:"urls"`

	// Default configuration values for this chart.
	ValuesConfig string `json:"valuesConfig,omitempty"`

	// JSON Schema for imposing a structure on the values.yaml file.
	ValuesSchema interface{} `json:"valuesSchema,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartDependency Dependency for the helm chart.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.fluxcd.helm.repository.chartmetadata.chart.Dependency
type VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartDependency struct {

	// Alias to be used for the chart.
	Alias string `json:"alias,omitempty"`

	// Name of the chart.
	ChartName string `json:"chartName,omitempty"`

	// Version of the chart.
	ChartVersion string `json:"chartVersion,omitempty"`

	// Yaml path that resolves to a boolean, used for enabling/disabling charts.
	Condition string `json:"condition,omitempty"`

	// Holds the mapping of source values to parent key to be imported.
	ImportValues []string `json:"importValues"`

	// Repository URL.
	Repository string `json:"repository,omitempty"`

	// Tags can be used to group charts for enabling/disabling together.
	Tags []string `json:"tags"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartDependency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartDependency) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartDependency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
