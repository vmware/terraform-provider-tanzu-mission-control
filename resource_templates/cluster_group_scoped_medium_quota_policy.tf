// Create/ Delete/ Update Tanzu Mission Control cluster group scoped medium namespace quota policy entry
resource "tanzu-mission-control_namespace_quota_policy" "cluster_group_scoped_medium_quota_policy" {
  name = "<namespace-quota-policy-name>"

  scope {
    cluster_group {
      cluster_group = "<cluster-group-name>" // Required
    }
  }

  spec {
    input {
      medium {}
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
