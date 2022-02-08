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

# Attach Cluster

To use the **Tanzu Mission Control provider** to attach an existing conformant Kubernetes cluster,
you must have `cluster.admin` permissions on the cluster and `clustergroup.edit` permissions in Tanzu Mission Control.
For more information, please refer [Attach a Cluster.][attach-cluster]

[attach-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-6DF2CE3E-DD07-499B-BC5E-6B3B2E02A070.html

## Example Usage

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

  # The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}
```


# Attach Cluster with Kubeconfig

## Example Usage

```terraform
# Create Tanzu Mission Control attach cluster with k8s cluster kubeconfig provided
# The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tanzu-mission-control_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     # Default: attached
  provisioner_name        = "attached"     # Default: attached
  name                    = "demo-cluster" # Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config path>" # Required
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
```


# Attach Cluster with Proxy

## Example Usage

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


# Tanzu Kubernetes Grid Service Workload Cluster

To use the **Tanzu Mission Control provider** for creating a new cluster, you must have access to an existing Tanzu Kubernetes Grid management cluster with a provisioner namespace wherein the cluster needs to be created.
For more information, please refer [managing the Lifecycle of Tanzu Kubernetes Clusters][create-cluster]
and
[Cluster Lifecycle Management.][lifecycle-management]

You must also have the appropriate permissions:

- To provision a cluster, you must have admin permissions on the management cluster to provision resources within it.
- You must also have `clustergroup.edit` permissions on the cluster group in which you want to put the new cluster.

[create-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-1F847180-1F98-4F8F-9062-46DE9AD8F79D.html
[lifecycle-management]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-concepts/GUID-A6B0184F-269F-41D3-B7FE-5C4F96B3A099.html

## Example Usage

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
              # storage class is either `wcpglobal-storage-profile` or `gc-storage-profile`
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


# Tanzu Kubernetes Grid Vsphere Workload Cluster

To use the **Tanzu Mission Control provider** for creating a new cluster, you must have access to an existing Tanzu Kubernetes Grid management cluster with a provisioner namespace wherein the cluster needs to be created.
For more information, please refer [Managing the Lifecycle of Tanzu Kubernetes Clusters][create-cluster]
and
[Cluster Lifecycle Management.][lifecycle-management]

You must also have the appropriate permissions:

- To provision a cluster, you must have admin permissions on the management cluster to provision resources within it.
- You must also have `clustergroup.edit` permissions on the cluster group in which you want to put the new cluster.

[create-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-1F847180-1F98-4F8F-9062-46DE9AD8F79D.html
[lifecycle-management]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-concepts/GUID-A6B0184F-269F-41D3-B7FE-5C4F96B3A099.html

## Example Usage

```terraform
# Create a Tanzu Kubernetes Grid Vsphere workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkg_vsphere_cluster" {
  management_cluster_name = "tkgm-terraform"
  provisioner_name        = "default"
  name                    = "tkgm-workload-test"

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "default" # Default: default
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

          control_plane_end_point = "10.191.249.39"
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
            name        = "default-nodepool"
            description = "my nodepool"
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

- **name** (String) Name of this cluster

### Optional

- **attach_k8s_cluster** (Block List, Max: 1) (see [below for nested schema](#nestedblock--attach_k8s_cluster))
- **id** (String) The ID of this resource.
- **management_cluster_name** (String) Name of the management cluster
- **meta** (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))
- **provisioner_name** (String) Provisioner of the cluster
- **ready_wait_timeout** (String) Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero
- **spec** (Block List, Max: 1) Spec for the cluster (see [below for nested schema](#nestedblock--spec))

### Read-Only

- **status** (Map of String) Status of the cluster

<a id="nestedblock--attach_k8s_cluster"></a>
### Nested Schema for `attach_k8s_cluster`

Optional:

- **description** (String) Attach cluster description
- **kubeconfig_file** (String) Attach cluster KUBECONFIG path


<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- **description** (String) Description of the resource
- **labels** (Map of String) Labels for the resource

Read-Only:

- **annotations** (Map of String) Annotations for the resource
- **resource_version** (String) Resource version of the resource
- **uid** (String) UID of the resource


<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Optional:

- **cluster_group** (String) Name of the cluster group to which this cluster belongs
- **proxy** (String) Optional proxy name is the name of the Proxy Config to be used for the cluster
- **tkg_service_vsphere** (Block List, Max: 1) The Tanzu Kubernetes Grid Service (TKGs) cluster spec (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere))
- **tkg_vsphere** (Block List, Max: 1) The Tanzu Kubernetes Grid (TKGm) cluster spec (see [below for nested schema](#nestedblock--spec--tkg_vsphere))

<a id="nestedblock--spec--tkg_service_vsphere"></a>
### Nested Schema for `spec.tkg_service_vsphere`

Required:

- **distribution** (Block List, Min: 1, Max: 1) VSphere specific distribution (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--distribution))
- **settings** (Block List, Min: 1, Max: 1) VSphere related settings for workload cluster (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--settings))
- **topology** (Block List, Min: 1, Max: 1) Topology specific configuration (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology))

<a id="nestedblock--spec--tkg_service_vsphere--distribution"></a>
### Nested Schema for `spec.tkg_service_vsphere.distribution`

Required:

- **version** (String) Version of the cluster


<a id="nestedblock--spec--tkg_service_vsphere--settings"></a>
### Nested Schema for `spec.tkg_service_vsphere.settings`

Required:

- **network** (Block List, Min: 1, Max: 1) Network Settings specifies network-related settings for the cluster (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--settings--network))

<a id="nestedblock--spec--tkg_service_vsphere--settings--network"></a>
### Nested Schema for `spec.tkg_service_vsphere.settings.network`

Required:

- **pods** (Block List, Min: 1) Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16 (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--settings--network--pods))
- **services** (Block List, Min: 1) Service CIDR for kubernetes services defaults to 10.96.0.0/12 (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--settings--network--services))

