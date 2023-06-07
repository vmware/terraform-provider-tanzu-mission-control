// Create/ Delete/ Update Tanzu Mission Control workspace scoped allow-all network policy entry
resource "tanzu-mission-control_network_policy" "workspace_scoped_allow-all_network_policy" {
  name = "<network-policy-name>"

  scope {
    workspace {
      workspace = "<workspace-name>" // Required
    }
  }

  spec {
    input {
      allow_all {
        from_own_namespace = false // Default: false
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
