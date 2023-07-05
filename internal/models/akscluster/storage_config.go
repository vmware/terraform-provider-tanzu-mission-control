package models

import (
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AksclusterStorageConfig The storage config.
//
// swagger:model vmware.tanzu.manage.v1alpha1.akscluster.StorageConfig
type VmwareTanzuManageV1alpha1AksclusterStorageConfig struct {

	// Enable the azure disk CSI driver for the storage.
	EnableDiskCsiDriver bool `json:"enableDiskCsiDriver,omitempty"`

	// Enable the azure file CSI driver for the storage.
	EnableFileCsiDriver bool `json:"enableFileCsiDriver,omitempty"`

	// Enable the snapshot controller for the storage.
	EnableSnapshotController bool `json:"enableSnapshotController,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterStorageConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1AksclusterStorageConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AksclusterStorageConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
