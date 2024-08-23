---
Title: "Tanzu Kubernetes Cluster (Class-based Cluster)"
Description: |-
   Creating a unified Tanzu Kubernetes Grid cluster.
---

# Tanzu Kubernetes Cluster (Class-based Cluster) Resource

This resource "tanzu-mission-control_tanzu_kubernetes_cluster" enables users to create and manage a Tanzu Kubernetes Grid cluster for both Tanzu Kubernetes Grid Vsphere 2.x & Tanzu Kubernetes Grid Service 2.x.

For more information about creating or managing the workload clusters on TMC refer to [Provision a Cluster using a Cluster Class][provision-cluster-class-cluster].

Cluster variables and node pool overrides are determined by the cluster class defined in the resource.
For identifying the structure of the cluster variables supported in the cluster class, users can utilize the cluster class [data source][cluster-class-datasource].
In order to configure & reuse cluster variables and node pools overrides, it is recommended defining these values in a local variable named after the cluster type
and cluster class version.

[provision-cluster-class-cluster]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-C778E447-DDBB-49FC-B0B2-A8012AC56B0E.html
[cluster-class-datasource]: https://registry.terraform.io/providers/vmware/tanzu-mission-control/latest/docs/data-sources/cluster_class

```
locals {
  tkgm_v110_cluster_variables = { ... }
  tkgm_v110_nodepool_a_overrides = { ... }
  tkgs_cluster_variables = { ... }
  tkgs_nodepool_a_overrides = { ... }
}
```

# Tanzu Kubernetes Grid Vsphere

## Example Usage

```terraform
locals {
  tkgm_v110_cluster_variables = {
    vcenter = {
      server       = "demo-vc-01.server.demo"
      datacenter   = "/Demo-Datacenter"
      datastore    = "/Demo-Datacenter/datastore/vsanDatastore"
      resourcePool = "/Demo-Datacenter/host/Demo-Cluster/Resources"
      folder       = "/Demo-Datacenter/vm/LABS/folder"
      network      = "/Demo-Datacenter/network/tkg-static-ips"
      cloneMode    = "fullClone"
      template     = "/Demo-Datacenter/vm/Templates/TKG-M/Ubuntu/ubuntu-2004-efi-kube-v1.25.7+vmware.2"
    }

    identityRef = {
      kind = "VSphereClusterIdentity"
      name = "tkg-vc-default"
    }

    user = {
      sshAuthorizedKeys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC4WyexTrHlHayExsx5sv6D/9LV+EsKzK0rekJowoFAk00LGZBCtf5+KMcRN1sVyJNfXZ68uuZCp/FHzjTl9yq4ViZaYNv4MjM4rEtpPfV7nrJeCJnGLZZNhBC5iQ5rnDSIW0ReC5vMERU71twAOHFiA019MSD9LfJE6LA/nmTc7zmYBBLtWuKAPbQBeJNM8UweNvXf6tvH82tkBO/emh4HaaaOyrypPOq8SdAD485I1/5m2dhFX52kUIn6R99AzYXE4sUbqd4wa4lKfQNhxeLYVry7bP/A1jAUWft66GhReKcNox9epyhGCm/icAHcjCCbmMOu0cINeZMPF4BvUQ/q+1j+wGqviNwbOa4Rr2VeO64xFi/ChwPLjajiHdOssBnwXzsXkBRYpWYH/gPLD3Y6W4wrjU958oQcMo2LtJ1pB7MIO9X30mDX0Qn/yTRFHqJ3Tm6ZAN+2pEXWISEEGGg1j6QRZ2TZQKL7rDD1nDQSqU7nH46KsaZ01Y4LU9SXZmk= k8s@tce-admin-vm"]
    }

    aviAPIServerHAProvider = true
    vipNetworkInterface    = "eth0"

    worker = {
      machine = {
        diskGiB   = 50
        memoryMiB = 8192
        numCPUs   = 4
      }
      network = {
        mtu           = 0
        nameservers   = ["10.100.100.100"]
        searchDomains = ["server.demo"]
      }
    }

    controlPlane = {
      machine = {
        diskGiB   = 50
        memoryMiB = 8192
        numCPUs   = 4
      }
      network = {
        mtu           = 0
        nameservers   = ["10.100.100.100"]
        searchDomains = ["server.demo"]
      }
    }

    controlPlaneCertificateRotation = {
      activate   = true
      daysBefore = 90
    }

    network = {
      addressesFromPools = [
        {
          apiGroup = "ipam.cluster.x-k8s.io"
          kind     = "InClusterIPPool"
          name     = "inclusterippool"
        }
      ]
    }
  }

  tkgm_v110_nodepool_a_overrides = {
    nodePoolLabels = [
      {
        key   = "nodepool"
        value = "md-0"
      }
    ]

    worker = {
      machine = {
        diskGiB   = 40
        memoryMiB = 8192
        numCPUs   = 4
      }
      network = {
        nameservers   = ["10.100.100.100"]
        searchDomains = ["server.demo"]
      }
    }
  }
}
```
```terraform
resource "tanzu-mission-control_tanzu_kubernetes_cluster" "tkgm_cluster" {
  name                    = "CLS_NAME"
  management_cluster_name = "MANAGEMENT_CLS_NAME"
  provisioner_name        = "PROVISIONER_NAME"

  meta {
    description = "description of the cluster"
    labels      = { "key" : "changedvalues" }
  }

  spec {
    cluster_group_name = "default"

    topology {
      version           = "v1.26.5+vmware.2-tkg.1"
      cluster_class     = "tkg-vsphere-default-v1.1.0"
      cluster_variables = jsonencode(local.tkgm_v110_cluster_variables)

      control_plane {
        replicas = 1

        meta {
          labels      = { "key" : "value" }
          annotations = { "key1" : "annotation1" }
        }

        os_image {
          name    = "ubuntu"
          version = "2004"
          arch    = "amd64"
        }
      }

      nodepool {
        name        = "md-0"
        description = "simple small md"

        spec {
          worker_class = "tkg-worker"
          replicas     = 1
          overrides    = jsonencode(local.tkgm_v110_nodepool_a_overrides)

          meta {
            labels      = { "key" : "value" }
            annotations = { "key1" : "annotation1" }
          }

          os_image {
            name    = "ubuntu"
            version = "2004"
            arch    = "amd64"
          }
        }
      }

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
    timeout             = 60
    wait_for_kubeconfig = true
    fail_on_timeout     = true
  }
}
```

