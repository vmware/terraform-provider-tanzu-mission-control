// Tanzu Mission Control Namespace
// Operations supported : Read, Create, Update & Delete

// Read Tanzu Mission Control namespace : fetch namespace details
data "tanzu-mission-control_namespace" "read_namespace" {
  name                    = "<namespace-name>"     // Required
  cluster_name            = "<cluster-name>"       // Required
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
}

// Create Tanzu Mission Control namespace entry
resource "tanzu-mission-control_namespace" "namespace" {
  name                    = "<namespace-name>"     // Required
  cluster_name            = "<cluster_name>"       // Required
  provisioner_name        = "<prov-name>"          // Default: attached
  management_cluster_name = "<management-cluster>" // Default: attached

  meta {
    description = "description of the namespace" // Optional
    labels      = { "key" : "value" }
  }

  spec {
    workspace_name = "<workspace-name>" // Default: default
    attach         = false              // Default: false
  }
}

// Create Tanzu Mission Control namespace entry with minimal information
resource "tanzu-mission-control_namespace" "namespace_min_info" {
  name         = "<namespace-name>" // Required
  cluster_name = "<cluster_name>"   // Required
}