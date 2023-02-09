// Create/ Delete/ Update Tanzu Mission Control workspace scoped allowed-name-tag image policy entry
resource "tanzu_mission_control_image_policy" "workspace_scoped_allowed-name-tag_image_policy" {
  name = "<image-policy-name>"

  scope {
    workspace {
      workspace = "<workspace-name>" // Required
    }
  }

  spec {
    input {
      allowed_name_tag {
        audit = false // Default: false
        rules {
          imagename = "<image-name>"
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
