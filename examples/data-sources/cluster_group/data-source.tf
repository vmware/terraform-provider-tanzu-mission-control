# Read Tanzu Mission Control cluster group : fetch cluster group details
data "tanzu_mission_control_cluster_group" "read_cluster_group" {
  name = "default"
}
