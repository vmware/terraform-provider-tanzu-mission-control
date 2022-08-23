// Create/ Delete/ Update Tanzu Mission Control cluster group scoped baseline security policy entry
resource "tanzu-mission-control_security_policy" "cluster_group_scoped_baseline_security_policy" {
  scope {
    cluster_group{
      cluster_group = "<cluster-group-name>" // Required
    }
  }

  spec {
    input {
      baseline {
        audit              = false // Default: false
        disable_native_psp = false // Default: false
      }
    }

    namespace_selector {
      match_expressions {
        key      = "<label-selector-requirement-key-1>"
        operator = "<label-selector-requirement-operator>"
        values   = [
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
