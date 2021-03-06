/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
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
		
		ready_wait_timeout = "3m"
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
		
		ready_wait_timeout = "3m"
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
		
		ready_wait_timeout = "3m"
	}

	data {{.ResourceName}} {{.DataSourceNameVar}} {
		management_cluster_name = {{.ResourceName}}.{{.ResourceNameVar}}.management_cluster_name
		provisioner_name        = {{.ResourceName}}.{{.ResourceNameVar}}.provisioner_name
		name                    = {{.ResourceName}}.{{.ResourceNameVar}}.name
	}
`
const testTKGmAWSClusterScript = `
 	resource {{.ResourceName}} {{.ResourceNameVar}} {
		management_cluster_name = "{{.ManagementClusterName}}"
		provisioner_name        = "{{.ProvisionerName}}"
		name                    = "{{.Name}}"

		spec {
			cluster_group = "default"
			tkg_aws {
				settings {
					network {
						cluster {
							pods {
							  cidr_blocks = "100.96.0.0/11" 
							}
				
							services {
							  cidr_blocks = "100.64.0.0/13" 
							}
							api_server_port = 6443
						}
						provider {
							subnets {
								availability_zone = "us-west-2a"
								cidr_block_subnet = "10.0.1.0/24"
								is_public = true
							}
							subnets {
								availability_zone = "us-west-2a"
								cidr_block_subnet = "10.0.0.0/24"
							}
							vpc {
								cidr_block_vpc = "10.0.0.0/16"
							}
						}
					}
					security {
						ssh_key = "jumper_ssh_key-sh-1643378-220418-062857"
					}
				}
		
				distribution {
					region = "us-west-2"
					version = "v1.21.2+vmware.1-tkg.2"
				}
		
				topology {
					control_plane {
						availability_zones = [
							"us-west-2a",
						]
						instance_type = "m5.large"
					}
					node_pools {
						spec {
							worker_node_count = "2"
							tkg_aws {
								availability_zone = "us-west-2a"
								instance_type = "m5.large"
								node_placement {
									aws_availability_zone = "us-west-2a"
								}
								version = "v1.21.2+vmware.1-tkg.2"
							}
						}
						info {
							name = "md-0"
						}
					}
				}
   			}
		}
	}
 `

const testTKGsClusterScript = `
	resource {{.ResourceName}} {{.ResourceNameVar}} {
		management_cluster_name = "{{.ManagementClusterName}}"
		provisioner_name        = "{{.ProvisionerName}}"
		name                    = "{{.Name}}"
		
		spec {
			cluster_group = "e2e-cvs-cg"
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
					storage {
						classes = [
							"wcpglobal-storage-profile",
						]
						default_class = "wcpglobal-storage-profile"
					}
				}
			
				distribution {
					version = "{{.Version}}"
				}
			
				topology {
					control_plane {
						class             = "best-effort-2xlarge"
						storage_class     = "{{.StorageClass}}"
						high_availability = false
						volumes {
							capacity          = 4
							mount_path        = "/var/lib/etcd"
							name              = "etcd-0"
							pvc_storage_class = "wcpglobal-storage-profile"
						}
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
								class = "best-effort-2xlarge"
								storage_class = "{{.StorageClass}}"
								volumes {
									capacity          = 4
									mount_path        = "/var/lib/etcd"
									name              = "etcd-0"
									pvc_storage_class = "wcpglobal-storage-profile"
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

const testTKGmVsphereClusterScript = `
 	resource {{.ResourceName}} {{.ResourceNameVar}} {
		management_cluster_name = "{{.ManagementClusterName}}"
		provisioner_name        = "{{.ProvisionerName}}"
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
					api_server_port = 6443
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
