/*
Workspace scoped Tanzu Mission Control network policy with deny-all-egress input recipe.
This policy is applied to a workspace with the deny-all-egress configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "workspace_scoped_deny-all-egress_network_policy" {
  name = "tf-network-test"

  scope {
    workspace {
      workspace = "tf-workspace"
    }
  }

  spec {
    input {
      deny_all_egress {}
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
