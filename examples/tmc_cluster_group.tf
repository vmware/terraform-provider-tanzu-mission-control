// TMC Cluster Group
// Operations supported : Create, read and delete

// Create TMC cluster group entry
resource "tmc_cluster_group" "create_cluster_group" {
  name = "<cluster-group--name>" // Required

  meta { // Optional
    description = "description of the cluster group"
    labels      = { "key" : "value" }
  }
}

// Create TMC cluster group entry with minimal information
resource "tmc_cluster_group" "cluster_group_create_min_info" {
  name = "<cluster-group-name>" // Required
}

// Read TMC cluster group : fetch cluster group details
data "tmc_cluster_group" "read_cluster_group" {
  name = "<cluster-name-group>" // Required
}
