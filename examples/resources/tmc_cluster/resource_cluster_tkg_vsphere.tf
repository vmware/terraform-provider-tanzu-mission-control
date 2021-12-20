# Create a TKGm Vsphere workload cluster entry
resource "tmc_cluster" "create_tkg_vsphere_cluster" {
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