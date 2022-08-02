// Tanzu Mission Control Cluster Group
// Operations supported : Create, Read, Update and Delete

// Create Tanzu Mission Control cluster group entry
resource "tanzu-mission-control_cluster_group" "create_cluster_group" {
  name = "<cluster-group-name>" // Required

  meta { // Optional
    description = "description of the cluster group"
    labels      = { "key" : "value" }
  }
}

// Create Tanzu Mission Control cluster group entry with minimal information
resource "tanzu-mission-control_cluster_group" "cluster_group_create_min_info" {
  name = "<cluster-group-name>" // Required
}

// Read Tanzu Mission Control cluster group : fetch cluster group details
data "tanzu-mission-control_cluster_group" "read_cluster_group" {
  name = "<cluster-name-group>" // Required
}
