---
Title: "Cluster Resource"
Description: |-
    Creating the Tanzu Kubernetes cluster resource.
---

# Cluster

The `tanzu-mission-control_cluster` resource enables you to attach conformant Kubernetes clusters for management through Tanzu Mission Control.
With Tanzu Kubernetes clusters, you can also provision resources to create new workload clusters.

A Tanzu Kubernetes cluster is an opinionated installation of Kubernetes open-source software that is built and supported by VMware.
It is part of a Tanzu Kubernetes Grid instance that includes the following components:

- **management cluster** - a Kubernetes cluster that performs the role of the primary management and operational center for the Tanzu Kubernetes Grid instance
- **provisioner** - a namespace on the management cluster that contains one or more workload clusters
- **workload cluster** - a Tanzu Kubernetes cluster that runs your application workloads

## Attach Cluster

To use the **Tanzu Mission Control provider** to attach an existing conformant Kubernetes cluster,
you must have `cluster.admin` permissions on the cluster and `clustergroup.edit` permissions in Tanzu Mission Control.
For more information, please refer [Attach a Cluster.][attach-cluster]

[attach-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-6DF2CE3E-DD07-499B-BC5E-6B3B2E02A070.html

### Example Usage

```terraform
# Create Tanzu Mission Control attach cluster entry
resource "tanzu-mission-control_cluster" "attach_cluster_without_apply" {
  management_cluster_name = "attached"         # Default: attached
  provisioner_name        = "attached"         # Default: attached
  name                    = "terraform-attach" # Required

  meta {
    description = "create attach cluster from terraform"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" # Default: default
  }

  ready_wait_timeout = "0s" # Shouldn't wait for the default time of 3m in this case

  # The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}
```


## Attach Cluster with Kubeconfig

### Example Usage

```terraform
# Create Tanzu Mission Control attach cluster with k8s cluster kubeconfig path provided
# The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tanzu-mission-control_cluster" "attach_cluster_with_kubeconfig_path" {
  management_cluster_name = "attached"     # Default: attached
  provisioner_name        = "attached"     # Default: attached
  name                    = "demo-cluster" # Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config-path>" # Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" # Default: default
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
}

# Create Tanzu Mission Control attach cluster with k8s cluster kubeconfig provided
# The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tanzu-mission-control_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     # Default: attached
  provisioner_name        = "attached"     # Default: attached
  name                    = "demo-cluster" # Required

  attach_k8s_cluster {
    kubeconfig_raw = var.kubeconfig # Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" # Default: default
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
}

variable "kubeconfig" {
  default = <<EOF
<config>
EOF
}
```


## Attach Cluster with Proxy

### Example Usage

```terraform
# Create Tanzu Mission Control attach cluster entry with proxy
resource "tanzu-mission-control_cluster" "attach_cluster_with_proxy" {
  management_cluster_name = "attached"               # Default: attached
  provisioner_name        = "attached"               # Default: attached
  name                    = "terraform-attach-proxy" # Required

  meta {
    description = "create attach cluster from terraform"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" # Default: default
    proxy         = "proxy-name"
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
}
```


## Tanzu Kubernetes Grid Service Workload Cluster

To use the **Tanzu Mission Control provider** for creating a new cluster, you must have access to an existing Tanzu Kubernetes Grid management cluster with a provisioner namespace wherein the cluster needs to be created.
For more information, please refer [managing the Lifecycle of Tanzu Kubernetes Clusters][create-cluster]
and
[Cluster Lifecycle Management.][lifecycle-management]

You must also have the appropriate permissions:

- To provision a cluster, you must have admin permissions on the management cluster to provision resources within it.
- You must also have `clustergroup.edit` permissions on the cluster group in which you want to put the new cluster.

[create-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-1F847180-1F98-4F8F-9062-46DE9AD8F79D.html
[lifecycle-management]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-concepts/GUID-A6B0184F-269F-41D3-B7FE-5C4F96B3A099.html

### Example Usage

```terraform
# Create Tanzu Mission Control Tanzu Kubernetes Grid Service workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkgs_workload" {
  management_cluster_name = "tkgs-terraform"
  provisioner_name        = "test-gc-e2e-demo-ns"
  name                    = "tkgs-workload"

  meta {
    labels = { "key" : "test" }
  }

  spec {
    cluster_group = "default"
    tkg_service_vsphere {
      settings {
        network {
          pods {
            cidr_blocks = [
              "172.20.0.0/16", # pods cidr block by default has the value `172.20.0.0/16`
            ]
          }
          services {
            cidr_blocks = [
              "10.96.0.0/16", # services cidr block by default has the value `10.96.0.0/16`
            ]
          }
        }
        storage {
          classes = [
            "wcpglobal-storage-profile",
          ]
          default_class = "tkgs-k8s-obj-policy"
        }
      }

      distribution {
        version = "v1.21.2+vmware.1-tkg.1.aad2fe1"
      }

      topology {
        control_plane {
          class         = "best-effort-xsmall"
          storage_class = "gc-storage-profile"
          # storage class is either `wcpglobal-storage-profile` or `gc-storage-profile`
          high_availability = false
          volumes {
            capacity          = 4
            mount_path        = "/var/lib/etcd"
            name              = "etcd-0"
            pvc_storage_class = "tkgs-k8s-obj-policy"
          }
        }
        node_pools {
          spec {
            worker_node_count = "1"
            cloud_label = {
              "key1" : "val1"
            }
            node_label = {
              "key2" : "val2"
            }

            tkg_service_vsphere {
              class         = "best-effort-xsmall"
              storage_class = "gc-storage-profile"
              failure_domain = "domain-c50"
              # storage class is either `wcpglobal-storage-profile` or `gc-storage-profile`
              volumes {
                capacity          = 4
                mount_path        = "/var/lib/etcd"
                name              = "etcd-0"
                pvc_storage_class = "tkgs-k8s-obj-policy"
              }
            }
          }
          info {
            name        = "default-nodepool" # default node pool name `default-nodepool`
            description = "tkgs workload nodepool"
          }
        }
      }
    }
  }
}
```


## Tanzu Kubernetes Grid Vsphere Workload Cluster

To use the **Tanzu Mission Control provider** for creating a new cluster, you must have access to an existing Tanzu Kubernetes Grid management cluster with a provisioner namespace wherein the cluster needs to be created.
For more information, please refer [Managing the Lifecycle of Tanzu Kubernetes Clusters][create-cluster]
and
[Cluster Lifecycle Management.][lifecycle-management]

You must also have the appropriate permissions:

- To provision a cluster, you must have admin permissions on the management cluster to provision resources within it.
- You must also have `clustergroup.edit` permissions on the cluster group in which you want to put the new cluster.

[create-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-1F847180-1F98-4F8F-9062-46DE9AD8F79D.html
[lifecycle-management]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-concepts/GUID-A6B0184F-269F-41D3-B7FE-5C4F96B3A099.html

### Example Usage

```terraform
# Create a Tanzu Kubernetes Grid Vsphere workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkg_vsphere_cluster" {
  management_cluster_name = "tkgm-terraform"
  provisioner_name        = "default"
  name                    = "tkgm-workload"

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" # Default: default
    tkg_vsphere {
      advanced_configs {
        key = "AVI_LABELS"
        value = "test"
      }
      settings {
        network {
          pods {
            cidr_blocks = [
              "172.20.0.0/16", # pods cidr block by default has the value `172.20.0.0/16`
            ]
          }

          services {
            cidr_blocks = [
              "10.96.0.0/16", # services cidr block by default has the value `10.96.0.0/16`
            ]
          }

          api_server_port = 6443
          control_plane_end_point = "10.191.249.39" # optional for AVI enabled option
        }

        security {
          ssh_key = "default"
        }
      }

      distribution {
        os_arch = "amd64"
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
            cloud_label = {
              "key1" : "val1"
            }
            node_label = {
              "key2" : "val2"
            }

            tkg_vsphere {
              vm_config {
                cpu       = "2"
                disk_size = "40"
                memory    = "4096"
              }
            }
          }

          info {
            name        = "default-nodepool" # default node pool name `default-nodepool`
            description = "my nodepool"
          }
        }
      }
    }
  }
}
```

## Tanzu Kubernetes Grid AWS Workload Cluster

To use the **Tanzu Mission Control provider** for creating a new cluster, you must have access to an existing Tanzu Kubernetes Grid management cluster with a provisioner namespace wherein the cluster needs to be created.
For more information, please refer [Managing the Lifecycle of Tanzu Kubernetes Clusters][create-cluster]
and
[Cluster Lifecycle Management.][lifecycle-management]

You must also have the appropriate permissions:

- To provision a cluster, you must have admin permissions on the management cluster to provision resources within it.
- You must also have `clustergroup.edit` permissions on the cluster group in which you want to put the new cluster.

[create-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-1F847180-1F98-4F8F-9062-46DE9AD8F79D.html
[lifecycle-management]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-concepts/GUID-A6B0184F-269F-41D3-B7FE-5C4F96B3A099.html

### Example Usage

```terraform
# Create a Tanzu Mission Control Tanzu Kubernetes Grid AWS workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkg_aws_cluster" {
  management_cluster_name = "tkgm-aws-terraform" // Default: attached
  provisioner_name        = "default"            // Default: attached
  name                    = "tkgm-aws-workload"  // Required

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" // Default: default
    tkg_aws {
      advanced_configs {
        key = "AWS_SECURITY_GROUP_BASTION"
        value = "sg-01376425482384"
      }
      settings {
        network {
          cluster {
            pods {
              cidr_blocks = "100.96.0.0/11" // Required
            }

            services {
              cidr_blocks = "100.64.0.0/13" // Required
            }
          }
          provider {
            subnets {
              availability_zone = "us-west-2a"
              cidr_block_subnet = "10.0.1.0/24"
              is_public         = true
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
          ssh_key = "jumper_ssh_key-sh-1585288-220404-010941" // Required
        }
      }

      distribution {
        os_arch = "amd64"
        os_name = "photon"
        os_version = "3"
        region  = "us-west-2"              // Required
        version = "v1.21.2+vmware.1-tkg.2" // Required
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
              instance_type     = "m5.large"
              node_placement {
                aws_availability_zone = "us-west-2a"
              }

              version = "v1.21.2+vmware.1-tkg.2"
            }
          }

          info {
            name = "md-0" // Required
          }
        }
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of this cluster

### Optional

- `attach_k8s_cluster` (Block List, Max: 1) (see [below for nested schema](#nestedblock--attach_k8s_cluster))
- `management_cluster_name` (String) Name of the management cluster
- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))
- `provisioner_name` (String) Provisioner of the cluster
- `ready_wait_timeout` (String) Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero. Should be set to 0 in case of simple attach cluster where kubeconfig input is not provided.
- `spec` (Block List, Max: 1) Spec for the cluster (see [below for nested schema](#nestedblock--spec))

### Read-Only

- `id` (String) The ID of this resource.
- `status` (Map of String) Status of the cluster

<a id="nestedblock--attach_k8s_cluster"></a>
### Nested Schema for `attach_k8s_cluster`

Optional:

- `description` (String) Attach cluster description
- `kubeconfig_file` (String) Attach cluster KUBECONFIG path
- `kubeconfig_raw` (String, Sensitive) Attach cluster KUBECONFIG


<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource


<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Optional:

- `cluster_group` (String) Name of the cluster group to which this cluster belongs
- `image_registry` (String) Optional image registry name is the name of the image registry to be used for the cluster
- `proxy` (String) Optional proxy name is the name of the Proxy Config to be used for the cluster
- `tkg_aws` (Block List, Max: 1) The Tanzu Kubernetes Grid (TKGm) AWS cluster spec (see [below for nested schema](#nestedblock--spec--tkg_aws))
- `tkg_service_vsphere` (Block List, Max: 1) The Tanzu Kubernetes Grid Service (TKGs) cluster spec (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere))
- `tkg_vsphere` (Block List, Max: 1) The Tanzu Kubernetes Grid (TKGm) vSphere cluster spec (see [below for nested schema](#nestedblock--spec--tkg_vsphere))

<a id="nestedblock--spec--tkg_aws"></a>
### Nested Schema for `spec.tkg_aws`

Required:

- `distribution` (Block List, Min: 1, Max: 1) Kubernetes version distribution for the cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--distribution))
- `settings` (Block List, Min: 1, Max: 1) AWS related settings for workload cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--settings))
- `topology` (Block List, Min: 1, Max: 1) Topology configuration of the cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--topology))

Optional:

- `advanced_configs` (Block List) Advanced configuration for TKGm cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--advanced_configs))

<a id="nestedblock--spec--tkg_aws--distribution"></a>
### Nested Schema for `spec.tkg_aws.distribution`

Required:

- `region` (String) Specifies region of the cluster
- `version` (String) Specifies version of the cluster

Optional:

- `os_arch` (String) Arch of the OS used for the cluster
- `os_name` (String) Name of the OS used for the cluster
- `os_version` (String) Version of the OS used for the cluster
- `provisioner_credential_name` (String) Specifies name of the account in which to create the cluster


<a id="nestedblock--spec--tkg_aws--settings"></a>
### Nested Schema for `spec.tkg_aws.settings`

Required:

- `network` (Block List, Min: 1, Max: 1) Network Settings specifies network-related settings for the cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--settings--network))
- `security` (Block List, Min: 1, Max: 1) Security Settings specifies security-related settings for the cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--settings--security))

<a id="nestedblock--spec--tkg_aws--settings--network"></a>
### Nested Schema for `spec.tkg_aws.settings.network`

Required:

- `cluster` (Block List, Min: 1, Max: 1) Cluster network specifies kubernetes network information for the cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--settings--network--cluster))
- `provider` (Block List, Min: 1) Provider Network specifies provider specific network information for the cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--settings--network--provider))

<a id="nestedblock--spec--tkg_aws--settings--network--cluster"></a>
### Nested Schema for `spec.tkg_aws.settings.network.cluster`

Required:

- `pods` (Block List, Min: 1) Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16 (see [below for nested schema](#nestedblock--spec--tkg_aws--settings--network--cluster--pods))
- `services` (Block List, Min: 1) Service CIDR for kubernetes services defaults to 10.96.0.0/12 (see [below for nested schema](#nestedblock--spec--tkg_aws--settings--network--cluster--services))

Optional:

- `api_server_port` (Number) APIServerPort specifies the port address for the cluster that defaults to 6443.

<a id="nestedblock--spec--tkg_aws--settings--network--cluster--pods"></a>
### Nested Schema for `spec.tkg_aws.settings.network.cluster.pods`

Required:

- `cidr_blocks` (String) CIDRBlocks specifies one or more of IP address ranges


<a id="nestedblock--spec--tkg_aws--settings--network--cluster--services"></a>
### Nested Schema for `spec.tkg_aws.settings.network.cluster.services`

Required:

- `cidr_blocks` (String) CIDRBlocks specifies one or more of IP address ranges



<a id="nestedblock--spec--tkg_aws--settings--network--provider"></a>
### Nested Schema for `spec.tkg_aws.settings.network.provider`

Required:

- `vpc` (Block List, Min: 1, Max: 1) AWS VPC configuration for the cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--settings--network--provider--vpc))

Optional:

- `subnets` (Block List) Optional list of subnets used to place the nodes in the cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--settings--network--provider--subnets))

<a id="nestedblock--spec--tkg_aws--settings--network--provider--vpc"></a>
### Nested Schema for `spec.tkg_aws.settings.network.provider.vpc`

Optional:

- `cidr_block_vpc` (String) CIDR for AWS VPC. A valid example is 10.0.0.0/16
- `vpc_id` (String) AWS VPC ID. The rest of the fields are ignored if this field is specified. Kindly add the VPC ID to the terraform script in case of existing VPC.


<a id="nestedblock--spec--tkg_aws--settings--network--provider--subnets"></a>
### Nested Schema for `spec.tkg_aws.settings.network.provider.subnets`

Optional:

- `availability_zone` (String) AWS availability zone e.g. us-west-2a
- `cidr_block_subnet` (String) CIDR for AWS subnet which must be in the range of AWS VPC CIDR block
- `is_public` (Boolean) Describes if it is public subnet or private subnet
- `subnet_id` (String) This is the subnet ID of AWS. The rest of the fields are ignored if this field is specified




<a id="nestedblock--spec--tkg_aws--settings--security"></a>
### Nested Schema for `spec.tkg_aws.settings.security`

Required:

- `ssh_key` (String) SSH key for provisioning and accessing the cluster VMs



<a id="nestedblock--spec--tkg_aws--topology"></a>
### Nested Schema for `spec.tkg_aws.topology`

Required:

- `control_plane` (Block List, Min: 1, Max: 1) AWS specific control plane configuration for workload cluster object (see [below for nested schema](#nestedblock--spec--tkg_aws--topology--control_plane))

Optional:

- `node_pools` (Block List) Nodepool specific configuration (see [below for nested schema](#nestedblock--spec--tkg_aws--topology--node_pools))

<a id="nestedblock--spec--tkg_aws--topology--control_plane"></a>
### Nested Schema for `spec.tkg_aws.topology.control_plane`

Required:

- `availability_zones` (List of String) List of availability zones for the control plane nodes
- `instance_type` (String) Control plane instance type

Optional:

- `high_availability` (Boolean) Flag which controls if the cluster needs to be highly available. HA cluster creates three controlplane machines, and non HA creates just one


<a id="nestedblock--spec--tkg_aws--topology--node_pools"></a>
### Nested Schema for `spec.tkg_aws.topology.node_pools`

Required:

- `info` (Block List, Min: 1, Max: 1) Info is the meta information of nodepool for cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--topology--node_pools--info))

Optional:

- `spec` (Block List, Max: 1) Spec for the cluster nodepool (see [below for nested schema](#nestedblock--spec--tkg_aws--topology--node_pools--spec))

<a id="nestedblock--spec--tkg_aws--topology--node_pools--info"></a>
### Nested Schema for `spec.tkg_aws.topology.node_pools.info`

Required:

- `name` (String) Name of the nodepool

Optional:

- `description` (String) Description of the nodepool


<a id="nestedblock--spec--tkg_aws--topology--node_pools--spec"></a>
### Nested Schema for `spec.tkg_aws.topology.node_pools.spec`

Optional:

- `tkg_aws` (Block List, Max: 1) Nodepool config for tkg aws (see [below for nested schema](#nestedblock--spec--tkg_aws--topology--node_pools--spec--tkg_aws))
- `worker_node_count` (String) Count is the number of nodes

<a id="nestedblock--spec--tkg_aws--topology--node_pools--spec--tkg_aws"></a>
### Nested Schema for `spec.tkg_aws.topology.node_pools.spec.tkg_aws`

Required:

- `instance_type` (String) Nodepool instance type whose potential values could be found using cluster:options api
- `node_placement` (Block List, Min: 1, Max: 1) List of Availability Zones to place the AWS nodes on. Please use this field to provision a nodepool for workload cluster on an attached TKG AWS management cluster (see [below for nested schema](#nestedblock--spec--tkg_aws--topology--node_pools--spec--tkg_aws--node_placement))
- `version` (String) Kubernetes version of the node pool

Optional:

- `availability_zone` (String) Availability zone for the nodepool that is to be used when you are creating a nodepool for cluster in TMC hosted AWS solution

Read-Only:

- `nodepool_subnet_id` (String) Subnet ID of the private subnet in which you want the nodes to be created in

<a id="nestedblock--spec--tkg_aws--topology--node_pools--spec--tkg_aws--node_placement"></a>
### Nested Schema for `spec.tkg_aws.topology.node_pools.spec.tkg_aws.node_placement`

Required:

- `aws_availability_zone` (String) The Availability Zone where the AWS nodes are placed






<a id="nestedblock--spec--tkg_aws--advanced_configs"></a>
### Nested Schema for `spec.tkg_aws.advanced_configs`

Required:

- `key` (String) The key of the advanced configuration parameters
- `value` (String) The value of the advanced configuration parameters



<a id="nestedblock--spec--tkg_service_vsphere"></a>
### Nested Schema for `spec.tkg_service_vsphere`

Required:

- `distribution` (Block List, Min: 1, Max: 1) VSphere specific distribution (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--distribution))
- `settings` (Block List, Min: 1, Max: 1) VSphere related settings for workload cluster (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--settings))
- `topology` (Block List, Min: 1, Max: 1) Topology specific configuration (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology))

<a id="nestedblock--spec--tkg_service_vsphere--distribution"></a>
### Nested Schema for `spec.tkg_service_vsphere.distribution`

Required:

- `version` (String) Version of the cluster

Optional:

- `os_arch` (String) Arch of the OS used for the cluster
- `os_name` (String) Name of the OS used for the cluster
- `os_version` (String) Version of the OS used for the cluster


<a id="nestedblock--spec--tkg_service_vsphere--settings"></a>
### Nested Schema for `spec.tkg_service_vsphere.settings`

Required:

- `network` (Block List, Min: 1, Max: 1) Network Settings specifies network-related settings for the cluster (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--settings--network))

Optional:

- `storage` (Block List, Max: 1) StorageSettings specifies storage-related settings for the cluster (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--settings--storage))

<a id="nestedblock--spec--tkg_service_vsphere--settings--network"></a>
### Nested Schema for `spec.tkg_service_vsphere.settings.network`

Required:

- `pods` (Block List, Min: 1) Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16 (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--settings--network--pods))
- `services` (Block List, Min: 1) Service CIDR for kubernetes services defaults to 10.96.0.0/12 (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--settings--network--services))

<a id="nestedblock--spec--tkg_service_vsphere--settings--network--pods"></a>
### Nested Schema for `spec.tkg_service_vsphere.settings.network.pods`

Required:

- `cidr_blocks` (List of String) CIDRBlocks specifies one or more ranges of IP addresses


<a id="nestedblock--spec--tkg_service_vsphere--settings--network--services"></a>
### Nested Schema for `spec.tkg_service_vsphere.settings.network.services`

Required:

- `cidr_blocks` (List of String) CIDRBlocks specifies one or more ranges of IP addresses



<a id="nestedblock--spec--tkg_service_vsphere--settings--storage"></a>
### Nested Schema for `spec.tkg_service_vsphere.settings.storage`

Optional:

- `classes` (List of String) Classes is a list of storage classes from the supervisor namespace to expose within a cluster. If omitted, all storage classes from the supervisor namespace will be exposed within the cluster.
- `default_class` (String) DefaultClass is the valid storage class name which is treated as the default storage class within a cluster. If omitted, no default storage class is set.



<a id="nestedblock--spec--tkg_service_vsphere--topology"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology`

Required:

- `control_plane` (Block List, Min: 1, Max: 1) Control plane specific configuration (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--control_plane))

Optional:

- `node_pools` (Block List) Nodepool specific configuration (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--node_pools))

<a id="nestedblock--spec--tkg_service_vsphere--topology--control_plane"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.control_plane`

Required:

- `class` (String) Control plane instance type
- `storage_class` (String) Storage Class to be used for storage of the disks which store the root filesystems of the nodes

Optional:

- `high_availability` (Boolean) High Availability or Non High Availability Cluster. HA cluster creates three controlplane machines, and non HA creates just one
- `volumes` (Block List) Configurable volumes for control plane nodes (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--control_plane--volumes))

<a id="nestedblock--spec--tkg_service_vsphere--topology--control_plane--volumes"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.control_plane.volumes`

Optional:

- `capacity` (Number) Volume capacity is in gib
- `mount_path` (String) It is the directory where the volume device is to be mounted
- `name` (String) It is the volume name
- `pvc_storage_class` (String) This is the storage class for PVC which in case omitted, default storage class will be used for the disks



<a id="nestedblock--spec--tkg_service_vsphere--topology--node_pools"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.node_pools`

Required:

- `info` (Block List, Min: 1, Max: 1) Info is the meta information of nodepool for cluster (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--node_pools--info))

Optional:

- `spec` (Block List, Max: 1) Spec for the cluster nodepool (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--node_pools--spec))

<a id="nestedblock--spec--tkg_service_vsphere--topology--node_pools--info"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.node_pools.info`

Optional:

- `description` (String) Description for the nodepool
- `name` (String) Name of the nodepool


<a id="nestedblock--spec--tkg_service_vsphere--topology--node_pools--spec"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.node_pools.spec`

Optional:

- `cloud_label` (Map of String) Cloud labels
- `node_label` (Map of String) Node labels
- `tkg_service_vsphere` (Block List, Max: 1) Nodepool config for tkg service vsphere (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--node_pools--spec--tkg_service_vsphere))
- `worker_node_count` (String) Count is the number of nodes

<a id="nestedblock--spec--tkg_service_vsphere--topology--node_pools--spec--tkg_service_vsphere"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.node_pools.spec.tkg_service_vsphere`

Required:

- `class` (String) Control plane instance type
- `storage_class` (String) Storage Class to be used for storage of the disks which store the root filesystems of the nodes

Optional:

- `failure_domain` (String) Configure the failure domain of node pool. The potential values could be found using cluster:options api.
- `volumes` (Block List) Configurable volumes for control plane nodes (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--node_pools--spec--tkg_service_vsphere--volumes))

<a id="nestedblock--spec--tkg_service_vsphere--topology--node_pools--spec--tkg_service_vsphere--volumes"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.node_pools.spec.tkg_service_vsphere.volumes`

Optional:

- `capacity` (Number) Volume capacity is in gib
- `mount_path` (String) It is the directory where the volume device is to be mounted
- `name` (String) It is the volume name
- `pvc_storage_class` (String) This is the storage class for PVC which in case omitted, default storage class will be used for the disks







<a id="nestedblock--spec--tkg_vsphere"></a>
### Nested Schema for `spec.tkg_vsphere`

Required:

- `distribution` (Block List, Min: 1, Max: 1) VSphere specific distribution (see [below for nested schema](#nestedblock--spec--tkg_vsphere--distribution))
- `settings` (Block List, Min: 1, Max: 1) VSphere related settings for workload cluster (see [below for nested schema](#nestedblock--spec--tkg_vsphere--settings))
- `topology` (Block List, Min: 1, Max: 1) Topology specific configuration (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology))

Optional:

- `advanced_configs` (Block List) Advanced configuration for TKGm cluster (see [below for nested schema](#nestedblock--spec--tkg_vsphere--advanced_configs))

<a id="nestedblock--spec--tkg_vsphere--distribution"></a>
### Nested Schema for `spec.tkg_vsphere.distribution`

Required:

- `version` (String) Version specifies the version of the Kubernetes cluster
- `workspace` (Block List, Min: 1, Max: 1) Workspace defines a workspace configuration for the vSphere cloud provider (see [below for nested schema](#nestedblock--spec--tkg_vsphere--distribution--workspace))

Optional:

- `os_arch` (String) Arch of the OS used for the cluster
- `os_name` (String) Name of the OS used for the cluster
- `os_version` (String) Version of the OS used for the cluster

<a id="nestedblock--spec--tkg_vsphere--distribution--workspace"></a>
### Nested Schema for `spec.tkg_vsphere.distribution.workspace`

Required:

- `datacenter` (String)
- `datastore` (String)
- `folder` (String)
- `resource_pool` (String)
- `workspace_network` (String)



<a id="nestedblock--spec--tkg_vsphere--settings"></a>
### Nested Schema for `spec.tkg_vsphere.settings`

Required:

- `network` (Block List, Min: 1, Max: 1) Network Settings specifies network-related settings for the cluster (see [below for nested schema](#nestedblock--spec--tkg_vsphere--settings--network))
- `security` (Block List, Min: 1, Max: 1) Security Settings specifies security-related settings for the cluster (see [below for nested schema](#nestedblock--spec--tkg_vsphere--settings--security))

<a id="nestedblock--spec--tkg_vsphere--settings--network"></a>
### Nested Schema for `spec.tkg_vsphere.settings.network`

Required:

- `pods` (Block List, Min: 1) Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16 (see [below for nested schema](#nestedblock--spec--tkg_vsphere--settings--network--pods))
- `services` (Block List, Min: 1) Service CIDR for kubernetes services defaults to 10.96.0.0/12 (see [below for nested schema](#nestedblock--spec--tkg_vsphere--settings--network--services))

Optional:

- `api_server_port` (Number) APIServerPort specifies the port address for the cluster that defaults to 6443.
- `control_plane_end_point` (String) ControlPlaneEndpoint specifies the control plane virtual IP address. The value should be unique for every create request, else cluster creation shall fail. This field is not needed when AVI enabled while creating a legacy cluster on TKGm.

<a id="nestedblock--spec--tkg_vsphere--settings--network--pods"></a>
### Nested Schema for `spec.tkg_vsphere.settings.network.pods`

Required:

- `cidr_blocks` (List of String) CIDRBlocks specifies one or more ranges of IP addresses


<a id="nestedblock--spec--tkg_vsphere--settings--network--services"></a>
### Nested Schema for `spec.tkg_vsphere.settings.network.services`

Required:

- `cidr_blocks` (List of String) CIDRBlocks specifies one or more ranges of IP addresses



<a id="nestedblock--spec--tkg_vsphere--settings--security"></a>
### Nested Schema for `spec.tkg_vsphere.settings.security`

Required:

- `ssh_key` (String) SSH key for provisioning and accessing the cluster VMs



<a id="nestedblock--spec--tkg_vsphere--topology"></a>
### Nested Schema for `spec.tkg_vsphere.topology`

Required:

- `control_plane` (Block List, Min: 1, Max: 1) VSphere specific control plane configuration for workload cluster object (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--control_plane))

Optional:

- `node_pools` (Block List) Nodepool specific configuration (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--node_pools))

<a id="nestedblock--spec--tkg_vsphere--topology--control_plane"></a>
### Nested Schema for `spec.tkg_vsphere.topology.control_plane`

Required:

- `vm_config` (Block List, Min: 1, Max: 1) VM specific configuration (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--control_plane--vm_config))

Optional:

- `high_availability` (Boolean) High Availability or Non High Availability Cluster. HA cluster creates three controlplane machines, and non HA creates just one

<a id="nestedblock--spec--tkg_vsphere--topology--control_plane--vm_config"></a>
### Nested Schema for `spec.tkg_vsphere.topology.control_plane.vm_config`

Optional:

- `cpu` (String) Number of CPUs per node
- `disk_size` (String) Root disk size in gigabytes for the VM
- `memory` (String) Memory associated with the node in megabytes



<a id="nestedblock--spec--tkg_vsphere--topology--node_pools"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools`

Required:

- `info` (Block List, Min: 1, Max: 1) Info is the meta information of nodepool for cluster (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--node_pools--info))

