// TMC Namespace:
// Operations supported : Read, Create, Update & Delete

// Create TMC namespace entry
resource "tmc_namespace" "namespace_create" {
  name                    = "<namespace-name>" // Required
  cluster_name            = "<cluster_name>" // Required
  provisioner_name        = "<prov-name>" // Default: attached
  management_cluster_name = "<management-cluster>" // Default: attached

  meta  {
    description    = "description of the namespace" // Optional
    labels         = { "key" : "value" }
  }

  spec  {
    workspace_name = "<workspace-name>" // Default: default
    attach         = false // Default: false
  }
}

// Create TMC namespace entry with minimal information
resource "tmc_namespace" "namespace_create_min_info" {
  name         = "<namespace-name>"     // Required
  cluster_name = "<cluster_name>" // Required
}