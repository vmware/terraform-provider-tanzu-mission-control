// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package tanzuekubernetesclustertests

import (
	"os"

	"github.com/pkg/errors"
)

type ClusterType string
type ClusterEnvVar string
type EnvVarRestriction string

const (
	// TKGM Env Vars.
	TKGMManagementClusterNameEnv ClusterEnvVar = "TKGM_MANAGEMENT_CLUSTER_NAME"
	TKGMProvisionerNameEnv       ClusterEnvVar = "TKGM_PROVISIONER_NAME"
	TKGMClusterVersionEnv        ClusterEnvVar = "TKGM_CLUSTER_VERSION"
	TKGMClusterClassEnv          ClusterEnvVar = "TKGM_CLUSTER_CLASS"
	TKGMClusterVariablesEnv      ClusterEnvVar = "TKGM_CLUSTER_VARIABLES"
	TKGMOSImageNameEnv           ClusterEnvVar = "TKGM_OS_IMAGE_NAME"
	TKGMOSImageVersionEnv        ClusterEnvVar = "TKGM_OS_IMAGE_VERSION"
	TKGMOSImageArchEnv           ClusterEnvVar = "TKGM_OS_IMAGE_ARCH"
	TKGMWorkerClassEnv           ClusterEnvVar = "TKGM_WORKER_CLASS"
	TKGMNodePoolOverridesEnv     ClusterEnvVar = "TKGM_NODE_POOL_OVERRIDES"

	// TKGS Env Vars.
	TKGSManagementClusterNameEnv ClusterEnvVar = "TKGS_MANAGEMENT_CLUSTER_NAME"
	TKGSProvisionerNameEnv       ClusterEnvVar = "TKGS_PROVISIONER_NAME"
	TKGSClusterVersionEnv        ClusterEnvVar = "TKGS_CLUSTER_VERSION"
	TKGSClusterClassEnv          ClusterEnvVar = "TKGS_CLUSTER_CLASS"
	TKGSClusterVariablesEnv      ClusterEnvVar = "TKGS_CLUSTER_VARIABLES"
	TKGSOSImageNameEnv           ClusterEnvVar = "TKGS_OS_IMAGE_NAME"
	TKGSOSImageVersionEnv        ClusterEnvVar = "TKGS_OS_IMAGE_VERSION"
	TKGSOSImageArchEnv           ClusterEnvVar = "TKGS_OS_IMAGE_ARCH"
	TKGSWorkerClassEnv           ClusterEnvVar = "TKGS_WORKER_CLASS"
	TKGSNodePoolOverridesEnv     ClusterEnvVar = "TKGS_NODE_POOL_OVERRIDES"
)

const (
	TKGMClusterType ClusterType = "TKGM"
	TKGSClusterType ClusterType = "TKGS"
)

const (
	RequiredEnvVar EnvVarRestriction = "Required"
	OptionalEnvVar EnvVarRestriction = "Optional"
)

var (
	ClusterEnvironmentVariables = map[ClusterType]map[EnvVarRestriction]map[ClusterEnvVar]bool{
		TKGMClusterType: {
			RequiredEnvVar: {
				TKGMManagementClusterNameEnv: true,
				TKGMProvisionerNameEnv:       true,
				TKGMClusterVersionEnv:        true,
				TKGMClusterClassEnv:          true,
				TKGMClusterVariablesEnv:      true,
				TKGMOSImageNameEnv:           true,
				TKGMOSImageVersionEnv:        true,
				TKGMOSImageArchEnv:           true,
				TKGMWorkerClassEnv:           true,
			},
			OptionalEnvVar: {
				TKGMNodePoolOverridesEnv: true,
			},
		},
		TKGSClusterType: {
			RequiredEnvVar: {
				TKGSManagementClusterNameEnv: true,
				TKGSProvisionerNameEnv:       true,
				TKGSClusterVersionEnv:        true,
				TKGSClusterClassEnv:          true,
				TKGSClusterVariablesEnv:      true,
				TKGSOSImageNameEnv:           true,
				TKGSOSImageVersionEnv:        true,
				TKGSOSImageArchEnv:           true,
				TKGSWorkerClassEnv:           true,
			},
			OptionalEnvVar: {
				TKGSNodePoolOverridesEnv: true,
			},
		},
	}
)

func ReadClusterEnvironmentVariables() (envVars map[ClusterType]map[ClusterEnvVar]string, errs []error) {
	envVars = make(map[ClusterType]map[ClusterEnvVar]string)
	errs = make([]error, 0)

	for key, value := range ClusterEnvironmentVariables {
		clusterTypeEnvVars := make(map[ClusterEnvVar]string)

		for k := range value[RequiredEnvVar] {
			envVarVal, exists := os.LookupEnv(string(k))

			if exists {
				clusterTypeEnvVars[k] = envVarVal
			} else {
				errs = append(errs, errors.Errorf("Environment variable '%s' is required!", k))
			}
		}

		for k := range value[OptionalEnvVar] {
			envVarVal, exists := os.LookupEnv(string(k))

			if exists {
				clusterTypeEnvVars[k] = envVarVal
			} else {
				clusterTypeEnvVars[k] = ""
			}
		}

		envVars[key] = clusterTypeEnvVars
	}

	if len(errs) > 0 {
		envVars = nil
	}

	return envVars, errs
}
