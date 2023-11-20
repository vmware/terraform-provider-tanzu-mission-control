/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helmchartorgmodel

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChart Represents Helm Chart.
//
// swagger:model vmware.tanzu.manage.v1alpha1.organization.fluxcd.helm.repository.chartmetadata.chart.Chart
type VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChart struct {

	// Full name for the Helm Chart.
	FullName *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartFullName `json:"fullName,omitempty"`

	// Metadata for the Helm Chart object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Helm Chart.
	Spec *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChartSpec `json:"spec,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChart) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1OrganizationFluxcdHelmRepositoryChartmetadataChart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
