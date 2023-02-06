/*
Cluster scoped Tanzu Mission Control namespace quota policy with small input recipe.
This policy is applied to a cluster with the small configuration option.
The defined scope and input blocks can be updated to change the policy's scope and recipe, respectively.
*/

resource "tanzu-mission-control_namespace_quota_policy" "cluster_scoped_small_quota_policy" {
  name = "tf-qt-test"

  scope {
    cluster {
      management_cluster_name = "attached"
      provisioner_name        = "attached"
      name                    = "tf-create-test"
    }
  }

  spec {
    input {
      small {} // Pre-defined parameters for Small Namespace quota Policy: CPU requests = 0.5 vCPU, Memory requests = 512 MB, CPU limits = 1 vCPU, Memory limits = 2 GB
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
