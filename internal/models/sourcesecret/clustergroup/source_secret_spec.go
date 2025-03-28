// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustergroupsourcesecret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/swag"

	spec "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/sourcesecret/cluster"
)

// VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec Spec for the Source Secret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.sourcesecret.Spec
type VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec struct {

	// Spec of the source secret defined at atomic level.
	AtomicSpec *spec.VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec `json:"atomicSpec,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
