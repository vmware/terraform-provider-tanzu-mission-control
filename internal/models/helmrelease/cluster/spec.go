// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helmreleaseclustermodel

import (
	"encoding/json"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec Spec of the Helm Release.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.release.Spec
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec struct {

	// Reference to the chart which will be installed.
	ChartRef *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef `json:"chartRef,omitempty"`

	// Inline values in yaml format.
	InlineConfiguration string `json:"inlineConfiguration,omitempty"`

	// Interval at which to reconcile the Helm release.
	Interval string `json:"interval,omitempty"`

	// Name of target namespace.
	TargetNamespace string `json:"targetNamespace,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef ChartRef of the helm release.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.release.ChartRef
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef struct {

	// Name/path of the chart in the helm/git repository.
	Chart string `json:"chart,omitempty"`

	// Repository name.
	RepositoryName string `json:"repositoryName,omitempty"`

	// Repository namespace.
	RepositoryNamespace string `json:"repositoryNamespace,omitempty"`

	// Repository type.
	RepositoryType *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType `json:"repositoryType,omitempty"`

	// Chart version, applicable for helm repository type.
	Version string `json:"version,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType RepositoryType specifies the type of repository.
//
//   - UNSPECIFIED: Repository type is unspecified.
//   - HELM: Helm repository.
//   - GIT: Git repository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.fluxcd.helm.release.RepositoryType
type VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType string

func NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(value VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType) *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType.
func (m VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType) Pointer() *VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeUNSPECIFIED captures enum value "UNSPECIFIED".
	VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeUNSPECIFIED VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType = "UNSPECIFIED"

	// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeHELM captures enum value "HELM".
	VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeHELM VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType = "HELM"

	// VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeGIT captures enum value "GIT".
	VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeGIT VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType = "GIT"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","HELM","GIT"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeEnum = append(vmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeEnum, v)
	}
}
