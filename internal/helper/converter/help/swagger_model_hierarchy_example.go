//go:build ignore
// +build ignore

package help

/*
######### Suggested Swagger API Model (Struct) Hierarchy #########
type RootObject struct {
	FullName FullNameObject `json:"fullName,omitempty"`
	Spec SpecObject `json:"spec,omitempty"`
}

type FullNameObject struct {
	Name string `json:"name,omitempty"`
	ProviderName string `json:"providerName,omitempty"`
}

type SpecObject struct {
	Bucket string `json:"bucket,omitempty"`
	CaCert string `json:"caCert,omitempty"`
	Region string `json:"region,omitempty"`
	AssignedGroups []AssignedGroupsObject `json:"assignedGroups,omitempty"`
	Config ConfigObject `json:"config,omitempty"`
	Credential CredentialsObject `json:"credential,omitempty"`
	TargetProvider *TargetProviderObject `json:"targetProvider,omitempty"`
}

type CredentialsObject struct {
	Name string `json:"name,omitempty"`
}

type TargetProviderObject string

type ConfigObject struct {
	AzureConfig AzureConfigObject `json:"azureConfig,omitempty"`
	S3Config S3ConfigObject `json:"s3Config,omitempty"`
}

type AssignedGroupsObject struct {
	Cluster ClusterObject `json:"cluster,omitempty"`
	Clustergroup  ClustergroupObject `json:"clustergroup,omitempty"`
}


type AzureConfigObject struct {
	ResourceGroup string `json:"resourceGroup,omitempty"`
	StorageAccount string `json:"storageAccount,omitempty"`
	SubscriptionID string `json:"subscriptionId,omitempty"`
}

type S3ConfigObject struct {
	PublicURL string `json:"publicUrl,omitempty"`
	S3ForcePathStyle bool `json:"s3ForcePathStyle,omitempty"`
	S3URL string `json:"s3Url,omitempty"`
}

type ClusterObject struct {
	ManagementClusterName string `json:"managementClusterName,omitempty"`
	Name string `json:"name,omitempty"`
	ProvisionerName string `json:"provisionerName,omitempty"`
}

type ClustergroupObject struct {
	Name string `json:"name,omitempty"`
}
*/
