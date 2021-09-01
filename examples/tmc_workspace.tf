// TMC Workspace:
// Operations supported : Read, Create & Delete

// Create TMC workspace entry
resource "tmc_workspace" "workspace_create" {
  name = "<workspace-name>" // Required
  meta {                    // Optional
    description = "description of the workspace"
    labels      = { "key" : "value" }
  }
}

// Create TMC workspace entry with minimal information
resource "tmc_workspace" "workspace_create_min_info" {
  name = "<workspace-name>" // Required
}

// Read TMC workspace entry
data "tmc_workspace" "workspace_read" {
  name = "<workspace-name>"
}
