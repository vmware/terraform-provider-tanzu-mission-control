---
Title: "Provisioning Node pool Resource"
Description: |-
    Adding the node pool resource to a cluster.
---

# Provisioning node pool on a cluster

For clusters that you create in Tanzu Mission Control, you can define a pool of worker nodes on which your workloads can run.
By default, each workload cluster that you create through Tanzu Mission Control has a node pool.
The `tanzu-mission-control_cluster_node_pool` resource allows you to create additional node pools, as well as read, update, and delete existing node pools in your clusters.
Because Tanzu Mission Control cannot provision additional resources in a cluster that is created elsewhere and subsequently attached, you cannot create a node pool in an attached cluster.

For provisioning of a cluster, refer to the `tanzu-mission-control_cluster` in guides.
For creating nodepool on the created cluster, one can use dependency in the terraform script.

You could create/manage a `node pool` for a cluster with the following config which shows a node pool resource with dependency on a TKGs cluster:

```terraform
/*
  NOTE: Creation of node pool depends on cluster creation
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create Tanzu Mission Control Tanzu Kubernetes Grid Service workload cluster entry
resource "tanzu-mission-control_cluster" "tkgs_workload_cluster" {
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
        version = "v1.21.6+vmware.1-tkg.1.b3d708a"
      }

      topology {
        control_plane {
          class             = "guaranteed-xsmall"
          storage_class     = "tkgs-k8s-obj-policy"
          high_availability = false
          volumes {
            capacity          = 4
            mount_path        = "/var/lib/etcd"
            name              = "etcd-0"
            pvc_storage_class = "tkgs-k8s-obj-policy"
          }
          volumes {
            capacity          = 4
            mount_path        = "/var/lib/etcd"
            name              = "etcd-1"
            pvc_storage_class = "tkgs-k8s-obj-policy"
          }
        }
        node_pools {
          spec {
            worker_node_count = "1"
            tkg_service_vsphere {
              class         = "best-effort-xsmall"
              storage_class = "gc-storage-profile"
              volumes {
                capacity          = 4
                mount_path        = "/var/lib/etcd"
                name              = "etcd-0"
                pvc_storage_class = "tkgs-k8s-obj-policy"
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

# Create Tanzu Mission Control nodepool entry
resource "tanzu-mission-control_cluster_node_pool" "node_pool" {

  management_cluster_name = "tkgs-terraform"
  provisioner_name        = "test-gc-e2e-demo-ns"
  cluster_name            = tanzu-mission-control_cluster.tkgs_workload_cluster.name
  name                    = "terraform-nodepool"

  spec {
    worker_node_count = "3"

    tkg_service_vsphere {
      class         = "best-effort-xsmall"
      storage_class = "gc-storage-profile"
    }
  }
}
```
