package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterAADConfig Configs for Azure Active Directory integration.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.AADConfig
type VmwareTanzuManageV1alpha1AksclusterAADConfig struct {

	// The list of AAD group object IDs that will have admin role of the cluster.
	AdminGroupObjectIds []string `json:"adminGroupObjectIds"`

	// Whether to enable Azure RBAC for Kubernetes authorization.
	EnableAzureRbac bool `json:"enableAzureRbac,omitempty"`

	// Whether to enable managed AAD.
	Managed bool `json:"managed,omitempty"`

	// The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription.
	TenantID string `json:"tenantId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAADConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterAADConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterAADConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
