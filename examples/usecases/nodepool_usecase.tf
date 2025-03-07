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
