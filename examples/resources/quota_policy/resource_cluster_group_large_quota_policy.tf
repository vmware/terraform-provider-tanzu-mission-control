/*
Cluster group scoped Tanzu Mission Control namespace quota policy with large input recipe.
This policy is applied to a cluster group with the large configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_group_scoped_large_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster_group {
      cluster_group = "tf-create-test"
    }
  }

  spec {
    input {
      large {} // Pre-defined parameters for Large Namespace quota Policy: CPU requests = 2 vCPU, Memory requests = 2 GB, CPU limits = 4 vCPU, Memory limits = 8 GB
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
