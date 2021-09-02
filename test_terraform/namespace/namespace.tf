terraform {
  required_providers {
    tmc = {
      source = "vmware/tanzu/tmc"
    }
  }
}
// Create namespace with attached set as default value.
resource "tmc_namespace" "create_namespace" {
  name                    = "tf-namespace" // Required
  cluster_name            = "testcluster"  // Required
  provisioner_name        = "attached"     // Default: attached
  management_cluster_name = "attached"     // Default: attached

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    workspace_name = "default" // Default: default
    attach         = false     // Default: false
  }
}

// Create namespace with attached set as 'true'
resource "tmc_namespace" "create_namespace_attached" {
  name                    = "tf-namespace" // Required
  cluster_name            = "testcluster"  // Required
  provisioner_name        = "attached"     // Default: attached
  management_cluster_name = "attached"     // Default: attached

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    workspace_name = "default" // Default: default
    attach         = true
  }
}
// Read TMC namespace : fetch namespace details
data "tmc_namespace" "read_namespace" {
  name                    = "tf-namespace" // Required
  cluster_name            = "testcluster"  // Required
  management_cluster_name = "attached"     // Default: attached
  provisioner_name        = "attached"     // Default: attached
}

// Output namespace resource
output "namespace" {
  value = tmc_namespace.create_namespace
}
// Get namespace resource
output "display_namespace" {
  value = data.tmc_namespace.read_namespace
}
