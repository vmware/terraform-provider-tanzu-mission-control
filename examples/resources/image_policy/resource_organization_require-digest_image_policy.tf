/*
Organization scoped Tanzu Mission Control image policy with require-digest input recipe.
This policy is applied to a organization with the require-digest configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu_mission_control_image_policy" "organization_scoped_require-digest_image_policy" {
  name = "tf-image-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      require_digest {
        audit = false
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
