# Read Tanzu Mission Control workspace : fetch workspace details
data "tanzu-mission-control_workspace" "read_workspace" {
  name = "default"
}
