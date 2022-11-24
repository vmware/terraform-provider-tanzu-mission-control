/*
Organization scoped Tanzu Mission Control custom policy with tmc-block-rolebinding-subjects input recipe.
This policy is applied to a organization with the tmc-block-rolebinding-subjects configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "organization_scoped_tmc-block-rolebinding-subjects_custom_policy" {
  name = "tf-custom-test"

  scope {
    organization {
      organization = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_block_rolebinding_subjects {
        audit = false
        parameters {
          disallowed_subjects {
            kind = "ServiceAccount"
            name = "subject-1"
          }
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
