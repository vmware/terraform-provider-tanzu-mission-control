package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec Spec of the cluster.
//
// swagger:model vmware.tanzu.manage.v1alpha1.manage.eks.providerekscluster.Spec
type VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec struct {

	// Name of the cluster in TMC.
	AgentName string `json:"agentName,omitempty"`

	// ARN of the EKS cluster.
	Arn string `json:"arn,omitempty"`

	// Name of the cluster group to which this cluster belongs.
	ClusterGroupName string `json:"clusterGroupName,omitempty"`

	// Optional proxy name is the name of the Proxy Config
	// to be used for the cluster.
	ProxyName string `json:"proxyName,omitempty"`

	// Tmc_managed field indicates whether the cluster is managed by TMC.
	TmcManaged bool `json:"tmcManaged,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
