// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package authctx

const (
	endpoint           = "endpoint"
	projectID          = "project_id"
	vmwCloudEndpoint   = "vmw_cloud_endpoint"
	vmwCloudAPIToken   = "vmw_cloud_api_token"
	defaultCSPEndpoint = "console.tanzu.broadcom.com"
	selfManaged        = "self_managed"
	oidcIssuer         = "oidc_issuer"
	smUsername         = "username"
	smPassword         = "password"

	// proxy configs.
	insecureAllowUnverifiedSSL = "insecure_allow_unverified_ssl"
	clientAuthCertFile         = "client_auth_cert_file"
	clientAuthKeyFile          = "client_auth_key_file"
	caFile                     = "ca_file"
	clientAuthCert             = "client_auth_cert"
	clientAuthKey              = "client_auth_key"
	caCert                     = "ca_cert"
)
