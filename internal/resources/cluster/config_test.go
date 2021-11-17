/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

const testDefaultAttachClusterScript = `
	resource {{.ResourceName}} {{.ResourceNameVar}} {
	  management_cluster_name = "attached"
	  provisioner_name        = "attached"
	  name                    = "{{.Name}}"

	  {{.Meta}}

	  spec {
		cluster_group = "default"
	  }

	  wait_until_ready = false
	}
`

const testAttachClusterWithKubeConfigScript = `
	resource {{.ResourceName}} {{.ResourceNameVar}} {
	  management_cluster_name = "attached"
	  provisioner_name        = "attached"
	  name                    = "{{.Name}}"

	  attach_k8s_cluster {
		kubeconfig_file = "{{.KubeConfigPath}}"
		description     = "optional description about the kube-config provided"
	  }

	  {{.Meta}}

	  spec {
		cluster_group = "default"
	  }

	  wait_until_ready = true
	}
`

const testDataSourceAttachClusterScript = `
	resource {{.ResourceName}} {{.ResourceNameVar}} {
	  management_cluster_name = "attached"
	  provisioner_name        = "attached"
	  name                    = "{{.Name}}"

      {{.Meta}}

	  spec {
		cluster_group = "default"
	  }

	  wait_until_ready = false
	}

	data {{.ResourceName}} {{.DataSourceNameVar}} {
		management_cluster_name = {{.ResourceName}}.{{.ResourceNameVar}}.management_cluster_name
		provisioner_name        = {{.ResourceName}}.{{.ResourceNameVar}}.provisioner_name
		name                    = {{.ResourceName}}.{{.ResourceNameVar}}.name
	}
`

const testTKGsClusterScript = `
	resource {{.ResourceName}} {{.ResourceNameVar}} {
	  management_cluster_name = "{{.ManagementClusterName}}"
	  provisioner_name        = "{{.ProvisionerName}}"
	  name                    = "{{.Name}}"

      spec {
		cluster_group = "default"
    	tkg_service_vsphere {
      	settings {
        	network {
          		pods {
            		cidr_blocks = [
              			"172.20.0.0/16",
            		]
          		}
          	services {
            	cidr_blocks = [
              		"10.96.0.0/16",
            		]
          		}
        	}
      	}

      	distribution {
        	version = "{{.Version}}"
      	}

      	topology {
        	control_plane {
          		class             = "best-effort-xsmall"
          		storage_class     = "{{.StorageClass}}"
          		high_availability = false
        	}
        	node_pools {
          		spec {
					cloud_label = {
						"key1": "val1"
					}
					node_label = {
						"key2": "val2"
					}
            		worker_node_count = "1"
            		tkg_service_vsphere {
						class = "best-effort-xsmall"
              			storage_class = "{{.StorageClass}}"
            		}
          		}
          		info {
            		name = "default-nodepool"
                }
              }
            }
          }
        }
      }
`
