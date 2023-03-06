/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package kubernetessecret

//nolint:all
const (
	NameKey                       = "name"
	NamespaceNameKey              = "namespace_name"
	OrgIDKey                      = "org_id"
	ExportKey                     = "export"
	DockerConfigjsonKey           = "docker_config_json"
	dockerConfigJSONSecretDataKey = ".dockerconfigjson"
	ImageRegistryURLKey           = "image_registry_url"
	UsernameKey                   = "username"
	DataSourceRead                = "dataSourceRead"
	PasswordKey                   = "password"
	DefaultSecretTypeValue        = "SECRET_TYPE_UNSPECIFIED"
	SecretPhaseKey                = "secret_phase"
	SecretExportPhaseKey          = "secret_export_phase"
	specKey                       = "spec"
	statusKey                     = "status"
	Ready                         = "Ready"
	ResourceName                  = "tanzu-mission-control_kubernetes_secret"
)
