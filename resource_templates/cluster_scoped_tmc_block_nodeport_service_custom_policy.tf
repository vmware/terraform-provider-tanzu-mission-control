// Create/ Delete/ Update Tanzu Mission Control cluster scoped tmc-block-nodeport-service custom policy entry
resource "tanzu_mission_control_custom_policy" "cluster_scoped_tmc-block-nodeport-service_custom_policy" {
  name = "<custom-policy-name>"

  scope {
    cluster {
      management_cluster_name = "<management-cluster>" // Required
      provisioner_name        = "<prov-name>"          // Required
      name                    = "<cluster-name>"       // Required
    }
  }

  spec {
    input {
      tmc_block_nodeport_service {
        audit = false // Default: false
        target_kubernetes_resources {
          api_groups = [
            "<api-group-1>",
            "<api-group-2>"
          ]
          kinds = [
            "<kind-1>",
            "<kind-2>"
          ]
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "<label-selector-requirement-key-1>"
        operator = "<label-selector-requirement-operator>"
        values = [
          "<label-selector-requirement-value-1>",
          "<label-selector-requirement-value-2>"
        ]
      }
      match_expressions {
        key      = "<label-selector-requirement-key-2>"
        operator = "<label-selector-requirement-operator>"
        values   = []
      }
    }
  }
}
