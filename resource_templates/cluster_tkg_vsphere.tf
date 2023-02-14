// Tanzu Mission Control Cluster Type: Tanzu Kubernetes Grid vSphere workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read Tanzu Mission Control Tanzu Kubernetes Grid vSphere workload cluster : fetch cluster details
data "tanzu-mission-control_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<management-cluster>" // Required
  provisioner_name        = "<prov-name>"          // Required
  name                    = "<cluster-name>"       // Required
}

// Create Tanzu Mission Control Tanzu Kubernetes Grid vSphere workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkg_vsphere_cluster" {
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
      advanced_configs {
        key   = "<key>"
        value = "<value>"
      }
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

          api_server_port         = api-server-port-default-value
          control_plane_end_point = "<end-point>" // Required
        }

        security {
          ssh_key = "<ssh-key>" // Required
        }
      }

      distribution {
        os_arch    = "<os-arch>"
        os_name    = "<os-name>"
        os_version = "<os-version>"
        version    = "<version>" // Required

        workspace {
          datacenter        = "<datacenter>"        // Required
          datastore         = "<datastore>"         // Required
          workspace_network = "<workspace_network>" // Required
          folder            = "<folder>"            // Required
          resource_pool     = "<resource_pool>"     // Required
        }
      }

      topology {
        control_plane {
          vm_config {
            cpu       = "<cpu>"       // Required
            disk_size = "<disk_size>" // Required
            memory    = "<memory>"    // Required
          }

          high_availability = false // Default: false
        }

        node_pools {
          spec {
            worker_node_count = "<worker-node-count>" // Required
            cloud_label = {
              "<key>" : "<val>"
            }
            node_label = {
              "<key>" : "<val>"
            }

            tkg_vsphere {
              vm_config {
                cpu       = "<cpu>"       // Required
                disk_size = "<disk_size>" // Required
                memory    = "<memory>"    // Required
              }
            }
          }

          info {
            name        = "<node-pool-name>" // Required
            description = "<node-pool-description>"
          }
        }
      }
    }
  }
}