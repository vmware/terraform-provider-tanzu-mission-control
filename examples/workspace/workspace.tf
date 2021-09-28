terraform {
  required_providers {
    tmc = {
      source = "vmware/tanzu/tmc"
    }
  }
}

// Create workspace
resource "tmc_workspace" "create_workspace" {
  name = "tf-workspace-test"
  meta {
    description = "Create workspace through terraform"
    labels = {
      "key1" : "value1",
      "key2" : "value2"
    }
  }
}

// Read workspace
data "tmc_workspace" "read_workspace" {
  name = "default"
}

// Output workspace resource
output "workspace" {
  value = tmc_workspace.create_workspace
}

output "display_workspace" {
  value = data.tmc_workspace.read_workspace
}
