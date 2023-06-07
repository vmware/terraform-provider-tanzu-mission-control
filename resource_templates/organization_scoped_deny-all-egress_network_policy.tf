// Create/ Delete/ Update Tanzu Mission Control organization scoped deny-all-egress network policy entry
resource "tanzu-mission-control_network_policy" "organization_scoped_deny-all-egress_network_policy" {
  name = "<network-policy-name>"

  scope {
    organization {
      organization = "<organization-id>" // Required
    }
  }

  spec {
    input {
      deny_all_egress {}
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
