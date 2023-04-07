package models

import (
	"fmt"

	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName Full name of the cluster. This includes the object name along
// with any parents or further identifiers.
//
// swagger:model vmware.tanzu.manage.v1alpha1.manage.eks.providerekscluster.FullName
type VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName struct {

	// Name of the credential.
	CredentialName string `json:"credentialName,omitempty"`

	// Name of this cluster in EKS.
	Name string `json:"name,omitempty"`

	// ID of Organization.
	OrgID string `json:"orgId,omitempty"`

	// Name of the region.
	Region string `json:"region,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName) ToString() string {
	if m == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s:%s:%s", m.OrgID, m.CredentialName, m.Region, m.Name)
}
