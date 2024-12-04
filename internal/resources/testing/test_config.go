// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

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
				tags           = { "testtag" : "testval", "newtesttag": "newtestval"}
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
					  "sg-0b77767aa25e20fec",
					]
					subnet_ids = [ // Forces new
					  	"subnet-0c285da60b373a4cc", "subnet-0be854d94fa197cb7", "subnet-04975d535cf761785", "subnet-0d50aa17c694457c9"
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
					ami_type       = "AL2_x86_64" // Forces New
					capacity_type  = "ON_DEMAND"
					root_disk_size = 40 // Default: 20GiB, forces New
					tags           = { "testnptag" : "testnptagvalue", "newtesttag": "testingtagvalue"}
					node_labels    = { "testnplabelkey" : "testnplabelvalue" }
					subnet_ids = [ // Required, forces new
						"subnet-0c285da60b373a4cc", "subnet-0be854d94fa197cb7", "subnet-04975d535cf761785", "subnet-0d50aa17c694457c9"
					]
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
					tags        = { "testnptag" : "testnptagvalue", "newtesttag": "testingtagvalue"}
					node_labels = { "testnplabelkey" : "testnplabelvalue" }
					subnet_ids = [ // Required, forces new
						"subnet-0c285da60b373a4cc", "subnet-0be854d94fa197cb7", "subnet-04975d535cf761785", "subnet-0d50aa17c694457c9"
					]
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

const testAttachClusterWithKubeConfigScriptImageRegistry = `
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
			image_registry = "{{.ImageRegistry}}"
		}

		ready_wait_timeout = "3m"
	}
`

const testAttachClusterWithKubeConfigScriptProxy = `
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
			proxy         = "{{.Proxy}}"
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
					storage {
						classes = [
							"{{.StorageClass}}",
						]
					}
				}

				distribution {
					version = "{{.Version}}"
				}

				topology {
					control_plane {
						class             = "best-effort-large"
						storage_class     = "{{.StorageClass}}"
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
								class = "best-effort-large"
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
				}
				security {
					ssh_key = "default"
				}
			}

				distribution {
					os_arch = "amd64"
					os_name = "ubuntu"
					os_version = "20.04"
					version = "v1.23.10+vmware.1-tkg.1"
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
