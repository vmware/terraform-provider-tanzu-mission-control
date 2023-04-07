package models

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster ProviderEksCluster is an EKS cluster resource identified in AWS.
// It includes all unmanaged EKS clusters and managed EKS clusters in a TMC credential associated with an AWS account.
//
// swagger:model vmware.tanzu.manage.v1alpha1.manage.eks.providerekscluster.ProviderEksCluster
type VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster struct {

	// Full name for the cluster.
	FullName *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterFullName `json:"fullName,omitempty"`

	// Metadata for the cluster object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the cluster.
	Spec *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterSpec `json:"spec,omitempty"`

	// Status for the cluster.
	Status *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManageEksProvidereksclusterProviderEksCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
