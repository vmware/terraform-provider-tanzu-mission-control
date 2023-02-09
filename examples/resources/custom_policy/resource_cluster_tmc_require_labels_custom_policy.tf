/*
Cluster scoped Tanzu Mission Control custom policy with tmc-require-labels input recipe.
This policy is applied to a cluster with the tmc-require-labels configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu_mission_control_custom_policy" "cluster_scoped_tmc-require-labels_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_require_labels {
        audit = false
        parameters {
          labels {
            key   = "test"
            value = "optional"
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
