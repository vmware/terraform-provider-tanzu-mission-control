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