# Read Tanzu Mission Control workspace : fetch workspace details
data "tanzu_mission_control_workspace" "read_workspace" {
  name = "default"
}
