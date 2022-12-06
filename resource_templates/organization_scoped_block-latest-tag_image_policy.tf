// Create/ Delete/ Update Tanzu Mission Control organization scoped block-latest-tag image policy entry
resource "tanzu-mission-control_image_policy" "organization_scoped_block-latest-tag_image_policy" {
  name = "<image-policy-name>"

  scope {
    organization {
      organization = "<organization-id>" // Required
    }
  }

  spec {
    input {
      block_latest_tag {
        audit = false // Default: false
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
