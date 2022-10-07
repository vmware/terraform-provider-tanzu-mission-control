// Create/ Delete/ Update Tanzu Mission Control cluster group scoped tmc-require-labels custom policy entry
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_tmc-require-labels_custom_policy" {
  name = "<custom-policy-name>"

  scope {
    cluster_group {
      cluster_group = "<cluster-group-name>" // Required
    }
  }

  spec {
    input {
      tmc_require_labels {
        audit = false // Default: false
        parameters {
          labels {
            key   = "<label-key>"
            value = "<label-value>"
          }
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
