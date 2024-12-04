// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package clusterclasstests

import (
	"fmt"

	clusterclassres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/clusterclass"
)

const (
	ClusterClassDataSourceName = "cluster_class_demo"
)

var (
	ClusterClassDataSourceFullName = fmt.Sprintf("data.%s.%s", clusterclassres.ResourceName, ClusterClassDataSourceName)
)

func GetClusterClassConfig(clusterClassEnvVars map[ClusterClassEnvVar]string) string {
	return fmt.Sprintf(`
		data "%s" "%s" {
		  management_cluster_name = "%s"
		  provisioner_name        = "%s"
          name                    = "%s"
		}
		`,
		clusterclassres.ResourceName,
		ClusterClassDataSourceName,
		clusterClassEnvVars[ManagementClusterNameEnv],
		clusterClassEnvVars[ProvisionerNameEnv],
		clusterClassEnvVars[ClusterClassNameEnv],
	)
}
