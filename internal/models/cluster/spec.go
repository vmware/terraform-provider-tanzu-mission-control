// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clustermodel

import (
	"github.com/go-openapi/swag"

	tkgawsmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgaws"
	tkgservicevspheremodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgservicevsphere"
	tkgvspheremodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster/tkgvsphere"
)

// VmwareTanzuManageV1alpha1ClusterSpec Spec of the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.Spec
type VmwareTanzuManageV1alpha1ClusterSpec struct {

	// Name of the cluster group to which this cluster belongs.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Optional proxy name is the name of the Proxy Config
	// to be used for the cluster.
	ProxyName string `json:"proxyName,omitempty"`

	// Optional image registry is the name of the Image Registry Config
	// to be used for the cluster.
	ImageRegistry string `json:"imageRegistry,omitempty"`

	// TKG AWS cluster spec.
	TkgAws *tkgawsmodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgawsSpec `json:"tkgAws,omitempty"`

	// TKG Service vSphere cluster spec.
	TkgServiceVsphere *tkgservicevspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgservicevsphereSpec `json:"tkgServiceVsphere,omitempty"`

	// TKG vSphere cluster spec.
	TkgVsphere *tkgvspheremodel.VmwareTanzuManageV1alpha1ClusterInfrastructureTkgvsphereSpec `json:"tkgVsphere,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
