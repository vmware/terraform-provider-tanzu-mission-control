/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package targetlocationmodels

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderSpecificConfig Provider specific configuration for backup location (https://github.com/heptio/velero/blob/master/docs/api-types/backupstoragelocation.md).
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.TargetProviderSpecificConfig.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderSpecificConfig struct {

	// Azure specific config.
	AzureConfig *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAzureStorageConfiguration `json:"azureConfig,omitempty"`

	// S3 and S3-compatible config.
	S3Config *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationS3Configuration `json:"s3Config,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderSpecificConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderSpecificConfig) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationTargetProviderSpecificConfig

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// ### Azure Config ###

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAzureStorageConfiguration Azure specific storage configuration details.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.AzureStorageConfiguration.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAzureStorageConfiguration struct {

	// Name of the resource group containing the storage account for this backup storage location.
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// Name of the storage account for this backup storage location.
	StorageAccount string `json:"storageAccount,omitempty"`

	// Subscription ID under which all the resources are being managed in azure. Optional.
	SubscriptionID string `json:"subscriptionId,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAzureStorageConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAzureStorageConfiguration) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationAzureStorageConfiguration

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// ### AWS Config ###

// VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationS3Configuration AWS S3 or other S3-compatible storage configuration details.
//
// swagger:model vmware.tanzu.manage.v1alpha1.dataprotection.provider.backuplocation.S3Configuration.
type VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationS3Configuration struct {

	// The service endpoint used for generating download URLs. This field is primarily for local storage services like Minio.
	PublicURL string `json:"publicUrl,omitempty"`

	// A flag for whether to force path style URLs for S3 objects. It is default to false and set it to true when.
	// using local storage service like Minio.
	S3ForcePathStyle bool `json:"s3ForcePathStyle"`

	// The service endpoint for non-AWS S3 storage solution.
	S3URL string `json:"s3Url,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationS3Configuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationS3Configuration) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1DataprotectionProviderBackuplocationS3Configuration

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
