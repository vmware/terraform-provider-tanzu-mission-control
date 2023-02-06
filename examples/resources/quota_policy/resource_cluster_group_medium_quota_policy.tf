/*
Cluster group scoped Tanzu Mission Control namespace quota policy with medium input recipe.
This policy is applied to a cluster group with the medium configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_group_scoped_medium_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      medium {} // Pre-defined parameters for Medium Namespace quota Policy: CPU requests = 1 vCPU, Memory requests = 1 GB, CPU limits = 2 vCPU, Memory limits = 4 GB
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
