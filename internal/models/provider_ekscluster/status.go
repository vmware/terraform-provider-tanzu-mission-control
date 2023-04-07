package models

import (
	"github.com/go-openapi/swag"

	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterStatus Status of the Provider Eks Cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.manage.eks.providerekscluster.Status
type VmwareTanzuManageV1alpha1ManageEksProvidereksclusterStatus struct {

	// Conditions of the cluster resource.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// Phase of the cluster resource.
	Phase *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterPhase `json:"phase,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManageEksProvidereksclusterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
