/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package testing

const testDefaultCreateEksClusterScript = `
	resource {{.ResourceName}} {{.ResourceNameVar}} {
		name	= "{{.Name}}"
		region	= "{{.Region}}"
		credential_name = "{{.CredentialName}}"
		
		{{.Meta}}
		
		spec {
			cluster_group = "{{.ClusterGroupName}}"
			config {
				kubernetes_version 	= "{{.KubernetesVersion}}"
				role_arn 			= "arn:aws:iam::{{.AWSAccountNumber}}:role/control-plane.{{.CloudFormationTemplateID}}.eks.tmc.cloud.vmware.com"
				kubernetes_network_config {
					service_cidr = "10.100.0.0/16" // Forces new
				}
				logging {
					api_server         = false
					audit              = true 
					authenticator      = true
					controller_manager =  true 
					scheduler          = true
				}
				vpc { // Required
					enable_private_access = true
					enable_public_access  = true
					public_access_cidrs = [
					  "0.0.0.0/0",
					]
					security_groups = [ // Forces new
					  "sg-0a6768722e9716768",
					]
					subnet_ids = [ // Forces new
					  "subnet-0a184f6302af32a86",
					  "subnet-0ed95d5c212ac62a1",
					  "subnet-0526ecaecde5b1bf7",
					  "subnet-06897e1063cc0cf4e",
					]
				}
			}
			nodepool {
				// could be flattened, but keeping it same for consistency
				info {
					name        = "first-np"
					description = "tf nodepool description"
				}
				spec {
					// Refer to nodepool's schema
					role_arn       = "arn:aws:iam::{{.AWSAccountNumber}}:role/worker.{{.CloudFormationTemplateID}}.eks.tmc.cloud.vmware.com"
					ami_type       = "CUSTOM" // Forces New
					ami_info {
						ami_id = "ami-2qu8409oisdfj0qw"
						override_bootstrap_cmd = "#!/bin/bash\n/etc/eks/bootstrap.sh tf-test-ami"
					}
					capacity_type  = "ON_DEMAND"
					root_disk_size = 40 // Default: 20GiB, forces New
					tags           = { "testnptag" : "testnptagvalue" }
					node_labels    = { "testnplabelkey" : "testnplabelvalue" }
					subnet_ids = [ // Required, forces new
						"subnet-0a184f6302af32a86",
						"subnet-0ed95d5c212ac62a1",
						"subnet-0526ecaecde5b1bf7",
						"subnet-06897e1063cc0cf4e",
					]
					remote_access  {     // Forces new
						ssh_key = "anshulc" // Required (for remote access)
						security_groups = [
							"sg-0a6768722e9716768",
						]
					}
					scaling_config  {
						desired_size = 4
						max_size     = 8
						min_size     = 1
					}
					update_config {
						max_unavailable_nodes = "2"
					}
					instance_types = [ // Forces new
						"t3.medium",
						"m3.large"
					]
				}
			  }
			  nodepool {
				// could be flattened, but keeping it same for consistency
				info  {
					name        = "second-np"
					description = "tf nodepool 2 description"
				}
				spec  {
					// Refer to nodepool's schema
					role_arn    = "arn:aws:iam::{{.AWSAccountNumber}}:role/worker.{{.CloudFormationTemplateID}}.eks.tmc.cloud.vmware.com"
					tags        = { "testnptag" : "testnptagvalue" }
					node_labels = { "testnplabelkey" : "testnplabelvalue" }
					subnet_ids = [ // Required, forces new
						"subnet-0a184f6302af32a86",
						"subnet-0ed95d5c212ac62a1",
						"subnet-0526ecaecde5b1bf7",
						"subnet-06897e1063cc0cf4e",
					]
					launch_template  {
						name    = "{{.LaunchTemplateName}}"
						version = "{{.LaunchTemplateVersion}}"
					}
					scaling_config  {
						desired_size = 4
						max_size     = 8
						min_size     = 1
					}
					update_config  {
						max_unavailable_percentage = "12"
					}
					taints {
						effect = "PREFER_NO_SCHEDULE"
						key    = "randomkey"
						value  = "randomvalue"
					}
				}
			}
		}
		
		ready_wait_timeout = "59m"
	}
`

const testDefaultCreateProviderEksClusterScript = `
	resource {{.ResourceName}} {{.ResourceNameVar}} {
		name	= "{{.Name}}"
		region	= "{{.Region}}"
		credential_name = "{{.CredentialName}}"

		{{.Meta}}

		spec {
			cluster_group = "{{.ClusterGroupName}}"

			eks_arn = "arn:aws:eks:{{.Region}}:{{.AWSAccountNumber}}:cluster/{{.Name}}"
			agent_name = "tf-test-cluster-3"
		}

		ready_wait_timeout = "59m"
	}
`

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
				advanced_configs {
					key = "key-1"
					value = "val-1"
				}
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
					os_arch = "amd"
 					os_name = "photon"
 					os_version = "3"
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
								failure_domain = ""
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
			advanced_configs {
				key = "key-1"
				value = "val-1"
			}
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
					os_arch = "amd"
 					os_name = "photon"
 					os_version = "3"
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
