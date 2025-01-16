// Tanzu Mission Control Workspace
// Operations supported : Read, Create, Update & Delete

// Create Tanzu Mission Control workspace entry
resource "tanzu-mission-control_workspace" "workspace" {
  name = "<workspace-name>" // Required
  meta {                    // Optional
    description = "description of the workspace"
    labels      = { "key" : "value" }
  }
}

// Create Tanzu Mission Control workspace entry with minimal information
resource "tanzu-mission-control_workspace" "workspace_min_info" {
  name = "<workspace-name>" // Required
}

// Read Tanzu Mission Control workspace entry : fetch workspace details
data "tanzu-mission-control_workspace" "workspace_read" {
  name = "<workspace-name>"
}
