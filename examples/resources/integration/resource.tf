# Create Tanzu Mission Control TSM Integration resource
resource "tanzu-mission-control_integration" "create_tsm-integration" {
  management_cluster_name = "attached"
  provisioner_name        = "attached"
  cluster_name            = "test-cluster"
  integration_name        = "tanzu-service-mesh"

  spec {
    configurations = jsonencode({
      enableNamespaceExclusions = true
      namespaceExclusions = [
        {
          match = "custom-namespace-1"
          type  = "EXACT"
          }, {
          match = "kube"
          type  = "START_WITH"
        }
      ]
    })
  }
}
