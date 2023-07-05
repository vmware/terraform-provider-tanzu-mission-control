package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterSpec Spec of the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.Spec
type VmwareTanzuManageV1alpha1AksclusterSpec struct {

	// Name of the cluster in TMC.
	AgentName string `json:"agentName,omitempty"`

	// Name of the cluster group to which this cluster belongs.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// AKS config for the cluster.
	Config *VmwareTanzuManageV1alpha1AksclusterClusterConfig `json:"config,omitempty"`

	// Optional proxy name is the name of the Proxy Config to be used for the cluster.
	ProxyName string `json:"proxyName,omitempty"`

	// Resource ID of the cluster in Azure.
	ResourceID string `json:"resourceId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
