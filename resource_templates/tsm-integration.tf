// Tanzu Mission Control TSM Integration
// Operations supported : Create, Read and Delete

// Create Tanzu Mission Control TSM Integration resource
resource "tanzu_mission_control_integration" "create_tsm-integration" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  cluster_name            = "<cluster_name>"       // Required
  integration_name        = "<integration-name>"

  spec {
    configurations = jsonencode({
      enableNamespaceExclusions = false // Default: false
      namespaceExclusions = [
        {
          match = "<namespace-match-1>"
          type  = "<namespace-type-1>"
          }, {
          match = "<namespace-match-2>"
          type  = "<namespace-type-2>"
        }
      ]
    })
  }
}

// Read Tanzu Mission Control TSM integration : fetch details
data "tanzu_mission_control_integration" "read_tsm-integration" {
  management_cluster_name = "<management-cluster>" // Default: attached
  provisioner_name        = "<prov-name>"          // Default: attached
  cluster_name            = "<cluster_name>"       // Required
  integration_name        = "<integration-name>"
}
