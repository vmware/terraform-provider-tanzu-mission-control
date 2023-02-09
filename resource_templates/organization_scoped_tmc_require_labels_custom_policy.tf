// Create/ Delete/ Update Tanzu Mission Control organization scoped tmc-require-labels custom policy entry
resource "tanzu_mission_control_custom_policy" "organization_scoped_tmc-require-labels_custom_policy" {
  name = "<custom-policy-name>"

  scope {
    organization {
      organization = "<organization-id>" // Required
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
            "<api-group>",
          ]
          kinds = [
            "<kind>",
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
