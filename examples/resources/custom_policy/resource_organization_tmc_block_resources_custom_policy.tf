/*
Organization scoped Tanzu Mission Control custom policy with tmc-block-resources input recipe.
This policy is applied to a organization with the tmc-block-resources configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "organization_scoped_tmc-block-resources_custom_policy" {
  name = "tf-custom-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      tmc_block_resources {
        audit = false
        target_kubernetes_resources {
          api_groups = [
            "apps",
          ]
          kinds = [
            "Event",
          ]
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
