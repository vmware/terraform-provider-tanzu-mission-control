/*
Organization scoped Tanzu Mission Control network policy with allow-all input recipe.
This policy is applied to a organization with the allow-all configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_network_policy" "organization_scoped_allow-all_network_policy" {
  name = "tf-network-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      allow_all {
        from_own_namespace = false
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
