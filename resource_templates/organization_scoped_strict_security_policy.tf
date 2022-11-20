// Create/ Delete/ Update Tanzu Mission Control organization scoped strict security policy entry
resource "tanzu-mission-control_security_policy" "organization_scoped_strict_security_policy" {
  name = "<security-policy-name>"

  scope {
    organization {
      organization = "<organization-id>" // Required
    }
  }

  spec {
    input {
      strict {
        audit              = false // Default: false
        disable_native_psp = false // Default: false
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
