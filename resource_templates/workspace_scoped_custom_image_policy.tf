// Create/ Delete/ Update Tanzu Mission Control workspace scoped custom image policy entry
resource "tanzu-mission-control_image_policy" "workspace_scoped_custom_image_policy" {
  name = "<image-policy-name>"

  scope {
    workspace {
      workspace = "<workspace-name>" // Required
    }
  }

  spec {
    input {
      custom {
        audit = false // Default: false
        rules {
          hostname = "<host-name>"
          imagename = "<image-name>"
          port = "<port-name>"
          requiredigest = false // Default: false
          tag {
            negate = false // Default: false
            value = "<tag-value>"
          }
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
