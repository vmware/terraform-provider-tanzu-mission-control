terraform {
  required_providers {
    tmc = {
      source = "vmware/tanzu/tmc"
    }
  }
}

// Create workspace
resource "tmc_workspace" "workspace" {
    workspace_name = "tf-workspace"
    meta  {
        description    = "Create workspace through terraform"
        labels         = {
            "key1" : "value1",
            "key2" : "value2"
        }
    }
}

// Output workspace resource
output "workspace" {
    value = tmc_workspace.workspace
}
