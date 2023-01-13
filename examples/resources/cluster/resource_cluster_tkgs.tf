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
            cloud_labels = {
              "key1" : "val1"
            }
            node_labels = {
              "key2" : "val2"
            }

            tkg_service_vsphere {
              class         = "best-effort-xsmall"
              storage_class = "gc-storage-profile"
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