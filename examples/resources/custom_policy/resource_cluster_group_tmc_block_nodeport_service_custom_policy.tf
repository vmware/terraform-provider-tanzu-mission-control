/*
Cluster group scoped Tanzu Mission Control custom policy with tmc-block-nodeport-service input recipe.
This policy is applied to a cluster group with the tmc-block-nodeport-service configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/
resource "tanzu-mission-control_custom_policy" "cluster_group_scoped_tmc-block-nodeport-service_custom_policy" {
  name = "tf-custom-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      tmc_block_nodeport_service {
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
