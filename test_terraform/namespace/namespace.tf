terraform {
  required_providers {
    tmc = {
      source = "vmware/tanzu/tmc"
    }
  }
}
// Create namespace with attached set as default value.
resource "tmc_namespace" "namespace" {
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
resource "tmc_namespace" "namespace" {
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

// Output namespace resource
output "namespace" {
  value = tmc_namespace.namespace
}
