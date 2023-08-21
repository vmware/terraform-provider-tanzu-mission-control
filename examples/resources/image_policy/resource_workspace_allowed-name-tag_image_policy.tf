/*
Workspace scoped Tanzu Mission Control image policy with allowed-name-tag input recipe.
This policy is applied to a workspace with the allowed-name-tag configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_image_policy" "workspace_scoped_allowed-name-tag_image_policy" {
  name = "tf-image-test"

  scope {
    workspace {
      workspace = "tf-workspace"
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
