/*
Organization scoped Tanzu Mission Control image policy with custom input recipe.
This policy is applied to a organization with the custom configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_image_policy" "organization_scoped_custom_image_policy" {
  name = "tf-image-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      custom {
        audit = true
        rules {
          hostname      = "foo"
          imagename     = "bar"
          port          = "80"
          requiredigest = false
          tag {
            negate = false
            value  = "test"
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
