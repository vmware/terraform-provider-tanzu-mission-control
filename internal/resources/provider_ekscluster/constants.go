/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package providerekscluster

const (
	ResourceName = "tanzu-mission-control_provider-ekscluster"

	credentialNameKey        = "credential_name" //nolint:gosec
	regionKey                = "region"
	nameKey                  = "name"
	specKey                  = "spec"
	statusKey                = "status"
	waitKey                  = "ready_wait_timeout"
	clusterGroupKey          = "cluster_group"
	clusterGroupDefaultValue = "default"
	proxyNameKey             = "proxy"
	agentNameKey             = "agent_name"
	eksARNKey                = "eks_arn"
)
