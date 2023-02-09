// Create/ Delete/ Update Tanzu Mission Control cluster scoped baseline security policy entry
resource "tanzu_mission_control_security_policy" "cluster_scoped_baseline_security_policy" {
  name = "<security-policy-name>"

  scope {
    cluster {
      management_cluster_name = "<management-cluster>" // Required
      provisioner_name        = "<prov-name>"          // Required
      name                    = "<cluster-name>"       // Required
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
