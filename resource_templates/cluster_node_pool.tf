// Tanzu Mission Control Nodepool
// Operations supported : Read, Create, Update & Delete

# Read Tanzu Mission Control cluster nodepool : fetch cluster nodepool details
data "tanzu-mission-control_cluster_node_pool" "read_node_pool" {
  management_cluster_name = "<existing-management-cluster>" // Default: attached
  provisioner_name        = "<existing-prov-name>"          // Default: attached
  cluster_name            = "<existing-cluster_name>"       // Required
  name                    = "<existing-node_pool-name>"     // Required
}

# Create Tanzu Mission Control cluster nodepool entry
resource "tanzu-mission-control_cluster_node_pool" "create_node_pool" {

  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  cluster_name            = "<cluster_name>"       // Required
  name                    = "<node_pool-name>"     // Required

  spec {
    worker_node_count = "<worker-nodes>"
    cloud_labels = {
      "<key>" : "<val>"
    }
    node_labels = {
      "<key>" : "<val>"
    }

    tkg_service_vsphere {
      class         = "<class>"        // Required
      storage_class = "<storage-class" // Required
    }
  }
}