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

const testTKGmVsphereClusterScript = `
 	resource {{.ResourceName}} {{.ResourceNameVar}} {
 	  management_cluster_name = "{{.ManagementClusterName}}"
 	  provisioner_name        = "default"
 	  name                    = "{{.Name}}"

	   spec {
		cluster_group = "default"
		tkg_vsphere {
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
				control_plane_end_point = "{{.ControlPlaneEndPoint}}"
			}
			security {
          		ssh_key = "default"
        	}
		}

       	distribution {
         	version = "v1.20.5+vmware.2-tkg.1"
			workspace {
			  datacenter        = "/dc0" 
			  datastore         = "/dc0/datastore/local-0" 
			  workspace_network = "/dc0/network/Avi Internal" 
			  folder            = "/dc0/vm" 
			  resource_pool     = "/dc0/host/cluster0/Resources" 
			}
       	}

       	topology {
         	control_plane {
           		vm_config {
					cpu       = "2" 
					disk_size = "20" 
					memory    = "4096" 
				}
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
             		tkg_vsphere {
 						vm_config {
							cpu       = "2" 
							disk_size = "20" 
							memory    = "4096" 
						}
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