Optional:

- `spec` (Block List, Max: 1) Spec for the cluster nodepool (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--node_pools--spec))

<a id="nestedblock--spec--tkg_vsphere--topology--node_pools--info"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools.info`

Required:

- `name` (String) Name of the nodepool

Optional:

- `description` (String) Description of the nodepool


<a id="nestedblock--spec--tkg_vsphere--topology--node_pools--spec"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools.spec`

Optional:

- `tkg_vsphere` (Block List, Max: 1) Nodepool config for tkgm vsphere (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--node_pools--spec--tkg_vsphere))
- `worker_node_count` (String) Count is the number of nodes

<a id="nestedblock--spec--tkg_vsphere--topology--node_pools--spec--tkg_vsphere"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools.spec.tkg_vsphere`

Required:

- `vm_config` (Block List, Min: 1, Max: 1) VM specific configuration (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--node_pools--spec--tkg_vsphere--vm_config))

<a id="nestedblock--spec--tkg_vsphere--topology--node_pools--spec--tkg_vsphere--vm_config"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools.spec.tkg_vsphere.vm_config`

Optional:

- `cpu` (String) Number of CPUs per node
- `disk_size` (String) Root disk size in gigabytes for the VM
- `memory` (String) Memory associated with the node in megabytes






<a id="nestedblock--spec--tkg_vsphere--advanced_configs"></a>
### Nested Schema for `spec.tkg_vsphere.advanced_configs`

Required:

- `key` (String) The key of the advanced configuration parameters
- `value` (String) The value of the advanced configuration parameters
