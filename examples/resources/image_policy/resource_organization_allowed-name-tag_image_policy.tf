/*
Organization scoped Tanzu Mission Control image policy with allowed-name-tag input recipe.
This policy is applied to a organization with the allowed-name-tag configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu_mission_control_image_policy" "organization_scoped_allowed-name-tag_image_policy" {
  name = "tf-image-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      allowed_name_tag {
        audit = true
        rules {
          imagename = "bar"
          tag {
            negate = true
            value = "test"
          }
        }
      }
    }

    namespace_selector {
      match_expressions {
        key      = "component"
        operator = "In"
        values = [
          "api-server",
          "agent-gateway"
        ]
      }
      match_expressions {
        key      = "not-a-component"
        operator = "DoesNotExist"
        values   = []
      }
    }
  }
}