<a id="nestedblock--spec--tkg_service_vsphere--settings--network--pods"></a>
### Nested Schema for `spec.tkg_service_vsphere.settings.network.pods`

Required:

- **cidr_blocks** (List of String) CIDRBlocks specifies one or more ranges of IP addresses


<a id="nestedblock--spec--tkg_service_vsphere--settings--network--services"></a>
### Nested Schema for `spec.tkg_service_vsphere.settings.network.services`

Required:

- **cidr_blocks** (List of String) CIDRBlocks specifies one or more ranges of IP addresses




<a id="nestedblock--spec--tkg_service_vsphere--topology"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology`

Required:

- **control_plane** (Block List, Min: 1, Max: 1) Control plane specific configuration (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--control_plane))

Optional:

- **node_pools** (Block List) Nodepool specific configuration (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--node_pools))

<a id="nestedblock--spec--tkg_service_vsphere--topology--control_plane"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.node_pools`

Required:

- **class** (String) Control plane instance type
- **storage_class** (String) Storage Class to be used for storage of the disks which store the root filesystems of the nodes

Optional:

- **high_availability** (Boolean) High Availability or Non High Availability Cluster. HA cluster creates three controlplane machines, and non HA creates just one


<a id="nestedblock--spec--tkg_service_vsphere--topology--node_pools"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.node_pools`

Required:

- **info** (Block List, Min: 1, Max: 1) Info is the meta information of nodepool for cluster (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--node_pools--info))

Optional:

- **spec** (Block List, Max: 1) Spec for the cluster nodepool (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--node_pools--spec))

<a id="nestedblock--spec--tkg_service_vsphere--topology--node_pools--info"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.node_pools.info`

Optional:

- **description** (String) Description for the nodepool
- **name** (String) Name of the nodepool


