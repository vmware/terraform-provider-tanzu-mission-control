/*
Organization scoped Tanzu Mission Control custom policy with tmc-external-ips input recipe.
This policy is applied to a organization with the tmc-external-ips configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu_mission_control_custom_policy" "organization_scoped_tmc-external-ips_custom_policy" {
  name = "tf-custom-test"

  scope {
    organization {
      organization = "dummy-id"
    }
  }

  spec {
    input {
      tmc_external_ips {
        audit = false
        parameters {
          allowed_ips = [
            "127.0.0.1",
          ]
        }
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