# Tanzu Kubernetes Grid Service

## Example Usage

```terraform
locals {
  tkgs_cluster_variables = {
    "controlPlaneCertificateRotation" : {
      "activate" : true,
      "daysBefore" : 30
    },
    "defaultStorageClass" : "k8s-storage-policy-vsan",
    "defaultVolumeSnapshotClass" : "volumesnapshotclass-delete",
    "nodePoolLabels" : [

    ],
    "nodePoolVolumes" : [
      {
        "capacity" : {
          "storage" : "20G"
        },
        "mountPath" : "/var/lib/containerd",
        "name" : "containerd",
        "storageClass" : "k8s-storage-policy-vsan"
      },
      {
        "capacity" : {
          "storage" : "20G"
        },
        "mountPath" : "/var/lib/kubelet",
        "name" : "kubelet",
        "storageClass" : "k8s-storage-policy-vsan"
      }
    ],
    "ntp" : "172.16.20.10",
    "storageClass" : "k8s-storage-policy-vsan",
    "storageClasses" : [
      "k8s-storage-policy-vsan"
    ],
    "vmClass" : "best-effort-medium"
  }

  tkgs_nodepool_a_overrides = {
    "nodePoolLabels" : [
      {
        "key" : "sample-worker-label",
        "value" : "value"
      }
    ],
    "storageClass" : "k8s-storage-policy-vsan",
    "vmClass" : "best-effort-medium"
  }
}
```
```terraform
resource "tanzu-mission-control_tanzu_kubernetes_cluster" "tkgs_cluster" {
  name                    = "CLS_NAME"
  management_cluster_name = "MANAGEMENT_CLS_NAME"
  provisioner_name        = "PROVISIONER_NAME"

  spec {
    cluster_group_name = "default"

    topology {
      version           = "v1.26.5+vmware.2-fips.1-tkg.1"
      cluster_class     = "tanzukubernetescluster"
      cluster_variables = jsonencode(local.tkgs_cluster_variables)

      control_plane {
        replicas = 1

        os_image {
          name    = "photon"
          version = "3"
          arch    = "amd64"
        }
      }

      nodepool {
        name        = "md-0"
        description = "simple small md"

        spec {
          worker_class = "node-pool"
          replicas     = 1
          overrides    = jsonencode(local.tkgs_nodepool_a_overrides)

          os_image {
            name    = "photon"
            version = "3"
            arch    = "amd64"
          }
        }
      }

      network {
        pod_cidr_blocks = [
          "100.96.0.0/11",
        ]
        service_cidr_blocks = [
          "100.64.0.0/13",
        ]
        service_domain = "cluster.local"
      }
    }
  }

  timeout_policy {
    timeout             = 60
    wait_for_kubeconfig = true
    fail_on_timeout     = true
  }
}
```

## Import Tanzu Kubernetes Grid Cluster
The resource ID for importing an existing Tanzu Kubernetes Grid 2.x cluster class based cluster should be comprised of a full cluster name separated by '/'.

