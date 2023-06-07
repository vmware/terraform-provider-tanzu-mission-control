// Create/ Delete/ Update Tanzu Mission Control workspace scoped allow-all-to-pods network policy entry
resource "tanzu-mission-control_network_policy" "workspace_scoped_allow-all-to-pods_network_policy" {
  name = "<network-policy-name>"

  scope {
    workspace {
      workspace = "<workspace-name>" // Required
    }
  }

  spec {
    input {
      allow_all_to_pods {
        from_own_namespace = false // Default: false
        to_pod_labels = {
          "<pod-label-key-1>" = "<pod-label-value-1>"
          "<pod-label-key-2>" = "<pod-label-value-2>"
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
