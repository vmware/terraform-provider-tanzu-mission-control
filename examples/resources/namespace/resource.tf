# Create Tanzu Mission Control namespace with attached set as default value.
resource "tanzu-mission-control_namespace" "create_namespace" {
  name                    = "tf-namespace" # Required
  cluster_name            = "testcluster"  # Required
  provisioner_name        = "attached"     # Default: attached
  management_cluster_name = "attached"     # Default: attached

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    workspace_name = "default" # Default: default
    attach         = false     # Default: false
  }
}

# Create Tanzu Mission Control namespace with attached set as 'true'
resource "tanzu-mission-control_namespace" "create_namespace_attached" {
  name                    = "tf-namespace" # Required
  cluster_name            = "testcluster"  # Required
  provisioner_name        = "attached"     # Default: attached
  management_cluster_name = "attached"     # Default: attached

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    workspace_name = "default" # Default: default
    attach         = true
  }
}
