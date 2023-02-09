// Create/ Delete/ Update Tanzu Mission Control cluster group scoped tmc-external-ips custom policy entry
resource "tanzu_mission_control_custom_policy" "cluster_group_scoped_tmc-external-ips_custom_policy" {
  name = "<custom-policy-name>"

  scope {
    cluster_group {
      cluster_group = "<cluster-group-name>" // Required
    }
  }

  spec {
    input {
      tmc_external_ips {
        audit = false // Default: false
        parameters {
          allowed_ips = [
            "<allowed-ip-address-1>",
            "<allowed-ip-address-2>"
          ]
        }
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
