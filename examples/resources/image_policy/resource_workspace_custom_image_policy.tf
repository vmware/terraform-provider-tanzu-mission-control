/*
Workspace scoped Tanzu Mission Control image policy with custom input recipe.
This policy is applied to a workspace with the custom configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu_mission_control_image_policy" "workspace_scoped_custom_image_policy" {
  name = "tf-image-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      custom {
        audit = true
        rules {
          hostname = "foo"
          imagename = "bar"
          port = "80"
          requiredigest = false
          tag {
            negate = false
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
