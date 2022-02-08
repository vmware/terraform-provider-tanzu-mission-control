# Read Tanzu Mission Control cluster group : fetch cluster group details
data "tanzu-mission-control_cluster_group" "read_cluster_group" {
  name = "default"
}