<a id="nestedblock--spec--tkg_service_vsphere--topology--node_pools--spec"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.node_pools.spec`

Optional:

- **cloud_label** (Map of String) Cloud labels
- **node_label** (Map of String) Node labels
- **tkg_service_vsphere** (Block List, Max: 1) Nodepool config for tkg service vsphere (see [below for nested schema](#nestedblock--spec--tkg_service_vsphere--topology--node_pools--spec--tkg_service_vsphere))
- **worker_node_count** (String) Count is the number of nodes

<a id="nestedblock--spec--tkg_service_vsphere--topology--node_pools--spec--tkg_service_vsphere"></a>
### Nested Schema for `spec.tkg_service_vsphere.topology.node_pools.spec.worker_node_count`

Required:

- **class** (String) Control plane instance type
- **storage_class** (String) Storage Class to be used for storage of the disks which store the root filesystems of the nodes






<a id="nestedblock--spec--tkg_vsphere"></a>
### Nested Schema for `spec.tkg_vsphere`

Required:

- **distribution** (Block List, Min: 1, Max: 1) VSphere specific distribution (see [below for nested schema](#nestedblock--spec--tkg_vsphere--distribution))
- **settings** (Block List, Min: 1, Max: 1) VSphere related settings for workload cluster (see [below for nested schema](#nestedblock--spec--tkg_vsphere--settings))
- **topology** (Block List, Min: 1, Max: 1) Topology specific configuration (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology))

<a id="nestedblock--spec--tkg_vsphere--distribution"></a>
### Nested Schema for `spec.tkg_vsphere.distribution`

Required:

- **version** (String) Version specifies the version of the Kubernetes cluster
- **workspace** (Block List, Min: 1, Max: 1) Workspace defines a workspace configuration for the vSphere cloud provider (see [below for nested schema](#nestedblock--spec--tkg_vsphere--distribution--workspace))

<a id="nestedblock--spec--tkg_vsphere--distribution--workspace"></a>
### Nested Schema for `spec.tkg_vsphere.distribution.workspace`

Required:

- **datacenter** (String)
- **datastore** (String)
- **folder** (String)
- **resource_pool** (String)
- **workspace_network** (String)



<a id="nestedblock--spec--tkg_vsphere--settings"></a>
### Nested Schema for `spec.tkg_vsphere.settings`

Required:

- **network** (Block List, Min: 1, Max: 1) Network Settings specifies network-related settings for the cluster (see [below for nested schema](#nestedblock--spec--tkg_vsphere--settings--network))
- **security** (Block List, Min: 1, Max: 1) Security Settings specifies security-related settings for the cluster (see [below for nested schema](#nestedblock--spec--tkg_vsphere--settings--security))

<a id="nestedblock--spec--tkg_vsphere--settings--network"></a>
### Nested Schema for `spec.tkg_vsphere.settings.security`

Required:

- **control_plane_end_point** (String) ControlPlaneEndpoint specifies the control plane virtual IP address. The value should be unique for every create request, else cluster creation shall fail
- **pods** (Block List, Min: 1) Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16 (see [below for nested schema](#nestedblock--spec--tkg_vsphere--settings--security--pods))
- **services** (Block List, Min: 1) Service CIDR for kubernetes services defaults to 10.96.0.0/12 (see [below for nested schema](#nestedblock--spec--tkg_vsphere--settings--security--services))

<a id="nestedblock--spec--tkg_vsphere--settings--security--pods"></a>
### Nested Schema for `spec.tkg_vsphere.settings.security.pods`

Required:

- **cidr_blocks** (List of String) CIDRBlocks specifies one or more ranges of IP addresses


<a id="nestedblock--spec--tkg_vsphere--settings--security--services"></a>
### Nested Schema for `spec.tkg_vsphere.settings.security.services`

Required:

- **cidr_blocks** (List of String) CIDRBlocks specifies one or more ranges of IP addresses



<a id="nestedblock--spec--tkg_vsphere--settings--security"></a>
### Nested Schema for `spec.tkg_vsphere.settings.security`

Required:

- **ssh_key** (String) SSH key for provisioning and accessing the cluster VMs



<a id="nestedblock--spec--tkg_vsphere--topology"></a>
### Nested Schema for `spec.tkg_vsphere.topology`

Required:

- **control_plane** (Block List, Min: 1, Max: 1) VSphere specific control plane configuration for workload cluster object (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--control_plane))

Optional:

- **node_pools** (Block List) Nodepool specific configuration (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--node_pools))

<a id="nestedblock--spec--tkg_vsphere--topology--control_plane"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools`

Required:

- **vm_config** (Block List, Min: 1, Max: 1) VM specific configuration (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--node_pools--vm_config))

Optional:

- **high_availability** (Boolean) High Availability or Non High Availability Cluster. HA cluster creates three controlplane machines, and non HA creates just one

<a id="nestedblock--spec--tkg_vsphere--topology--node_pools--vm_config"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools.vm_config`

Optional:

- **cpu** (String)
- **disk_size** (String)
- **memory** (String)



<a id="nestedblock--spec--tkg_vsphere--topology--node_pools"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools`

Required:

- **info** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--node_pools--info))

Optional:

- **spec** (Block List, Max: 1) Spec for the cluster nodepool (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--node_pools--spec))

<a id="nestedblock--spec--tkg_vsphere--topology--node_pools--info"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools.info`

Required:

- **name** (String)

Optional:

- **description** (String)


<a id="nestedblock--spec--tkg_vsphere--topology--node_pools--spec"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools.spec`

Optional:

- **cloud_label** (Map of String) Cloud labels
- **node_label** (Map of String) Node labels
- **tkg_vsphere** (Block List, Max: 1) Nodepool config for tkgm vsphere (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--node_pools--spec--tkg_vsphere))
- **worker_node_count** (String) Count is the number of nodes

<a id="nestedblock--spec--tkg_vsphere--topology--node_pools--spec--tkg_vsphere"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools.spec.worker_node_count`

Required:

- **vm_config** (Block List, Min: 1, Max: 1) VM specific configuration (see [below for nested schema](#nestedblock--spec--tkg_vsphere--topology--node_pools--spec--worker_node_count--vm_config))

<a id="nestedblock--spec--tkg_vsphere--topology--node_pools--spec--worker_node_count--vm_config"></a>
### Nested Schema for `spec.tkg_vsphere.topology.node_pools.spec.worker_node_count.vm_config`

Optional:

- **cpu** (String)
- **disk_size** (String)
- **memory** (String)