/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tanzuekubernetesclustertests

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	tanzukubernetesclusteres "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/tanzukubernetescluster"
)

const (
	TKGMClusterResourceName = "test_tkgm_cluster"
	TKGSClusterResourceName = "test_tkgs_cluster"
)

var (
	TKGMClusterResourceFullName = fmt.Sprintf("%s.%s", tanzukubernetesclusteres.ResourceName, TKGMClusterResourceName)
	TKGSClusterResourceFullName = fmt.Sprintf("%s.%s", tanzukubernetesclusteres.ResourceName, TKGSClusterResourceName)
	TKGMClusterName             = acctest.RandomWithPrefix("test-tkgm-cls")
	TKGSClusterName             = acctest.RandomWithPrefix("test-tkgs-cls")
)

type ResourceTFConfigBuilder struct {
	NodePoolDefinition string
}

func InitResourceTFConfigBuilder() *ResourceTFConfigBuilder {
	tfConfigBuilder := &ResourceTFConfigBuilder{
		NodePoolDefinition: `
		  nodepool {
			name        = "md-%d"
			description = "simple small md"
	
			spec {
			  worker_class = "%s"
			  replicas     = 1
			  overrides    = jsonencode(%s)
	
			  meta {
				labels      = { "key" : "value" }
				annotations = { "key1" : "annotation1" }
			  }
	
			  os_image {
				name    = "%s"
				version = "%s"
				arch    = "%s"
			  }
			}
		  }
		`,
	}

	return tfConfigBuilder
}

func (builder *ResourceTFConfigBuilder) GetTKGMClusterConfig(tkgmEnvVars map[ClusterEnvVar]string, nodePoolsNum int) string {
	nodePools := builder.BuildNodePools(tkgmEnvVars[TKGMWorkerClassEnv], tkgmEnvVars[TKGMNodePoolOverridesEnv],
		tkgmEnvVars[TKGMOSImageNameEnv], tkgmEnvVars[TKGMOSImageVersionEnv], tkgmEnvVars[TKGMOSImageArchEnv],
		nodePoolsNum)

	return fmt.Sprintf(`	
		resource "%s" "%s" {
		  name                    = "%s"
		  management_cluster_name = "%s"
		  provisioner_name        = "%s"
		
		  spec {	
			topology {
			  version           = "%s"
			  cluster_class     = "%s"
			  cluster_variables = jsonencode(%s)
		
			  control_plane {
				replicas = 1
		
				meta {
				  labels      = { "key" : "value" }
				  annotations = { "key1" : "annotation1" }
				}
		
				os_image {
				  name    = "%s"
				  version = "%s"
				  arch    = "%s"
				}
			  }
		
			  %s
		
			  network {
				pod_cidr_blocks = [
				  "100.96.0.0/11",
				]
				service_cidr_blocks = [
				  "100.64.0.0/13",
				]
			  }
		
			  core_addon {
				type     = "cni"
				provider = "antrea"
			  }
		
			  core_addon {
				type     = "cpi"
				provider = "vsphere-cpi"
			  }
		
			  core_addon {
				type     = "csi"
				provider = "vsphere-csi"
			  }
			}
		  }

          timeout_policy {
			timeout = 0
          }
		}
		`,
		tanzukubernetesclusteres.ResourceName,
		TKGMClusterResourceName,
		TKGMClusterName,
		tkgmEnvVars[TKGMManagementClusterNameEnv],
		tkgmEnvVars[TKGMProvisionerNameEnv],
		tkgmEnvVars[TKGMClusterVersionEnv],
		tkgmEnvVars[TKGMClusterClassEnv],
		tkgmEnvVars[TKGMClusterVariablesEnv],
		tkgmEnvVars[TKGMOSImageNameEnv],
		tkgmEnvVars[TKGMOSImageVersionEnv],
		tkgmEnvVars[TKGMOSImageArchEnv],
		nodePools,
	)
}

func (builder *ResourceTFConfigBuilder) GetTKGSClusterConfig(tkgsEnvVars map[ClusterEnvVar]string, nodePoolsNum int) string {
	nodePools := builder.BuildNodePools(tkgsEnvVars[TKGSWorkerClassEnv], tkgsEnvVars[TKGSNodePoolOverridesEnv],
		tkgsEnvVars[TKGSOSImageNameEnv], tkgsEnvVars[TKGSOSImageVersionEnv], tkgsEnvVars[TKGSOSImageArchEnv],
		nodePoolsNum)

	return fmt.Sprintf(`	
		resource "%s" "%s" {
		  name                    = "%s"
		  management_cluster_name = "%s"
		  provisioner_name        = "%s"
		
		  spec {	
			topology {
			  version           = "%s"
			  cluster_class     = "%s"
			  cluster_variables = jsonencode(%s)
		
			  control_plane {
				replicas = 1
		
				meta {
				  labels      = { "key" : "value" }
				  annotations = { "key1" : "annotation1" }
				}
		
			    os_image {
				  name    = "%s"
				  version = "%s"
				  arch    = "%s"
			    }
			  }

              %s
		
			  network {
				pod_cidr_blocks = [
				  "100.96.0.0/11",
				]
				service_cidr_blocks = [
				  "100.64.0.0/13",
				]
			  }
			}
		  }

          timeout_policy {
			timeout = 0
          }
		}
		`,
		tanzukubernetesclusteres.ResourceName,
		TKGSClusterResourceName,
		TKGSClusterName,
		tkgsEnvVars[TKGSManagementClusterNameEnv],
		tkgsEnvVars[TKGSProvisionerNameEnv],
		tkgsEnvVars[TKGSClusterVersionEnv],
		tkgsEnvVars[TKGSClusterClassEnv],
		tkgsEnvVars[TKGSClusterVariablesEnv],
		tkgsEnvVars[TKGSOSImageNameEnv],
		tkgsEnvVars[TKGSOSImageVersionEnv],
		tkgsEnvVars[TKGSOSImageArchEnv],
		nodePools,
	)
}

func (builder *ResourceTFConfigBuilder) BuildNodePools(workerClass string, overrides string, osImageName string,
	osImageVersion string, osImageArch string, nodePoolsNum int) string {
	nodePools := ""

	for i := 0; i < nodePoolsNum; i++ {
		np := fmt.Sprintf(builder.NodePoolDefinition, i, workerClass, overrides, osImageName, osImageVersion, osImageArch)
		nodePools = fmt.Sprintf("%s\n%s", nodePools, np)
	}

	return nodePools
}
