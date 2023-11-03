/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package inspectionstests

import (
	"fmt"

	inspectionsres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/inspections"
)

const (
	InspectionListDataSourceName    = "inspection_list_Test"
	InspectionResultsDataSourceName = "inspection_result_Test"
)

var (
	InspectionListDataSourceFullName    = fmt.Sprintf("data.%s.%s", inspectionsres.ResourceNameInspections, InspectionListDataSourceName)
	InspectionResultsDataSourceFullName = fmt.Sprintf("data.%s.%s", inspectionsres.ResourceNameInspectionResults, InspectionResultsDataSourceName)
)

func GetInspectionListConfig(inspectionsEnvVars map[ClusterClassEnvVar]string) string {
	return fmt.Sprintf(`	
		data "%s" "%s" {
		  management_cluster_name = "%s"
		  provisioner_name        = "%s"
          cluster_name            = "%s"
		}
		`,
		inspectionsres.ResourceNameInspections,
		InspectionListDataSourceName,
		inspectionsEnvVars[ManagementClusterNameEnv],
		inspectionsEnvVars[ProvisionerNameEnv],
		inspectionsEnvVars[ClusterNameEnv],
	)
}

func GetInspectionResultsConfig(inspectionsEnvVars map[ClusterClassEnvVar]string) string {
	return fmt.Sprintf(`	
		data "%s" "%s" {
		  management_cluster_name = "%s"
		  provisioner_name        = "%s"
          cluster_name            = "%s"
          name                    = "%s"
		}
		`,
		inspectionsres.ResourceNameInspectionResults,
		InspectionResultsDataSourceName,
		inspectionsEnvVars[ManagementClusterNameEnv],
		inspectionsEnvVars[ProvisionerNameEnv],
		inspectionsEnvVars[ClusterNameEnv],
		inspectionsEnvVars[InspectionNameEnv],
	)
}
