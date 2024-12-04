// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package credential

const (
	ResourceName = "tanzu-mission-control_credential"

	NameKey              = "name"
	specKey              = "spec"
	statusKey            = "status"
	capabilityKey        = "capability"
	providerKey          = "provider"
	dataKey              = "data"
	genericCredentialKey = "generic_credential"
	// nolint:gosec
	awsCredentialKey            = "aws_credential"
	keyValueKey                 = "key_value"
	awsAccountIDKey             = "account_id"
	awsIAMRoleKey               = "iam_role"
	iamRoleARNKey               = "arn"
	iamRoleExtIDKey             = "ext_id"
	typeKey                     = "type"
	azureCredentialKey          = "azure_credential" // nolint:gosec
	servicePrincipalKey         = "service_principal"
	servicePrincipalWithCertKey = "service_principal_with_certificate"
	subscriptionIDKey           = "subscription_id"
	tenantIDKey                 = "tenant_id"
	resourceGroupKey            = "resource_group"
	clientIDKey                 = "client_id"
	clientSecretKey             = "client_secret"
	azureCloudNameKey           = "azure_cloud_name"
	clientCertificateKey        = "client_certificate"
	managedSubscriptionsKey     = "managed_subscriptions"
	waitKey                     = "ready_wait_timeout"
)