```bash
terraform import tanzu-mission-control_tanzu_kubernetes_cluster.demo_cluster MANAGEMENT_CLUSTER_NAME/PROVISIONER_NAME/CLUSTER_NAME
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `management_cluster_name` (String) Management cluster name
- `name` (String) Cluster name
- `provisioner_name` (String) Cluster provisioner name
- `spec` (Block List, Min: 1, Max: 1) Spec of tanzu kubernetes cluster. (see [below for nested schema](#nestedblock--spec))

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))
- `timeout_policy` (Block List, Max: 1) Timeout policy for Tanzu Kubernetes cluster. (see [below for nested schema](#nestedblock--timeout_policy))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Required:

- `topology` (Block List, Min: 1, Max: 1) The cluster topology. (see [below for nested schema](#nestedblock--spec--topology))

Optional:

- `cluster_group_name` (String) Name of the cluster group to which this cluster belongs.
- `image_registry` (String) Name of the image registry configuration to use.
- `proxy_name` (String) Name of the proxy configuration to use.

Read-Only:

- `kubeconfig` (String) Cluster's kubeconfig.
- `tmc_managed` (Boolean) TMC-managed flag indicates if the cluster is managed by tmc.
(Default: False)

<a id="nestedblock--spec--topology"></a>
### Nested Schema for `spec.topology`

Required:

- `cluster_variables` (String) Variables configuration for the cluster.
- `control_plane` (Block List, Min: 1, Max: 1) Control plane specific configuration. (see [below for nested schema](#nestedblock--spec--topology--control_plane))
- `nodepool` (Block List, Min: 1) (Repeatable Block) Node pool definition for the cluster. (see [below for nested schema](#nestedblock--spec--topology--nodepool))
- `version` (String) Kubernetes version of the cluster.

Optional:

- `cluster_class` (String) The name of the cluster class for the cluster.
- `core_addon` (Block List) (Repeatable Block) The core addons. (see [below for nested schema](#nestedblock--spec--topology--core_addon))
- `network` (Block List, Max: 1) Network specific configuration. (see [below for nested schema](#nestedblock--spec--topology--network))

<a id="nestedblock--spec--topology--control_plane"></a>
### Nested Schema for `spec.topology.control_plane`

Required:

- `replicas` (Number) Number of replicas

Optional:

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--spec--topology--control_plane--meta))
- `os_image` (Block List, Max: 1) OS image block (see [below for nested schema](#nestedblock--spec--topology--control_plane--os_image))

<a id="nestedblock--spec--topology--control_plane--meta"></a>
### Nested Schema for `spec.topology.control_plane.meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource


<a id="nestedblock--spec--topology--control_plane--os_image"></a>
### Nested Schema for `spec.topology.control_plane.os_image`

Required:

- `arch` (String) The architecture of the OS image.
- `name` (String) The name of the OS image.
- `version` (String) The version of the OS image.



<a id="nestedblock--spec--topology--nodepool"></a>
### Nested Schema for `spec.topology.nodepool`

Required:

- `name` (String) Name of the node pool.
- `spec` (Block List, Min: 1) Spec for the node pool. (see [below for nested schema](#nestedblock--spec--topology--nodepool--spec))

Optional:

- `description` (String) Description of the node pool.

<a id="nestedblock--spec--topology--nodepool--spec"></a>
### Nested Schema for `spec.topology.nodepool.spec`

Required:

- `replicas` (Number) Number of replicas
- `worker_class` (String) The name of the machine deployment class used to create the node pool.

Optional:

- `auto_scaling` (Block List, Max: 1) Autoscaling block (see [below for nested schema](#nestedblock--spec--topology--nodepool--spec--auto_scaling))
- `failure_domain` (String) The failure domain the machines will be created in.
- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--spec--topology--nodepool--spec--meta))
- `os_image` (Block List, Max: 1) OS image block (see [below for nested schema](#nestedblock--spec--topology--nodepool--spec--os_image))
- `overrides` (String) Overrides can be used to override cluster level variables.

<a id="nestedblock--spec--topology--nodepool--spec--meta"></a>
### Nested Schema for `spec.topology.nodepool.spec.meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `labels` (Map of String) Labels for the resource


<a id="nestedblock--spec--topology--nodepool--spec--os_image"></a>
### Nested Schema for `spec.topology.nodepool.spec.os_image`

Required:

- `arch` (String) The architecture of the OS image.
- `name` (String) The name of the OS image.
- `version` (String) The version of the OS image.


<a id="nestedblock--spec--topology--nodepool--spec--auto_scaling"></a>
### Nested Schema for `spec.topology.nodepool.spec.auto_scaling`

Required:

- `enabled` (Boolean) Autoscaling config.
- `min_count` (Number) The minimum number of nodes for autoscaling.
- `max_count` (Number) The maximum number of nodes for autoscaling.




<a id="nestedblock--spec--topology--core_addon"></a>
### Nested Schema for `spec.topology.core_addon`

Required:

- `provider` (String) Provider of core add on
- `type` (String) Type of core add on


<a id="nestedblock--spec--topology--network"></a>
### Nested Schema for `spec.topology.network`

Optional:

- `pod_cidr_blocks` (List of String) Pod CIDR for Kubernetes pods defaults to 192.168.0.0/16.
- `service_cidr_blocks` (List of String) Service CIDR for kubernetes services defaults to 10.96.0.0/12.
- `service_domain` (String) Domain name for services.




<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource


<a id="nestedblock--timeout_policy"></a>
### Nested Schema for `timeout_policy`

Optional:

- `fail_on_timeout` (Boolean) Fail on timeout if timeout is reached and cluster is not ready. (Default = true)
- `timeout` (Number) Timeout in minutes for tanzu kubernetes creation process. A value of 0 means that no timeout is set. (Default: 60)
- `wait_for_kubeconfig` (Boolean) Wait for kubeconfig. (Default = true)
