// TMC Cluster Group
// Operations supported : Create and delete

// Create TMC cluster group entry
resource "tmc_cluster_group" "cluster_group_create" {
    name = "<cluster-group--name>"     // Required

    meta {  // Optional
        description = "description of the cluster group"
        labels      = { "key" : "value" }
        }
    }

// Create TMC cluster group entry with minimal information
resource "tmc_cluster_group" "cluster_group_create_min_info" {
  name = "<cluster-group-name>" // Required
}
