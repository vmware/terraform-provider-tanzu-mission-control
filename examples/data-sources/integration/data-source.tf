data "tanzu-mission-control_cluster" "default" {
  management_cluster_name = "attached"       # Default: attached
  provisioner_name        = "attached"       # Default: attached
  name                    = "terraform-test" # Required
}

data "tanzu-mission-control_integration" "default" {
  management_cluster_name = "attached"
  provisioner_name        = "attached"
  cluster_name            = tanzu-mission-control_cluster.default.name
  integration_name        = "tanzu-service-mesh"
}
