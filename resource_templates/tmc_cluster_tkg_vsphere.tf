// TMC Cluster Type: TKGVsphere workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read TMC TKG Vsphere workload cluster : fetch cluster details
data "tmc_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<management-cluster>" // Required
  provisioner_name        = "<prov-name>"          // Required
  name                    = "<cluster-name>"       // Required
}

// Create TMC TKG Vsphere workload cluster entry
resource "tmc_cluster" "create_tkg_vsphere_cluster" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  name                    = "<cluster-name>"       // Required

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = "<cluster-group>" // Default: default
    tkg_vsphere {
      settings {
        network {
          pods {
            cidr_blocks = [
              "<pods-cidr-blocks>", // Required
            ]
          }

          services {
            cidr_blocks = [
              "<services-cidr-blocks>", // Required
            ]
          }

          control_plane_end_point = "<end-point>" // Required
        }

        security {
          ssh_key = "<ssh-key>" // Required
        }
      }

      distribution {
        version = "<version>" // Required

        workspace {
          datacenter        = "<datacenter>" // Required
          datastore         = "<datastore>" // Required
          workspace_network = "<workspace_network>" // Required
          folder            = "<folder>" // Required
          resource_pool     = "<resource_pool>" // Required
        }
      }

      topology {
        control_plane {
          vm_config {
            cpu       = "<cpu>" // Required
            disk_size = "<disk_size>" // Required
            memory    = "<memory>" // Required
          }

          high_availability = false // Default: false
        }

        node_pools {
          node_pool_spec {
            worker_node_count = "<worker-node-count>" // Required

            tkg_vsphere {
              vm_config {
                cpu       = "<cpu>" // Required
                disk_size = "<disk_size>" // Required
                memory    = "<memory>" // Required
              }
            }
          }

          node_pool_info {
            node_pool_name        = "<node-pool-name>" // Required
            node_pool_description = "<node-pool-description>"
          }
        }
      }
    }
  }
}