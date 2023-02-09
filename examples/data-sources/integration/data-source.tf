# Read Tanzu Mission Control TSM integration : fetch details
data "tanzu_mission_control_integration" "read_tsm-integration" {
  management_cluster_name = "attached"
  provisioner_name        = "attached"
  cluster_name            = "test-cluster"
  integration_name        = "tanzu-service-mesh"
}
