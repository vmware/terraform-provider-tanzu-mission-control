// Tanzu Mission Control Cluster Type: Tanzu Kubernetes Grid Service workload.
// Operations supported : Read, Create, Update & Delete (except nodepools)

// Read Tanzu Mission Control Tanzu Kubernetes Grid Service workload cluster : fetch cluster details
data "tanzu-mission-control_cluster" "ready_only_cluster_view" {
  management_cluster_name = "<management-cluster>" // Required
  provisioner_name        = "<prov-name>"          // Required
  name                    = "<cluster-name>"       // Required
}

# Create Tanzu Mission Control Tanzu Kubernetes Grid Service workload cluster entry
resource "tanzu-mission-control_cluster" "create_tkgs_workload" {
  management_cluster_name = "<management-cluster>"
  provisioner_name        = "<prov-name>"
  name                    = "<cluster-name>"

  meta {
    labels = { "key" : "test" }
  }

  spec {
    cluster_group = "<cluster-group>" // Default: default
    tkg_service_vsphere {
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
        }
        storage {
          classes = [
            "<storage-classes>",
          ]
          default_class = "<default-storage-class>"
        }
      }

      distribution {
        version = "<version>" // Required
      }

      topology {
        control_plane {
          class             = "<class>"         // Required
          storage_class     = "<storage-class>" // Required
          high_availability = false             // Default: false
          volumes {
            capacity          = volume-capacity
            mount_path        = "<mount-path>"
            name              = "<volume-name>"
            pvc_storage_class = "<storage-class>"
          }
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
            tkg_service_vsphere {
              class          = "<class>"         // Required
              storage_class  = "<storage-class>" // Required
              failure_domain = "<failure-domain>"
              volumes {
                capacity          = volume-capacity
                mount_path        = "<mount-path>"
                name              = "<volume-name>"
                pvc_storage_class = "<storage-class>"
              }
            }
          }
          info {
            name = "<node-pool-name>" // Required
          }
        }
      }
    }
  }
}
