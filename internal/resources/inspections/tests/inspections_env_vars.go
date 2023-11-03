/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionstests

import (
	"os"

	"github.com/pkg/errors"
)

type ClusterClassEnvVar string

const (
	// TKGM Env Vars.
	ManagementClusterNameEnv ClusterClassEnvVar = "MANAGEMENT_CLUSTER_NAME"
	ProvisionerNameEnv       ClusterClassEnvVar = "PROVISIONER_NAME"
	ClusterNameEnv           ClusterClassEnvVar = "CLUSTER_NAME"
	InspectionNameEnv        ClusterClassEnvVar = "INSPECTION_NAME"
)

var ClusterEnvironmentVariables = map[ClusterClassEnvVar]bool{
	ManagementClusterNameEnv: true,
	ProvisionerNameEnv:       true,
	ClusterNameEnv:           true,
	InspectionNameEnv:        true,
}

func ReadClusterEnvironmentVariables() (envVars map[ClusterClassEnvVar]string, errs []error) {
	envVars = make(map[ClusterClassEnvVar]string)
	errs = make([]error, 0)

	for k := range ClusterEnvironmentVariables {
		envVarVal, exists := os.LookupEnv(string(k))

		if exists {
			envVars[k] = envVarVal
		} else {
			errs = append(errs, errors.Errorf("Environment variable '%s' is required!", k))
		}
	}

	if len(errs) > 0 {
		envVars = nil
	}

	return envVars, errs
}
