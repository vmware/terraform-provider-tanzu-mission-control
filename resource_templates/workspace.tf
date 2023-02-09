// Tanzu Mission Control Workspace
// Operations supported : Read, Create, Update & Delete

// Create Tanzu Mission Control workspace entry
resource "tanzu_mission_control_workspace" "create_workspace" {
  name = "<workspace-name>" // Required
  meta {                    // Optional
    description = "description of the workspace"
    labels      = { "key" : "value" }
  }
}

// Create Tanzu Mission Control workspace entry with minimal information
resource "tanzu_mission_control_workspace" "create_workspace_min_info" {
  name = "<workspace-name>" // Required
}

// Read Tanzu Mission Control workspace entry : fetch workspace details
data "tanzu_mission_control_workspace" "workspace_read" {
  name = "<workspace-name>"
}
